package internal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SignupRepo struct {
	db *gorm.DB
}

func NewSignupRepo(db *gorm.DB) *SignupRepo {
	return &SignupRepo{db: db}
}


// save info to database
func (s *SignupRepo) SaveInfo(userInfo User) error {
	// save into DB
	userInfo.ID = uuid.New()
	res := s.db.Create(&userInfo)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (s *SignupRepo) Migration() {
	s.db.AutoMigrate(&User{})
}


