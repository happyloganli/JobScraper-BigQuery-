// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type JobPost struct {
	ID                      uuid.UUID
	DetailUrl               string
	Title                   string
	Company                 string
	Location                string
	Level                   string
	ApplyUrl                string
	MinimumQualifications   []string
	PreferredQualifications []string
	AboutJob                []string
	Responsibilities        []string
	CreatedAt               time.Time
}
