package diamonds

import (
  "testing"
)

func TestShapeString (t *testing.T) {
  shapeStrings := []string{"none", "Round", "Princess", "Cushion", "Radiant", "Asscher",
    "Emerald", "Pear", "Heart", "Oval", "Marquise", "Baguette", "Trillion"}

  for i, v := range shapeStrings {
     if s := Shape(i); s.String() != v {
       t.Error("String for shape did not match expected value. Expected: " + v +
     ". Got: " + s.String())
     }
  }
}
