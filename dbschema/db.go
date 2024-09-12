package dbschema

import (
	"Blog/internal/db"
	"Blog/pkg/logger"
	"context"
	"database/sql"
	_ "embed"
	"os"
)

//go:embed schema.sql
var ddl string

func ConnectCreateDoShitWithDB() (*db.Queries, *sql.DB) {
	blogDB, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		logger.Fatal("Failed to open db: %v", err)
	}

	if _, err := os.Stat("./blog.db"); os.IsNotExist(err) {
		if _, err := blogDB.ExecContext(context.Background(), ddl); err != nil {
			logger.Fatal("Failed to create tables: %v", err)
		}
	}

	queries := db.New(blogDB)
	return queries, blogDB
}
