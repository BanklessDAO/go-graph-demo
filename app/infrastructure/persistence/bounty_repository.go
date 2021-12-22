package persistence

import (
	"context"

	"go-graph-demo/app/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type bountyService struct {
	db *gorm.DB
}

func NewBountyRepository(db *gorm.DB) *bountyService {
	return &bountyService{
		db,
	}
}

//We implement the interface defined in the domain
var _ BountyService = &bountyService{}

func (s *bountyService) CreateBounty(ctx context.Context, input *models.NewBounty) (*models.Bounty, error) {
	bounty := &models.Bounty{
		Text:   input.Text,
		UserID: input.UserID,
	}

	err := s.db.WithContext(ctx).Where(models.Bounty{Text: input.Text}).FirstOrCreate(&bounty).Error
	if err != nil {
		return nil, err
	}

	tx := s.db.WithContext(ctx).Model(bounty).Where(models.Bounty{Text: input.Text}).Find(&bounty)
	if tx.Error != nil {
		log.Errorf("caught error retrieving bounty document, err %+v", tx.Error)
		return nil, tx.Error
	}

	return bounty, nil

}

func (s *bountyService) GetBountyByID(ctx context.Context, id int) (*models.Bounty, error) {

	bounty := &models.Bounty{}

	err := s.db.WithContext(ctx).Where("id = ?", id).Take(&bounty).Error
	if err != nil {
		return nil, err
	}

	return bounty, nil
}

func (s *bountyService) GetBounties(ctx context.Context) ([]*models.Bounty, error) {

	var bounties []*models.Bounty

	err := s.db.WithContext(ctx).Preload(clause.Associations).Find(&bounties).Error
	if err != nil {
		return nil, err
	}

	return bounties, nil
}
