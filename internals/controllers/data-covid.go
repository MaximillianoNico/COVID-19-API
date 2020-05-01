package controllers

import (
	"context"
	"encoding/json"
	"encoding/csv"
	"net/http"
	"fmt"
	// "strconv"
	// "log"

	"time"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

func readData (url string) ([][] string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'

	data, err := reader.ReadAll()
	
	if err != nil {
		return nil, err
	}

	return data, nil
}

func filterDataCovid19 (dateStart string, dateEndtime string, country string) {
	evStart := time.Date(dateStart)
	evEnd := time.Date(dateEndtime)

	diffDay := evEnd.Sub(evStart).Hours() / 24

	urlIFilter := `https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/01-22-2020.csv`
	res, err := readData()
}

func GetDataCovid (w http.ResponseWriter, r *http.Request) {
	data := [] struct {
        Name string
        Age  int
    } {
        { "Richard Grayson", 24 },
        { "Jason Todd", 23 },
        { "Tim Drake", 22 },
        { "Damian Wayne", 21 },
	}
	
	jsonInBytes, err := json.Marshal(data)

	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
	
}