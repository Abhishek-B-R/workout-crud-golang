package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Abhishek-B-R/workout-crud/internals/api"
	"github.com/Abhishek-B-R/workout-crud/internals/store"
)

type Application struct{
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB *sql.DB
}

func NewApplication() (*Application, error){
	pgDb, err := store.Open()
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	workoutHandler := api.NewWorkoutHandler()
	app := &Application{
		Logger: logger, 
		WorkoutHandler: workoutHandler,
		DB: pgDb,
	}
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Server is working pretty fine")
}