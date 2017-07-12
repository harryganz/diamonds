package diamonds

// A Shape is the shape of a diamond
type Shape uint32

const (
	None     Shape = iota // 0
	Round                 // 1
	Princess              // 2
	Cushion               // 3
	Radiant               // 4
	Asscher               // 5
	Emerald               // 6
	Pear                  // 7
	Heart                 // 8
	Oval                  // 9
	Marquise              // 10
	Baguette              // 11
	Trillion              // 12
)

// String returns the string representation of a Shape
func (s Shape) String() string {
	switch s {
	case None:
		return "none"
	case Round:
		return "Round"
	case Princess:
		return "Princess"
	case Cushion:
		return "Cushion"
	case Radiant:
		return "Radiant"
	case Asscher:
		return "Asscher"
	case Emerald:
		return "Emerald"
	case Heart:
		return "Heart"
	case Pear:
		return "Pear"
	case Oval:
		return "Oval"
	case Marquise:
		return "Marquise"
	case Baguette:
		return "Baguette"
	case Trillion:
		return "Trillion"
	default:
		panic("Shape was not recognized")
	}
}
