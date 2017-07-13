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
	NumShapes                  // 13
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
		panic("Shape was not recognized")
	}
}
