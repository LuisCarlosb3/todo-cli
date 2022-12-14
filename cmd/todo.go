package main

import (
	"todo/cmd/comands"
	"todo/internal/task/database"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/spf13/cobra"
)

func createDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		panic("error on connection with database")
	}
	repo:=database.NewTaskRepository(db)
	repo.CreateTaskTableIfNoExists()
	return db
}


func main() {
	db := createDatabase()
	
	rootCmd := cobra.Command{
		Use: "task",
	}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Version = "1.0.0"
	rootCmd.AddCommand(comands.AddTaskComand(db))
	rootCmd.AddCommand(comands.CompleteTaskComand(db))
	rootCmd.AddCommand(comands.DeleteTaskComand(db))
	rootCmd.AddCommand(comands.ListTasksComand(db))
	rootCmd.AddCommand(comands.NextTasksComand(db))
	rootCmd.Execute()
}
