package diamonds

import (
	"reflect"
	"strconv"
)

type Parameters struct {
	// Unexported fields
	shape      string
	minCarat   float64
	maxCarat   float64
	minColor   int
	maxColor   int
	minPrice   float64
	maxPrice   float64
	minCut     int
	maxCut     int
	minClarity int
	maxClarity int
	minDepth   float64
	maxDepth   float64
	minWidth   float64
	maxWidth   float64
	gia        int
	ags        int
	egl        int
	oth        int
	currency   string
	sortCol    string
	sortDir    string
	rowStart   int
}

// NewParameters returns a new parameters instance
// with default values
func NewParameters() Parameters {
	return Parameters{
		shape:      "none",
		minCarat:   0.5,
		maxCarat:   2.5,
		minColor:   1,
		maxColor:   9,
		minPrice:   100.0,
		maxPrice:   5000,
		minCut:     5,
		maxCut:     1,
		minClarity: 1,
		maxClarity: 10,
		minDepth:   0.01,
		maxDepth:   90.00,
		minWidth:   0.01,
		maxWidth:   90.00,
		gia:        1,
		ags:        1,
		egl:        0,
		oth:        0,
		currency:   "USD",
		sortCol:    "price",
		sortDir:    "ASC",
		rowStart:   0,
	}
}

func (p *Parameters) SetRow(n int) {
	p.rowStart = n
}

func (p Parameters) ToMap() map[string]string {
	out := make(map[string]string)
	str := reflect.ValueOf(p)

	for n := 0; n < str.NumField(); n++ {
		name := str.Type().Field(n).Name
		value := str.Field(n)

		switch value.Type().String() {
		case "string":
			out[name] = value.String()
		case "int":
			out[name] = strconv.FormatInt(value.Int(), 10)
		case "float64":
			out[name] = strconv.FormatFloat(value.Float(), 'f', -1, 64)
		}
	}

	return out
}
