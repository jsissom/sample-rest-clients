package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Closed         bool
	DocumentNumber string
}

func main() {
	universityFiscalYear := os.Args[1]
	universityFiscalPeriodCode := os.Args[2]

	url := "http://localhost:8080/kfs-dev/coa/accounting_periods/close"

	jsonStr := "{ \"description\": \"Closing Time\",\"universityFiscalYear\": " +
		universityFiscalYear + ", \"universityFiscalPeriodCode\": \"" + universityFiscalPeriodCode + "\" }"

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(jsonStr))
	req.Header.Set("authorization", "2e6e3b939593ff732b9482e80efe405c")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	dec := json.NewDecoder(strings.NewReader(string(body)))

	var m Message
	dec.Decode(&m)

	if m.Closed {
		fmt.Printf("The accounting period was closed.  Edoc %s was created.\n", m.DocumentNumber)
	} else {
		fmt.Println("The accounting period was already closed. No edoc was created.")
	}
}
