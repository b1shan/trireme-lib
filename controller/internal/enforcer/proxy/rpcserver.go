package enforcerproxy

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.aporeto.io/trireme-lib/v11/collector"
	"go.aporeto.io/trireme-lib/v11/common"
	"go.aporeto.io/trireme-lib/v11/controller/internal/enforcer/utils/rpcwrapper"
)

// ProxyRPCServer This struct is a receiver for Statsserver and maintains a handle to the RPC ProxyRPCServer.
type ProxyRPCServer struct {
	collector   collector.EventCollector
	rpchdl      rpcwrapper.RPCServer
	secret      string
	tokenIssuer common.ServiceTokenIssuer
}

// PostStats is the function called from the remoteenforcer when it has new flow events to publish.
func (r *ProxyRPCServer) PostStats(req rpcwrapper.Request, resp *rpcwrapper.Response) error {

	if !r.rpchdl.ProcessMessage(&req, r.secret) {
		return errors.New("message sender cannot be verified")
	}

	payload := req.Payload.(rpcwrapper.StatsPayload)

	for _, record := range payload.Flows {
		r.collector.CollectFlowEvent(record)
	}
	payload.Flows = nil

	for _, record := range payload.Users {
		r.collector.CollectUserEvent(record)
	}
	payload.Users = nil

	return nil
}

// PostPacketEvent is called from the remote to post multiple records from the remoteenforcer
func (r *ProxyRPCServer) PostPacketEvent(req rpcwrapper.Request, resp *rpcwrapper.Response) error {
	if !r.rpchdl.ProcessMessage(&req, r.secret) {
		return errors.New("message sender cannot be verified")
	}

	payload := req.Payload.(rpcwrapper.DebugPacketPayload)
	for _, record := range payload.PacketRecords {
		r.collector.CollectPacketEvent(record)
	}
	payload.PacketRecords = nil

	return nil
}

// PostCounterEvent is called from the remote to post multiple counter records from the remoteenforcer
func (r *ProxyRPCServer) PostCounterEvent(req rpcwrapper.Request, resp *rpcwrapper.Response) error {
	if !r.rpchdl.ProcessMessage(&req, r.secret) {
		return errors.New("message sender cannot be verified")
	}

	payload := req.Payload.(rpcwrapper.CounterReportPayload)
	for _, record := range payload.CounterReports {
		r.collector.CollectCounterEvent(record)
	}

	payload.CounterReports = nil
	return nil
}

// DNSReports is called from the remote to post dns requests
func (r *ProxyRPCServer) DNSReports(req rpcwrapper.Request, resp *rpcwrapper.Response) error {
	if !r.rpchdl.ProcessMessage(&req, r.secret) {
		return errors.New("message sender cannot be verified")
	}

	payload := req.Payload.(rpcwrapper.DNSReportPayload)
	r.collector.CollectDNSRequests(payload.Report)

	payload.Report = nil
	return nil
}

// PostPingEvent is called from the remote to post ping events
func (r *ProxyRPCServer) PostPingEvent(req rpcwrapper.Request, resp *rpcwrapper.Response) error {
	if !r.rpchdl.ProcessMessage(&req, r.secret) {
		return errors.New("message sender cannot be verified")
	}

	payload := req.Payload.(rpcwrapper.PingReportPayload)
	r.collector.CollectPingEvent(payload.Report)

	payload.Report = nil
	return nil
}

// RetrieveToken propagates the master request to the token retriever and returns a token.
func (r *ProxyRPCServer) RetrieveToken(req rpcwrapper.Request, resp *rpcwrapper.Response) error {

	if !r.rpchdl.ProcessMessage(&req, r.secret) {
		return errors.New("message sender cannot be verified")
	}

	payload, ok := req.Payload.(rpcwrapper.TokenRequestPayload)
	if !ok {
		return errors.New("invalid request payload for token request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	token, err := r.tokenIssuer.Issue(ctx, payload.ContextID, payload.ServiceTokenType, payload.Audience, payload.Validity)
	if err != nil {
		resp.Status = "error"
		return fmt.Errorf("control plane failed to issue token: %s", err)
	}

	resp.Status = "ok"
	resp.Payload = &rpcwrapper.TokenResponsePayload{
		Token: token,
	}

	return nil
}
