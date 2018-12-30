package multi

import (
	"fmt"
	"time"

	"github.com/exmonitor/chronos"
	"github.com/exmonitor/exclient/database/spec/status"
	"github.com/pkg/errors"
)

// **************************************************
// ELASTIC SEARCH
///--------------------------------------------------
func (c *Client) ES_GetFailedServices(from time.Time, to time.Time, interval int) ([]*status.ServiceStatus, error) {
	// just dummy record return
	fmt.Printf("ES_GetFailedServices - NOT IMPLEMENTED\n")

	return nil, nil
}

func (c *Client) ES_SaveServiceStatus(s *status.ServiceStatus) error {
	t := chronos.New()

	// insert data to elasticsearch db
	fmt.Printf("DEBUG: saving %#v\n",*s)
	_, err := c.esClient.Index().Index(esStatusIndex).Type(esStatusDocName).BodyJson(s).Do(c.ctx)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to save service status for id %d", s.Id))
	}

	t.Finish()
	if c.timeProfiling {
		c.logger.LogDebug("TIME_PROFILING: executed ES_SaveServiceStatus:ID:%d in %sms", s.Id, t.StringMilisec())
	}
	return nil
}
