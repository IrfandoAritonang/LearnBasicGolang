package main

import (
	"fmt"
	"math"
)

func main() {
	//2// inflasi
	const inflationRate = 2.5
	//1// invesment calculator
		var invesmentAmount float64 // -->  bisa juga dibuat tanpa nilai tapi haru kasih tipe data nya|||| var invesmentAmount float64
		expectedReturnRate := 5.5 //float64
		var years float64

	

	//3// menerima inputan, dan menimpa nilai awal jika ada
	//tambahkan simbol pointer & ke variabel untuk menangkap nilai variabel

	//nilai investasi
	fmt.Print("Input invesment amount here: ")
	fmt.Scan(&invesmentAmount)

 //nilai expectedReturnRate
	fmt.Print("Input expected Return Rate here: ")
	fmt.Scan(&expectedReturnRate)
	
  //nilai tahun
		fmt.Print("Input years here: ")
		fmt.Scan(&years)

		//1//ubah invesmentAmount ke float64
	//1// rumus keuntungan investasi
	futureValue := invesmentAmount  * math.Pow (1 + expectedReturnRate / 100, years)
	//2//rumus inflasi
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

		fmt.Println(futureValue)
		fmt.Println(futureRealValue)
} 
