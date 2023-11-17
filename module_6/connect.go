package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Task represents a task in the database.
type Task struct {
	ID          int
	Description string
}

func main() {
	// Create a database connection pool
	pool, err := NewDatabasePool()
	if err != nil {
		log.Fatalf("Failed to create database pool: %v", err)
	}

	// Defer closing the database pool to ensure it's closed when done
	defer pool.Close()

	// Example: Fetch all tasks from the database
	tasks, err := GetAllTasks(pool)
	if err != nil {
		log.Fatalf("Failed to fetch tasks: %v", err)
	}

	InsertTask(pool, "Task 2", "Description 2")

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s\n", task.ID, task.Description)
	}
}

// NewDatabasePool creates a PostgreSQL database connection pool.
func NewDatabasePool() (*pgxpool.Pool, error) {
	// Replace with your PostgreSQL connection string
	connStr := "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

// GetAllTasks fetches all tasks from the database.
func GetAllTasks(db *pgxpool.Pool) ([]Task, error) {
	var tasks []Task

	query := "SELECT id, description FROM tasks"

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func InsertTask(db *pgxpool.Pool, title, description string) (int, error) {
	var id int

	query := "INSERT INTO tasks (title, description, created_at) VALUES ($1, $2, NOW()) RETURNING id"
	err := db.QueryRow(context.Background(), query, title, description).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateTask(db *pgxpool.Pool, id int, title, description string) error {
	query := "UPDATE tasks SET title = $1, description = $2 WHERE id = $3"
	_, err := db.Exec(context.Background(), query, title, description, id)
	return err
}

func DeleteTask(db *pgxpool.Pool, id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := db.Exec(context.Background(), query, id)
	return err
}
