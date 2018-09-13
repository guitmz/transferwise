[![Build Status](https://travis-ci.org/guitmz/n26.svg?branch=master)](https://travis-ci.org/guitmz/transferwise) [![Go Report Card](https://goreportcard.com/badge/github.com/guitmz/n26)](https://goreportcard.com/report/github.com/guitmz/transferwise) [![](https://images.microbadger.com/badges/image/guitmz/n26.svg)](https://microbadger.com/images/guitmz/transferwise "Get your own image badge on microbadger.com")

# transferwise
A transferwise command line tool to get quotes

# Installation
`$ go get -u github.com/guitmz/transferwise`
macOS: Available via Homebrew. Just run brew install guitmz/n26/n26
Linux: You can manually build this project or download a binary release.
You can also install with go get -u github.com/guitmz/n26/cmd/n26 (make sure you have your Go env setup correctly).

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