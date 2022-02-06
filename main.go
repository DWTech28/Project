package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davidwarth/Project"
)

func main() {
	client := finnhub.NewClient("")
	symbol := strings.ToUpper(os.Args[1])
	quote, err := client.Quote(symbol)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("Symbol: %s\n", symbol)
	fmt.Printf("  Close: %f\n", quote.Previous)
	fmt.Printf("  Open:  %f\n", quote.Open)
	fmt.Printf("  High:  %f\n", quote.High)
	fmt.Printf("  Low:   %f\n", quote.Low)
	fmt.Printf("  Price: %f\n", quote.Current)
}