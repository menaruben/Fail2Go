package SqlHandling

// package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// creates an sql database file
func Sql_create_database(database_name string) *sql.DB {
	database, err := sql.Open("sqlite3", fmt.Sprintf("%s.db", database_name))
	check(err)

	return database
}

// creates an sql table
func Sql_create_table(database *sql.DB, table_name string) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (ip_addr TEXT, release_date TEXT)", table_name)
	statement, err := database.Prepare(query)
	check(err)

	statement.Exec()
}

// inserts values into an sql database
func Sql_insert_values(database *sql.DB, table_name string, ip_addr string, release_date string) {
	query := fmt.Sprintf("INSERT INTO %s (ip_addr, release_date) VALUES (?, ?)", table_name)
	statement, err := database.Prepare(query)
	check(err)

	statement.Exec(ip_addr, release_date)
}

// gets the values of ip_addr and release_date columns of sql database
func Sql_get_values(database *sql.DB, table_name string) ([]string, []string) {
	query := fmt.Sprintf("SELECT ip_addr, release_date FROM %s", table_name)
	rows, err := database.Query(query)
	check(err)

	var ip_addrs []string
	var release_dates []string
	var ip_addr string
	var release_date string

	for rows.Next() {
		rows.Scan(&ip_addr, &release_date)
		ip_addrs = append(ip_addrs, ip_addr)
		release_dates = append(release_dates, release_date)
	}

	return ip_addrs, release_dates
}

// func main() {
// 	database := Sql_create_database("sshjail")

// 	var table_name string = "sshjail"
// 	Sql_create_table(database, table_name)

// 	Sql_insert_values(database, table_name, "10.81.72.12", "30.04.2023")

// 	ip_addrs, release_dates := Sql_get_values(database, table_name)
// 	fmt.Println(ip_addrs, release_dates)
// }
