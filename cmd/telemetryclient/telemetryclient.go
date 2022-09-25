package telemetryclient

import (
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type TelemetryClient struct {
	client appinsights.TelemetryClient
}

func (t *TelemetryClient) TrackEvent(eventname string, properties map[string]string) { 
	event := appinsights.NewEventTelemetry(eventname)
	
	for key, value := range properties {
		event.Properties[key] = value
	}
	t.client.Track(event)
}

func NewTelemetryClient() *TelemetryClient {
	// InstrumentationKey=c764f176-19a5-4949-825d-9f30db2f14e8;IngestionEndpoint=https://germanywestcentral-1.in.applicationinsights.azure.com
	client := appinsights.NewTelemetryClient("c764f176-19a5-4949-825d-9f30db2f14e8")
	return &TelemetryClient{client: client}
}