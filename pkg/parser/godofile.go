package parser

type GodoFile struct {
	Commands map[string]Command `yaml:"commands"`
}

type Command struct {
	Run         Run     `yaml:"run"`
	Where       *string `yaml:"where"`
	Description *string `yaml:"description"`
}
