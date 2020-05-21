package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func main2() {
	fmt.Println("Downloading data file...")
	url := "https://raw.githubusercontent.com/lyndadotcom/LPO_weatherdata/master/Environmental_Data_Deep_Moor_2012.txt"
	content := downloadFileData(url)
	fmt.Println("Downloaded file")
	r := csv.NewReader(content)
	defer content.Close()

	r.Comma = '\t'
	r.TrimLeadingSpace = true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record, len(record))
		for a, d := range record {
			fmt.Println(a, d)
		}
	}
}

func downloadFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func downloadFileData(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.Body
}

func main() {
	fmt.Println("Downloading data file...")
	url := "https://raw.githubusercontent.com/lyndadotcom/LPO_weatherdata/master/Environmental_Data_Deep_Moor_2012.txt"
	records := downloadCSVData(url)
	dataOnly := records[1:]
	airTemps := converToFloat64(dataOnly, 1)
	barometricPressures := converToFloat64(dataOnly, 2)
	windSpeeds := converToFloat64(dataOnly, 7)
	fmt.Printf("Number of Records: %v\n", len(dataOnly))
	fmt.Printf("Air_Temp         -> Mean: %0.2f Median: %0.2f\n", mean(airTemps), median(airTemps))
	fmt.Printf("Barometric_Press -> Mean: %0.2f Median: %0.2f\n", mean(barometricPressures), median(barometricPressures))
	fmt.Printf("Wind_Speed       -> Mean: %0.2f Median: %0.2f\n", mean(windSpeeds), median(windSpeeds))
}

func downloadCSVData(url string) [][]string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Read contents of the file using csv
	r := csv.NewReader(resp.Body)
	r.Comma = '\t'
	r.TrimLeadingSpace = true
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}

func converToFloat64(records [][]string, pos int) []float64 {
	converted := make([]float64, len(records))
	for i, record := range records {
		converted[i], _ = strconv.ParseFloat(record[pos], 64)
	}
	return converted
}

func mean(values []float64) float64 {
	var sum float64
	for _, data := range values {
		sum += data
	}
	return sum / float64(len(values))
}

func median(values []float64) float64 {
	sort.Float64s(values)
	length := len(values)
	mid := length / 2
	if length%2 == 0 {
		low := values[mid-1]
		high := values[mid]
		return (low + high) / 2.0
	} else {
		return values[mid]
	}
}
