package gohttp

import (
	"testing"
)

func TestVersion(t *testing.T) {
	tearDown := setupServer(t)
	defer tearDown(t)

	request, err := NewRequest("http://localhost:1234/path")
	if err != nil {
		t.Fatal(err.Error())
	}
	request.version = "1.0"
	response, err := GET(request)
	if err != nil {
		t.Fatal(err.Error())
	}
	if response.version != "1.0" {
		t.Fatalf("HTTP VERSION IS WRONG")
	}

	headerValue, exists := response.GetHeader("TestHeader")
	if response.StatusCode != STATUS_OK || !exists || headerValue != "Hello" {
		t.FailNow()
	}

	bodyBuffer := make([]byte, 1024)
	bodyLength, _ := response.Body.Read(bodyBuffer)
	if string(bodyBuffer[:bodyLength]) != "Hello World!\n" {
		t.FailNow()
	}

	request, err = NewRequest("http://localhost:1234/")
	if err != nil {
		t.Fatal(err.Error())
	}
	request.version = "1.1"
	response, err = GET(request)
	if err != nil {
		t.Fatal(err.Error())
	}

	if response.version != "1.1" {
		t.Fatalf("HTTP VERSION IS WRONG")
	}

	headerValue, exists = response.GetHeader("TestHeader")
	if response.StatusCode != STATUS_OK || !exists || headerValue != "Hello" {
		t.FailNow()
	}

	bodyLength, _ = response.Body.Read(bodyBuffer)
	if string(bodyBuffer[:bodyLength]) != "Hello World!\n" {
		t.FailNow()
	}

	request, err = NewRequest("http://localhost:1234/resource")
	if err != nil {
		t.Fatal(err.Error())
	}
	request.version = "2.0"
	response, err = GET(request)
	if err != nil {
		t.Fatal(err.Error())
	}
	if response.version != "2.0" {
		t.Fatalf("HTTP VERSION IS WRONG")
	}

	if response.StatusCode != STATUS_NOT_IMPLEMENTED {
		t.FailNow()
	}

}
