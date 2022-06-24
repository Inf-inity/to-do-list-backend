package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name    string
	Owners  []*User `gorm:"many2many:team_owners;"`
	Members []*User `gorm:"many2many:team_members;"`
	Tasks   []*Task `gorm:"many2many:task_teams;"`
}

type TeamInput struct {
	Name    string
	Members []uint
	Owners  []uint
	Tasks   []uint
}

type TeamUpdate struct {
	Name    *string
	Members []uint
	Owners  []uint
	Tasks   []uint
}
