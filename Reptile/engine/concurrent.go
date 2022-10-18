package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler //Scheduler的来源
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		creatWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	// 收out
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: #%d : %v", itemCount, item)
			itemCount++
		}

		// request 送给Scheduler
		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}

func creatWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
