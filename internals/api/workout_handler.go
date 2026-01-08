package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Abhishek-B-R/workout-crud/internals/store"
	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{
	workoutStore store.WorkoutStore
}

func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
	}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request){
	paramsWorkoutId := chi.URLParam(r, "id")
	if paramsWorkoutId == "" {
		http.NotFound(w,r)
		return
	}

	workoutID,err := strconv.ParseInt(paramsWorkoutId, 10, 64)
	if err != nil{
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w, "This is the workout id %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request){
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create workout", http.StatusInternalServerError)
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(createdWorkout)
}