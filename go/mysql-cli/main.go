package main

import (
	"bufio"
	"database/sql"
	"fmt"
    "github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"log"
)
func main() {

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

	reader := bufio.NewReader(os.Stdin)

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

        writeHTML(rows)

		fmt.Println("Generated index.html successfully")
	}

	db.Close()
	fmt.Println("Connection closed")
}

func writeHTML(rows *sql.Rows) {
    cols, _ := rows.Columns()

    htmlFile, _ := os.Create("index.html")
    defer htmlFile.Close()

    htmlFile.WriteString(`<!DOCTYPE html>
    <html>
    <head>
    <style>
    /* Style the table */
    table {
        border-collapse: collapse;
        width: 100%;
    }

    /* Style the table cells */
    th, td {
        border: 1px solid black;
        text-align: left;
        padding: 8px;
    }

    /* Style the table headers */
    th {
        background-color: #f2f2f2;
    }
    </style>
    </head>
    <body>
    <table>
    `)

    htmlFile.WriteString("<tr>\n")
    for _, column := range cols {
        htmlFile.WriteString("<th>" + column + "</th>")
    }
    htmlFile.WriteString("</tr>\n")

    vals := make([]interface{}, len(cols))
    for i := 0; i < len(cols); i++ {
        vals[i] = new(sql.RawBytes)
    }

    for rows.Next() {
        err := rows.Scan(vals...)
        if err != nil {
            log.Fatal(err)
        }
        htmlFile.WriteString("<tr>\n")
        for _, val := range vals {
            str := string(*val.(*sql.RawBytes)) // Convert byte array to string
            htmlFile.WriteString("<td>" + str + "</td>")
        }
        htmlFile.WriteString("</tr>\n")
    }

    htmlFile.WriteString(`
    </table>
    </body>
    </html>`)
}

