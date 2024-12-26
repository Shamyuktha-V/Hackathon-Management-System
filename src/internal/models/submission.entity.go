package models

import (
	"github.com/google/uuid"
)

type Submission struct {
	ID               uuid.UUID `gorm:"primaryKey;column:id"` // Primary key for submission
	TeamID           uuid.UUID `gorm:"size:36;not null;column:team_id"`
	Team             Team      `gorm:"foreignKey:TeamID;references:ID"`
	Hackathon        Hackathon `gorm:"foreignKey:HackathonID;references:ID"`
	HackathonID      uuid.UUID `gorm:"size:36;not null;column:hackathon_id"` // Foreign key referencing Hackathon table
	GithubLink       string    `gorm:"size:255;column:github_link"`          // URL to the Github repository
	DocumentURL      string    `gorm:"size:255;column:document_url"`         // URL to the submitted document
	PresentationURL  string    `gorm:"size:255;column:presentation_url"`     // URL to the presentation file
	SubmittedAt      string    `gorm:"default:NULL;column:submitted_at"`     // Timestamp of when the submission was made
	IsSubmitted      bool      `gorm:"not null;column:is_submitted"`         // Boolean indicating if the submission is complete
	KeyFeatures      string    `gorm:"type:text;column:key_features"`        // Key features of the submission
	Feedback         string    `gorm:"type:text;column:feedback"`            // Feedback on the submission
	Adherence        string    `gorm:"type:text;column:adherence"`           // Adherence score (e.g., 0-10 scale)
	InnovationScore  float64   `gorm:"column:innovation_score"`              // Innovation score (e.g., 0-10 scale)
	FeasibilityScore float64   `gorm:"column:feasibility_score"`             // Feasibility score (e.g., 0-10 scale)
	ImpactScore      float64   `gorm:"column:impact_score"`                  // Impact score (e.g., 0-10 scale)
	Summary          string    `gorm:"type:text;column:summary"`             // Summary of the submission
}

func (s Submission) TableName() string {
	return "submissions"
}
