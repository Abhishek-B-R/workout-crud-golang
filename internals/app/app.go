package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Abhishek-B-R/workout-crud/internals/api"
	"github.com/Abhishek-B-R/workout-crud/internals/store"
	"github.com/Abhishek-B-R/workout-crud/migrations"
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

	err = store.MigrateFS(pgDb, migrations.FS, ".")
	if err != nil { 
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutStore := store.NewPostgresWorkoutStore(pgDb)
	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)
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