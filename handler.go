package wasmbus

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/nats-io/nats.go"
)

// EventHandlers holds the event handler functions
type EventHandlers struct {
	onHandleComponentScaled       HandleComponentScaled
	onHandleComponentScaledFailed HandleComponentScaledFailed
	onHandleLinkdefSet            HandleLinkdefSet
	onHandleLinkdefDeleted        HandleLinkdefDeleted
	onHandleProviderStarted       HandleProviderStarted
	onHandleProviderStartFailed   HandleProviderStartFailed
	onHandleProviderStopped       HandleProviderStopped
	onHandleHealthCheckPassed     HandleHealthCheckPassed
	onHandleHealthCheckFailed     HandleHealthCheckFailed
	onHandleHealthCheckStatus     HandleHealthCheckStatus
	onHandleConfigSet             HandleConfigSet
	onHandleConfigDeleted         HandleConfigDeleted
	onHandleLabelsChanged         HandleLabelsChanged
	onHandleHostHeartbeat         HandleHostHeartbeat
	onHandleHostStarted           HandleHostStarted
	onHandleHostStopped           HandleHostStopped
}

// NewEventHandlers creates a *EventHandlers with noop handler functions
func NewEventHandlers() *EventHandlers {
	return &EventHandlers{
		onHandleComponentScaled:       func(BaseEvent, ComponentScaled) {},
		onHandleComponentScaledFailed: func(BaseEvent, ComponentScaledFailed) {},
		onHandleLinkdefSet:            func(BaseEvent, LinkdefSet) {},
		onHandleLinkdefDeleted:        func(BaseEvent, LinkdefDeleted) {},
		onHandleProviderStarted:       func(BaseEvent, ProviderStarted) {},
		onHandleProviderStartFailed:   func(BaseEvent, ProviderStartFailed) {},
		onHandleProviderStopped:       func(BaseEvent, ProviderStopped) {},
		onHandleHealthCheckPassed:     func(BaseEvent, HealthCheckPassed) {},
		onHandleHealthCheckFailed:     func(BaseEvent, HealthCheckFailed) {},
		onHandleHealthCheckStatus:     func(BaseEvent, HealthCheckStatus) {},
		onHandleConfigSet:             func(BaseEvent, ConfigSet) {},
		onHandleConfigDeleted:         func(BaseEvent, ConfigDeleted) {},
		onHandleLabelsChanged:         func(BaseEvent, LabelsChanged) {},
		onHandleHostHeartbeat:         func(BaseEvent, HostHeartbeat) {},
		onHandleHostStarted:           func(BaseEvent, HostStarted) {},
		onHandleHostStopped:           func(BaseEvent, HostStopped) {},
	}
}

type subjectHandler func(BaseEvent) error

func (w *Wasmbus) handleEvent(m *nats.Msg) {
	evt := BaseEvent{}
	err := json.Unmarshal(m.Data, &evt)
	if err != nil {
		slog.Error("failed to unmarshal incoming event", "error", err)
		return
	}

	eventType, found := strings.CutPrefix(evt.Type, "com.wasmcloud.lattice.")
	if !found {
		slog.Error("invalid prefix for incoming event", "event_type", evt.Type)
		return
	}

	handler := w.getEventHandler(eventType)
	if handler == nil {
		slog.Warn("handler not found for incoming event", "event_type", evt.Type)
		return
	}

	err = handler(evt)
	if err != nil {
		slog.Error("handler function failed", "error", err)
	}
}

func (w *Wasmbus) getEventHandler(sub string) subjectHandler {
	switch sub {
	case "component_scaled":
		return w.handleComponentScaled
	case "component_scale_failed":
		return w.handleComponentScaledFailed
	case "linkdef_set":
		return w.handleLinkdefSet
	case "linkdef_deleted":
		return w.handleLinkdefDeleted
	case "provider_started":
		return w.handleProviderStarted
	case "provider_start_failed":
		return w.handleProviderStartFailed
	case "provider_stopped":
		return w.handleProviderStopped
	case "health_check_passed":
		return w.handleHealthCheckPassed
	case "health_check_failed":
		return w.handleHealthCheckFailed
	case "health_check_status":
		return w.handleHealthCheckStatus
	case "config_set":
		return w.handleConfigSet
	case "config_deleted":
		return w.handleConfigDeleted
	case "labels_changed":
		return w.handleLabelsChanged
	case "host_heartbeat":
		return w.handleHostHeartbeat
	case "host_started":
		return w.handleHostStarted
	case "host_stopped":
		return w.handleHostStopped
	default:
		return nil
	}
}

func (w *Wasmbus) handleComponentScaled(evt BaseEvent) error {
	data := ComponentScaled{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleComponentScaled(evt, data)
	return nil
}

func (w *Wasmbus) handleComponentScaledFailed(evt BaseEvent) error {
	data := ComponentScaledFailed{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleComponentScaledFailed(evt, data)
	return nil
}

func (w *Wasmbus) handleLinkdefSet(evt BaseEvent) error {
	data := LinkdefSet{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleLinkdefSet(evt, data)
	return nil
}

func (w *Wasmbus) handleLinkdefDeleted(evt BaseEvent) error {
	data := LinkdefDeleted{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleLinkdefDeleted(evt, data)
	return nil
}

func (w *Wasmbus) handleProviderStarted(evt BaseEvent) error {
	data := ProviderStarted{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleProviderStarted(evt, data)
	return nil
}

func (w *Wasmbus) handleProviderStartFailed(evt BaseEvent) error {
	data := ProviderStartFailed{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleProviderStartFailed(evt, data)
	return nil
}

func (w *Wasmbus) handleProviderStopped(evt BaseEvent) error {
	data := ProviderStopped{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleProviderStopped(evt, data)
	return nil
}

func (w *Wasmbus) handleHealthCheckPassed(evt BaseEvent) error {
	data := HealthCheckPassed{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleHealthCheckPassed(evt, data)
	return nil
}

func (w *Wasmbus) handleHealthCheckFailed(evt BaseEvent) error {
	data := HealthCheckFailed{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleHealthCheckFailed(evt, data)
	return nil
}

func (w *Wasmbus) handleHealthCheckStatus(evt BaseEvent) error {
	data := HealthCheckStatus{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleHealthCheckStatus(evt, data)
	return nil
}

func (w *Wasmbus) handleConfigSet(evt BaseEvent) error {
	data := ConfigSet{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleConfigSet(evt, data)
	return nil
}

func (w *Wasmbus) handleConfigDeleted(evt BaseEvent) error {
	data := ConfigDeleted{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleConfigDeleted(evt, data)
	return nil
}

func (w *Wasmbus) handleLabelsChanged(evt BaseEvent) error {
	data := LabelsChanged{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleLabelsChanged(evt, data)
	return nil
}

func (w *Wasmbus) handleHostHeartbeat(evt BaseEvent) error {
	data := HostHeartbeat{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleHostHeartbeat(evt, data)
	return nil
}

func (w *Wasmbus) handleHostStarted(evt BaseEvent) error {
	data := HostStarted{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleHostStarted(evt, data)
	return nil
}

func (w *Wasmbus) handleHostStopped(evt BaseEvent) error {
	data := HostStopped{}
	err := json.Unmarshal(evt.Data, &data)
	if err != nil {
		return err
	}
	w.handlers.onHandleHostStopped(evt, data)
	return nil
}
