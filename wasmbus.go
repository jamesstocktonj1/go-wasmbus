package wasmbus

import (
	"github.com/nats-io/nats.go"
)

type Wasmbus struct {
	cfg      *Config
	conn     *nats.Conn
	handlers EventHandlers
	sub      *nats.Subscription
}

// New creates a new Wasmbus instance based on the supplied Config, or returns an error
func New(cfg *Config) (*Wasmbus, error) {
	w := Wasmbus{
		handlers: *NewEventHandlers(),
	}

	cfg, err := defaultConfig(cfg)
	if err != nil {
		return nil, err
	}
	w.cfg = cfg

	conn, err := nats.Connect(w.cfg.NatsUrl)
	if err != nil {
		return nil, err
	}
	w.conn = conn

	return &w, nil
}

// Start subscribes to the "wasmbus.evt."
func (w *Wasmbus) Start() error {
	sub, err := w.conn.Subscribe("wasmbus.evt.*.>", w.handleEvent)
	if err != nil {
		return err
	}
	w.sub = sub
	return nil
}

// Close clears the subscription and closes the Nats connection
func (w *Wasmbus) Close() error {
	err := w.sub.Drain()
	if err != nil {
		return err
	}
	w.conn.Close()
	return nil
}

// OnComponentScaled subscribes a HandleComponentScaled function to the Wasmbus
func (w *Wasmbus) OnComponentScaled(handler HandleComponentScaled) {
	w.handlers.onHandleComponentScaled = handler
}

// OnComponentScaledFailed subscribes a HandleComponentScaledFailed function to the Wasmbus
func (w *Wasmbus) OnComponentScaledFailed(handler HandleComponentScaledFailed) {
	w.handlers.onHandleComponentScaledFailed = handler
}

// OnLinkdefSet subscribes a HandleLinkdefSet function to the Wasmbus
func (w *Wasmbus) OnLinkdefSet(handler HandleLinkdefSet) {
	w.handlers.onHandleLinkdefSet = handler
}

// OnLinkdefDeleted subscribes a HandleLinkdefDeleted function to the Wasmbus
func (w *Wasmbus) OnLinkdefDeleted(handler HandleLinkdefDeleted) {
	w.handlers.onHandleLinkdefDeleted = handler
}

// OnProviderStarted subscribes a HandleProviderStarted function to the Wasmbus
func (w *Wasmbus) OnProviderStarted(handler HandleProviderStarted) {
	w.handlers.onHandleProviderStarted = handler
}

// OnProviderStartFailed subscribes a HandleProviderStartFailed function to the Wasmbus
func (w *Wasmbus) OnProviderStartFailed(handler HandleProviderStartFailed) {
	w.handlers.onHandleProviderStartFailed = handler
}

// OnProviderStopped subscribes a HandleProviderStopped function to the Wasmbus
func (w *Wasmbus) OnProviderStopped(handler HandleProviderStopped) {
	w.handlers.onHandleProviderStopped = handler
}

// OnHealthCheckPassed subscribes a HandleHealthCheckPassed function to the Wasmbus
func (w *Wasmbus) OnHealthCheckPassed(handler HandleHealthCheckPassed) {
	w.handlers.onHandleHealthCheckPassed = handler
}

// OnHealthCheckFailed subscribes a HandleHealthCheckFailed function to the Wasmbus
func (w *Wasmbus) OnHealthCheckFailed(handler HandleHealthCheckFailed) {
	w.handlers.onHandleHealthCheckFailed = handler
}

// OnHealthCheckStatus subscribes a HandleHealthCheckStatus function to the Wasmbus
func (w *Wasmbus) OnHealthCheckStatus(handler HandleHealthCheckStatus) {
	w.handlers.onHandleHealthCheckStatus = handler
}

// OnConfigSet subscribes a HandleConfigSet function to the Wasmbus
func (w *Wasmbus) OnConfigSet(handler HandleConfigSet) {
	w.handlers.onHandleConfigSet = handler
}

// OnConfigDeleted subscribes a HandleConfigDeleted function to the Wasmbus
func (w *Wasmbus) OnConfigDeleted(handler HandleConfigDeleted) {
	w.handlers.onHandleConfigDeleted = handler
}

// OnLabelsChanged subscribes a HandleLabelsChanged function to the Wasmbus
func (w *Wasmbus) OnLabelsChanged(handler HandleLabelsChanged) {
	w.handlers.onHandleLabelsChanged = handler
}

// OnHostHeartbeat subscribes a HandleHostHeartbeat function to the Wasmbus
func (w *Wasmbus) OnHostHeartbeat(handler HandleHostHeartbeat) {
	w.handlers.onHandleHostHeartbeat = handler
}

// OnHostStarted subscribes a HandleHostStarted function to the Wasmbus
func (w *Wasmbus) OnHostStarted(handler HandleHostStarted) {
	w.handlers.onHandleHostStarted = handler
}

// OnHostStopped subscribes a HandleHostStopped function to the Wasmbus
func (w *Wasmbus) OnHostStopped(handler HandleHostStopped) {
	w.handlers.onHandleHostStopped = handler
}
