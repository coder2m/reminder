/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:17
 */
package database

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Client struct {
	db *gorm.DB
}

func NewDatabaseClient(o *Options, stopCh <-chan struct{}) (c *Client, err error) {
	db, err := gorm.Open(o.Type, o.GetDSN())

	if err != nil {
		log.Printf("unable to connect to database", err)
		return nil, err
	}
	go func() {
		<-stopCh
		if err := db.Close(); err != nil {
			log.Printf("error happened during closing database connection", err)
		}
	}()

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTablePrefix + defaultTableName
	}

	if o.Debug {
		db = db.Debug()
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(o.MaxIdleConnections)
	db.DB().SetMaxOpenConns(o.MaxOpenConnections)
	return &Client{db: db}, nil

}

func (c *Client) DB() *gorm.DB {
	if c == nil {
		log.Print("database client is nil")
		return nil
	}
	return c.db
}
