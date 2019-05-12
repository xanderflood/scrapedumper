package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bipol/scrapedumper/pkg/dumper"
	"github.com/bipol/scrapedumper/pkg/martaapi"
	"go.uber.org/zap"
)

type WorkPoller interface {
	Poll(ctx context.Context, errC chan error) error
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ScrapeAndDumpClient
type ScrapeAndDumpClient struct {
	dumper         dumper.Dumper
	scheduleFinder martaapi.ScheduleFinder
	pollTime       time.Duration
	logger         *zap.Logger
}

func New(dumper dumper.Dumper, api martaapi.ScheduleFinder, pollTime time.Duration, logger *zap.Logger) ScrapeAndDumpClient {
	return ScrapeAndDumpClient{
		dumper,
		api,
		pollTime,
		logger,
	}
}

func (c ScrapeAndDumpClient) Poll(ctx context.Context, errC chan error) {
	fmt.Println("starting to poll")
	go func() {
		for {
			select {
			case <-ctx.Done():
				c.logger.Info("exiting poll")
				return
			default:
			}
			err := c.scrapeAndDump(ctx)
			if err != nil {
				errC <- err
				return
			}
			time.Sleep(c.pollTime)
		}
	}()
}

func (c ScrapeAndDumpClient) scrapeAndDump(ctx context.Context) error {
	c.logger.Debug("scrape and dumping")
	schedules, err := c.scheduleFinder.FindSchedules(ctx)
	if err != nil {
		return err
	}
	b, err := json.Marshal(schedules)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	err = c.dumper.Dump(ctx, r, time.Now().Format(time.RFC3339)+".json")
	if err != nil {
		return err
	}
	return nil
}
