package main

import "fmt"

type PrintJob struct {
	Id string
}

func (job *PrintJob) Execute() JobResult {
	fmt.Println("Successs!", job.Id)
	return JobResult{
		JobID:   job.Id,
		Success: true,
	}
}
func main() {
	processor := NewProcessor(3)
	go processor.Start()
	p := PrintJob{
		Id: "123",
	}
	processor.Process("Job1", JobMeta{
		JobName:    "J1",
		RetryCount: 2,
		Job:        &p,
	})
	for {

	}
}
