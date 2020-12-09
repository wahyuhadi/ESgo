// author rwahyuhadi
// i hope this small function can make your live easier

package es

import (
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// PushData Function
func PushData(c *elasticsearch.Client, indexName string, data io.Reader) (*esapi.Response, error) {
	res, err := c.Index(
		indexName,
		data,
	)

	// Error handling when input data to elastic serach
	if err != nil {
		return nil, errors.New("Error when add data to elastic")
	}

	return res, err
}
