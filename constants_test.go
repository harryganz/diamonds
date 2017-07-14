package diamonds_test

import (
	. "github.com/harryganz/diamonds"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String method for Shape constants", func() {
	table := []struct {
		input   Shape
		output  string
		context string
		outcome string
	}{
		{NoneShape, "none", "when shape is NoneShape", "returns none"},
		{RoundShape, "Round", "when shape is RoundShape", "returns Round"},
		{PrincessShape, "Princess", "when shape is PrincessShape", "returns Princess"},
		{CushionShape, "Cushion", "when shape is CushionShape", "returns Cushion"},
		{RadiantShape, "Radiant", "when shape is RadiantShape", "returns Radiant"},
		{AsscherShape, "Asscher", "when shape is AsscherShape", "returns Asscher"},
		{EmeraldShape, "Emerald", "when shape is EmeraldShape", "returns Emerald"},
		{PearShape, "Pear", "when shape is PearShape", "returns Pear"},
		{HeartShape, "Heart", "when shape is HeartShape", "returns Heart"},
		{OvalShape, "Oval", "when shape is OvalShape", "returns Oval"},
		{MarquiseShape, "Marquise", "when shape is MarquiseShape", "returns Marquise"},
		{BaguetteShape, "Baguette", "when shape is BaguetteShape", "returns Baguette"},
		{TrillionShape, "Trillion", "when shape is TrillionShape", "returns Trillion"},
	}

	for _, entry := range table {
		Context(entry.context, func() {
			It(entry.outcome, func() {
				Expect(entry.input.String()).To(Equal(entry.output))
			})
		})
	}
})

var _ = Describe("String method for Color constants", func() {
	table := []struct {
		input   Color
		output  string
		context string
		outcome string
	}{
		{DColor, "D", "when color is DColor", "returns D"},
		{EColor, "E", "when color is EColor", "returns E"},
		{FColor, "F", "when color is FColor", "returns F"},
		{GColor, "G", "when color is GColor", "returns G"},
		{HColor, "H", "when color is GColor", "returns H"},
		{IColor, "I", "when color is IColor", "returns I"},
		{JColor, "J", "when color is JColor", "returns J"},
		{KColor, "K", "when color is KColor", "returns K"},
		{LColor, "L", "when color is LColor", "returns L"},
	}

	for _, entry := range table {
		Context(entry.context, func () {
			It(entry.outcome, func () {
				Expect(entry.input.String()).To(Equal(entry.output))
			})
		})
	}
})
