package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/leekchan/accounting"
	"github.com/logrusorgru/aurora"
)

const (
	appVersion = "1.0.0"
)

func getTempQuote(sourceAamount float64, fromCurrency, toCurrency string, jsonOutput bool) {
	params := map[string]string{
		"rateType":     "FIXED",
		"sourceAmount": strconv.FormatFloat(sourceAamount, 'f', -1, 64),
		"source":       fromCurrency,
		"target":       toCurrency,
	}
	body := transferwiseRequest("/v1/quotes", params)

	temporaryQuote := temporaryQuote{}
	json.Unmarshal(body, &temporaryQuote)

	if jsonOutput {
		prettyPrint(temporaryQuote)
	} else {
		estimateDelivery := temporaryQuote.DeliveryEstimate.Format("Mon Jan _2 15:04:05 2006")
		createdTime := temporaryQuote.CreatedTime.Format("Mon Jan _2 15:04:05 2006")
		ac := accounting.Accounting{Precision: 2, Thousand: ".", Decimal: ","}

		fmt.Printf(aurora.Sprintf("\nTransferWise Quota (%s -> %s | Rate: %.5f):\n\n", fromCurrency, toCurrency, aurora.Bold(aurora.Blue(temporaryQuote.Rate))))
		fmt.Printf(aurora.Sprintf("Amount to send (%s): %s\n", fromCurrency, aurora.Bold(aurora.Red(ac.FormatMoney(sourceAamount)))))
		fmt.Printf(aurora.Sprintf("Amount to receive (%s): %s\n", toCurrency, aurora.Bold(aurora.Green(ac.FormatMoney(temporaryQuote.TargetAmount)))))
		fmt.Printf(aurora.Sprintf("Fee (%s): %s\n", fromCurrency, aurora.Bold(aurora.Brown(ac.FormatMoney(temporaryQuote.Fee)))))
		fmt.Printf(aurora.Sprintf("Quote requested at: %s\n", aurora.Bold(createdTime)))
		fmt.Printf(aurora.Sprintf("Delivery estimate: %s\n\n", aurora.Bold(estimateDelivery)))
	}
}

func main() {
	sourceAmountPtr := flag.Float64("amount", 1000.00, "Amount you want to transfer from the source currency")
	fromCurrencyPtr := flag.String("from", "", "Source currency acronym")
	toCurrencyPtr := flag.String("to", "", "Destination currency acronym")
	jsonOutputPtr := flag.Bool("json", false, "Outputs in JSON format.")
	versionPtr := flag.Bool("version", false, "Prints current version.")
	flag.Parse()

	if *versionPtr {
		fmt.Println("transferwise command line tool version", appVersion)
		os.Exit(0)
	}

	if *fromCurrencyPtr == "" || *toCurrencyPtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	getTempQuote(*sourceAmountPtr, *fromCurrencyPtr, *toCurrencyPtr, *jsonOutputPtr)
}
