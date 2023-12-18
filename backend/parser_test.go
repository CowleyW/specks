package main

import "testing"

func TestParseCreateStatement(t *testing.T) {
	sql := "CREATE TABLE Users (" +
		"user_id INT AUTO_INCREMENT PRIMARY KEY," +
		"username VARCHAR(60) NOT NULL" +
		"email VARCHAR(100) UNIQUE" +
		"birthdate DATE," +
		"is_active BOOLEAN DEFAULT true" +
		");"

	table, err := parseStatement(sql)
	if err != nil {
		t.Fatal(err)
	}

	if table.name != "Users" {
		t.Fatalf("Expected table name \"Users\", got \"%s\".", table.name)
	}

	if len(table.columns) != 5 {
		t.Fatalf("Expected 5 columns, got %d.", len(table.columns))
	}
}
