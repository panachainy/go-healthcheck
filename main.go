package main

import (
	"fmt"
	"go-healthcheck/healthz/services"
	"go-healthcheck/utils"
	"os"
)

func main() {
	utils.LoadConfigLog("logs.txt")
	utils.LoadConfiguration(".env")

	// Receive csvPath from argument.
	if len(os.Args) < 2 {
		panic("Require csv path at first argument\nUsage: go-healthcheck test.csv")
	}
	csvPath := os.Args[1]

	// Read csv file.
	healths, err := services.GetHealthFromFile(csvPath)
	if err != nil {
		panic(fmt.Sprintf("Error getHealthFromFile: %v", err))
	}

	fmt.Println()
	fmt.Println("Perform website checking...")
	summary := services.GetHealthSummary(healths)

	fmt.Println("Done!")
	fmt.Println()

	services.SubmitReport(summary)

	fmt.Printf("Checked webistes: %v\n", summary.TotalWebsites)
	fmt.Printf("Successful websites: %v\n", summary.Success)
	fmt.Printf("Failure websites: %v\n", summary.Failure)
	fmt.Printf("Total times to finished checking website:Total times to finished checking website: %v\n", summary.TotalTime)
}
