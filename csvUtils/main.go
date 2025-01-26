package csvUtils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type CreditCardStatememt struct {
	date    string
	name    string
	credit  float64
	debit   float64
	balance float64
}

func GetDataFromCsvFile(filePath string) []CreditCardStatememt {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening csv file: %v\n", err)
		return []CreditCardStatememt{}
	}
	defer file.Close()

	csv_reader := csv.NewReader(file)
	records, err := csv_reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading csv file: %v\n", err)
		return []CreditCardStatememt{}
	}

	data := make([]CreditCardStatememt, 0)

	for _, record := range records {
		credit, _ := strconv.ParseFloat(record[2], 64)
		debit, _ := strconv.ParseFloat(record[3], 64)
		balance, _ := strconv.ParseFloat(record[4], 64)

		statement := CreditCardStatememt{
			date:    record[0],
			name:    record[1],
			credit:  credit,
			debit:   debit,
			balance: balance,
		}

		data = append(data, statement)
	}

	return data
}

func (c CreditCardStatememt) Equals(other CreditCardStatememt) bool {
	return c.date == other.date &&
		c.name == other.name &&
		c.credit == other.credit &&
		c.debit == other.debit &&
		c.balance == other.balance
}

func includes(slice []CreditCardStatememt, item CreditCardStatememt) bool {
	for _, element := range slice {
		if element.Equals(item) {
			return true
		}
	}
	return false
}

func removeDuplicates(slice1, slice2 []CreditCardStatememt) []CreditCardStatememt {
	result := []CreditCardStatememt{}

	for _, item := range slice2 {
		if !includes(slice1, item) {
			result = append(result, item)
		}
	}

	return result
}

func MergeCsvFiles() {
	files, err := os.ReadDir("./files")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	if len(files) != 2 {
		fmt.Println("There should be exactly 2 files in the directory")
		return
	}

	file1path := "./files/" + files[0].Name()
	file1Data := GetDataFromCsvFile(file1path)

	file2path := "./files/" + files[1].Name()
	file2Data := GetDataFromCsvFile(file2path)

	var mergedData []CreditCardStatememt
	if len(file2Data) > len(file1Data) {
		mergedData = removeDuplicates(file1Data, file2Data)
	} else {
		mergedData = removeDuplicates(file2Data, file1Data)
	}

	currentDate := time.Now()
	formattedDate := currentDate.Format("2006-01-02")
	outputFileName := "./output/" + formattedDate + "-merged.csv"

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for _, record := range mergedData {
		err := writer.Write([]string{
			record.date,
			record.name,
			strconv.FormatFloat(record.credit, 'f', 2, 64),
			strconv.FormatFloat(record.debit, 'f', 2, 64),
			strconv.FormatFloat(record.balance, 'f', 2, 64),
		})
		if err != nil {
			fmt.Printf("Error writing record to csv: %v\n", err)
			return
		}
	}

	fmt.Println("Merged data has been written to merged.csv")
}
