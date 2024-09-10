package dboperations

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnect(t *testing.T) {
	// Use sqlmock to create a mock database connection
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock *sql.DB in a SQLDB instance
	// sqlDB := &SQLDB{DB: db}

	// Test the Connect function
	tests := []struct {
		name    string
		dbURL   string
		wantErr bool
	}{
		{
			name:    "Successful connection",
			dbURL:   "file:test.db?mode=memory&cache=shared",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Connect(tt.dbURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuery_helper(t *testing.T) {
	// Use sqlmock to create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock *sql.DB in a SQLDB instance
	sqlDB := &SQLDB{DB: db}

	type TestStruct struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "John").
		AddRow(2, "Jane")

	mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(rows)

	tests := []struct {
		name    string
		query   string
		want    []TestStruct
		wantErr bool
	}{
		{
			name:  "Valid query",
			query: "SELECT id, name FROM users",
			want: []TestStruct{
				{ID: 1, Name: "John"},
				{ID: 2, Name: "Jane"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Query_helper[TestStruct](sqlDB, tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query_helper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query_helper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExec_helper(t *testing.T) {
	// Use sqlmock to create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock *sql.DB in a SQLDB instance
	sqlDB := &SQLDB{DB: db}

	mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))

	tests := []struct {
		name    string
		query   string
		args    []interface{}
		want    Exec
		wantErr bool
	}{
		{
			name:    "Valid insert",
			query:   "INSERT INTO users (name) VALUES (?)",
			args:    []interface{}{"John"},
			want:    Exec{LastInsertId: 1, RowsAffected: 1},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Exec_helper[struct{}](sqlDB, tt.query, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec_helper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec_helper() = %v, want %v", got, tt.want)
			}
		})
	}
}
