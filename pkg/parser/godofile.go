package parser

const (
	SHELL = "shell"
	RAW   = "raw"
	PATH  = "path"
)

type GodoFile struct {
	Commands map[string]Command `yaml:"commands"`
}

type Command struct {
	Run         *Run      `yaml:"run"`
	Where       *string   `yaml:"where"`
	Times       *int32    `yaml:"times"`
	Type        *string   `yaml:"type"`
	Description *string   `yaml:"description"`
	Variants    []Variant `yaml:"variants"`
}

type Variant struct {
	Run      Run     `yaml:"run"`
	Platform string  `yaml:"platform"`
	Type     *string `yaml:"type"`
}
