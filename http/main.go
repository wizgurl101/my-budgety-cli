package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type SetBudgetRequest struct {
	UserId string  `json:"userId"`
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Amount float64 `json:"budget_amount"`
}

type Response struct {
	Message string `json:"message"`
}

func SetYearBudgetAmount(year int, start_month int, amount float64) {
	url := "http://localhost:5000/monthlyBudget"
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
		return
	}

	userId := os.Getenv("USER_ID")
	if userId == "" {
		fmt.Println("The USER_ID environment variable is not set.")
		return
	}

	for i := start_month; i <= 12; i++ {
		requestBody := SetBudgetRequest{
			UserId: userId,
			Year:   year,
			Month:  i,
			Amount: amount,
		}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v\n", err)
			continue
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request for month %d: %v\n", i, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 201 {
			fmt.Printf("Request successful for month %d\n", i)
		} else {
			fmt.Printf("Request failed for month %d with status code: %d\n", i, resp.StatusCode)
		}
	}
}
