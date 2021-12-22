package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graph-demo/app/generated"
	"go-graph-demo/app/models"
	"net/http"
)

func (r *mutationResolver) CreateBounty(ctx context.Context, input models.NewBounty) (*models.BountyResponse, error) {
	results, err := r.Repositories.Bounties.CreateBounty(ctx, &input)
	if err != nil {
		return &models.BountyResponse{
			Message: "failed to create bounty",
			Status:  http.StatusInternalServerError,
		}, err
	}

	return &models.BountyResponse{
		Message: "successfully created bounty",
		Status:  http.StatusOK,
		Data:    []*models.Bounty{results},
	}, nil
}

func (r *queryResolver) Bounties(ctx context.Context) (*models.BountyResponse, error) {
	results, err := r.Repositories.Bounties.GetBounties(ctx)
	if err != nil {
		return &models.BountyResponse{
			Message: "failed to retrieve bounties",
			Status:  http.StatusInternalServerError,
		}, err
	}

	return &models.BountyResponse{
		Message: "successfully retrieved bounties",
		Status:  http.StatusOK,
		Data:    results,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
