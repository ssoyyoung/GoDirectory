package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch"
	M "github.com/ssoyyoung.p/GoDirectory/models"
	U "github.com/ssoyyoung.p/GoDirectory/utils"
)

// TODO
// type Decoder struct
// func NewDecoder(r io.Reader) *Decoder
// func (dec *Decoder) Decode(v interface{}) error

// ConnectES func
func ConnectES() *elasticsearch.Client {

	data, err := os.Open("elastic/elastic_auth.json")
	U.CheckErr(err)

	var auth M.AuthElastic

	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &auth)

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://" + auth.HOST_IP + auth.PORT,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	U.CheckErr(err)
	log.Println(es.Info())

	return es
}

func main() {
	ConnectES()
}

/* // GetAvgAttc func
func GetAvgAttc() {
	// connect elasticsearch
	es := ConnectES
	var (
		buf bytes.Buffer
		r   map[string]interface{}
	)

	// define query
	query := map[string]interface{}{
		"size": 0,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"_uniq.keyword": "twitchwoowakgood"}},
					map[string]interface{}{
						"match": map[string]interface{}{
							"onLive": true}},
				},
			},
		},
		"aggs": map[string]interface{}{
			"date_his": map[string]interface{}{
				"aggs": map[string]interface{}{
					"average_attc": map[string]interface{}{
						"avg": map[string]interface{}{
							"field": "liveAttdc",
						},
					},
				},
				"date_histogram": map[string]interface{}{
					"field":    "updateDate",
					"interval": "day",
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	// TODO fix
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("tracking_streamer"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
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

	log.Println(strings.Repeat("=", 37))
}
*/
