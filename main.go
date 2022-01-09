package main

import (
	"fmt"
	"go-healthcheck/healthz/services"
)

func main() {
	// // TODO: Uncomment after development
	// // Receive csvPath from argument.
	// if len(os.Args) < 2 {
	// 	panic("Require csv path at first argument\nUsage: go-healthcheck test.csv")
	// }

	// csvPath := os.Args[1]

	// Mock instead real argument.
	csvPath := "test.csv"

	// Read csv file.
	healths, err := services.GetHealthFromFile(csvPath)
	if err != nil {
		panic(fmt.Sprintf("Error getHealthFromFile: %v", err))
	}

	fmt.Println("Perform website checking...")
	summary, err2 := services.GetHealthSummary(healths)
	if err2 != nil {
		panic(fmt.Sprintf("Error GetHealthSummary: %v", err2))
	}
	fmt.Println("Done!")

	// TODO: Uncomment after development
	// Submit report
	// services.SubmitReport(summary)

	fmt.Printf("Checked webistes: %v\n", summary.TotalWebsites)
	fmt.Printf("Successful websites: %v\n", summary.Success)
	fmt.Printf("Failure websites: %v\n", summary.Failure)
	fmt.Printf("Total times to finished checking website:Total times to finished checking website: %v\n", summary.TotalTime)
}
