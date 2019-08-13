package main

import (
	// "encoding/json"
	"log"

	"github.com/YaleOpenLab/openclimate/blockchain"
	"github.com/YaleOpenLab/openclimate/database"
	"github.com/YaleOpenLab/openclimate/server"
	"github.com/YaleOpenLab/openclimate/oracle"
	//"github.com/Varunram/essentials/ipfs"
	//"github.com/YaleOpenLab/openclimate/notif"
)

func main() {
	// Interact with the blockchain and check token balance
	_, err := oracle.RetrNoaaGlobalTrendDailyCO2()
	if err != nil {
		log.Fatal(err)
	}
	blockchain.CheckTokenBalance()
	database.FlushDB()
	database.CreateHomeDir()
	log.Println("flushed and created new db")
	server.StartServer("8001", true)
}
