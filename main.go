package main

import (
	"encoding/csv"
	"fmt"
	"go-healthcheck/dto"
	"go-healthcheck/externals"
	"os"
	"strings"
)

func main() {
	// // Receive csvPath from argument.
	// if len(os.Args) < 2 {
	// 	panic("Require csv path at first argument\nUsage: go-healthcheck test.csv")
	// }

	// csvPath := os.Args[1]

	// Mock instead real argument.
	csvPath := "test.csv"

	// Read csv file.
	healths, err := getHealthFromFile(csvPath)
	if err != nil {
		panic(fmt.Sprintf("Error getHealthFromFile: %v", err))
	}

	fmt.Println("Perform website checking...")
	summary, err2 := externals.GetHealthSummary(healths)
	if err2 != nil {
		panic(fmt.Sprintf("Error GetHealthSummary: %v", err2))
	}
	fmt.Println("Done!")

	// Submit report
	externals.SubmitReport(summary)

	fmt.Printf("Checked webistes: %v\n", summary.TotalWebsites)
	fmt.Printf("Successful websites: %v\n", summary.Success)
	fmt.Printf("Failure websites: %v\n", summary.Failure)
	fmt.Printf("Total times to finished checking website:Total times to finished checking website: %v\n", summary.TotalTime)
}

func getHealthFromFile(csvPath string) ([]dto.Health, error) {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	var healthData []dto.Health
	for i, line := range csvLines {
		// Check header
		if i == 0 {
			if strings.ToLower(line[0]) != "url" {
				return nil, fmt.Errorf("Invalid csv file")
			}
			continue
		}

		healthData = append(healthData, dto.Health{URL: line[0]})
	}

	return healthData, nil
}
