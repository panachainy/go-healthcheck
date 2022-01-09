package dto

type Summary struct {
	// Number of websites that have been checked
	TotalWebsites int `json:"total_websites"`
	// Number of websites that could reach
	Success int `json:"success"`
	// Number of websites that could not reach
	Failure int `json:"failure"`
	// Total time that used to check all websites (unix nano)
	TotalTime int64 `json:"total_time"`
}
