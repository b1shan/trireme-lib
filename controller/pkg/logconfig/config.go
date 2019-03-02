package logconfig

import (
	"os"

	"go.aporeto.io/trireme-lib/controller/constants"
	"go.aporeto.io/trireme-lib/controller/internal/processmon"
	"go.aporeto.io/trireme-lib/controller/pkg/claimsheader"
)

// SetLogParameters sets up environment to be passed to the remote trireme instances.
func SetLogParameters(logToConsole, logWithID bool, logLevel string, logFormat string, compressedTags claimsheader.CompressionType) {

	h := processmon.GetProcessManagerHdl()
	if h == nil {
		panic("Unable to find process manager handle")
	}

	h.SetLogParameters(logToConsole, logWithID, logLevel, logFormat, compressedTags)
}

// GetLogParameters retrieves log parameters for Remote Enforcer.
func GetLogParameters() (logToConsole bool, logID string, logLevel string, logFormat string, compressedTagsVersion claimsheader.CompressionType) {

	logLevel = os.Getenv(constants.EnvLogLevel)
	if logLevel == "" {
		logLevel = "info"
	}
	logFormat = os.Getenv(constants.EnvLogFormat)
	if logLevel == "" {
		logFormat = "json"
	}

	if console := os.Getenv(constants.EnvLogToConsole); console == constants.EnvLogToConsoleEnable {
		logToConsole = true
	}

	logID = os.Getenv(constants.EnvLogID)

	compressedTagsVersion = claimsheader.CompressionTypeNone
	if console := os.Getenv(constants.EnvCompressedTags); console != string(claimsheader.CompressionTypeNone) {
		if console == string(claimsheader.CompressionTypeV1) {
			compressedTagsVersion = claimsheader.CompressionTypeV1
		} else if console == string(claimsheader.CompressionTypeV2) {
			compressedTagsVersion = claimsheader.CompressionTypeV2
		}
	}
	return
}