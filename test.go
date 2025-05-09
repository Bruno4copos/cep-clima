package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidCEP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/weather?cep=abc", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(weatherHandler)
	handler.ServeHTTP(rr, req)
	if rr.Code != 422 {
		t.Errorf("esperado 422, recebeu %d", rr.Code)
	}
}

func TestMissingCEP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/weather", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(weatherHandler)
	handler.ServeHTTP(rr, req)
	if rr.Code != 422 {
		t.Errorf("esperado 422, recebeu %d", rr.Code)
	}
}
