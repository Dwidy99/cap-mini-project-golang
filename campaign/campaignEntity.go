package campaign

import (
	"time"
)

type Campaign struct {
	ID               int `json:"campaign_id" gorm:"column:campaign_id"`
	UserID           int `json:"user_id" gorm:"column:user_id"`
	CampaignName             string `json:"campaign_name"`
	ShortDescription string `json:"short_description"`
	Description string `json:"description"`
	Perks            string `json:"perks"`
	BeckerCount      int `json:"backer_count" gorm:"column:backer_count"`
	GoalAmount       int `json:"goal_amount"`
	CurrentAmount    int `json:"current_amount"`
	Slug             string `json:"slug"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CampaignImages []CampaignImage `json:"campaign_images"`
	// User user.User `json:"user"`
}

type CampaignImage struct {
	ID int `json:"id" gorm:"column:campaign_image_id"`
	CampaignID int `json:"campaign_id" gorm:"column:campaign_id"`
	FileName string `json:"file_name"`
	IsPrimary int `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}