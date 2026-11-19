package main

type JobResult struct {
	Success bool
	JobID   string
}

type JobMeta struct {
	RetryCount int
	Success    bool
	JobName    string
	Job        Job
}
