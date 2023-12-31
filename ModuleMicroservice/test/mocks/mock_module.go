package mocks

import (
	"Module/graph/model"
	"time"
)

var Description = "This is a sample module."
var UpdatedDescription = "Updated Description"
var Difficulty = "A2"
var Category = "Grammatica"
var MadeBy = "Sample User"
var Private = false
var Key = "sample-key"
var timestamp = time.Now().String()
var SoftDeleted = false
var IsSoftDeleted = true

var MockCreateInput = model.ModuleInputCreate{
	Name:        "Sample Module",
	Description: Description,
	Difficulty:  model.LanguageLevel(Difficulty),
	Category:    model.Category(Category),
	Private:     Private,
	Key:         &Key,
}

var MockUpdateInput = model.ModuleInputUpdate{
	Name:        "Sample Module",
	Description: UpdatedDescription,
	Difficulty:  model.LanguageLevel(Difficulty),
	Category:    model.Category(Category),
	Private:     Private,
	Key:         &Key,
}

var MockModule = model.Module{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	Name:        "Sample Module",
	Description: Description,
	Difficulty:  model.LanguageLevel(Difficulty),
	Category:    model.Category(Category),
	MadeBy:      MadeBy,
	Private:     Private,
	Key:         &Key,
	CreatedAt:   &timestamp,
	SoftDeleted: &SoftDeleted,
}

var SoftDeletedMockModule = model.Module{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	Name:        "Sample Module",
	Description: Description,
	Difficulty:  model.LanguageLevel(Difficulty),
	Category:    model.Category(Category),
	MadeBy:      MadeBy,
	Private:     Private,
	Key:         &Key,
	CreatedAt:   &timestamp,
	SoftDeleted: &IsSoftDeleted,
}

var MockModuleInfo = model.ModuleInfo{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	Name:        "Sample Module",
	Description: Description,
	Difficulty:  model.LanguageLevel(Difficulty),
	Category:    model.Category(Category),
	MadeBy:      MadeBy,
	Private:     Private,
}

var MockUpdatedModule = model.Module{
	ID:          "3a3bd756-6353-4e29-8aba-5b3531bdb9ed",
	Name:        "Sample Module",
	Description: UpdatedDescription,
	Difficulty:  model.LanguageLevel(Difficulty),
	Category:    model.Category(Category),
	MadeBy:      MadeBy,
	Private:     Private,
	Key:         &Key,
	CreatedAt:   &timestamp,
	SoftDeleted: &SoftDeleted,
}
