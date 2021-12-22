package persistence

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Bounties BountyService
	Users    UserService
	db       *gorm.DB
}
