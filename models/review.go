package models

type Review struct {
	Id                      uint `gorm:"primaryKey;" json:"-"`
	OverallRating           uint8 `json:"overall_rating"`
	OverallRatingReviewText string `json:"overall_rating_review_text"`
	RideQuality             uint8 `json:"ride_quality"`
	RideQualityReviewText   string `json:"ride_quality_review_text"`
	Engine                  uint8 `json:"engine"`
	EngineReviewText        string `json:"engine_review_text"`
	Reliability             uint8 `json:"reliability"`
	ReliabilityReviewText   string `json:"reliability_review_text"`
	Value                   uint8 `json:"value"`
	ValueReviewText         string `json:"value_review_text"`
	Equipment               uint8 `json:"equipment"`
	EquipmentReviewText     string `json:"equipment_review_text"`
}
