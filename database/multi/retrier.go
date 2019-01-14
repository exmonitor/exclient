package multi

import (
	"context"
	"net/http"
	"time"

	"github.com/exmonitor/exlogger"
	"github.com/olivere/elastic"
)

const (
	maxRetries = 15
)

type ElasticRetrier struct {
	backoff elastic.Backoff
	logger  *exlogger.Logger
}

func (e *ElasticRetrier) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {

	// Stop after 10 retries
	if retry >= maxRetries {
		e.logger.LogError(executionFailedError, "elasticSearchRetrier failed  after %d retries", maxRetries)
		return 0, false, executionFailedError
	}

	e.logger.Log("retrying elasticSearch db request  %d/%d", retry, maxRetries)
	// Let the backoff strategy decide how long to wait and whether to stop
	wait, stop := e.backoff.Next(retry)
	return wait, stop, nil
}

func NewElasticRetrier(logger *exlogger.Logger) *ElasticRetrier {
	return &ElasticRetrier{
		backoff: elastic.NewExponentialBackoff(250*time.Millisecond, 30*time.Second),
		logger:  logger,
	}
}
