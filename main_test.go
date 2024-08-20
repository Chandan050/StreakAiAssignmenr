package main

import (
	"bytes"
	"encoding/json"
	"myapp/controllers"
	"myapp/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	router := Router()
	router.POST("/fair-pairs", controllers.FindPairAPi)
	gin.SetMode(gin.TestMode)

	
	testCases := []struct {
		Name string
		Request models.RequestData
		ExpectedResponse string
		Code int
	}{
		{"Test 01", models.RequestData{[]int{1,2,3,4,5},6},`{"solutions":[[0,4],[1,3]]}`,200 },
		{"Test 02", models.RequestData{[]int{1,2,3,4,4,5,5},6},`{"solutions":[[0,5],[0,6],[1,3],[1,4]]}`,200 },
		{"Test 03", models.RequestData{[]int{1,2,3},9},`{"solutions":[]}`,200},
		{"Test 04", models.RequestData{[]int{},10},`{"error":"Numbers Should be more than equal 2"}`,400},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Arrange
			inputdata, err := json.Marshal(tc.Request)
			if err != nil {
				t.Fatalf("Error while parsing to json %v",err)
			}

			//Request the Api
			req, err := http.NewRequest("POST", "/find-pairs", bytes.NewBuffer(inputdata))
			if err != nil {
				t.Fatalf("Error while parsing to json %v",err)
			}
			
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			//Compares the given status code and response code
			assert.Equal(t,tc.Code, rr.Code)

			//compares the expected the response and response from the API
			assert.Equal(t, tc.ExpectedResponse, rr.Body.String())
			

		})

	}
}
