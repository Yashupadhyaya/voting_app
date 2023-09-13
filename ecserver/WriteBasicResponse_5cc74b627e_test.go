// Test generated by RoostGPT for test demo56 using AI Type Open AI and AI Model gpt-4

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BasicResponse struct {
	Status int
	Body   string
}

func writeBasicResponse(w http.ResponseWriter, resp *BasicResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	respJson, err := json.Marshal(resp)
	if err != nil {
		log.Println("error marshaling response to vote request. error: ", err)
	}
	w.Write(respJson)
}

func TestWriteBasicResponse_5cc74b627e(t *testing.T) {
	// Test case 1: Successful response
	resp1 := &BasicResponse{Status: http.StatusOK, Body: "Success"}
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	writeBasicResponse(rr, resp1)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Status":200,"Body":"Success"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Test case 2: Marshaling error
	resp2 := &BasicResponse{Status: http.StatusOK, Body: string([]byte{0x80, 0x81})}
	rr2 := httptest.NewRecorder()
	writeBasicResponse(rr2, resp2)

	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedErr := `{"Status":200,"Body":"\u0080\u0081"}`
	if rr2.Body.String() != expectedErr {
		t.Errorf("handler returned unexpected body: got %v want %v", rr2.Body.String(), expectedErr)
	}
}
