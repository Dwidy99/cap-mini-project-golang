package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(int, Pagination) (Pagination, error)
	// GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
	DeleteCampaign(ID int, userID int) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int, pagin Pagination) (Pagination, error) {
	if userID != 0 {
		pagin, err := s.repository.GetCampaigns(pagin)
		if err != nil {
			return pagin, err
		}
		return pagin, nil
	}
	return pagin, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.ID == 0 {
		return campaign, errors.New("No Campaign Found!")
	}

	return campaign, nil
}

func (s *service) DeleteCampaign(ID int, userID int) (Campaign, error) {
	campaignDeleted, err := s.repository.DeleteCampaign(ID)
	if err != nil {
		return campaignDeleted, err
	}
	return campaignDeleted, nil

}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.CampaignName = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID
	slugCampaign := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	// Pembuatan slug
	campaign.Slug = slug.Make(slugCampaign)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not a User of Owner Campaign")
	}

	campaign.CampaignName = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return campaign, err
	}

	return updateCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return CampaignImage{}, errors.New("Not a User of Owner Campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}	
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}