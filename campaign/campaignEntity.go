package campaign

import "time"

type Campaign struct {
	ID               int `json:"campaign_id" gorm:"column:campaign_id"`
	UserID           int
	CampaignName             string
	ShortDescription string
	Description string
	Perks            string
	BeckerCount      int `json:"backer_count" gorm:"column:backer_count"`
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages []CampaignImage
}

type CampaignImage struct {
	ID int `json:"id" gorm:"column:campaign_image_id"`
	CampaignID int `json:"campaign_id" gorm:"column:campaign_id"`
	FileName string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}