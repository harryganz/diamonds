package diamonds

import (
	"testing"
	"reflect"
)

func TestNewParametersFromBytes(t *testing.T) {
	in := []byte(`shape: Round`)
	exp := Parameters{
		shape: "Round",
	}

	if obs := NewParametersFromBytes(in); !reflect.DeepEqual(exp, obs) {
		t.Errorf("Did not return expected parameters. Expected: %s, Got: %s", exp.String(),
	obs.String())
	}
}
