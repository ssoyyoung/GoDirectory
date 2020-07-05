package elastic

import (
	"fmt"
	"reflect"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/elastic/go-elasticsearch"
	U "github.com/ssoyyoung.p/GoDirectory/utils"
	M "github.com/ssoyyoung.p/GoDirectory/models"
)

// TODO
// type Decoder struct 
// func NewDecoder(r io.Reader) *Decoder 
// func (dec *Decoder) Decode(v interface{}) error 

func ConnectES() string{

	data, err := os.Open("elastic/elastic_auth.json")
	U.CheckErr(err)

	var auth M.AuthElastic

	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &auth)

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://"+auth.HOST_IP + auth.PORT,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	U.CheckErr(err)

	//fmt.Println(es.Info())
	
	esInfo, err := es.Info()
	U.CheckErr(err)
	fmt.Println(reflect.TypeOf(esInfo))

	// mapResp := make(map[string]interface{})

	// if err := json.NewDecoder(esInfo).Decode(&mapResp); err != nil {
	// 	log.Fatalf("Error parsing the response body: %s", err)
	// }	
	//dec := json.NewDecoder(esInfo)
	//dec.Decode(&res)
	//fmt.Println(reflect.TypeOf(info))
	
	return  "done"
}