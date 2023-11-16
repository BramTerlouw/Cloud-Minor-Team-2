package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"example/graph/model"
)

// CreateSchool is the resolver for the createSchool field.
func (r *mutationResolver) CreateSchool(ctx context.Context, input model.SchoolInput) (*model.School, error) {
	//headers := ctx.Value("headers").(http.Header)
	//
	//// Access tokens from the headers, e.g., for Bearer token
	//accessToken := headers.Get("Authorization")

	school, err := r.Service.CreateSchool(input)
	if err != nil {
		return nil, err
	}

	return school, nil
}

// UpdateSchool is the resolver for the updateSchool field.
func (r *mutationResolver) UpdateSchool(ctx context.Context, id string, input model.SchoolInput) (*model.School, error) {
	school, err := r.Service.UpdateSchool(id, input)
	if err != nil {
		return nil, err
	}

	return school, nil
}

// DeleteSchool is the resolver for the deleteSchool field.
func (r *mutationResolver) DeleteSchool(ctx context.Context, id string) (*string, error) {
	err := r.Service.DeleteSchool(id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

// GetSchool is the resolver for the getSchool field.
func (r *queryResolver) GetSchool(ctx context.Context, id string) (*model.School, error) {
	school, err := r.Service.GetSchoolById(id)
	if err != nil {
		return nil, err
	}

	return school, nil
}

// ListSchools is the resolver for the listSchools field.
func (r *queryResolver) ListSchools(ctx context.Context) ([]*model.School, error) {
	// Retrieve a list of all modules using the repository.
	schools, err := r.Service.ListSchools()
	if err != nil {
		return nil, err
	}

	return schools, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
