// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Category struct {
	ID   string `json:"id"`
	Name string `json:"Name"`
}

type CreateCategoryInput struct {
	Name string `json:"Name"`
}

type CreateHackathonInput struct {
	JudgeID          string    `json:"JudgeID"`
	Name             string    `json:"Name"`
	ProblemStatement string    `json:"ProblemStatement"`
	StartDate        time.Time `json:"StartDate"`
	EndDate          time.Time `json:"EndDate"`
	Duration         int       `json:"Duration"`
	FromDate         time.Time `json:"FromDate"`
	ToDate           time.Time `json:"ToDate"`
	CategoryID       string    `json:"CategoryID"`
}

type CreateSubmissionInput struct {
	TeamID      string `json:"TeamID"`
	HackathonID string `json:"HackathonID"`
}

type CreateTeamInput struct {
	TeamName string `json:"TeamName"`
	LeaderID string `json:"LeaderID"`
	TeamSize int    `json:"TeamSize"`
}

type CreateTeamMemberInput struct {
	TeamID string `json:"TeamID"`
	UserID string `json:"UserID"`
}

type CreateTeamMemberInputFrontend struct {
	TeamID string `json:"TeamID"`
}

type CreateUserInput struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Role  *Role  `json:"Role,omitempty"`
}

type Hackathon struct {
	ID               string `json:"id"`
	JudgeID          string `json:"JudgeID"`
	Name             string `json:"Name"`
	ProblemStatement string `json:"ProblemStatement"`
	StartDate        string `json:"StartDate"`
	EndDate          string `json:"EndDate"`
	Duration         int    `json:"Duration"`
	FromDate         string `json:"FromDate"`
	ToDate           string `json:"ToDate"`
	CategoryID       string `json:"CategoryID"`
}

type Mutation struct {
}

type Query struct {
}

type Submission struct {
	ID               string     `json:"ID"`
	TeamID           string     `json:"TeamID"`
	HackathonID      string     `json:"HackathonID"`
	GithubLink       string     `json:"GithubLink"`
	DocumentURL      string     `json:"DocumentURL"`
	PresentationURL  string     `json:"PresentationURL"`
	SubmittedAt      string     `json:"SubmittedAt"`
	IsSubmitted      bool       `json:"IsSubmitted"`
	KeyFeatures      string     `json:"KeyFeatures"`
	Feedback         string     `json:"Feedback"`
	Adherence        string     `json:"Adherence"`
	InnovationScore  float64    `json:"InnovationScore"`
	FeasibilityScore float64    `json:"FeasibilityScore"`
	ImpactScore      float64    `json:"ImpactScore"`
	Summary          string     `json:"Summary"`
	Hackathon        *Hackathon `json:"Hackathon"`
	Team             *Team      `json:"Team"`
}

type Team struct {
	ID       string `json:"ID"`
	TeamName string `json:"TeamName"`
	LeaderID string `json:"LeaderID"`
	TeamSize int    `json:"TeamSize"`
}

type TeamMember struct {
	ID     string `json:"ID"`
	TeamID string `json:"TeamID"`
	UserID string `json:"UserID"`
}

type UpdateCategoryInput struct {
	Name *string `json:"Name,omitempty"`
}

type UpdateHackathonInput struct {
	JudgeID          *string    `json:"JudgeID,omitempty"`
	Name             *string    `json:"Name,omitempty"`
	ProblemStatement *string    `json:"ProblemStatement,omitempty"`
	StartDate        *time.Time `json:"StartDate,omitempty"`
	EndDate          *time.Time `json:"EndDate,omitempty"`
	Duration         *int       `json:"Duration,omitempty"`
	FromDate         *time.Time `json:"FromDate,omitempty"`
	ToDate           *time.Time `json:"ToDate,omitempty"`
	CategoryID       *string    `json:"CategoryID,omitempty"`
}

type UpdateSubmissionInput struct {
	TeamID           *string  `json:"TeamID,omitempty"`
	HackathonID      *string  `json:"HackathonID,omitempty"`
	GithubLink       *string  `json:"GithubLink,omitempty"`
	DocumentURL      *string  `json:"DocumentURL,omitempty"`
	PresentationURL  *string  `json:"PresentationURL,omitempty"`
	SubmittedAt      *string  `json:"SubmittedAt,omitempty"`
	IsSubmitted      *bool    `json:"IsSubmitted,omitempty"`
	KeyFeatures      *string  `json:"KeyFeatures,omitempty"`
	Feedback         *string  `json:"Feedback,omitempty"`
	Adherence        *string  `json:"Adherence,omitempty"`
	InnovationScore  *float64 `json:"InnovationScore,omitempty"`
	FeasibilityScore *float64 `json:"FeasibilityScore,omitempty"`
	ImpactScore      *float64 `json:"ImpactScore,omitempty"`
	Summary          *string  `json:"Summary,omitempty"`
}

type UpdateTeamInput struct {
	TeamName *string `json:"TeamName,omitempty"`
	LeaderID *string `json:"LeaderID,omitempty"`
	TeamSize *int    `json:"TeamSize,omitempty"`
}

type UpdateTeamMemberInput struct {
	TeamID *string `json:"TeamID,omitempty"`
	UserID *string `json:"UserID,omitempty"`
}

type UpdateUserInput struct {
	Name  *string `json:"Name,omitempty"`
	Email *string `json:"Email,omitempty"`
	Role  *Role   `json:"Role,omitempty"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Role  Role   `json:"Role"`
}

type Role string

const (
	RoleAdmin       Role = "ADMIN"
	RoleParticipant Role = "PARTICIPANT"
	RoleJudge       Role = "JUDGE"
)

var AllRole = []Role{
	RoleAdmin,
	RoleParticipant,
	RoleJudge,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleParticipant, RoleJudge:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
