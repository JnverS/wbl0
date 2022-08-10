package apiserver

type Config struct {
	Addr              string  `toml:"addr"`
	DatabaseURL       string  `toml:"database_url"`
	NatsClusterId     string  `toml:"nats_cluster_id"`
	NatsClientId      string  `toml:"nats_client_id"`
	NatsPubliserId    string  `toml:"nats_publisher_id"`
	Durable           string  `toml:"durable"`
	Unsubscribe       bool  `toml:"unsubscribe"`
}

func NewConfig() *Config {
	return &Config{
		Addr: "localhost:8080",
	}
}
