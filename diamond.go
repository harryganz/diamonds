package diamonds

// The Diamond struct contains the properties
// for a single diamond.
type Diamond struct {
	// Shape of a diamond, e.g. Round, Emerald, Pear, etc.
	Shape string
	// Carat weight of a diamond as a floating point number
	Carat float64
	// Cut quality of a diamond from Ideal to Good
	Cut string
	// Color grade of a diamond from D (Colorless) to L (Yellow)
	Color string
	// Clarity of a diamond from FL (Flawless) to I2 (Many visible intrusions)
	Clarity string
	// Width of of the table of the diamond as a percent of its max width
	Width float64
	// Depth of a diamond as as percent of its max width
	Depth float64
	// Certification agency, e.g. GIA, AGS, EGL
	Certification string
	// Dimensions of a diamond in millimeters (length * width * height)
	Dimensions string
	// Price of a diamond given in the currency used in the search, e.g. U.S. Dollars
	Price float64
}
