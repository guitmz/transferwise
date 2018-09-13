[![Build Status](https://travis-ci.org/guitmz/transferwise.svg?branch=master)](https://travis-ci.org/guitmz/transferwise) [![Go Report Card](https://goreportcard.com/badge/github.com/guitmz/transferwise)](https://goreportcard.com/report/github.com/guitmz/transferwise)

# transferwise
A transferwise command line tool to get quotes (with colors!)

# Installation
Available via:

- `go get -u github.com/guitmz/transferwise`
- Homebrew: `brew install guitmz/transferwise/transferwise`
- GitHub precompiled releases

# Usage
```
$ transferwise -help
Usage of transferwise:
  -amount float
    	Amount you want to transfer from the source currency (default 1000)
  -from string
    	Source currency acronym
  -json
    	Outputs in JSON format.
  -to string
    	Destination currency acronym
  -version
    	Prints current version.
```

Regular output:
```
$ transferwise -amount 1000.00 -from EUR -to BRL

TransferWise Quota (EUR -> BRL | Rate: 4.87295):

Amount to send (EUR): 1.000,00
Amount to receive (BRL): 4.815,55
Fee (EUR): 11,78
Quote requested at: Thu Sep 13 15:39:30 2018
Delivery estimate: Fri Sep 14 14:42:30 2018
```

JSON output:
```
transferwise -amount 1000.00 -from EUR -to BRL -json
{
  "source": "EUR",
  "target": "BRL",
  "sourceAmount": 1000,
  "targetAmount": 4815.55,
  "rate": 4.87295,
  "createdTime": "2018-09-13T15:40:19.102Z",
  "deliveryEstimate": "2018-09-14T14:43:19.03Z",
  "fee": 11.78
}
```
