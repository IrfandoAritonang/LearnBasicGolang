package main

import (
	"fmt"
	"math"
)

func main() {
		var invesmentAmount = 1000
		var expectedReturnRate = 5.5 //float64
		var years = 10

	//ubah invesmentAmount ke float64
		var futureValue = float64(invesmentAmount)  * math.Pow (1 + expectedReturnRate / 100, float64(years))
		fmt.Println(futureValue)
} 
