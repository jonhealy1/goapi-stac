package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/jonhealy1/goapi-stac/pg-api/responses"
)

func TestSearchItems(t *testing.T) {
	jsonBody := []byte(`{"ids": ["S2B_1CCV_20181004_0_L2A"]}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 1
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "S2B_1CCV_20181004_0_L2A"
	if searchResponse.Features[0].Id != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestSearchCollections(t *testing.T) {
	jsonBody := []byte(`{"collections": ["sentinel-s2-l2a-cogs-test"]}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 50
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}
	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestSearchNoCollections(t *testing.T) {
	jsonBody := []byte(`{"collections": ["sentinel-s2-l2a-cogs-test-test"]}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 0
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}
}

func TestSearchGeometry(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"geometry": {
			"type": "Polygon",
        	"coordinates": [[
				[170.8515625, -74.14512718337613],
				[178.35937499999999, -74.14512718337613],
				[178.35937499999999, -70.15296965617042],
				[170.8515625, -70.15296965617042],
				[170.8515625, -74.14512718337613]
          	]]
      	}
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 50
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestSearchGeometryLimit(t *testing.T) {
	jsonBody := []byte(`{
		"limit": 1,
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"geometry": {
			"type": "Polygon",
        	"coordinates": [[
				[170.8515625, -74.14512718337613],
				[178.35937499999999, -74.14512718337613],
				[178.35937499999999, -70.15296965617042],
				[170.8515625, -70.15296965617042],
				[170.8515625, -74.14512718337613]
          	]]
      	}
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 1
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 1
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestSearchNoGeometry(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"geometry": {
			"type": "Polygon",
        	"coordinates": [[
				[70.8515625, -74.14512718337613],
				[78.35937499999999, -74.14512718337613],
				[78.35937499999999, -70.15296965617042],
				[70.8515625, -70.15296965617042],
				[70.8515625, -74.14512718337613]
          	]]
      	}
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 0
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}
}

func TestSearchPoint(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"geometry": {
			"type": "Point",
			"coordinates": [177.064544, -72.690647]
      	}
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 49
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestSearchNoPoint(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"geometry": {
			"type": "Point",
			"coordinates": [77.064544, -72.690647]
      	}
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 0
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}
}

func TestSearchLine(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"geometry": {
			"type": "LineString",
                "coordinates": [
                    [
                        179.85156249999997,
                        -70.554563528593656
                    ],
                    [
                        171.101642, 
                        -75.690647
                    ]
                ]
      	}
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 49
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestSearchNoLine(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"geometry": {
			"type": "LineString",
                "coordinates": [
                    [
                        79.85156249999997,
                        -70.554563528593656
                    ],
                    [
                        71.101642, 
                        -75.690647
                    ]
                ]
      	}
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 0
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}
}

func TestPostSearchBbox(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"bbox": [97.504892,-75.254738,179.321298,-65.431580]
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 50
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestPostSearchBboxSort(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"bbox": [97.504892,-75.254738,179.321298,-65.431580],
		"sortby": [{"field": "id", "direction": "DESC"}]
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 50
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}

	expectedId = "S2B_1CCV_20201222_0_L2A"
	if searchResponse.Features[0].Id != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestPostSearchBboxSortDatetime(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"bbox": [97.504892,-75.254738,179.321298,-65.431580],
		"sortby": [{"field": "properties.datetime", "direction": "DESC"}]
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 50
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}

	expectedId = "S2B_1CCV_20201222_0_L2A"
	if searchResponse.Features[0].Id != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestPostSearchBbox3d(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"bbox": [97.504892,-75.254738, 0, 179.321298,-65.431580, 0]
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 50
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestPostSearchBboxNoResults(t *testing.T) {
	jsonBody := []byte(`{
		"collections": ["sentinel-s2-l2a-cogs-test"],
		"bbox": [17.504892,-75.254738,19.321298,-65.431580]
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	app := Setup()
	req, _ := http.NewRequest("POST", "/search", bodyReader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 0
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}
}

func TestGetSearchBbox(t *testing.T) {
	app := Setup()
	req, _ := http.NewRequest("GET", "/search?bbox=97.504892,-75.254738,179.321298,-65.431580", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 50
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestGetSearchBboxLimit(t *testing.T) {
	app := Setup()
	req, _ := http.NewRequest("GET", "/search?bbox=97.504892,-75.254738,179.321298,-65.431580&limit=10", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 10
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 10
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestGetSearchBboxNoResults(t *testing.T) {
	app := Setup()
	req, _ := http.NewRequest("GET", "/search?bbox=17.504892,-75.254738,99.321298,-65.431580", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 0
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}
}

func TestGetSearchLine(t *testing.T) {
	app := Setup()

	url := "/search?geometry=" + url.QueryEscape(`{"type": "LineString","coordinates": [[179.85156249999997,-70.554563528593656],[171.101642,-75.690647]]}`)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 49
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}

	expectedId := "sentinel-s2-l2a-cogs-test"
	if searchResponse.Features[0].Collection != expectedId {
		t.Errorf("Expected id %s, but got %s", expectedId, searchResponse.Features[0].Id)
	}
}

func TestGetSearchLineNoResults(t *testing.T) {
	app := Setup()

	url := "/search?geometry=" + url.QueryEscape(`{"type": "LineString","coordinates": [[19.85156249999997,-70.554563528593656],[11.101642,-75.690647]]}`)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var searchResponse responses.SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	expectedType := "FeatureCollection"
	if searchResponse.Type != expectedType {
		t.Errorf("Expected type %s, but got %s", expectedType, searchResponse.Type)
	}

	expectedLimit := 100
	if searchResponse.Context.Limit != expectedLimit {
		t.Errorf("Expected limit %d, but got %d", expectedLimit, searchResponse.Context.Limit)
	}

	expectedReturned := 0
	if searchResponse.Context.Returned != expectedReturned {
		t.Errorf("Expected returned %d, but got %d", expectedReturned, searchResponse.Context.Returned)
	}
}
