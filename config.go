package hummingbird

// Config has HTTP server configurations
type Config struct {
	Port string
}

var ServerConfig Config

func init() {
	ServerConfig = setDefaultConfig()
}

func setDefaultConfig() Config {
	return Config{
		Port: "8080",
	}
}
