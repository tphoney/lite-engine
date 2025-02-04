package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config provides the system configuration.
type Config struct {
	Debug      bool   `envconfig:"DEBUG"`
	Trace      bool   `envconfig:"TRACE"`
	ServerName string `envconfig:"SERVER_NAME" default:"drone"`

	Server struct {
		Bind       string `envconfig:"HTTPS_BIND" default:":9079"`
		CertFile   string `envconfig:"SERVER_CERT_FILE" default:"/tmp/certs/server-cert.pem"` // Server certificate PEM file
		KeyFile    string `envconfig:"SERVER_KEY_FILE" default:"/tmp/certs/server-key.pem"`   // Server key PEM file
		CACertFile string `envconfig:"CLIENT_CERT_FILE" default:"/tmp/certs/ca-cert.pem"`     // CA certificate file
	}

	Client struct {
		Bind       string `envconfig:"HTTPS_BIND" default:":9079"`
		CertFile   string `envconfig:"CLIENT_CERT_FILE" default:"/tmp/certs/server-cert.pem"` // Server certificate PEM file
		KeyFile    string `envconfig:"CLIENT_KEY_FILE" default:"/tmp/certs/server-key.pem"`   // Server Key PEM file
		CaCertFile string `envconfig:"CA_CERT_FILE" default:"/tmp/certs/ca-cert.pem"`         // CA certificate file
	}
}

// Load loads the configuration from the environment.
func Load() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	return cfg, err
}
