package multi

import (
	"fmt"
	"github.com/exmonitor/exclient/database/spec/notification"
	"github.com/exmonitor/exclient/database/spec/service"
	"github.com/pkg/errors"
)

// ********************************************
// MARIA DB
//----------------------------------------------

// intervals table
/*
| intervalSec | CREATE TABLE `intervalSec` (
  `id_interval` int(5) NOT NULL AUTO_INCREMENT,
  `value` int(5) NOT NULL,
  PRIMARY KEY (`id_interval`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 |

*/
func (c *Client) SQL_GetIntervals() ([]int, error) {
	q := "SELECT id,value FROM `intervalSec`"
	// create sql query
	rows, err := c.sqlClient.Query(q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute SQL_GetIntervals")
	}

	var intervals []int
	// read result
	for rows.Next() {
		var id, value int
		err := rows.Scan(id, value)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan values in SQL_GetIntervals")
		}
		intervals = append(intervals, value)
	}

	return intervals, nil
}

func (c *Client) SQL_GetUsersNotificationSettings(serviceID int) ([]*notification.UserNotificationSettings, error) {

	return nil, nil
}

func (c *Client) SQL_GetServices(interval int) ([]*service.Service, error) {
	var services []*service.Service

	fmt.Printf("SQL_GetServices - NOT IMPLEMENTED\n")

	return services, nil
}

func (c *Client) SQL_GetServiceDetails(serviceID int) (*service.Service, error) {

	fmt.Printf("SQL_GetServiceDetails - NOT IMPLEMENTED\n")

	return nil, nil
}
