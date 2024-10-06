package main

import (
	"fmt"
	"math"
)
const inflationRate  = 2.5

func main() {



 var investmentAmount  float64
 expectedReturnRate := 5.5
var years  float64

// fmt.Print("Investmento Amounto: ")
outputText(("Investmento Amounto: "))
fmt.Scan(&investmentAmount)

fmt.Print("Expected return: ")
fmt.Scan(&expectedReturnRate)

fmt.Print("years: ")
fmt.Scan(&years)

futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
//  futureValue := investmentAmount * math.Pow((1 + (expectedReturnRate) / 100), years) 
// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)


formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
formattedRFV := fmt.Sprintf("future value (Adjusted for inflation: %.1f\n", futureRealValue)

fmt.Print(formattedFV, formattedRFV)
}


func outputText(text string ){
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv =  math.Pow((1 + (expectedReturnRate) / 100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

return fv, rfv 

}