// dboperations/operations.go
package dboperations

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Exec struct {
	LastInsertId int64 `json:"lastInsertId"`
	RowsAffected int64 `json:"rowsAffected"`
}

func Connect(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	return db, nil
}

func Query_helper[T any](db *sql.DB, query string, args ...interface{}) ([]T, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	// rows.Close()

	var objects []T

	for rows.Next() {
		var object T

		s := reflect.ValueOf(&object).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		objects = append(objects, object)
		// fmt.Println(user.ID, user.Name)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}
	return objects, nil
}

func Exec_helper[T any](db *sql.DB, query string, args ...interface{}) (Exec, error) {
	exec, err := db.Exec(query, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	// rows.Close()

	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	rows, err := exec.RowsAffected()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	result := Exec{
		LastInsertId: id,
		RowsAffected: rows,
	}
	exec.LastInsertId()
	exec.RowsAffected()

	return result, nil
}
