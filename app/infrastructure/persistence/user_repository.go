package persistence

import (
	"context"
	"errors"
	"strings"

	"go-graph-demo/app/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userService {
	return &userService{
		db,
	}
}

//We implement the interface defined in the domain
var _ UserService = &userService{}

func (s *userService) CreateUser(ctx context.Context, input *models.NewUser) (*models.User, error) {

	user := &models.User{
		Username:  input.Username,
		DiscordID: input.DiscordID,
	}

	err := s.db.WithContext(ctx).Where(models.User{Username: input.Username}).Attrs(&user).FirstOrCreate(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("username already taken")
		}
		return nil, err
	}

	tx := s.db.WithContext(ctx).Model(&user).Where(&models.User{Username: input.Username}).Find(&user)
	if tx.Error != nil {
		log.Errorf("createBounty: caught error getting bounty document, err %+v", tx.Error)
		return nil, tx.Error
	}

	return &models.User{}, nil

}

func (s *userService) GetUserByID(ctx context.Context, id int) (*models.User, error) {

	user := &models.User{}

	err := s.db.WithContext(ctx).Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUsers(ctx context.Context) ([]*models.User, error) {

	var users []*models.User

	err := s.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
