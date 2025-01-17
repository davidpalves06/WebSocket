package gohttp

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestLargeFiles(t *testing.T) {
	tearDown := setupServer(t)
	defer tearDown(t)

	file, err := os.Open("testdata/lusiadasTest.txt")
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	request, err := NewRequestWithBody("http://localhost:1234/large", body)
	if err != nil {
		t.Fatal(err.Error())
	}
	response, err := POST(request)
	if err != nil {
		t.Fatal(err.Error())
	}

	headerValue, exists := response.GetHeader("TestHeader")
	if response.StatusCode != STATUS_OK || !exists || headerValue != "Hello" {
		t.FailNow()
	}
	headerLength, exists := response.GetHeader("Content-Length")
	if !exists || headerLength != "362128" {
		t.Fatalf("Body length is incorrect")
	}

	bodyBuffer := make([]byte, 1024)
	var totalRead int
	for {
		read, err := response.Body.Read(bodyBuffer)
		if err != nil {
			break
		}
		totalRead += read
	}
	if totalRead != 362128 {
		t.Fatalf("Bad body")
	}
}

func TestSmallerContentLength(t *testing.T) {
	tearDown := setupServer(t)
	defer tearDown(t)

	file, err := os.Open("testdata/lusiadasTest.txt")
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	request, err := NewRequestWithBody("http://localhost:1234/large", body)
	if err != nil {
		t.Fatal(err.Error())
	}
	request.SetHeader("Content-Length", "10000")
	response, err := POST(request)
	if err != nil {
		t.Fatal(err.Error())
	}

	headerValue, exists := response.GetHeader("TestHeader")
	if response.StatusCode != STATUS_OK || !exists || headerValue != "Hello" {
		t.FailNow()
	}

	headerLength, exists := response.GetHeader("Content-Length")
	if !exists || headerLength != "362128" {
		t.Fatalf("Body length is incorrect")
	}

	bodyBuffer := make([]byte, 1024)
	var totalRead int
	for {
		read, err := response.Body.Read(bodyBuffer)
		if err != nil {
			break
		}
		totalRead += read
	}
	if totalRead != 362128 {
		t.Fatalf("Bad body")
	}
}

func TestBiggerContentLength(t *testing.T) {
	tearDown := setupServer(t)
	defer tearDown(t)

	file, err := os.Open("testdata/lusiadasTest.txt")
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	request, err := NewRequestWithBody("http://localhost:1234/large", body)
	if err != nil {
		t.Fatal(err.Error())
	}
	request.SetHeader("Content-Length", "1000000")
	response, err := POST(request)
	if err != nil {
		t.Fatal(err.Error())
	}

	headerValue, exists := response.GetHeader("TestHeader")
	if response.StatusCode != STATUS_OK || !exists || headerValue != "Hello" {
		t.FailNow()
	}

	headerLength, exists := response.GetHeader("Content-Length")
	if !exists || headerLength != "362128" {
		t.Fatalf("Body length is incorrect")
	}

	bodyBuffer := make([]byte, 1024)
	var totalRead int
	for {
		read, err := response.Body.Read(bodyBuffer)
		if err != nil {
			break
		}
		totalRead += read
	}
	if totalRead != 362128 {
		t.Fatalf("Bad body")
	}
}