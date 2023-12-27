// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type InputResult struct {
	ExerciseID string `json:"exercise_id"`
	UserID     string `json:"user_id"`
	ClassID    string `json:"class_id"`
	ModuleID   string `json:"module_id"`
	Input      string `json:"input"`
	Result     string `json:"result"`
}

type Paginator struct {
	Amount int `json:"amount"`
	Step   int `json:"step"`
}

type Result struct {
	ID          string `json:"id"`
	ExerciseID  string `json:"exercise_id"`
	UserID      string `json:"user_id"`
	ClassID     string `json:"class_id"`
	ModuleID    string `json:"module_id"`
	Input       string `json:"input"`
	Result      string `json:"result"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	SoftDeleted bool   `json:"soft_deleted"`
}

type ResultFilter struct {
	SoftDelete *bool   `json:"softDelete,omitempty"`
	ExerciseID *string `json:"exerciseId,omitempty"`
	UserID     *string `json:"userId,omitempty"`
	ClassID    *string `json:"classId,omitempty"`
	ModuleID   *string `json:"moduleId,omitempty"`
	Input      *string `json:"input,omitempty"`
	Result     *string `json:"result,omitempty"`
}
