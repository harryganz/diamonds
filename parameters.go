package main



const (
  defaultShape = "Round"

)

// Params holds the query parameters for
// making requests to the diamonds search engine
type Params struct {
  // shape of the diamond, e.g. round, princess, cushion
  shape string
  // Minimum size in carats
  minCarat float64
  // Maximum size in carats
  maxCarat float64
  // Minimum color value (1-9) where 1 is equivalent to D and 9 to L
  // to L
  minColor int
  // Maximum color value (1-9) where 1 is equivalent to D and 9 to L
  maxColor int
  // Minimum price in currency units (e.g. USD, GBP)
  minPrice float64
  // Maximum price in currency units (e.g. USD, GBP)
  maxPrice float64
  // Minimum cut grade (3-5), where
  // 3 is ideal, 4 is very good, and 5 is good
  // minCut is the minimum grade (higher number)
  minCut int
  // Maximum cut grade (3-5), where
  // 3 is ideal, 4 is very good, and 5 is good
  // maxCut is the highest grade (lower number)
  maxCut int
  // Minimum clarity (1-10), where 1 is FL (Flawless) and
  // 10 is I2 (Large and numerous inclusions)
  minClarity int
  // Maximum clarity (1-10), where 1 is FL (Flawless) and
  // 10 is I2 (Large and numerous inclusions)
  maxClarity int
  // Minimum depth in mm
  minDepth float64
  // Maximum depth in mm
  maxDepth float64
  // Minimum width (table) in mm
  minWidth float64
  // Maximum width (table) in mm
  maxWidth float64
  // Include GIA certification?
  gia bool
  // Include AGS certification?
  ags bool
  // Include EGL certification?
  egl bool
  // Include other certifications?
  oth bool
  // Which currency to use
  currency string
  // Which column to sort by
  sortCol string
  // Which direction to sort in
  sortDir string
}

func NewParams() {
}
