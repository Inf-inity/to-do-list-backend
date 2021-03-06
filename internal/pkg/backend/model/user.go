package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string
	OwnedTasks    []*Task `gorm:"foreignKey:OwnerID"`
	Teams         []*Team `gorm:"many2many:team_members;"`
	OwnedTeams    []*Team `gorm:"many2many:team_owners;"`
	AssignedTasks []*Task `gorm:"many2many:task_assignees;"`
}

type UserInput struct {
	Name  string
	Teams []uint
}

type UserUpdate struct {
	Name            string
	Teams           []uint
	AssignedTaskIDs []uint
	OwnedTasks      []uint
}
