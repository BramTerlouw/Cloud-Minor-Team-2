// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type SagaFilter struct {
	SoftDeleted bool            `json:"soft_deleted"`
	ObjectID    string          `json:"object_id"`
	ObjectType  SagaObjectTypes `json:"object_type"`
}

type SagaObject struct {
	ID           string                    `json:"id"`
	Copy         *SagaObject               `json:"copy,omitempty"`
	ObjectID     string                    `json:"object_id"`
	ObjectType   SagaObjectTypes           `json:"object_type"`
	CreatedAt    string                    `json:"created_at"`
	UpdatedAt    *string                   `json:"updated_at,omitempty"`
	Status       *SagaObjectStatusProgress `json:"status,omitempty"`
	ObjectStatus SagaObjectStatus          `json:"object_status"`
	ActionDoneBy string                    `json:"action_done_by"`
	ParentID     *string                   `json:"parent_id,omitempty"`
	Children     []*SagaObject             `json:"children,omitempty"`
}

type SuccessMessage struct {
	ID         string           `json:"id"`
	Text       string           `json:"text"`
	Status     SagaObjectStatus `json:"status"`
	ObjectID   string           `json:"object_id"`
	ObjectType SagaObjectTypes  `json:"object_type"`
}

type SagaObjectStatus string

const (
	SagaObjectStatusExist   SagaObjectStatus = "Exist"
	SagaObjectStatusDeleted SagaObjectStatus = "Deleted"
)

var AllSagaObjectStatus = []SagaObjectStatus{
	SagaObjectStatusExist,
	SagaObjectStatusDeleted,
}

func (e SagaObjectStatus) IsValid() bool {
	switch e {
	case SagaObjectStatusExist, SagaObjectStatusDeleted:
		return true
	}
	return false
}

func (e SagaObjectStatus) String() string {
	return string(e)
}

func (e *SagaObjectStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SagaObjectStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SagaObjectStatus", str)
	}
	return nil
}

func (e SagaObjectStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SagaObjectStatusProgress string

const (
	SagaObjectStatusProgressStarted    SagaObjectStatusProgress = "Started"
	SagaObjectStatusProgressInProgress SagaObjectStatusProgress = "InProgress"
	SagaObjectStatusProgressFinished   SagaObjectStatusProgress = "Finished"
)

var AllSagaObjectStatusProgress = []SagaObjectStatusProgress{
	SagaObjectStatusProgressStarted,
	SagaObjectStatusProgressInProgress,
	SagaObjectStatusProgressFinished,
}

func (e SagaObjectStatusProgress) IsValid() bool {
	switch e {
	case SagaObjectStatusProgressStarted, SagaObjectStatusProgressInProgress, SagaObjectStatusProgressFinished:
		return true
	}
	return false
}

func (e SagaObjectStatusProgress) String() string {
	return string(e)
}

func (e *SagaObjectStatusProgress) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SagaObjectStatusProgress(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SagaObjectStatusProgress", str)
	}
	return nil
}

func (e SagaObjectStatusProgress) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SagaObjectTypes string

const (
	SagaObjectTypesSchool   SagaObjectTypes = "School"
	SagaObjectTypesClass    SagaObjectTypes = "Class"
	SagaObjectTypesExercise SagaObjectTypes = "Exercise"
	SagaObjectTypesResult   SagaObjectTypes = "Result"
	SagaObjectTypesModule   SagaObjectTypes = "Module"
)

var AllSagaObjectTypes = []SagaObjectTypes{
	SagaObjectTypesSchool,
	SagaObjectTypesClass,
	SagaObjectTypesExercise,
	SagaObjectTypesResult,
	SagaObjectTypesModule,
}

func (e SagaObjectTypes) IsValid() bool {
	switch e {
	case SagaObjectTypesSchool, SagaObjectTypesClass, SagaObjectTypesExercise, SagaObjectTypesResult, SagaObjectTypesModule:
		return true
	}
	return false
}

func (e SagaObjectTypes) String() string {
	return string(e)
}

func (e *SagaObjectTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SagaObjectTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SagaObjectTypes", str)
	}
	return nil
}

func (e SagaObjectTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
