// author rwahyuhadi
// i hope this small function can make your live easier

package es

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func constructQuery(q string, size int) *strings.Reader {

	// Build a query string from string passed to function
	var query = `{"query": {`

	// Concatenate query string with string passed to method call
	query = query + q

	// Use the strconv.Itoa() method to convert int to string
	query = query + `}, "size": ` + strconv.Itoa(size) + `}`

	// Check for JSON errors
	isValid := json.Valid([]byte(query)) // returns bool

	// Default query is "{}" if JSON is invalid
	if isValid == false {
		log.Println("constructQuery() ERROR: query string not valid:", query)
		log.Println("Using default match_all query")
		query = "{}"
	} else {
		log.Println("constructQuery() valid JSON:", isValid)
	}

	// Build a new string from JSON query
	var b strings.Builder
	b.WriteString(query)

	// Instantiate a *strings.Reader object from string
	read := strings.NewReader(b.String())

	// Return a *strings.Reader object
	return read
}

// Search data function
func GetData(c *elasticsearch.Client, indexName string, query string) (*esapi.Response, error) {

	read := constructQuery(query, 2)
	var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		return nil, errors.New("json.NewEncoder() ERROR:")
	}
	res, err := c.Search(
		c.Search.WithContext(context.Background()),
		c.Search.WithIndex(indexName),
		c.Search.WithBody(read),
		c.Search.WithTrackTotalHits(true),
		c.Search.WithPretty(),
	)

	if err != nil {
		return nil, errors.New("Error when get data ")
	}

	return res, err
}
