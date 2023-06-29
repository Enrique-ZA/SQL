package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error connecting to the MySQL Database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the MySQL Database: ", err)
	}

	fmt.Println("Connection established successfully")

	// Create a reader to read from stdin
	reader := bufio.NewReader(os.Stdin)

	// Open the log file
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		fmt.Print("Enter query or EXIT to end: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "EXIT" {
			break
		}

		// Write the query to the log file
		if _, err := file.WriteString(text + "\n"); err != nil {
			log.Println("Error writing to log file:", err)
			continue
		}

		rows, err := db.Query(text)
		if err != nil {
			fmt.Println("Error executing query: ", err)
			continue
		}
		defer rows.Close()

		columns, _ := rows.Columns()

		for rows.Next() {
			columnsData := make([]interface{}, len(columns))
			columnPointers := make([]interface{}, len(columns))
			for i := range columnsData {
				columnPointers[i] = &columnsData[i]
			}

			if err := rows.Scan(columnPointers...); err != nil {
				fmt.Println("Error scanning row: ", err)
				continue
			}

			m := make(map[string]string)
			for i, colName := range columns {
				val := columnPointers[i].(*interface{})
				str, ok := (*val).([]byte)
				if ok {
					m[colName] = string(str)
				} else {
					m[colName] = fmt.Sprintf("%v", *val)
				}
			}

			// Print column names
			for _, colName := range columns {
				fmt.Printf("| %15s ", colName)
			}
			fmt.Println("|")

			// Print row data
			for _, colName := range columns {
				fmt.Printf("| %15s ", m[colName])
			}
			fmt.Println("|")

			fmt.Println(strings.Repeat("-", len(columns)*17+1))
		}

		if rows.Err() != nil {
			fmt.Println("Error fetching rows: ", err)
			continue
		}
	}

	db.Close()
	fmt.Println("Connection closed")
}

