package dockermonitor

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/aporeto-inc/trireme/collector"
	"github.com/aporeto-inc/trireme/constants"
	"github.com/aporeto-inc/trireme/monitor"
	"github.com/aporeto-inc/trireme/policy"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"

	dockerClient "github.com/docker/docker/client"
)

// DockerEvent is the type of various docker events.
type DockerEvent string

const (
	// DockerEventCreate represents the Docker "create" event.
	DockerEventCreate DockerEvent = "create"

	// DockerEventStart represents the Docker "start" event.
	DockerEventStart DockerEvent = "start"

	// DockerEventDie represents the Docker "die" event.
	DockerEventDie DockerEvent = "die"

	// DockerEventDestroy represents the Docker "destroy" event.
	DockerEventDestroy DockerEvent = "destroy"

	// DockerEventPause represents the Docker "destroy" event.
	DockerEventPause DockerEvent = "pause"

	// DockerEventUnpause represents the Docker "destroy" event.
	DockerEventUnpause DockerEvent = "unpause"

	// DockerEventConnect represents the Docker "connect" event.
	DockerEventConnect DockerEvent = "connect"

	// DockerClientVersion is the version sent out as the client
	DockerClientVersion = "v1.23"
)

// A DockerEventHandler is type of docker event handler functions.
type DockerEventHandler func(event *events.Message) error

// A DockerMetadataExtractor is a function used to extract a *policy.PURuntime from a given
// docker ContainerJSON.
type DockerMetadataExtractor func(*types.ContainerJSON) (*policy.PURuntime, error)

func contextIDFromDockerID(dockerID string) (string, error) {

	if dockerID == "" {
		return "", fmt.Errorf("Empty DockerID String")
	}

	if len(dockerID) < 12 {
		return "", fmt.Errorf("dockerID smaller than 12 characters")
	}

	return dockerID[:12], nil
}

func initDockerClient(socketType string, socketAddress string) (*dockerClient.Client, error) {

	var socket string

	switch socketType {
	case "tcp":
		socket = "https://" + socketAddress

	case "unix":
		// Sanity check that this path exists
		if _, oserr := os.Stat(socketAddress); os.IsNotExist(oserr) {
			return nil, oserr
		}
		socket = "unix://" + socketAddress

	default:
		return nil, fmt.Errorf("Bad socket type %s", socketType)
	}

	defaultHeaders := map[string]string{"User-Agent": "engine-api-dockerClient-1.0"}
	dockerClient, err := dockerClient.NewClient(socket, DockerClientVersion, nil, defaultHeaders)

	if err != nil {
		return nil, fmt.Errorf("Error creating Docker Client %s", err.Error())
	}

	return dockerClient, nil
}

func defaultDockerMetadataExtractor(info *types.ContainerJSON) (*policy.PURuntime, error) {

	tags := policy.NewTagsMap(map[string]string{
		"@sys:image": info.Config.Image,
		"@sys:name":  info.Name,
	})

	for k, v := range info.Config.Labels {
		tags.Add("@usr:"+k, v)
	}

	ipa := policy.NewIPMap(map[string]string{
		"bridge": info.NetworkSettings.IPAddress,
	})

	return policy.NewPURuntime(info.Name, info.State.Pid, tags, ipa, constants.ContainerPU, nil), nil
}

// dockerMonitor implements the connection to Docker and monitoring based on events
type dockerMonitor struct {
	dockerClient       *dockerClient.Client
	metadataExtractor  DockerMetadataExtractor
	handlers           map[DockerEvent]func(event *events.Message) error
	eventnotifications chan *events.Message
	stopprocessor      chan bool
	stoplistener       chan bool
	syncAtStart        bool
	syncHandler        monitor.SynchronizationHandler

	collector collector.EventCollector
	puHandler monitor.ProcessingUnitsHandler
}

// NewDockerMonitor returns a pointer to a DockerMonitor initialized with the given
// socketType ('tcp' or 'unix') and socketAddress (a port for 'tcp' or
// a socket file for 'unix').
//
// After creating a new DockerMonitor, call addHandler to install one
// or more callback handlers for the events to monitor. Then call Start.
func NewDockerMonitor(
	socketType string,
	socketAddress string,
	p monitor.ProcessingUnitsHandler,
	m DockerMetadataExtractor,
	l collector.EventCollector,
	syncAtStart bool,
	s monitor.SynchronizationHandler,
) monitor.Monitor {

	cli, err := initDockerClient(socketType, socketAddress)

	if err != nil {
		log.WithFields(log.Fields{
			"package": "monitor",
			"error":   err.Error(),
		}).Fatal("Unable to initialize Docker client")
	}

	d := &dockerMonitor{
		puHandler:          p,
		collector:          l,
		eventnotifications: make(chan *events.Message, 1000),
		handlers:           make(map[DockerEvent]func(event *events.Message) error),
		stoplistener:       make(chan bool),
		stopprocessor:      make(chan bool),
		metadataExtractor:  m,
		dockerClient:       cli,
		syncAtStart:        syncAtStart,
		syncHandler:        s,
	}

	// Add handlers for the events that we know how to process
	d.addHandler(DockerEventCreate, d.handleCreateEvent)
	d.addHandler(DockerEventStart, d.handleStartEvent)
	d.addHandler(DockerEventDie, d.handleDieEvent)
	d.addHandler(DockerEventDestroy, d.handleDestroyEvent)
	d.addHandler(DockerEventPause, d.handlePauseEvent)
	d.addHandler(DockerEventUnpause, d.handleUnpauseEvent)

	return d
}

// addHandler adds a callback handler for the given docker event.
// Interesting event names include 'start' and 'die'. For more on events see
// https://docs.docker.com/engine/reference/api/docker_remote_api/
// under the section 'Docker Events'.
func (d *dockerMonitor) addHandler(event DockerEvent, handler DockerEventHandler) {
	d.handlers[event] = handler
}

// Start will start the DockerPolicy Enforcement.
// It applies a policy to each Container already Up and Running.
// It listens to all ContainerEvents
func (d *dockerMonitor) Start() error {

	log.WithFields(log.Fields{
		"package": "monitor",
	}).Debug("Starting the docker monitor")

	// Starting the eventListener First.
	go d.eventListener()

	//Syncing all Existing containers depending on MonitorSetting
	if d.syncAtStart {
		err := d.syncContainers()

		if err != nil {
			log.WithFields(log.Fields{
				"package": "monitor",
				"error":   err.Error(),
			}).Error("Error Syncing existingContainers")
		}
	}

	// Processing the events received duringthe time of Sync.
	go d.eventProcessor()

	return nil
}

// Stop monitoring docker events.
func (d *dockerMonitor) Stop() error {

	log.WithFields(log.Fields{
		"package": "monitor",
	}).Debug("Stopping the docker monitor")

	d.stoplistener <- true
	d.stopprocessor <- true

	return nil
}

// eventProcessor processes docker events
func (d *dockerMonitor) eventProcessor() {

	for {
		select {
		case event := <-d.eventnotifications:
			if event.Action != "" {
				f, present := d.handlers[DockerEvent(event.Action)]
				if present {
					log.WithFields(log.Fields{
						"package": "monitor",
					}).Debug("Handling docker event")

					err := f(event)

					if err != nil {
						log.WithFields(log.Fields{
							"package": "monitor",
							"action":  event.Action,
							"error":   err.Error(),
						}).Error("Error while handling event")
					}
				} else {
					log.WithFields(log.Fields{
						"package": "monitor",
					}).Debug("Docker event not handled.")
				}
			}
		case <-d.stopprocessor:
			return
		}
	}
}

// eventListener listens to Docker events from the daemon and passes to
// to the processor through a buffered channel. This minimizes the chances
// that we will miss events because the processor is delayed
func (d *dockerMonitor) eventListener() {

	options := types.EventsOptions{}
	options.Filters = filters.NewArgs()
	options.Filters.Add("type", "container")

	messages, errs := d.dockerClient.Events(context.Background(), options)
	for {
		select {
		case message := <-messages:
			log.WithFields(log.Fields{
				"package": "monitor",
				"message": message,
			}).Debug("Got message from docker client")

			d.eventnotifications <- &message

		case err := <-errs:
			if err != nil && err != io.EOF {
				log.WithFields(log.Fields{
					"package": "monitor",
					"error":   err.Error(),
				}).Debug("Received docker event error")
			}
		case stop := <-d.stoplistener:
			if stop {
				return
			}
		}
	}
}

// syncContainers resyncs all the existing containers on the Host, using the
// same process as when a container is initially spawn up
func (d *dockerMonitor) syncContainers() error {

	log.WithFields(log.Fields{
		"package": "monitor",
	}).Debug("Syncing all existing containers")

	options := types.ContainerListOptions{All: true}
	containers, err := d.dockerClient.ContainerList(context.Background(), options)

	if err != nil {
		return fmt.Errorf("Error Getting ContainerList: %s", err)
	}

	if d.syncHandler != nil {
		for _, c := range containers {
			container, err := d.dockerClient.ContainerInspect(context.Background(), c.ID)

			if err != nil {
				log.WithFields(log.Fields{
					"package": "monitor",
					"error":   err.Error(),
				}).Error("Error Syncing existing Container")
				continue
			}

			contextID, _ := contextIDFromDockerID(container.ID)

			PURuntime, _ := d.extractMetadata(&container)

			var state monitor.State
			if container.State.Running {
				if !container.State.Paused {
					state = monitor.StateStarted
				} else {
					state = monitor.StatePaused
				}
			} else {
				state = monitor.StateStopped
			}
			if err := d.syncHandler.HandleSynchronization(contextID, state, PURuntime, monitor.SynchronizationTypeInitial); err != nil {
				log.WithFields(log.Fields{
					"package": "monitor",
					"error":   err.Error(),
				}).Error("Error Syncing existing Container")
			}
		}

		d.syncHandler.HandleSynchronizationComplete(monitor.SynchronizationTypeInitial)
	}

	for _, c := range containers {
		container, err := d.dockerClient.ContainerInspect(context.Background(), c.ID)

		if err != nil {
			log.WithFields(log.Fields{
				"package": "monitor",
				"error":   err.Error(),
			}).Error("Error Syncing existing Container")
			continue
		}

		if err := d.startDockerContainer(&container); err != nil {
			log.WithFields(log.Fields{
				"package": "monitor",
				"error":   err.Error(),
			}).Error("Error Syncing existing Container")
		}
	}

	return nil
}

func (d *dockerMonitor) startDockerContainer(dockerInfo *types.ContainerJSON) error {

	timeout := time.Second * 0

	if !dockerInfo.State.Running {

		return nil
	}

	contextID, err := contextIDFromDockerID(dockerInfo.ID)

	if err != nil {
		return fmt.Errorf("Couldn't generate ContextID: %s", err)
	}

	runtimeInfo, err := d.extractMetadata(dockerInfo)
	if err != nil {
		return fmt.Errorf("Error getting some of the Docker primitives: %s", err)
	}

	if err := d.puHandler.SetPURuntime(contextID, runtimeInfo); err != nil {
		return err
	}

	errorChan := d.puHandler.HandlePUEvent(contextID, monitor.EventStart)

	if err := <-errorChan; err != nil {
		if err := d.dockerClient.ContainerStop(context.Background(), dockerInfo.ID, &timeout); err != nil {
			log.WithFields(log.Fields{
				"package": "monitor",
				"error":   err.Error(),
			}).Warn("Failed to stop bad container")
		}
		return fmt.Errorf("Policy cound't be set - container was killed")
	}

	return nil
}

func (d *dockerMonitor) stopDockerContainer(dockerID string) error {

	contextID, err := contextIDFromDockerID(dockerID)

	if err != nil {
		return fmt.Errorf("Couldn't generate ContextID: %s", err)
	}

	errChan := d.puHandler.HandlePUEvent(contextID, monitor.EventStop)
	return <-errChan
}

// ExtractMetadata generates the RuntimeInfo based on Docker primitive
func (d *dockerMonitor) extractMetadata(dockerInfo *types.ContainerJSON) (*policy.PURuntime, error) {

	if dockerInfo == nil {
		return nil, fmt.Errorf("DockerInfo is empty")
	}

	if d.metadataExtractor != nil {
		return d.metadataExtractor(dockerInfo)
	}

	return defaultDockerMetadataExtractor(dockerInfo)
}

// handleCreateEvent generates a create event type.
func (d *dockerMonitor) handleCreateEvent(event *events.Message) error {
	dockerID := event.ID

	contextID, err := contextIDFromDockerID(dockerID)
	if err != nil {
		return fmt.Errorf("Error Generating ContextID: %s", err)
	}

	// Send the event upstream
	errChan := d.puHandler.HandlePUEvent(contextID, monitor.EventCreate)

	return <-errChan
}

// handleStartEvent will notify the agent immediately about the event in order
//to start the implementation of the functions. The agent must query
//the policy engine for details on what to do with this container.
func (d *dockerMonitor) handleStartEvent(event *events.Message) error {

	timeout := time.Second * 0
	dockerID := event.ID
	contextID, err := contextIDFromDockerID(dockerID)

	if err != nil {
		return fmt.Errorf("Error Generating ContextID: %s", err)
	}

	info, err := d.dockerClient.ContainerInspect(context.Background(), dockerID)

	if err != nil {
		//If we see errors, we will kill the container for security reasons.
		if err := d.dockerClient.ContainerStop(context.Background(), dockerID, &timeout); err != nil {
			log.WithFields(log.Fields{
				"package": "monitor",
				"error":   err.Error(),
			}).Warn("Failed to stop illegal container")
		}

		d.collector.CollectContainerEvent(&collector.ContainerRecord{
			ContextID: contextID,
			IPAddress: "N/A",
			Tags:      nil,
			Event:     collector.ContainerFailed,
		})

		return fmt.Errorf("Cannot read container information. Killing container. ")
	}

	if info.HostConfig.NetworkMode == "host" {
		log.WithFields(log.Fields{
			"package":                  "monitor",
			"Host namespace container": contextID,
		}).Warn("Ignore host namespace container - do nothing ")
		return nil
	}

	return d.startDockerContainer(&info)
}

//handleDie event is called when a container dies. It generates a "Stop" event.
func (d *dockerMonitor) handleDieEvent(event *events.Message) error {

	return d.stopDockerContainer(event.ID)
}

// handleDestroyEvent handles destroy events from Docker. It generated a "Destroy event"
func (d *dockerMonitor) handleDestroyEvent(event *events.Message) error {

	dockerID := event.ID
	contextID, err := contextIDFromDockerID(dockerID)

	if err != nil {
		return fmt.Errorf("Error Generating ContextID: %s", err)
	}

	// Send the event upstream
	errChan := d.puHandler.HandlePUEvent(contextID, monitor.EventDestroy)
	return <-errChan
}

// handlePauseEvent generates a create event type.
func (d *dockerMonitor) handlePauseEvent(event *events.Message) error {
	dockerID := event.ID
	contextID, err := contextIDFromDockerID(dockerID)
	if err != nil {
		return fmt.Errorf("Error Generating ContextID: %s", err)
	}

	errChan := d.puHandler.HandlePUEvent(contextID, monitor.EventPause)
	return <-errChan
}

// handleCreateEvent generates a create event type.
func (d *dockerMonitor) handleUnpauseEvent(event *events.Message) error {
	dockerID := event.ID
	contextID, err := contextIDFromDockerID(dockerID)
	if err != nil {
		return fmt.Errorf("Error Generating ContextID: %s", err)
	}

	// Send the event upstream
	errChan := d.puHandler.HandlePUEvent(contextID, monitor.EventUnpause)
	return <-errChan
}
