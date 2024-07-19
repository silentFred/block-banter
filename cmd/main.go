package main

import (
	"block-banter/api"
	"block-banter/background"
	"block-banter/database"
	"sync"
)

func main() {
	database.Init()

	var wg sync.WaitGroup
	wg.Add(2)

	go indexTransactions(&wg)
	go startWebServer(&wg)

	wg.Wait()
}

func indexTransactions(wg *sync.WaitGroup) {
	defer wg.Done()
	background.IndexTransactions()
}

func startWebServer(wg *sync.WaitGroup) {
	defer wg.Done()
	api.StartWebServer()
}

// TODO
// [x] connect to chain and pull events for particular erc 20 token
// [x] write token data to file
// [x] write api to fetch token data in format that d3 can use
// [x] write d3 code to display token data - serve the static page using go
// [x] use a database to store transactions instead of a file to learn how to use sql connectors in go? (perhaps look into graph databases)
// refactor the app into a layered architecture? (api, background, database, models, etc)
// pause and reflect on what we have learned so far
// - structs, gorm,

// if we have a database, we can pull balance for each address in the chart and display it in a heatmap style for colouring?
// we can get dapp names for common addresses and display them in the chart

// introduce messaging with grpc and protobuf to communicate between the different components
// use websockets to communicate with the frontend instead of polling the api

// write tests for the different components
// write a dockerfile to package the application
// publish the app on github
// make the ui more interactive and pretty
// suppliment the chart with more information about the token
// make the chart clickable and go to etherscan page to address / node clicked

// write a kubernetes deployment file to deploy the application
