package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	m "github.com/ssoyyoung.p/GoDirectory/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database Name
const dbName = "meerkatonair"

// Collection name
const colNameLive = "live_list"
const colNameCrawl = "crawl_target"

// get MongoDB Authorization info
func getAuth() m.Auth {
	data, err := os.Open("mongo/mongodb_auth.json")
	checkErr(err)

	var auth m.Auth
	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &auth)

	return auth
}

// checkErr function
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// connect to MongoDB
func connectDB() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	// Timeout 설정을 위한 Context생성
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)

	Authrization := getAuth()

	// Auth에러 처리를 위한 client option 구성
	clientOptions := options.Client().ApplyURI("mongodb://" + Authrization.Hostname + Authrization.Port).SetAuth(options.Credential{
		Username: Authrization.Username,
		Password: Authrization.Password,
	})

	// MongoDB 연결
	client, err := mongo.Connect(ctx, clientOptions)
	checkErr(err)

	// MongoDB 연결 검증
	checkErr(client.Ping(ctx, readpref.Primary()))

	return client, ctx, cancel
}

// get Crawl_target collection
func getCollectionCrawl(client *mongo.Client) *mongo.Collection {
	return client.Database(dbName).Collection(colNameCrawl)
}

// get Live_live collection
func getCollectionLive(client *mongo.Client) *mongo.Collection {
	return client.Database(dbName).Collection(colNameLive)
}

// define bson.M type data
var datas []bson.M

// jsonData marshal to string func
func jsonMarshalString(datas []bson.M) string {
	jsonBytes, err := json.Marshal(datas)
	checkErr(err)
	jsonString := string(jsonBytes)

	return jsonString
}

// string ID convert to OjectID
func convertID(id string) primitive.ObjectID {
	docID, err := primitive.ObjectIDFromHex(id)
	checkErr(err)

	return docID
}

// CrawlList func
func CrawlList() string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	res, err := getCollectionCrawl(client).Find(ctx, bson.M{})
	checkErr(err)

	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	return jsonMarshalString(datas)
}

// SearchDBbyID func
func SearchDBbyID(id string) string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	res, _ := getCollectionCrawl(client).Find(ctx, bson.M{"_id": convertID(id)})
	if err := res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	return jsonMarshalString(datas)
}

// DeleteDBbyID func
func DeleteDBbyID(id string) string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	_, err := getCollectionCrawl(client).DeleteOne(ctx, bson.M{"_id": convertID(id)})
	checkErr(err)

	return "Delete!"
}

//UpdateDBbyID func
func UpdateDBbyID(id, platform, channel, channelID string) string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"_id": convertID(id)}
	update := bson.D{
		{"$set", bson.D{
			{"platform", platform},
			{"channel", channel},
			{"channelID", channelID},
		},
		},
	}
	_, err := getCollectionCrawl(client).UpdateOne(ctx, filter, update)
	checkErr(err)

	return "Update!"
}

// CreateDB func
func CreateDB(platform, channel, channelID string) string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	newData := m.CrawlTarget{
		Platform:  platform,
		Channel:   channel,
		ChannelID: channelID,
	}
	_, err := getCollectionCrawl(client).InsertOne(ctx, newData)
	checkErr(err)

	return "create!"
}
