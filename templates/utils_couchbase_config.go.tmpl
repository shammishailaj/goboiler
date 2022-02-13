package utils

import (
	"fmt"
	"github.com/couchbase/gocb"
)

type CouchbaseConfig struct {
	HostName   string // Hostname of the SQL server
	Port       int64  // Port number of the SQL server
	UserName   string // Username part of the SQL server credentials
	Password   string // Password for the a fore-defined UserName
	Bucket     string // bucket to open
	BucketPass string // bucket to open
}

func (c *CouchbaseConfig) Values(hostName string, port int64, userName, password, bucket, bucketPass string) {

	c.HostName = hostName
	c.Port = port
	c.UserName = userName
	c.Bucket = bucket
	c.Password = password
	c.BucketPass = bucketPass
}

func (c *CouchbaseConfig) URL() string {
	return fmt.Sprintf("http://%s:%d", c.HostName, c.Port)
}

func (c *CouchbaseConfig) String() string {
	return fmt.Sprintf("CouchbaseConfig = %#v", c)
}

func (c *CouchbaseConfig) Connect() (*gocb.Cluster, error) {

	cluster, err := gocb.Connect(c.URL())
	if err != nil {
		fmt.Printf("ERROR CONNECTING TO CLUSTER: %s", err.Error())
	}
	err = cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: c.UserName,
		Password: c.Password,
	})

	if err != nil {
		fmt.Printf("ERROR Authentication BUCKET: %s", err.Error())
	}

	return cluster, err
}
