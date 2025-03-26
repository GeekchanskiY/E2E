package main

import (
	"database/sql"
	"fmt"
	"os"

	"finworker/internal/app"

	_ "github.com/lib/pq"
)

func main() {
	var (
		userInput int
		confirm   string
		err       error
	)

	fmt.Print("E2E scripts \n\n")
	fmt.Println(`Choose script to run:
	1 - create superuser
	2 - clear all tables
	3 - drop all tables`)

	_, err = fmt.Scanf("%d", &userInput)
	if err != nil {
		panic(err)
	}

	if userInput == 2 {
		fmt.Println(`Are you sure you want to clear all data? y/n`)
		_, err = fmt.Scanf("%s", &confirm)
		if err != nil {
			panic(err)
		}

		if !confirmed(confirm) {
			fmt.Println("Cancelled")

			os.Exit(0)
		}

		fmt.Println("Clearing all data...")
		clearAllTables()
	}

	if userInput == 3 {
		fmt.Println(`Are you sure you want to drop all tables? y/n`)
		_, err = fmt.Scanf("%s", &confirm)
		if err != nil {
			panic(err)
		}

		if !confirmed(confirm) {
			fmt.Println("Cancelled")

			os.Exit(0)
		}

		fmt.Println("Clearing all data...")
		dropAllTables()
	}
}

func clearAllTables() {
	config := app.NewConfig()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = db.Exec(`
		delete from operations;
		delete from operation_groups;
		delete from user_permission;
		delete from users;
		delete from distributors;
		delete from wallets;
		delete from currency_states;
		delete from banks;
		delete from permission_groups;
	`)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
	os.Exit(0)
}

func dropAllTables() {
	config := app.NewConfig()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = db.Exec(`
		drop table operations;
		drop table operation_groups;
		drop table user_permission;
		drop table users;
		drop table distributors;
		drop table wallets;
		drop table currency_states;
		drop table banks;
		drop table permission_groups;
		drop table migrations;

		drop type gender;
		drop type access_level;
	`)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
	os.Exit(0)
}

func confirmed(s string) bool {
	if s == "y" || s == "Y" || s == "yes" || s == "Yes" || s == "YES" {
		return true
	}

	return false
}
