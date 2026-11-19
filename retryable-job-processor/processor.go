package main

import (
	"time"
)

type Job interface {
	Execute() JobResult
}

type Processor struct {
	JobQueue []Job
	Cc       int
	JobMap   map[string]JobMeta
}

func NewProcessor(cc int) *Processor {
	return &Processor{
		JobQueue: make([]Job, 0),
		Cc:       cc,
		JobMap:   make(map[string]JobMeta),
	}
}

/*
background job processor
- input job
- output void
*/
func (p *Processor) Process(jobID string, jobMeta JobMeta) {
	p.JobQueue = append(p.JobQueue, jobMeta.Job)
	p.JobMap[jobID] = jobMeta
}

func (p *Processor) Start() {
	ticker := time.NewTicker(time.Second)
	jobResults := make(chan JobResult, 1)
	for {
		select {
		case <-ticker.C:
			if len(p.JobQueue) != 0 && p.Cc > 0 {
				job := p.JobQueue[0]
				p.JobQueue = p.JobQueue[1:]
				p.Cc--
				go func() {
					jobResults <- job.Execute()
				}()
			}
		case jr := <-jobResults:
			meta := p.JobMap[jr.JobID]
			if jr.Success {
				p.Cc++
				meta.Success = true
			} else {
				if meta.RetryCount > 0 {
					meta.RetryCount--
					p.JobQueue = append(p.JobQueue, meta.Job)
				}
			}
			p.JobMap[jr.JobID] = meta
		}
	}
}
