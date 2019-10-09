package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "arch", //TODO
		User:       "userarch",
		Password:   "test",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://userarch:test@localhost/arch?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}