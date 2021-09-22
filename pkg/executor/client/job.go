package client

import (
	"fmt"
	"time"

	"github.com/kubeshop/kubtest/pkg/api/kubtest"
	"github.com/kubeshop/kubtest/pkg/jobs"
)

func NewJobExecutorClient() (client JobExecutorClient, err error) {
	jobClient, err := jobs.NewJobClient()
	if err != nil {
		return client, fmt.Errorf("can't get k8s jobs client: %w", err)
	}

	return JobExecutorClient{
		Client: jobClient,
	}, nil
}

type JobExecutorClient struct {
	Client *jobs.JobClient
}

// Watch will get valid execution after async Execute, execution will be returned when success or error occurs
// Worker should set valid state for success or error after script completion
// TODO add timeout
func (c JobExecutorClient) Watch(id string) (events chan ResultEvent) {
	events = make(chan ResultEvent)

	go func() {
		ticker := time.NewTicker(WatchInterval)
		for range ticker.C {
			execution, err := c.Get(id)

			events <- ResultEvent{
				Result: execution,
				Error:  err,
			}

			if err != nil || execution.IsCompleted() {
				close(events)
				return
			}
		}

	}()

	return events
}

func (c JobExecutorClient) Get(id string) (execution kubtest.Result, err error) {
	// TODO Get Logs ? Update Execution
	return
}

// Execute starts new external script execution, reads data and returns ID
// Execution is started asynchronously client can check later for results
func (c JobExecutorClient) Execute(options ExecuteOptions) (result kubtest.Result, err error) {
	// TODO move to mapper
	execution := kubtest.NewExecutionWithID(options.ID, options.ScriptSpec.Type_, options.ScriptSpec.Name)
	execution.ScriptContent = options.ScriptSpec.Content
	execution.Repository = (*kubtest.Repository)(options.ScriptSpec.Repository)
	execution.Params = options.Request.Params

	return c.Client.LaunchK8sJob(options.ID, options.ExecutorSpec.Image, execution)
}

func (c JobExecutorClient) Abort(id string) error {
	c.Client.AbortK8sJob(id)
	return nil
}
