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
