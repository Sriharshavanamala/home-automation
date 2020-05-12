// Code generated by jrpc. DO NOT EDIT.

package devicedef

import (
	"encoding/json"

	"github.com/jakewright/home-automation/libraries/go/firehose"
	"github.com/jakewright/home-automation/libraries/go/oops"
)

// Publish publishes the event to the Firehose
func (m *DeviceStateChangedEvent) Publish() error {
	if err := m.Validate(); err != nil {
		return err
	}

	return firehose.Publish("device-state-changed", m)
}

// DeviceStateChangedEventHandler implements the necessary functions to be a Firehose handler
type DeviceStateChangedEventHandler func(*DeviceStateChangedEvent) firehose.Result

// EventName returns the Firehose channel name
func (h DeviceStateChangedEventHandler) EventName() string {
	return "device-state-changed"
}

// HandleEvent handles the Firehose event
func (h DeviceStateChangedEventHandler) HandleEvent(e *firehose.Event) firehose.Result {
	var body DeviceStateChangedEvent
	if err := json.Unmarshal(e.Payload, &body); err != nil {
		return firehose.Discard(oops.WithMessage(err, "failed to unmarshal payload"))
	}
	return h(&body)
}
