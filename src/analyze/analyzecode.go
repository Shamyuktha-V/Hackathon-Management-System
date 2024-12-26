package analyze

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jpoz/groq"
)

type GitHubAPIResponse struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Type    string `json:"type"`
	Content string `json:"content"` // Base64 encoded content
}

// Fetch files from the GitHub repository
func fetchFilesFromRepo(owner, repo, path string) ([]GitHubAPIResponse, error) {
	fmt.Printf("inside fetch")
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", owner, repo, path)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var files []GitHubAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return nil, err
	}

	return files, nil
}

// Process each file using Groq API
func processFileWithGroq(problemStatement string, file GitHubAPIResponse, client *groq.Client, wg *sync.WaitGroup, ch chan<- string) {
	fmt.Printf("inside process")
	defer wg.Done()

	// Decode Base64 content
	content, err := base64.StdEncoding.DecodeString(file.Content)
	if err != nil {
		ch <- fmt.Sprintf("Error decoding file %s: %v", file.Name, err)
		return
	}

	// Construct the prompt for summarizing, providing feedback, analyzing problem-statement adherence, scoring, and generating a concise summary for judges
	prompt := fmt.Sprintf(`
    Please analyze the following code and provide:
    1. A summary of its key features including technology , language used
    2. Constructive feedback on the code.
    3. An analysis of its adherence to the problem statement.
    4. A score based on the following criteria:
       - Innovation (out of 10)
       - Feasibility (out of 10)
       - Impact (out of 10)
    5. A concise summary for judges that includes:
       - The purpose of the code.
       - Key functions and their roles.
       - Any dependencies or libraries used.

    Code:
    %s

    Problem Statement:
    %s
`, string(content), problemStatement)

	// Create the chat completion request
	response, err := client.CreateChatCompletion(groq.CompletionCreateParams{
		Model: "llama3-8b-8192", // Specify the model
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	})

	if err != nil {
		ch <- fmt.Sprintf("Error processing file %s: %v", file.Name, err)
		return
	}
	fmt.Printf("-------------------------------------------------------------------------------\n")
	fmt.Printf("Response for file %s: %s\n", file.Name, response.Choices[0].Message.Content)

	// Send the AI's response to the channel
	ch <- fmt.Sprintf("Summary, Feedback, Problem-Statement Adherence, Scores, and Judges' Summary for %s:\n%s", file.Name, response.Choices[0].Message.Content)
}

// Check if the file is a code file based on its extension
func isCodeFile(filename string) bool {
	// Define a list of code file extensions
	codeExtensions := []string{".go", ".py", ".js", ".java", ".c", ".cpp", ".cs", ".rb", ".php", ".html", ".css", ".sh", ".sql"}

	// Get the file extension
	ext := strings.ToLower(filepath.Ext(filename))

	// Check if the extension is in the list of code extensions
	for _, codeExt := range codeExtensions {
		if ext == codeExt {
			return true
		}
	}
	return false
}

func AnalyzeCode(url string, problemStatement string) string {
	// url := "https://github.com/kavushik07/Portfolio"

	// problemStatement := "This project aims to create a personal portfolio website to showcase an individual’s skills, projects, and achievements. The website will feature a responsive design, easy navigation, and sections for bio, work experience, and contact details, providing a professional online presence for career growth and networking opportunities.before storing them in the database, thereby safeguarding sensitive information against unauthorized access.By integrating these features, the application aims to deliver a secure, intuitive, and reliable banking experience for users.This project aims to create a personal portfolio website to showcase an individual’s skills, projects, and achievements. The website will feature a responsive design, easy navigation, and sections for bio, work experience, and contact details, providing a professional online presence for career growth and networking opportunities."

	trimmedURL := strings.TrimPrefix(url, "https://github.com/")

	// Split the remaining part by "/"
	parts := strings.Split(trimmedURL, "/")

	// The first part is the owner, and the second part is the repository name
	owner := parts[0]
	repo := strings.TrimSuffix(parts[1], ".git")
	resultString := ""
	// Initialize Groq client
	client := groq.NewClient(groq.WithAPIKey("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")) // Replace with your actual API key

	// Channel to collect results
	ch := make(chan string)

	var wg sync.WaitGroup

	// Fetch files from the root directory of the repo
	files, err := fetchFilesFromRepo(owner, repo, "")
	if err != nil {
		log.Fatal("Error fetching files:", err)
	}

	// Process each file concurrently
	for _, file := range files {
		if file.Type == "file" && isCodeFile(file.Name) { // Check if the file is a code file
			wg.Add(1)
			go processFileWithGroq(problemStatement, file, client, &wg, ch)
		} else if file.Type == "dir" {
			// Fetch files from subdirectories
			subFiles, err := fetchFilesFromRepo(owner, repo, file.Path)
			if err != nil {
				log.Printf("Error fetching files from directory %s: %v", file.Path, err)
				continue
			}

			// Process subdirectory files concurrently
			for _, subFile := range subFiles {
				if isCodeFile(subFile.Name) { // Check if the subfile is a code file
					wg.Add(1)
					go processFileWithGroq(problemStatement, subFile, client, &wg, ch)
				}
			}
		}
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Collect results from the channel and print
	for result := range ch {
		resultString += result
	}
	print("Result :: ", resultString)
	return resultString

}
