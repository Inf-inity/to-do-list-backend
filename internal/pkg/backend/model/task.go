package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Assignees     []*User `gorm:"many2many:task_assignees;"`
	OwnerID       uint
	Title         string
	Description   string
	InvolvedTeams []*Team `gorm:"many2many:task_teams;"`
	State         TaskState
	Priority      Priority
}

type TaskState uint64

type Priority int64

type TaskInput struct {
	Assignees     []uint
	Owner         uint
	Title         string
	Description   string
	InvolvedTeams []uint
	Priority      Priority
}

type TaskUpdate struct {
	Assignees     []uint
	Owner         *uint
	Title         *string
	Description   *string
	InvolvedTeams []uint
	Priority      *Priority
}
