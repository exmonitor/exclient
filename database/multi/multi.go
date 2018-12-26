package multi

import (
	"github.com/exmonitor/exclient/database"

	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	sqlDriver = "mysql"
)

func DBDriverName() string {
	return "multi"
}

// config for multi db client with mariaDB and elastic search
type Config struct {
	// elastic search
	ElasticConnection string
	// maria db
	MariaConnection   string
	MariaUser         string
	MariaPassword     string
	MariaDatabaseName string
}

type Client struct {
	sqlClient *sql.DB

	// implement client db interface
	database.ClientInterface
}

func New(conf Config) (*Client, error) {
	// create sql connection string
	sqlConnectionString := mysqlConnectionString(conf.MariaConnection, conf.MariaUser, conf.MariaPassword, conf.MariaDatabaseName)
	// init sql connection
	db, err := sql.Open(sqlDriver, sqlConnectionString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create sql connection")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping sql connection")
	}



	newClient := &Client{
		sqlClient:db,
	}
	return newClient, nil
}

func mysqlConnectionString(mariaConnection string, mariaUser string, mariaPassword string, mariaDatabaseName string) string {
	return fmt.Sprintf("%s:%s@%s/%s", mariaUser, mariaPassword, mariaConnection, mariaDatabaseName)
}

// close db connections
func (c *Client) Close() {
	c.sqlClient.Close()
}