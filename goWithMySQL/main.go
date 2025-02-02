// Reference: https://gowebexamples.com/mysql-database/
// Solved: https://stackoverflow.com/questions/47577385/error-non-standard-import-github-com-go-sql-driver-mysql-in-standard-package/67431068#67431068
// Run this Commands in Terminal to Install Required Dependies and Libraries
// go install github.com/go-sql-driver/mysql@latest
// go mod init github.com/go-sql-driver/mysql@latest
// go mod tidy

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() (*sql.DB, error) {
	// MySQL Default username: root and password is ''
	// Name of My Database is music_db
	dsn := "root:@tcp(127.0.0.1:3306)/music_db?parseTime=true"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database Connected Successfully..")
	return db, nil
}

func CreateTable(db *sql.DB) {
	query := `
        CREATE TABLE IF NOT EXISTS instruments (
            instrument_id INT AUTO_INCREMENT,
            instrument_name VARCHAR(40) NOT NULL,
            created_at DATETIME,
            PRIMARY KEY (instrument_id)
        );
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalln("Error creating table: ", err)
	}

	fmt.Println("Table 'instruments' created successfully.")
}

func insertInstrumentsData(instrument_name string, db *sql.DB) {
	query := `
		INSERT INTO instruments (instrument_name, created_at)
		VALUES ( ? , NOW());
	`

	_, err := db.Exec(query, instrument_name)

	if err != nil {
		log.Fatalln("Error Inserting Value: ", err)
	}

	fmt.Println("Instrument Added Successfully...")
}

func updateInstrumentsData(instrument_name string, instrument_id int, db *sql.DB) {
	query := `
		UPDATE instruments 
		SET instrument_name = ?
		WHERE
		instrument_id = ?;
	`

	_, err := db.Exec(query, instrument_name, instrument_id)
	if err != nil {
		log.Fatalln("Error Updating Value: ", err)
	}

	fmt.Println("Instrument Name updated Successfully...")
}

func deleteInstrument(instrument_id int, db *sql.DB) {
	query := `
		DELETE FROM instruments 
		WHERE 
		instrument_id = ?
	`

	_, err := db.Exec(query, instrument_id)
	if err != nil {
		log.Fatalln("Error Deleting Instrument: ", err)
	}

	fmt.Println("Instrument Deleted Successfully...")
}

func printInstruments(db *sql.DB) {
	query := `
		SELECT * FROM instruments;
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln("Error Fetching Instruments: ", err)
	}
	defer rows.Close()

	fmt.Println("******************* Instrument Table *******************")
	for rows.Next() {
		var (
			instrument_id   int
			instrument_name string
			created_at      string
		)

		err := rows.Scan(&instrument_id, &instrument_name, &created_at)
		if err != nil {
			log.Fatalln("Error Fetching Instruments: ", err)
		}

		fmt.Printf("ID: %d, Name: %s, Created At: %s\n", instrument_id, instrument_name, created_at)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	db, err := ConnectToDatabase()
	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
		fmt.Println("Database Connection Error! Software Stopping..")
		return
	}
	defer db.Close()

	var userChoice int

	for {
		fmt.Println()
		fmt.Println("******* Welcome to Golang Database Connectivity *******")
		fmt.Println("1. Insert Instrument")
		fmt.Println("2. Delete Instrument")
		fmt.Println("3. Update Instrument")
		fmt.Println("4. Display All Instruments")
		fmt.Println("5. To Exit")

		fmt.Scanln(&userChoice)

		switch userChoice {
		case 1:
			var instrument_name string
			fmt.Println("Enter Instrument Name: ")
			fmt.Scanln(&instrument_name)
			insertInstrumentsData(instrument_name, db)
		case 2:
			var instrument_id int
			fmt.Println("Enter Instrument Id: ")
			fmt.Scanln(&instrument_id)
			deleteInstrument(instrument_id, db)
		case 3:
			var (
				instrument_id   int
				instrument_name string
			)
			fmt.Println("Enter Instrument Id: ")
			fmt.Scanln(&instrument_id)
			fmt.Println("Enter Instrument Name: ")
			fmt.Scanln(&instrument_name)
			updateInstrumentsData(instrument_name, instrument_id, db)
		case 4:
			printInstruments(db)
		case 5:
			fmt.Println("Until Next Time, Have a Good Time!")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
