package dtos

import "github.com/SwipEats/SwipEats/server/internal/models"

type GroupRestaurantResponseDto struct {
	ID            uint   `json:"id"`
	GroupID       uint   `json:"group_id"`
	RestaurantID  uint   `json:"restaurant_id,omitempty"`
	Restaurant    models.Restaurant `json:"restaurant"`
	DistanceInKM  float64 `json:"distance_in_km,omitempty"`
}