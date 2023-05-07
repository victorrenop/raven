package domain

// Config is the main data object of our domain
// Defines the fundamental representation of a config dataset
type Config struct {
	ConfigVersion     int
	ConfigProjectName string
	ConfigEnv         string
	ConfigCreatedAt   string
	ConfigState       string
	ConfigData        map[string]interface{}
}
