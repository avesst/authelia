package schema

import (
	"net/url"
	"time"
)

// Telemetry represents the telemetry config.
type Telemetry struct {
	Metrics TelemetryMetrics `koanf:"metrics" yaml:"metrics,omitempty" toml:"metrics,omitempty" json:"metrics,omitempty" jsonschema:"title=Metrics" jsonschema_description:"The telemetry metrics server configuration."`
}

// TelemetryMetrics represents the telemetry metrics config.
type TelemetryMetrics struct {
	Enabled bool        `koanf:"enabled" yaml:"enabled" toml:"enabled" json:"enabled" jsonschema:"default=false,title=Enabled" jsonschema_description:"Enables the metrics server."`
	Address *AddressTCP `koanf:"address" yaml:"address,omitempty" toml:"address,omitempty" json:"address,omitempty" jsonschema:"default=tcp://:9959/,title=Address" jsonschema_description:"The address for the metrics server to listen on."`

	Buffers  ServerBuffers  `koanf:"buffers" yaml:"buffers,omitempty" toml:"buffers,omitempty" json:"buffers,omitempty" jsonschema:"title=Buffers" jsonschema_description:"The server buffers configuration for the metrics server."`
	Timeouts ServerTimeouts `koanf:"timeouts" yaml:"timeouts,omitempty" toml:"timeouts,omitempty" json:"timeouts,omitempty" jsonschema:"title=Timeouts" jsonschema_description:"The server timeouts configuration for the metrics server."`
}

// DefaultTelemetryConfig is the default telemetry configuration.
var DefaultTelemetryConfig = Telemetry{
	Metrics: TelemetryMetrics{
		Address: &AddressTCP{Address{true, false, -1, 9959, nil, &url.URL{Scheme: AddressSchemeTCP, Host: ":9959", Path: "/metrics"}}},
		Buffers: ServerBuffers{
			Read:  4096,
			Write: 4096,
		},
		Timeouts: ServerTimeouts{
			Read:  time.Second * 6,
			Write: time.Second * 6,
			Idle:  time.Second * 30,
		},
	},
}
