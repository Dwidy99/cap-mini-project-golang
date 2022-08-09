package campaign

import (
	"fmt"
	"funding-api/user"
	"math"

	"gorm.io/gorm"
)

type Repository interface {
	GetCampaigns(Pagination) (Pagination, error)
	// GetCampaigns() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campiagnImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
	DeleteCampaign(ID int) (Campaign, error)
	FindById(ID int) (Campaign, error)
	GetByUserID(int) (user.User, error) 
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDB(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaigns(pagination Pagination) (Pagination, error) {
	var pagin Pagination
	
	tr, totalPages, fromRow, toRow := 0, 0, 0, 0
	offset := pagination.Page * pagination.Limit
	
	var campaigns []Campaign
	var campaign Campaign
	
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Limit(pagination.Limit).Offset(offset).Find(&campaigns).Error
	if err != nil {
		fmt.Println("ERROR: ", err)
		return pagin, err
	}

	pagination.Rows = campaigns

	// count all data
	totalRows := int64(tr)
	errCount := r.db.Model(campaign).Count(&totalRows).Error
	if errCount != nil {
		return pagin, errCount
	}

	pagination.TotalRows = totalRows

	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > tr {
		// set to row with total rows
		toRow = tr
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return pagination, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) GetByUserID(ID int) (user.User, error) {
	var user user.User

	err := r.db.Where("user_id = ?", ID).Find(&user).Error
	if err != nil {
		return user, nil
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign
	
	err := r.db.Preload("CampaignImages").Where("campaign_id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) FindById(ID int) (Campaign, error) {
	var campaign Campaign
	
	err := r.db.Where("user_id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) DeleteCampaign(ID int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Where("campaign_id = ?", ID).Delete(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}	

	return true, nil
}