package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// type params []interface{}

func main() {

	start := time.Now()

	query := "select id, name from challenges_v2 where end_time < ? and start_time > ?"
	params := []interface{}{"2018-05-01", "2017-05-20"}

	rows, err := runQuery(query, params)
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}

	fmt.Println("query response", rows)

	time.Sleep(10 * time.Second)
	elapsed := time.Since(start)
	log.Printf("%s took %s", start, elapsed)

}

func runQuery(query string, params []interface{}) ([]interface{}, error) {

	if len(query) <= 0 {
		return nil, errors.New("Query is empty")
	}

	db, err := sql.Open("mysql", "root:123456@tcp(192.168.99.100:3306)/activity")
	if err != nil {
		log.Fatal(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error()) // proper error handling instead of panic in your app
	}

	op := "read"

	switch op {
	case "read":
		rows, err := readRows(db, query, params)
		if err != nil {
			// log.Fatal("Error occured ", err)
			return nil, err
		}
		return rows, nil

	default:
		return []interface{}{}, nil
	}
}

func readRows(db *sql.DB, query string, params []interface{}) ([]interface{}, error) {

	if strings.ContainsAny(query, "?") && len(params) != strings.Count(query, "?") {
		return nil, errors.New("Required params is not the same as the params mentioned in the query")
	}

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		return nil, err
		// log.Println(err)
	}

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	var records []interface{}
	// returnValues := []interface{}{}

	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		result := make(map[string]interface{}, count)

		for i, _ := range columns {

			var v interface{}
			val := values[i]
			b, ok := val.([]byte)

			if ok {
				v = string(b)
			} else {
				v = val
			}

			result[columns[i]] = v
		}
		records = append(records, result)
	}
	return records, nil
}
