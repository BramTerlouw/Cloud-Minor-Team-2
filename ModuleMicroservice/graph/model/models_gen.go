// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Module struct {
	ID          string        `json:"id"`
	SchoolID    string        `json:"school_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Difficulty  LanguageLevel `json:"difficulty"`
	Category    Category      `json:"category"`
	MadeBy      string        `json:"made_by"`
	Private     bool          `json:"private"`
	Key         *string       `json:"key,omitempty"`
	CreatedAt   *string       `json:"created_at,omitempty"`
	UpdatedAt   *string       `json:"updated_at,omitempty"`
	SoftDeleted *bool         `json:"soft_deleted,omitempty"`
}

type ModuleFilter struct {
	SchoolID   *string        `json:"school_id,omitempty"`
	MadeBy     *string        `json:"made_by,omitempty"`
	SoftDelete *bool          `json:"softDelete,omitempty"`
	Name       *NameFilter    `json:"name,omitempty"`
	Difficulty *LanguageLevel `json:"difficulty,omitempty"`
	Category   *Category      `json:"category,omitempty"`
	Private    *bool          `json:"private,omitempty"`
}

type ModuleInfo struct {
	ID          string        `json:"id"`
	SchoolID    string        `json:"school_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Difficulty  LanguageLevel `json:"difficulty"`
	Category    Category      `json:"category"`
	MadeBy      string        `json:"made_by"`
	Private     bool          `json:"private"`
}

type ModuleInputCreate struct {
	Name        string        `json:"name"`
	SchoolID    string        `json:"school_id"`
	Description string        `json:"description"`
	Difficulty  LanguageLevel `json:"difficulty"`
	Category    Category      `json:"category"`
	Private     bool          `json:"private"`
	Key         *string       `json:"key,omitempty"`
}

type ModuleInputUpdate struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Difficulty  LanguageLevel `json:"difficulty"`
	Category    Category      `json:"category"`
	Private     bool          `json:"private"`
	Key         *string       `json:"key,omitempty"`
}

type NameFilter struct {
	Input []*string       `json:"input"`
	Type  NameFilterTypes `json:"type"`
}

type Paginator struct {
	Amount int `json:"amount"`
	Step   int `json:"Step"`
}

type Category string

const (
	CategoryGrammatica            Category = "Grammatica"
	CategorySpelling              Category = "Spelling"
	CategoryWoordenschat          Category = "Woordenschat"
	CategoryUitdrukkingen         Category = "Uitdrukkingen"
	CategoryInterpunctie          Category = "Interpunctie"
	CategoryWerkwoordvervoegingen Category = "Werkwoordvervoegingen"
	CategoryFastTrack             Category = "Fast_Track"
)

var AllCategory = []Category{
	CategoryGrammatica,
	CategorySpelling,
	CategoryWoordenschat,
	CategoryUitdrukkingen,
	CategoryInterpunctie,
	CategoryWerkwoordvervoegingen,
	CategoryFastTrack,
}

func (e Category) IsValid() bool {
	switch e {
	case CategoryGrammatica, CategorySpelling, CategoryWoordenschat, CategoryUitdrukkingen, CategoryInterpunctie, CategoryWerkwoordvervoegingen, CategoryFastTrack:
		return true
	}
	return false
}

func (e Category) String() string {
	return string(e)
}

func (e *Category) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Category(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Category", str)
	}
	return nil
}

func (e Category) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LanguageLevel string

const (
	LanguageLevelA1 LanguageLevel = "A1"
	LanguageLevelA2 LanguageLevel = "A2"
	LanguageLevelB1 LanguageLevel = "B1"
	LanguageLevelB2 LanguageLevel = "B2"
	LanguageLevelC1 LanguageLevel = "C1"
	LanguageLevelC2 LanguageLevel = "C2"
)

var AllLanguageLevel = []LanguageLevel{
	LanguageLevelA1,
	LanguageLevelA2,
	LanguageLevelB1,
	LanguageLevelB2,
	LanguageLevelC1,
	LanguageLevelC2,
}

func (e LanguageLevel) IsValid() bool {
	switch e {
	case LanguageLevelA1, LanguageLevelA2, LanguageLevelB1, LanguageLevelB2, LanguageLevelC1, LanguageLevelC2:
		return true
	}
	return false
}

func (e LanguageLevel) String() string {
	return string(e)
}

func (e *LanguageLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LanguageLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LanguageLevel", str)
	}
	return nil
}

func (e LanguageLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NameFilterTypes string

const (
	NameFilterTypesEq     NameFilterTypes = "eq"
	NameFilterTypesNe     NameFilterTypes = "ne"
	NameFilterTypesStarts NameFilterTypes = "starts"
	NameFilterTypesEnds   NameFilterTypes = "ends"
)

var AllNameFilterTypes = []NameFilterTypes{
	NameFilterTypesEq,
	NameFilterTypesNe,
	NameFilterTypesStarts,
	NameFilterTypesEnds,
}

func (e NameFilterTypes) IsValid() bool {
	switch e {
	case NameFilterTypesEq, NameFilterTypesNe, NameFilterTypesStarts, NameFilterTypesEnds:
		return true
	}
	return false
}

func (e NameFilterTypes) String() string {
	return string(e)
}

func (e *NameFilterTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NameFilterTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NameFilterTypes", str)
	}
	return nil
}

func (e NameFilterTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
