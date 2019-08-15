package main

import (
	// "encoding/json"
	"log"

	"github.com/YaleOpenLab/openclimate/blockchain"
	"github.com/YaleOpenLab/openclimate/database"
	"github.com/YaleOpenLab/openclimate/oracle"
	"github.com/YaleOpenLab/openclimate/server"
	//"github.com/Varunram/essentials/ipfs"
	//"github.com/YaleOpenLab/openclimate/notif"
)

func main() {

	oracle.Schedule()

	blockchain.CheckTokenBalance()
	database.FlushDB()
	database.CreateHomeDir()
	log.Println("flushed and created new db")
	server.StartServer("8001", true)
}
