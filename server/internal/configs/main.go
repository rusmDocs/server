package configs

type AppConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Salt string `yaml:"salt"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
}

type ServerConfig struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
}
