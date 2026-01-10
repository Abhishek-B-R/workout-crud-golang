package authz

import "github.com/Abhishek-B-R/workout-crud/internals/store"

func CanDeleteWorkout(user *store.User, workout *store.Workout) bool {
	if user.IsAnonymous() {
		return false
	}

	return workout.UserID == user.ID
}

func CanUpdateWorkout(user *store.User, workout *store.Workout) bool {
	if user.IsAnonymous() {
		return false
	}

	return workout.UserID == user.ID
}