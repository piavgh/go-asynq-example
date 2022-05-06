package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

// HandleWelcomeEmailTask handler for welcome email task.
func HandleWelcomeEmailTask(c context.Context, t *asynq.Task) error {
	// Get user ID from given task.
	var p welcomeEmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	// Dummy message to the worker's output.
	log.Printf(" [*] Send Welcome Email to User %d", p.UserID)

	return nil
}

// HandleReminderEmailTask for reminder email task.
func HandleReminderEmailTask(c context.Context, t *asynq.Task) error {
	// Get user ID from given task.
	var p reminderEmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	// Dummy message to the worker's output.
	fmt.Printf("Send Reminder Email to User ID %d\n", p.UserID)
	fmt.Printf("Reason: time is up (%v)\n", p.SendIn)

	return nil
}
