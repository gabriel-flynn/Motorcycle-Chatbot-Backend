package models

type Review struct {
	Id                      uint `gorm:"primaryKey;"`
	OverallRating           uint8
	OverallRatingReviewText string
	RideQuality             uint8
	RideQualityReviewText   string
	Engine                  uint8
	EngineReviewText        string
	Reliability             uint8
	ReliabilityReviewText   string
	Value                   uint8
	ValueReviewText         string
	Equipment               uint8
	EquipmentReviewText     string
}
