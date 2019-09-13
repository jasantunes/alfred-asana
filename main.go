package main

import (
	"fmt"
	"os"
	"strings"

	aw "github.com/deanishe/awgo"
	asana "github.com/odeke-em/asana/v1"
)

var (
	wf          *aw.Workflow
	client      *asana.Client
	taskName    string
	assignee    string
	workspaceID string
)

func init() {
	wf = aw.New()
	wf.Configure(aw.TextErrors(true))
	client, _ = asana.NewClient()
	assignee = os.Getenv("ASANA_ASSIGNEE")
	workspaceID = os.Getenv("ASANA_WORKSPACE_ID")
}

func run() {
	if args := os.Args[1:]; len(args) > 0 {
		taskName = strings.Join(args, " ")
		fmt.Printf("Task added")
	}

	client.CreateTask(&asana.TaskRequest{
		Assignee:  assignee,
		Name:      taskName,
		Workspace: workspaceID,
	})
}

func main() {
	wf.Run(run)
}
