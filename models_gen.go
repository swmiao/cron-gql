// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package cron_gql

type AddJobInput struct {
	CronExp string `json:"cronExp"`
	RootDir string `json:"rootDir"`
	Cmd     string `json:"cmd"`
	Args    string `json:"args"`
}

type Job struct {
	EntryID      string  `json:"entryId"`
	NextTime     *string `json:"NextTime"`
	PreviousTime *string `json:"PreviousTime"`
}
