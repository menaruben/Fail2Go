package main

import (
	"fail2go/SSHLogParsing"
	"fail2go/SqlHandling"
	"fail2go/winfw"
	"fmt"
	"log"
	"os"
)

func main() {
	// Set up logging
	logFile, err := os.OpenFile("Fail2Go.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Can't open logfile: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	// must run as administrator!
	bannable_hosts := SSHLogParsing.Get_bannable_hosts(2)
	fmt.Println(bannable_hosts)

	for i := 0; i < len(bannable_hosts); i++ {
		winfw.Ban_IP(bannable_hosts[i])
	}

	/*
		If you have McAfee installed then you might encounter a problem with the sql part
		where McAfee will stop the program from running
	*/

	database := SqlHandling.Sql_create_database("sshjail")

	var table_name string = "sshjail"
	SqlHandling.Sql_create_table(database, table_name)

	SqlHandling.Sql_insert_values(database, table_name, "10.81.72.12", "30.04.2023")

	ip_addrs, release_dates := SqlHandling.Sql_get_values(database, table_name)
	fmt.Println(ip_addrs, release_dates)

	/* TODO:
	- add Release date from sshjail
	- save banned hosts to sql or toml ("ip = release date" form maybe)
	- import banned hosts at start */

	// winfw.Unban_IP("")
}
