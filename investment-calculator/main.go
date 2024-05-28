package main

import (
    "fmt"
    "math"
)

const inflationRate = 2.5

func main() {
    var investmentAmount float64
    var expectedReturnRate float64
    var years float64

    fmt.Print("Investment Amount: ")
    fmt.Scan(&investmentAmount)

    fmt.Print("Expected Return Rate: ")
    fmt.Scan(&expectedReturnRate)

    fmt.Print("Years: ")
    fmt.Scan(&years)

    futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

    fmt.Printf("Future Value: %.2f\n", futureValue)
    fmt.Printf("Future Value (Adjusted for Inflation): %.2f", futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
    fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
    rfv = fv / math.Pow(1+inflationRate/100, years)
    return fv, rfv
}
