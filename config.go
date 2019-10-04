package hummingbird

// Config has HTTP server configurations
type Config struct {
	Port string
}

var config Config

func init() {
	config = setDefaultConfig()
}

func setDefaultConfig() Config {
	return Config{
		Port: "8080",
	}
}
