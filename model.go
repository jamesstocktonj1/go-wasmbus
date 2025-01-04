package wasmbus

import "encoding/json"

// BaseEvent contains the base event for
type BaseEvent struct {
	SpecVersion     string          `json:"specversion"`
	ID              string          `json:"id"`
	Type            string          `json:"type"`
	Source          string          `json:"source"`
	DataContentType string          `json:"datacontenttype"`
	Time            string          `json:"time"`
	Data            json.RawMessage `json:"data"`
}

// Claims contains the claims data for a ComponentScaled or ProviderStarted event
type Claims struct {
	ExpiresHuman   string   `json:"expires_human"`
	Issuer         string   `json:"issuer"`
	Name           string   `json:"name"`
	NotBeforeHuman string   `json:"not_before_human"`
	Revision       int      `json:"revision"`
	Tags           []string `json:"tags"`
	Version        string   `json:"version"`
}

// Component contains the component data for a HostHeartbeat event
type Component struct {
	Anotations   map[string]string `json:"annotations"`
	ID           string            `json:"id"`
	ImageRef     string            `json:"image_ref"`
	MaxInstances int               `json:"max_instances"`
	Name         string            `json:"name"`
	Revision     int               `json:"revision"`
}

// Provider contains the provider data for a HostHeartbeat event
type Provider struct {
	Annotations map[string]string `json:"annotations"`
	ID          string            `json:"id"`
	ImageRef    string            `json:"image_ref"`
	Name        string            `json:"name"`
	Revision    int               `json:"revision"`
}
