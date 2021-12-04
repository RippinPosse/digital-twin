package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"api/internal/graph/model"
	model1 "api/internal/model"
)

func (r *mutationResolver) Authorize(ctx context.Context, input model.AuthorizeUserInput) (*model.AuthorizeUserResult, error) {
	token, err := r.service.User().Authorize(input.Username, input.Password)
	if err != nil {
		return nil, fmt.Errorf("authorize failed: %w", err)
	}

	result := &model.AuthorizeUserResult{
		Success:     true,
		Message:     "authorized",
		AccessToken: token,
	}

	return result, nil
}

func (r *queryResolver) Users(ctx context.Context, id string) (*model1.User, error) {
	panic(fmt.Errorf("not implemented"))
}
