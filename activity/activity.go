package activity

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Event struct {
	Type      string    `json:"type"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Ref     string   `json:"ref"`
	RefType string   `json:"ref_type"`
	Action  string   `json:"action"`
	Commits []Commit `json:"commits"`
}

type Commit struct {
	Message string `json:"message"`
}

func FetchActivity(username string) ([]Event, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events?per_page=10", username))
	if err != nil {
		return nil, fmt.Errorf("error fetching user activity: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user %s not found. Please check the username and try again", username)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch user activity: Status Code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var events []Event
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return events, nil
}

func DisplayActivity(events []Event, user string) {
	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	fmt.Printf("Recent Activity for %s\n", user)
	fmt.Println("================================")

	for _, event := range events {
		switch event.Type {
		case "CreateEvent":
			fmt.Printf("- Created %s in %s\n", event.Payload.RefType, event.Repo.Name)
		case "ForkEvent":
			fmt.Printf("- Forked %s\n", event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- %s an issue in %s\n", event.Payload.Action, event.Repo.Name)
		case "PushEvent":
			commitCount := len(event.Payload.Commits)
			if commitCount == 1 {
				fmt.Printf("- Pushed 1 commit to %s\n", event.Repo.Name)
				fmt.Printf("  Commit: %s\n", event.Payload.Commits[0].Message)
			} else {
				fmt.Printf("- Pushed %d commits to %s\n", commitCount, event.Repo.Name)
			}
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		default:
			fmt.Printf("- %s in %s\n", event.Type, event.Repo.Name)
		}
		fmt.Printf("  Date: %s\n\n", event.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
