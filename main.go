package main

import (
	"fmt"
	"math"

	"github.com/mitchfen/learning_go/pkg/outputHelpers"
)

func main() {

	var investmentAmount = 22000
	var returnRateAsPercent = 8.0
	var yearsHeld = 40
	var futureValue = float64(investmentAmount) * math.Pow(1+returnRateAsPercent/100, float64(yearsHeld))

	var s = fmt.Sprintf("Initial invesment: %d\nFuture value after %d years: %d", investmentAmount, yearsHeld, int(futureValue))
	outputHelpers.PrettyPrint(s)
}
