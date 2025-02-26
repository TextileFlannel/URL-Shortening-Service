package config

type Config struct {
    ServerPort  string
    DatabaseURL string
}

func LoadConfig() Config {
    return Config{
        ServerPort:  ":8080",
        DatabaseURL: "test.db",
    }
}