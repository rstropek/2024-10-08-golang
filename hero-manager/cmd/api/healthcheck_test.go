package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheck(t *testing.T) {
	app := &application{
		config: config{env: "abc"},
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.healthcheckHandler(rr, req)

	rs := rr.Result()
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	healthResult := make(map[string]string)
	json.NewDecoder(rs.Body).Decode(&healthResult)

	if healthResult["environment"] != "abc" {
		t.Errorf("want environment available; got %s", healthResult["environment"])
	}
}

func TestHealtcheck_EndToEnd(t *testing.T) {
	app := &application{
		config: config{env: "abc"},
	}

	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/healthcheck")
	if err != nil {
		t.Fatal(err)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	healthResult := make(map[string]string)
	json.NewDecoder(rs.Body).Decode(&healthResult)

	if healthResult["environment"] != "abc" {
		t.Errorf("want environment available; got %s", healthResult["environment"])
	}
}
