package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerProjetoLooper(t *testing.T) {
	req, err := http.NewRequest("GET", "/projeto-looper", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerProjetoLooper)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response RespostaLooper
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	expectedNome := "Projeto Looper"
	if response.Nome != expectedNome {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Nome, expectedNome)
	}

	if response.Horario == "" {
		t.Errorf("expected non-empty horario field")
	}
}
