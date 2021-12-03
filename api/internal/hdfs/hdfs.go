package hdfs

import (
	"fmt"
	"net"
	"strconv"

	"github.com/colinmarc/hdfs"
)

type HDFS struct {
	client *hdfs.Client
}

func New(host string, port int) (*HDFS, error) {
	address := net.JoinHostPort(host, strconv.Itoa(port))

	client, err := hdfs.New(address)
	if err != nil {
		return nil, fmt.Errorf("new client: %w", err)
	}

	h := &HDFS{
		client: client,
	}

	return h, nil
}
