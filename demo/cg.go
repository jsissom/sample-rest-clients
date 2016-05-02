package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Response      bool
	AccountNumber string
	FinCoaCode    string
	AccountName   string
}

func main() {
	chartOfAccountsCode := os.Args[1]
	accountNumber := os.Args[2]

	url := "http://localhost:8080/kfs-dev/coa/account_validation/is_cg_account/" + accountNumber + "/" + chartOfAccountsCode

	req, err := http.NewRequest("GET", url, nil)
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

	if resp.StatusCode == 200 {
		dec := json.NewDecoder(strings.NewReader(string(body)))

		var m Message
		dec.Decode(&m)

		if m.Response {
			fmt.Printf("The Account is a C&G Account\n")
		} else {
			fmt.Printf("The Account is not a C&G Account\n")
		}
		fmt.Printf("Account Name: %s\n", m.AccountName)
	} else {
		fmt.Printf("The account number is invalid\n")
	}
}
