// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package cron_gql

type AddJobInput struct {
	// Cron expression for scheduling e.g. '0 * * * *'
	CronExp string `json:"cronExp"`
	// Root directory to run the command
	RootDir string `json:"rootDir"`
	// Terminal-based command
	Cmd string `json:"cmd"`
	// Command arguments
	Args *string `json:"args"`
	// Tags for easier job retrieval
	Tags []*string `json:"tags"`
}

type Job struct {
	// Unique ID for the job (generated)
	JobID int `json:"jobID"`
	// Cron expression used for scheduling e.g. '0 * * * *'
	CronExp string `json:"cronExp"`
	// Root directory to run the command
	RootDir string `json:"rootDir"`
	// Terminal-based command
	Cmd string `json:"cmd"`
	// Command arguments
	Args *string `json:"args"`
	// Tags for easier job retrieval
	Tags []*string `json:"tags"`
	// Last scheduled execution time (human friendly)
	LastRun *string `json:"lastRun"`
	// Next scheduled execution time (human friendly)
	NextRun *string `json:"nextRun"`
}

type JobsInput struct {
	// Return job with unique jobID
	JobID *int `json:"jobID"`
	// Return all jobs which match at least one of the tags
	Tags []*string `json:"tags"`
}
