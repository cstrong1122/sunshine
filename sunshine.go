package main

import (
	"fmt"

	"github.com/cstrong1122/foobank"
	"github.com/cstrong1122/moonlight"
)

func main() {

	fooProvider := foobank.FooBankRateProvider{}
	fooParameters := &moonlight.RateOptionParameters{
		FICO:  750,
		CLTV:  80,
		LTV:   80,
		HCLTV: 90}

	fmt.Println("####### START RATE OPTION PARAMETERS #######")
	fmt.Println("FICO:", fooParameters.FICO)
	fmt.Println("LTV:", fooParameters.LTV)
	fmt.Println("CLTV:", fooParameters.CLTV)
	fmt.Println("HCLTV:", fooParameters.HCLTV)
	fmt.Print("####### END RATE OPTION PARAMETERS #######", "\n\n\n")

	fooOptions, err := fooProvider.GetRateOptions(fooParameters)
	if err == nil {
		for _, option := range fooOptions {
			fmt.Println("######## START RATE OPTION ########")
			fmt.Println("Rate:", option.InterestRate)
			fmt.Println("Origination Points:", option.OriginationPoints)
			fmt.Println("Discount Points:", option.DiscountPoints)
			fmt.Print("######## END RATE OPTION ########", "\n\n")
		}
	} else {
		fmt.Println("Error caught:", err)
	}

}
