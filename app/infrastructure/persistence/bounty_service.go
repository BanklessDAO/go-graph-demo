package persistence

import (
	"context"

	"go-graph-demo/app/models"
)

type BountyService interface {
	CreateBounty(ctx context.Context, input *models.NewBounty) (model *models.Bounty, err error)
	GetBountyByID(ctx context.Context, input int) (model *models.Bounty, err error)
	GetBounties(ctx context.Context) (model []*models.Bounty, err error)
}
