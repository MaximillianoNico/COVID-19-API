package controllers

import (
	// "context"
	"encoding/csv"
	"strconv"
	"time"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MaximillianoNico/COVID-19-API/pkg/e"
	app "github.com/MaximillianoNico/COVID-19-API/pkg/formatter"
)

type SearchType struct {
	Datetime string `form:"datetime" json:"datetime"`
	City     string `form:"city" json:"city"`
}

type countryList struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	LangId string `json:"langId`
}

type Report struct {
	Province   string `json:"province"`
	Country    string `json:"country"`
	LastUpdate string `json:"last_update"`
	Confirmed  string `json:"confirmed"`
	Deaths     int    `json:"death"`
	Recovered  int    `json:"recovered"`
}

func GetDataCsvToJSON(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	// reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// @Summary GetAll
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /statistic/latest [get]
func GetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	// var reports Report
	var reportCovid []Report
	t := time.Now()
	formatted := fmt.Sprintf("%02d-%02d-%02d", t.Month(), t.Day(), t.Year())

	fmt.Println(formatted)

	urlIFilter := "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/" + formatted + ".csv"

	resp, err := GetDataCsvToJSON(urlIFilter)
	if err != nil {
		fmt.Println(err)
	}

	for idx, row := range resp {
		if idx == 0 {
			continue
		}
		totalDeath, _ := strconv.Atoi(row[8])
		totalRecovered, _ := strconv.Atoi(row[9])
		dataConvert := Report{
			Province:   row[2],
			Country:    row[3],
			LastUpdate: row[4],
			Confirmed:  row[7],
			Deaths:     totalDeath,
			Recovered:  totalRecovered,
		}
		reportCovid = append(reportCovid, dataConvert)
	}
	payload := reportCovid
	if len(reportCovid) == 0 {
		payload = []Report{}
	}
	appG.Response(http.StatusOK, e.SUCCESS, "Success", payload)
}

// @Summary Search Data COVID-19
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /statistic/search [get]
func SearchData(c *gin.Context) {
	appG := app.Gin{C: c}

	var searchType SearchType
	var reportCovid []Report

	if c.Bind(&searchType) == nil {
		var hasCity bool

		if hasCity = false; searchType.City != "" {
			hasCity = true
		}

		urlIFilter := "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/" + searchType.Datetime + ".csv"

		resp, err := GetDataCsvToJSON(urlIFilter)
		if err != nil {
			fmt.Println(err)
		}

		for idx, row := range resp {
			if idx == 0 {
				continue
			}

			totalDeath, _ := strconv.Atoi(row[8])
			totalRecovered, _ := strconv.Atoi(row[9])
			if hasCity == true {
				if row[3] == searchType.City {
					dataConvert := Report{
						Province:   row[2],
						Country:    row[3],
						LastUpdate: row[4],
						Confirmed:  row[7],
						Deaths:     totalDeath,
						Recovered:  totalRecovered,
					}
					reportCovid = append(reportCovid, dataConvert)
				}
			} else {
				dataConvert := Report{
					Province:   row[2],
					Country:    row[3],
					LastUpdate: row[4],
					Confirmed:  row[7],
					Deaths:     totalDeath,
					Recovered:  totalRecovered,
				}
				reportCovid = append(reportCovid, dataConvert)
			}
		}
		payload := reportCovid
		if len(reportCovid) == 0 {
			payload = []Report{}
		}

		appG.Response(http.StatusOK, e.SUCCESS, "Success", payload)
		return

	}

	appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "Error", gin.H{
		"message": "invalid parameter",
	})

	return

}

func GetStatistic(c *gin.Context) {
	appG := app.Gin{C: c}

	country := c.Param("city")

	appG.Response(http.StatusOK, e.SUCCESS, "Success", map[string]string{
		"country": country,
		"message": "get data covid-19 for country " + country,
	})
}

func GetCountryList(c *gin.Context) {
	appG := app.Gin{C: c}
	// country := c.Param("country")

	// countryLangId := country

	countries := []countryList{
		{
			Id:     "1",
			Name:   "Indonesia",
			LangId: "ID",
		},
		{
			Id:     "2",
			Name:   "English",
			LangId: "EN",
		},
	}

	fmt.Println(countries)

	appG.Response(http.StatusOK, e.SUCCESS, "Success", countries)
}
