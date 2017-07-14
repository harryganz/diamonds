package diamonds

// A Shape is the shape of a diamond
type Shape uint32

// Constants for the different shape names
// Last holds the total number of shapes
const (
	NoneShape     Shape = iota // 0
	RoundShape                 // 1
	PrincessShape              // 2
	CushionShape               // 3
	RadiantShape               // 4
	AsscherShape               // 5
	EmeraldShape               // 6
	PearShape                  // 7
	HeartShape                 // 8
	OvalShape                  // 9
	MarquiseShape              // 10
	BaguetteShape              // 11
	TrillionShape              // 12
)

// Diamond colors from D (colorless) to L (faint-yellow)
type Color uint32

// Constants for diamond colors
const (
	_ Color = iota // 0
	DColor// 1
	EColor // 2
	FColor // 3
	GColor // 4
	HColor // 5
	IColor // 6
	JColor // 7
	KColor // 8
	LColor // 9
)

// String returns the string representation of a Shape
func (s Shape) String() string {
	switch s {
	case NoneShape:
		return "none"
	case RoundShape:
		return "Round"
	case PrincessShape:
		return "Princess"
	case CushionShape:
		return "Cushion"
	case RadiantShape:
		return "Radiant"
	case AsscherShape:
		return "Asscher"
	case EmeraldShape:
		return "Emerald"
	case HeartShape:
		return "Heart"
	case PearShape:
		return "Pear"
	case OvalShape:
		return "Oval"
	case MarquiseShape:
		return "Marquise"
	case BaguetteShape:
		return "Baguette"
	case TrillionShape:
		return "Trillion"
	default:
		return ""
	}
}

func (c Color) String() string {
	switch c {
	case DColor:
		return "D"
	case EColor:
		return "E"
	case FColor:
		return "F"
	case GColor:
		return "G"
	case HColor:
		return "H"
	case IColor:
		return "I"
	case JColor:
		return "J"
	case KColor:
		return "K"
	case LColor:
		return "L"
	default:
		return ""
	}
}
