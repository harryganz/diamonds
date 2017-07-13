package diamonds

import (
	"gopkg.in/yaml.v2"
)

// The Parameters struct handles parameters sent to the diamond search engine.
// These parameters are used in constructing the query string for searching
// diamonds
type Parameters struct {
	shape string
}

// GetShape returns the shape field for the parameter instance
func (p *Parameters) GetShape() string {
	return p.shape
}

// NewParameters returns a parameter instance with the passed-in values.
// shape: A Shape, the shape of the diamond (e.g. Round, Radiant, etc.)
func NewParameters(shape Shape) Parameters {
	return Parameters{
		shape: shape.String(),
	}
}

// NewParametersFromYaml returns a parameter instance from a YAML byte array.
// The expected fields are the same as for NewParameters
// Any missing or invalid fields will be replaced by default values
func NewParametersFromYaml(config []byte) Parameters {
	params := struct {
		Shape string `yaml:"shape"`
	}{}

	err := yaml.Unmarshal(config, &params)
	if err != nil {
		panic("Could not parse YAML: " + err.Error())
	}

	return NewParameters(shapeFromString(params.Shape))
}

// Converts string to valid shape instance
// returns default shape if invalid or missing
func shapeFromString(s string) Shape {
	for i := 0; i < int(NumShapes); i++ {
		if Shape(i).String() == s {
			return Shape(i)
		}
	}

	return Shape(0)
}
