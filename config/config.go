package config

// Config holds the global configuration settings
type Config struct {
	Verbose     *bool
	Profile     *string
	Region      *string
	NameFile    *string
	Eps         EpsConfiguration
	Eventbridge EventbridgeConfiguration
	Ecs         EcsConfiguration
	ErrorCode   bool
}

// EpsConfiguration config struct that holds config values for EPS
type EpsConfiguration struct {
	Serviceid string
}

// EventbridgeConfiguration config struct that holds config values for Eventbridge
type EventbridgeConfiguration struct {
	Nameprefix string
}

type EcsConfiguration struct {
	Cluster string
}
