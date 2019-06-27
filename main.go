package main

import (
	"os"
	"strings"

	aw "github.com/deanishe/awgo"
	asana "github.com/odeke-em/asana/v1"
)

var (
	wf          *aw.Workflow
	client      *asana.Client
	taskName    string
	taskNote    string
	assignee    string
	workspaceID string
)

func init() {
	wf = aw.New()
	client, _ = asana.NewClient()
	assignee = os.Getenv("ASANA_ASSIGNEE")
	workspaceID = os.Getenv("ASANA_WORKSPACE_ID")
}

func run() {
	if args := os.Args[1:]; len(args) > 0 {
		taskName = args[0]
		if len(args) > 1 {
			taskNote = strings.Join(args[1:], "\n")
		}
	}

	client.CreateTask(&asana.TaskRequest{
		Assignee:  assignee,
		Notes:     taskNote,
		Name:      taskName,
		Workspace: workspaceID,
	})
}

func main() {
	wf.Run(run)
}
