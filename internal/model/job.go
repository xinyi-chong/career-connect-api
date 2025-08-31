package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type PostCreateJobInput struct {
	Title         string     `json:"title"`
	CompanyID     string     `json:"company_id"`
	Tag           string     `json:"tag"`
	Description   string     `json:"description"`
	Level         string     `json:"level"`
	Salary        string     `json:"salary"`
	PostedAt      gtime.Time `json:"posted_at"`
	Location      string     `json:"location"`
	IsRemote      bool       `json:"is_remote"`
	IsHybrid      bool       `json:"is_hybrid"`
	Expiry        gtime.Time `json:"expiry"`
	JobQuestion   *string    `json:"job_question"`
	CreatedBy     string     `json:"created_by"`
	CreatedByType string     `json:"created_by_type"`
	UpdatedBy     string     `json:"updated_by"`
	UpdatedByType string     `json:"updated_by_type"`
}
