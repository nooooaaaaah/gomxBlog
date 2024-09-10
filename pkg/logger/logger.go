package logger

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db         *sql.DB
	once       sync.Once
	loggerInit bool
)

type LogLevel int

const (
	INFO LogLevel = iota
	ERROR
	DEBUG
	WARN
	FATAL
)

func (l LogLevel) String() string {
	return [...]string{"INFO", "ERROR", "DEBUG", "WARN", "FATAL"}[l]
}

func InitLogger(dirPath, dbName string) error {
	var err error
	once.Do(func() {
		err = initDB(dirPath, dbName)
		if err == nil {
			loggerInit = true
		}
	})
	return err
}

func initDB(dirPath, dbName string) error {
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("could not create directory: %v", err)
	}

	dbPath := filepath.Join(dirPath, dbName)
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("could not open database: %v", err)
	}

	db = database

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
			level TEXT,
			message TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("could not create logs table: %v", err)
	}

	return nil
}

func logMessage(level LogLevel, message string) {
	if !loggerInit {
		log.Printf("Logger not initialized. Message: [%s] %s", level, message)
		return
	}

	_, err := db.Exec("INSERT INTO logs (level, message) VALUES (?, ?)", level.String(), message)
	if err != nil {
		log.Printf("Failed to log message: %v", err)
	}

	if level == FATAL {
		os.Exit(1)
	}
}

func Info(message string, args ...any) {
	logMessage(INFO, fmt.Sprintf(message, args...))
}

func Error(message string, args ...any) {
	logMessage(ERROR, fmt.Sprintf(message, args...))
}

func Debug(message string, args ...any) {
	logMessage(DEBUG, fmt.Sprintf(message, args...))
}

func Warn(message string, args ...any) {
	logMessage(WARN, fmt.Sprintf(message, args...))
}

func Fatal(message string, args ...any) {
	logMessage(FATAL, fmt.Sprintf(message, args...))
}

func QueryLogs(level LogLevel, limit int) ([]string, error) {
	if !loggerInit {
		return nil, fmt.Errorf("logger not initialized")
	}

	rows, err := db.Query("SELECT timestamp, message FROM logs WHERE level = ? ORDER BY timestamp DESC LIMIT ?", level.String(), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []string
	for rows.Next() {
		var timestamp, message string
		if err := rows.Scan(&timestamp, &message); err != nil {
			return nil, err
		}
		logs = append(logs, fmt.Sprintf("[%s] %s", timestamp, message))
	}

	return logs, nil
}
