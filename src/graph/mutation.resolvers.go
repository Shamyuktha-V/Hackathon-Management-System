package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"Hackathon-Management-System/src/analyze"
	"Hackathon-Management-System/src/auth"
	"Hackathon-Management-System/src/graph/model"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// UpdateUser is the resolver for the UpdateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	user, err := auth.GetUser(ctx, r.AppConfig)
	if err != nil {
		return nil, err
	}
	userID, _ := uuid.Parse(user.ID)
	user, err = r.UserService.UpdateUser(ctx, userID, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// RegisterToHackathon is the resolver for the RegisterToHackathon field.
func (r *mutationResolver) RegisterToHackathon(ctx context.Context, input *model.CreateSubmissionInput) (*string, error) {
	_, err := r.SubmissionService.CreateSubmission(ctx, *input)
	if err != nil {
		return nil, err
	}
	status := "Registered successfully"
	return &status, nil
}

// CreateHackathon is the resolver for the CreateHackathon field.
func (r *mutationResolver) CreateHackathon(ctx context.Context, input *model.CreateHackathonInput) (*model.Hackathon, error) {
	user, err := auth.GetUser(ctx, r.AppConfig)
	if err != nil {
		return nil, err
	}
	if user.Role != "admin" {
		return nil, fmt.Errorf("You are not authorized to perform this action")
	} else {
		hackathon, err := r.HackathonService.CreateHackathon(ctx, *input)
		if err != nil {
			return nil, err
		}
		return hackathon, nil
	}
}

// CreateCategories is the resolver for the CreateCategories field.
func (r *mutationResolver) CreateCategories(ctx context.Context, input *model.CreateCategoryInput) (*model.Category, error) {
	user, err := auth.GetUser(ctx, r.AppConfig)
	if err != nil {
		return nil, err
	}
	if user.Role != "admin" {
		return nil, fmt.Errorf("You are not authorized to perform this action")
	} else {
		category, err := r.CategoryService.CreateCategory(ctx, *input)
		if err != nil {
			return nil, err
		}
		return category, nil
	}
}

// AddToTeam is the resolver for the AddToTeam field.
func (r *mutationResolver) AddToTeam(ctx context.Context, inputFrontend *model.CreateTeamMemberInputFrontend) (*string, error) {
	user, err := auth.GetUser(ctx, r.AppConfig)
	if err != nil {
		return nil, err
	}
	if user.Role != "participant" {
		return nil, fmt.Errorf("You are not authorized to perform this action")
	} else {
		teamID, _ := uuid.Parse(inputFrontend.TeamID)
		team, err := r.TeamService.GetTeam(ctx, teamID)
		if err != nil {
			return nil, err
		}
		input := &model.CreateTeamMemberInput{
			TeamID: inputFrontend.TeamID,
			UserID: user.ID,
		}
		_, err = r.TeamMemberService.CreateTeamMember(ctx, *input)
		if err != nil {
			return nil, err
		}
		teamSize := team.TeamSize + 1
		updateInput := &model.UpdateTeamInput{
			TeamSize: &teamSize,
		}
		_, err = r.TeamService.UpdateTeam(ctx, teamID, *updateInput)
		if err != nil {
			return nil, err
		}
		status := "Added to team successfully"
		return &status, nil
	}
}

// CreateTeam is the resolver for the CreateTeam field.
func (r *mutationResolver) CreateTeam(ctx context.Context, teamName string) (*string, error) {
	user, err := auth.GetUser(ctx, r.AppConfig)
	if err != nil {
		return nil, err
	}
	if user.Role != "participant" {
		return nil, fmt.Errorf("You are not authorized to perform this action")
	} else {
		input := &model.CreateTeamInput{
			TeamName: teamName,
			LeaderID: user.ID,
			TeamSize: 1,
		}
		team, err := r.TeamService.CreateTeam(ctx, *input)
		if err != nil {
			return nil, err
		}
		status := "team created successfully"
		print("Team ID :: ", team.ID)
		inputTeamMember := &model.CreateTeamMemberInput{
			TeamID: team.ID,
			UserID: user.ID,
		}
		_, err = r.TeamMemberService.CreateTeamMember(ctx, *inputTeamMember)
		if err != nil {
			return nil, err
		}
		return &status, nil
	}
}

// SubmitCode is the resolver for the SubmitCode field.
func (r *mutationResolver) SubmitCode(ctx context.Context, id string, githubLink string, documentURL graphql.Upload, presentationURL graphql.Upload) (*string, error) {
	re := regexp.MustCompile(`^https://github.com/([^/]+)/([^/]+)/?$`)
	if !re.MatchString(githubLink) {
		return nil, fmt.Errorf("Invalid github link")
	}
	boolValue := true
	submittedAt := time.Now().Format("2006-01-02 15:04:05")
	uploadDir := "./uploads"

	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not create upload directory: %v", err)
	}
	presentationFileName := id + "_presentation_" + presentationURL.Filename

	filePath := filepath.Join(uploadDir, presentationFileName)

	out, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not create file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, presentationURL.File)
	if err != nil {
		return nil, fmt.Errorf("could not save file: %v", err)
	}

	if err != nil {
		return nil, fmt.Errorf("could not create upload directory: %v", err)
	}
	documentFileName := id + "_document_" + documentURL.Filename

	filePath = filepath.Join(uploadDir, documentFileName)

	out, err = os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not create file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, documentURL.File)
	if err != nil {
		return nil, fmt.Errorf("could not save file: %v", err)
	}
	updateInput := model.UpdateSubmissionInput{
		GithubLink:      &githubLink,
		PresentationURL: &presentationFileName,
		DocumentURL:     &documentFileName,
		IsSubmitted:     &boolValue,
		SubmittedAt:     &submittedAt,
	}
	inputID, _ := uuid.Parse(id)
	_, err = r.SubmissionService.UpdateSubmission(ctx, inputID, updateInput)
	if err != nil {
		return nil, err
	}
	status := "Code submitted successfully"
	return &status, nil
}

// AnalyzeCode is the resolver for the AnalyzeCode field.
func (r *mutationResolver) AnalyzeCode(ctx context.Context, githubLink string, problemStatement string) (string, error) {
	re := regexp.MustCompile(`^https://github.com/([^/]+)/([^/]+)/?$`)
	if !re.MatchString(githubLink) {
		return "", fmt.Errorf("Invalid github link")
	}
	results := analyze.AnalyzeCode(githubLink, problemStatement)
	return results, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
