package domain

type Config struct {
	ConfigVersion     int
	ConfigProjectName string
	ConfigEnv         string
	ConfigCreatedAt   string
	ConfigState       string
	ConfigData        map[string]interface{}
}
