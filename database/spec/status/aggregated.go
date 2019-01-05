package status

import "time"

type AgregatedServiceStatus struct {
	Id            string        `json:"id"`
	ServiceID     int           `json:"service_id"`
	Interval      int           `json:"interval"`
	AvgDuration   time.Duration `json:"avg_duration"`
	Aggregated    int           `json:"aggregated"`
	Result        bool          `json:"result"`
	TimestampFrom time.Time     `json:"@timestamp_from"`
	TimestampTo   time.Time     `json:"@timestamp_to"`
}
