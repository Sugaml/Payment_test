package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPaymentSetting(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/payment/paymentsetting/1", nil)
	if err != nil {
		t.Fatalf("Could not created request :%v", err)
	}
	rec := httptest.NewRecorder()
	server.CreatePaymentSetting(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected satus ok, got %v", res.StatusCode)
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read responses :%v", err)
	}

}
