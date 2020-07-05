package elastic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/elastic/go-elasticsearch"
	"github.com/ssoyyoung.p/GoDirectory/utils"
	"github.com/ssoyyoung.p/GoDirectory/models"
)

// TODO
// type Decoder struct 
// func NewDecoder(r io.Reader) *Decoder 
// func (dec *Decoder) Decode(v interface{}) error 

func ConnectES() string{

	data, err := os.Open("elastic/elastic_auth.json")
	utils.CheckErr(err)

	var auth models.AuthElastic

	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &auth)

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://"+auth.HOST_IP + auth.PORT,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	utils.CheckErr(err)

	fmt.Println(es.Info())
	
	//esInfo := es.Info()
	// var u ES
	// dec := json.NewDecoder(esInfo)
	// dec.Decode(&ES)
	// fmt.Println(reflect.TypeOf(info))
	
	return  "done"
}