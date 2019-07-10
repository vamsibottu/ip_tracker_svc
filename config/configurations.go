package config

import (
	"github.com/spf13/viper"
)

// all configuration keys
const (
	ListeningPort    = "api.handlers.listeningport"
	IPStackURL       = "service.ipstack.url"
	IPStackAccessKey = "service.ipstack.accesskey"
)

func init() {
	viper.SetDefault(ListeningPort, ":8090")
	viper.SetDefault(IPStackURL, "http://api.ipstack.com/")
	// TODO(vbottu): need to add access key in artifacts-secrets when kublets come into the play
	viper.SetDefault(IPStackAccessKey, "8359ab14f0fdbc6ec03db1b12d122ebe")
}
