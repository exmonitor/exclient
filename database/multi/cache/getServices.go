package cache

import (
	"time"

	"github.com/exmonitor/exclient/database/spec/service"
)

type SQL_GetServices struct {
	Cache map[int]SQL_GetServices_Record
}

type SQL_GetServices_Record struct {
	Age  time.Time
	Data []*service.Service
}

// check if cache is still valid
// returns false in case there is no cache or cache is already expired
func (s *SQL_GetServices) IsCacheValid(interval int, ttl time.Duration) bool {
	if r, ok := s.Cache[interval]; ok {
		if r.Age.IsZero() {
			// cache age is not set, cache is not balid
			return false
		} else {
			return time.Now().After(r.Age.Add(ttl))
		}
	} else {
		// no cache for this record, so cache is not valid
		return false
	}
}

// fetch data from cache
func (s *SQL_GetServices) GetData(interval int) []*service.Service {
	if d, ok := s.Cache[interval]; ok {
		return d.Data
	}
	// cached data not found
	return nil
}

// save data to cache
func (s *SQL_GetServices) CacheData(interval int, d []*service.Service) {
	r := SQL_GetServices_Record{
		Age:  time.Now(),
		Data: d,
	}
	s.Cache[interval] = r
}
