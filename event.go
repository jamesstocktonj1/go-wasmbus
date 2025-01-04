package wasmbus

// The component_scaled event is emitted when a component is scaled up or down.
type ComponentScaled struct {
	Annotations  map[string]string `json:"annotations"`
	Claims       Claims            `json:"claims"`
	ComponentID  string            `json:"component_id"`
	HostID       string            `json:"host_id"`
	ImageRef     string            `json:"image_ref"`
	MaxInstances int               `json:"max_instances"`
	PublicKey    string            `json:"public_key"`
}

// The component_scale_failed event is emitted when a component fails to scale up or down.
type ComponentScaledFailed struct {
	Annotations  map[string]string `json:"annotations"`
	ComponentID  string            `json:"component_id"`
	Error        string            `json:"error"`
	HostID       string            `json:"host_id"`
	ImageRef     string            `json:"image_ref"`
	MaxInstances int               `json:"max_instances"`
}

// The linkdef_set event is emitted when a link definition is set.
type LinkdefSet struct {
	Interfaces   []string `json:"interfaces"`
	Name         string   `json:"name"`
	SourceConfig []string `json:"source_config"`
	SourceID     string   `json:"source_id"`
	TargetConfig []string `json:"target_config"`
	TargetID     string   `json:"target"`
	WitNamespace string   `json:"wit_namespace"`
	WitPackage   string   `json:"wit_package"`
}

// The linkdef_deleted event is emitted when a link definition is deleted.
type LinkdefDeleted struct {
	Interfaces   []string `json:"interfaces"`
	Name         string   `json:"name"`
	SourceID     string   `json:"source_id"`
	TargetID     string   `json:"target"`
	WitNamespace string   `json:"wit_namespace"`
	WitPackage   string   `json:"wit_package"`
}

// The provider_started event is emitted when a provider is started.
type ProviderStarted struct {
	Annotations map[string]string `json:"annotations"`
	Claims      Claims            `json:"claims"`
	HostID      string            `json:"host_id"`
	ImageRef    string            `json:"image_ref"`
	InstanceID  string            `json:"instance_id"`
	LinkName    string            `json:"link_name"`
	ProviderID  string            `json:"provider_id"`
	PublicKey   string            `json:"public_key"`
}

// The provider_start_failed event is emitted when a provider fails to start.
type ProviderStartFailed struct {
	Error      string `json:"error"`
	LinkName   string `json:"link_name"`
	ImageRef   string `json:"provider_ref"`
	ProviderID string `json:"provider_id"`
}

// The provider_stopped event is emitted when a provider is stopped.
type ProviderStopped struct {
	Annotations map[string]string `json:"annotations"`
	HostID      string            `json:"host_id"`
	InstanceID  string            `json:"instance_id"`
	LinkName    string            `json:"link_name"`
	ProviderID  string            `json:"provider_id"`
	PublicKey   string            `json:"public_key"`
	Reason      string            `json:"stop"`
}

// The health_check_passed event is emitted when a health check has passed.
type HealthCheckPassed struct {
	HostID     string `json:"host_id"`
	ProviderID string `json:"provider_id"`
}

// The health_check_failed event is emitted when a health check has failed.
type HealthCheckFailed struct {
	HostID     string `json:"host_id"`
	ProviderID string `json:"provider_id"`
}

// The health_check_status event is emitted when the status of a health check has been reported.
type HealthCheckStatus struct {
	HostID     string `json:"host_id"`
	ProviderID string `json:"provider_id"`
}

// The config_set event is emitted when a configuration is set.
type ConfigSet struct {
	ConfigName string `json:"config_name"`
}

// The config_deleted event is emitted when a configuration is deleted.
type ConfigDeleted struct {
	ConfigName string `json:"config_name"`
}

// The labels_changed event is emitted when the labels on a host have changed.
type LabelsChanged struct {
	HostID string            `json:"host_id"`
	Labels map[string]string `json:"labels"`
}

// The host_heartbeat event is emitted when a host has sent a heartbeat.
type HostHeartbeat struct {
	Components    []Component       `json:"components"`
	FriendlyName  string            `json:"friendly_name"`
	HostID        string            `json:"host_id"`
	Labels        map[string]string `json:"labels"`
	Providers     []Provider        `json:"providers"`
	UptimeHuman   string            `json:"uptime_human"`
	UptimeSeconds int               `json:"uptime_seconds"`
	Version       string            `json:"version"`
}

// The host_started event is emitted when a host has started.
type HostStarted struct {
	FriendlyName  string            `json:"friendly_name"`
	Labels        map[string]string `json:"labels"`
	UptimeSeconds int               `json:"uptime_seconds"`
	Version       string            `json:"version"`
}

// The host_stopped event is emitted when a host has stopped.
type HostStopped struct {
	Labels map[string]string `json:"labels"`
}
