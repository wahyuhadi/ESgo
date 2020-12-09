### elasticsearch golang client https://github.com/elastic/go-elasticsearch 

### how to use 

```go
import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/wahyuhadi/ESgo/es"
)

var (
	clusterURLs = []string{"http://127.0.0.1:9200"}
	username    = "foo"
	password    = "bar"
)

// Model data
type MyData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func main() {

	cfg := elasticsearch.Config{
		Addresses: clusterURLs,
		// Username:  username, // if ES need this
		// Password:  password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
				// ...
			},
		},
	}

	c, _ := elasticsearch.NewClient(cfg)
	// PushexamplePushData(c)
	exampleGetData(c)

}
```

* how to push data 

```go
func examplePushData(c *elasticsearch.Client) {
	// Example push data to elastic server
	doc := MyData{
		Name:    "bb",
		Email:   "hrahmatwsahyu@gmail.com",
		Address: "DKI Jakartas",
	}

	// parsing with esutil from elastic
	data := esutil.NewJSONReader(&doc)
	// Push data to elastic
	response, err := es.PushData(c, "test-data", data)
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
}
```

* example search data

```go
func exampleGetData(c *elasticsearch.Client) {
	var r map[string]interface{}

	//Create a new query string for the Elasticsearch method call
	var query = `
		"match_all":{
		}
	`
	res, err := es.GetData(c, "test-data", query)
	if err != nil {
		log.Println(err)
	}

	// Close the result body when the function call is complete
	defer res.Body.Close()

	// Decode the JSON response and using a pointer
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

}

```

* for details please see example.go


* author rwahyuhadi
* i hope this small function can make your live easier

&copy; [Rahmat Wahyu Hadi](https://github.com/wahyuhadi/) - end of 2020 
