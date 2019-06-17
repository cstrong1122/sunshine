package main

import (
	"fmt"

	"github.com/cstrong1122/foobank"
	"github.com/cstrong1122/moonlight"
)

func main() {

	provider := moonlight.IRateOptionProviderAdapter(foobank.FooBankRateProvider{})

	parameters := &moonlight.RateOptionParameters{
		FICO:  750,
		CLTV:  80,
		LTV:   80,
		HCLTV: 90}

	fmt.Println("####### START RATE OPTION PARAMETERS #######")
	fmt.Println("FICO:", parameters.FICO)
	fmt.Println("LTV:", parameters.LTV)
	fmt.Println("CLTV:", parameters.CLTV)
	fmt.Println("HCLTV:", parameters.HCLTV)
	fmt.Print("####### END RATE OPTION PARAMETERS #######", "\n\n\n")

	rateOptions, err := provider.GetRateOptions(parameters)

	if err == nil {
		for _, option := range rateOptions {
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
