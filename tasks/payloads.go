package tasks

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

const (
	// TypeWelcomeEmail is a name of the task type
	// for sending a welcome email.
	TypeWelcomeEmail = "email:welcome"

	// TypeReminderEmail is a name of the task type
	// for sending a reminder email.
	TypeReminderEmail = "email:reminder"
)

type welcomeEmailTaskPayload struct {
	// ID for the email recipient.
	UserID int
}

type reminderEmailTaskPayload struct {
	// ID for the email recipient.
	UserID int
	SendIn string
}

// NewWelcomeEmailTask task payload for a new welcome email.
func NewWelcomeEmailTask(id int) (*asynq.Task, error) {
	payload, err := json.Marshal(welcomeEmailTaskPayload{UserID: id})

	if err != nil {
		return nil, err
	}

	// Return a new task with given type and payload.
	return asynq.NewTask(TypeWelcomeEmail, payload), nil
}

// NewReminderEmailTask task payload for a reminder email.
func NewReminderEmailTask(id int, ts time.Time) (*asynq.Task, error) {
	payload, err := json.Marshal(reminderEmailTaskPayload{
		UserID: id,
		SendIn: ts.String(),
	})

	if err != nil {
		return nil, err
	}

	// Return a new task with given type and payload.
	return asynq.NewTask(TypeReminderEmail, payload), nil
}
