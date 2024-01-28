package repo

import (
	"gorm.io/gorm"
)

type PlanStatus string

const (
	PlanFinish  PlanStatus = "finish"
	PlanRunning PlanStatus = "running"
)

type PlanCreatorTask struct {
	ID                uint       `gorm:"primarykey" json:"id"`
	Start             int        `json:"count"`
	ColOrderBy        string     `json:"col_order_by"`
	PlanCount         int        `json:"plan_count"`
	CreatorCount      int        `json:"creator_count"`
	LimitCountCreator int64      `json:"limit_count_creator"`
	Email             string     `json:"email"`
	Status            PlanStatus `json:"status"`
}

type PlanCreatorReport struct {
	Task *PlanCreatorTask
	db   *gorm.DB
}

func NewPlanCreatorReport(db *gorm.DB, email string, limit int64) (*PlanCreatorReport, error) {
	task := PlanCreatorTask{
		Start:             0,
		Email:             email,
		PlanCount:         1,
		LimitCountCreator: limit,
		Status:            PlanRunning,
	}

	err := db.Save(&task).Error

	report := PlanCreatorReport{
		Task: &task,
		db:   db,
	}

	return &report, err
}

func (report *PlanCreatorReport) SetCount(c int) error {
	report.Task.CreatorCount = c
	return report.db.Save(report.Task).Error
}

func (report *PlanCreatorReport) SetPlanCount(c int) error {
	report.Task.PlanCount = c
	return report.db.Save(report.Task).Error
}

func (report *PlanCreatorReport) SetLimit(c int64) error {
	report.Task.LimitCountCreator = c
	return nil
}
func (report *PlanCreatorReport) SetFinish() error {
	report.Task.Status = PlanFinish
	return report.db.Save(report.Task).Error
}
func (report *PlanCreatorReport) SetRunning() error {
	report.Task.Status = PlanRunning
	return report.db.Save(report.Task).Error
}
func (report *PlanCreatorReport) SetError(err error) error {
	return nil
}
