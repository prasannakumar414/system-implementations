package main

import (
	"sync"
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
	var mutex sync.Mutex
	go func() {
		for range ticker.C {
			if len(p.JobQueue) != 0 && p.Cc > 0 {
				job := p.JobQueue[0]
				mutex.Lock()
				p.JobQueue = p.JobQueue[1:]
				p.Cc--
				mutex.Unlock()
				go func() {
					jobResults <- job.Execute()
				}()
			}
		}
	}()
	for jr := range jobResults {
		meta := p.JobMap[jr.JobID]
		if jr.Success {
			meta.Success = true
		} else {
			if meta.RetryCount > 0 {
				meta.RetryCount--
				p.JobQueue = append(p.JobQueue, meta.Job)
			}
		}
		mutex.Lock()
		p.Cc++
		p.JobMap[jr.JobID] = meta
		mutex.Unlock()
	}
}
