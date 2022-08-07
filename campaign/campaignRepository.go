package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetCampaigns() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDB(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaigns() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}