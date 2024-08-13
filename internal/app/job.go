package app

import (
	"context"
	"time"
)

func (a *App) addJob(name, schedule string, fn func(context.Context) error) error {
	job, err := a.scheduler.Cron(schedule).Do(func() {
		ctx := context.Background()
		start := time.Now()
		a.log.Info("starting job", "job", name)
		err := fn(ctx)
		if err != nil {
			a.log.Error("error reporting data: ", err)
			return
		}
		a.log.Info("job finished", "job", name, "duration", time.Since(start))
		// report next run
		nextJob, nextJobTime := a.scheduler.NextRun()
		a.log.Info("next job", "job", nextJob.GetName(), "time", nextJobTime)
	})
	if err != nil {
		return err
	}

	job.Name(name)
	a.log.Info("job added", "name", job.GetName(), "schedule", schedule)
	return nil
}

func (a *App) StartJobs() {
	a.log.Info("scheduler started")
	a.scheduler.StartAsync()
}

func (a *App) StopJobs() {
	a.log.Info("scheduler stopped")
	a.scheduler.Stop()
}
