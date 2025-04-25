package parser


type GodoFile struct {
	Commands map[string]Command `yaml:"commands"`
}

type Command struct {
	Run         Run     `yaml:"run"`
	Description *string `yaml:"description"`
}
