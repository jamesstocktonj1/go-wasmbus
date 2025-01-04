package wasmbus

// HandleComponentScaled is called when a component_scaled event is emitted.
type HandleComponentScaled func(BaseEvent, ComponentScaled)

// HandleComponentScaledFailed is called when a component_scale_failed event is emitted.
type HandleComponentScaledFailed func(BaseEvent, ComponentScaledFailed)

// HandleLinkdefSet is called when a linkdef_set event is emitted.
type HandleLinkdefSet func(BaseEvent, LinkdefSet)

// HandleLinkdefDeleted is called when a linkdef_deleted event is emitted.
type HandleLinkdefDeleted func(BaseEvent, LinkdefDeleted)

// HandleProviderStarted is called when a provider_started event is emitteded.
type HandleProviderStarted func(BaseEvent, ProviderStarted)

// HandleProviderStartFailed is called when a provider_start_failed event is emitted.
type HandleProviderStartFailed func(BaseEvent, ProviderStartFailed)

// HandleProviderStopped is called when a provider_stopped event is emitted.
type HandleProviderStopped func(BaseEvent, ProviderStopped)

// HandleHealthCheckPassed is called when a health_check_passed event is emitted.
type HandleHealthCheckPassed func(BaseEvent, HealthCheckPassed)

// HandleHealthCheckFailed is called when a health_check_failed event is emitted.
type HandleHealthCheckFailed func(BaseEvent, HealthCheckFailed)

// HandleHealthCheckStatus is called when a health_check_status event is emitted.
type HandleHealthCheckStatus func(BaseEvent, HealthCheckStatus)

// HandleConfigSet is called when a config_set event is emitted.
type HandleConfigSet func(BaseEvent, ConfigSet)

// HandleConfigDeleted is called when a config_deleted event is emitted.
type HandleConfigDeleted func(BaseEvent, ConfigDeleted)

// HandleLabelsChanged is called when a labels_changed event is emitted.
type HandleLabelsChanged func(BaseEvent, LabelsChanged)

// HandleHostHeartbeat is called when a host_heartbeat event is emitted.
type HandleHostHeartbeat func(BaseEvent, HostHeartbeat)

// HandleHostStarted is called when a host_started event is emitted.
type HandleHostStarted func(BaseEvent, HostStarted)

// HandleHostStopped is called when a host_stopped event is emitted.
type HandleHostStopped func(BaseEvent, HostStopped)
