package diamonds

import (
  "fmt"
	"gopkg.in/yaml.v2"
)

type Parameters struct {
	shape string
}

func (p Parameters) String() string {
  return fmt.Sprintf("{shape: %s}", p.shape)
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

// NewParametersFromBytes returns a parameter instance from a YAML byte array.
// The expected fields are the same as for NewParameters
// Any missing or invalid fields will be replaced by default values
func NewParametersFromBytes(in []byte) Parameters {
	params := struct {
		Shape string `yaml:"shape"`
	}{}

	err := yaml.Unmarshal(in, &params)
	if err != nil {
		panic("Could not parse YAML ")
	}

	return NewParameters(shapeFromString(params.Shape))
}

// Converts string to valid shape instance
// returns default shape if invalid or missing
func shapeFromString(s string) Shape {
	for i := 0; i < int(NUM_SHAPES); i++ {
		if Shape(i).String() == s {
			return Shape(i)
		}
	}

	return Shape(0)
}
