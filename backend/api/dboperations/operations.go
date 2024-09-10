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

type SQLDB struct {
	DB *sql.DB
}

type DBHandler interface {
	Query_helper(query string, args ...interface{}) (any, error)
	Exec_helper(query string, args ...interface{}) (Exec, error)
}

// Connect opens a connection to the database and wraps it in an SQLDB struct
func Connect(dbUrl string) (*SQLDB, error) {
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return &SQLDB{DB: db}, nil // Wrap in SQLDB
}



func Query_helper[T any](db *SQLDB, query string, args ...interface{}) ([]T, error) {
	rows, err := db.DB.Query(query, args...)
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

func  Exec_helper[T any](db *SQLDB, query string, args ...interface{}) (Exec, error) {
	exec, err := db.DB.Exec(query, args...)
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