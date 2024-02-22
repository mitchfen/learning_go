package main

import (
	"fmt"
	"math"

	"github.com/mitchfen/learning_go/pkg/outputHelpers"
)

func main() {

	var investmentAmount float64 = 22000
	var returnRateAsPercent float64 = 8.0
	var yearsHeld = 40
	var futureValue = investmentAmount * math.Pow(1+returnRateAsPercent/100, float64(yearsHeld))

	var s = fmt.Sprintf("Initial invesment: %f\nFuture value after %d years: %f", investmentAmount, yearsHeld, futureValue)
	outputHelpers.PrettyPrint(s)
}
