package main

import (
	"fmt"
	"math"

	"github.com/mitchfen/learning_go/pkg/outputHelpers"
)

func main() {

	investmentAmount := 22000
	returnRateAsPercent := 8.0
	yearsHeld := 40
	var futureValue = float64(investmentAmount) * math.Pow(1+returnRateAsPercent/100, float64(yearsHeld))

	// Sprintf and Printf are just the same as Print and Sprint but with formatting functions included
	var s = fmt.Sprintf("Initial invesment: %v\nFuture value after %v years: %0.2f", investmentAmount, yearsHeld, futureValue)
	outputHelpers.PrettyPrint(s)
}
