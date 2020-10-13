package controllers

import (
	// "context"
	// "encoding/json"
	// "net/http"
	// "fmt"
	// "strconv"
	// "log"

	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/MaximillianoNico/COVID-19-API/pkg/e"
	app "github.com/MaximillianoNico/COVID-19-API/pkg/formatter"
)

type question struct {
	Id string `json:"id"`
	// interests []string
	Value string `json:"value"`
}

func GetQuestionSymptoms(c *gin.Context) {
	appG := app.Gin{C: c}
	// country := c.Param("country")

	// countryLangId := country

	QuestionV1 := []question{
		{
			Id:    "1",
			Value: "Do you have Fever?",
		},
		{
			Id:    "2",
			Value: "Are you coughing and /or have shortness of breath?",
		},
		{
			Id:    "3",
			Value: "Are you Sneezing?",
		},
		{
			Id:    "4",
			Value: "Are you Aches and pains?",
		},
		{
			Id:    "5",
			Value: "Are you Runny or stuffy nose?",
		},
		{
			Id:    "6",
			Value: "Are you Diarrhea?",
		},
		{
			Id:    "7",
			Value: "Are you Diarrhea?",
		},
		{
			Id:    "8",
			Value: "Are you feel Headaches?",
		},
		{
			Id:    "3",
			Value: "Have you been in close contact with someone confirmed to have Coronavirus (COVID-19)?",
		},
		{
			Id:    "4",
			Value: "Have you traveled outside of the country in the last 14 days?",
		},
	}

	appG.Response(http.StatusOK, e.SUCCESS, "Success", QuestionV1)
}
