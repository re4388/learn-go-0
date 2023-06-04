package main

import (
	// httpclients "hello/httpclients"

	"fmt"
	"hello/db"
	"hello/learn"
	"hello/server"
)

////////////////////////////////////////

func main() {
	// log
	learn.Run_LearnLog()
	server.Run_serverLog()
	db.Run_dbLog()
	fmt.Println("==========\n")

	///////////
	// acc.RunAccount1()
	// learn.LearnTypes()
	// learn.LearnInterfaces()
	// learn.LearnConcurrency()
	// learn.LearnErrorHandling()
	// learn.LearnControlFlow()
	// learn.RunTime()
	// learn.PrintCircle()
	// learn.UseComponent()
	// learn.RunUrlPool()
	// learn.Run_Cli()
	// learn.Run_wordCount()
	// learn.Run_rot13()
	// learn.Is_the_same_tree()
	// learn.Run_WebCrawler()
	// learn.RUN_CO4()

	///////////
	// learn.RUN_interface1()
	db.RUN_sqlite0()

	// server.RUN_GIN_01()
	// server.RunWebSocketServer()

	// httpclients.RUN_GoHttpClient()
	// httpclients.RUN_restry()

}
