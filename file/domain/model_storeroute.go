package domain

import "gorm.io/gorm"

type RouteStore struct {
	gorm.Model
	RouteName string `json:"route_name" dynamodbav:"route_name" gorm:"not null"`
	Store     string `json:"store" dynamodbav:"store" gorm:"not null"`
}
