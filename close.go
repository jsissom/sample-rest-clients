package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8080/kfs-dev/coa/accounting_periods/close"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"description":"Document: The Final Frontier","universityFiscalYear": 2016,"universityFiscalPeriodCode": "03"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("authorization", "NSA_this_is_for_you")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
