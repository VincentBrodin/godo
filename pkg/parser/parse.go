package parser

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Run []string

func Parse(data []byte) (*GodoFile, error) {
	godoFile := GodoFile{}
	if err := yaml.Unmarshal(data, &godoFile); err != nil {
		return nil, err
	}
	return &godoFile, nil
}

func (r *Run) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.ScalarNode {
		*r = []string{value.Value}
		return nil
	}
	if value.Kind == yaml.SequenceNode {
		var result []string
		for _, node := range value.Content {
			result = append(result, node.Value)
		}
		*r = result
		return nil
	}
	return fmt.Errorf("invalid run format")
}

func (r *Run) Add(value string) {
	*r = append(*r, value)
}
