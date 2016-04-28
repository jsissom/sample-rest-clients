#!/bin/sh

curl -X POST --data "{ \"description\": \"My Document 2: Even More Document!\",\"universityFiscalYear\": 2009, \"universityFiscalPeriodCode\": \"02\" }" --header "Authorization: NSA_this_is_for_you" --header "Content-Type: application/json" --header "Accept: application/json" -v http://localhost:8080/kfs-dev/coa/accounting_periods/close

