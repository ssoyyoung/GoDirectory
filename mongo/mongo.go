package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	m "github.com/ssoyyoung.p/GoDirectory/models"
	U "github.com/ssoyyoung.p/GoDirectory/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database Name
const dbName = "meerkatonair"

// Collection name
const colNameLive = "live_list"
const colNameCrawl = "crawl_target"
const colNameUser = "user_info"

// get Collection
func getCollection(client *mongo.Client, colName string) *mongo.Collection {
	return client.Database(dbName).Collection(colName)
}

// get MongoDB Authorization info7
func getAuth() m.Auth {
	data, err := os.Open("mongo/mongodb_auth.json")
	U.CheckErr(err)

	var auth m.Auth
	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &auth)

	return auth
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
	U.CheckErr(err)

	// MongoDB 연결 검증
	U.CheckErr(client.Ping(ctx, readpref.Primary()))

	return client, ctx, cancel
}

// define bson.M type data
var datas []bson.M

// LiveTrueList func
func LiveTrueList() string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"liveAttdc", -1}})

	res, err := getCollection(client, colNameLive).Find(ctx, bson.M{"onLive": true}, findOptions)
	U.CheckErr(err)

	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	return U.JSONMarshalString(datas)
}

// LiveAllList func
func LiveAllList() string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	res, err := getCollection(client, colNameCrawl).Find(ctx, bson.M{})
	U.CheckErr(err)

	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	return U.JSONMarshalString(datas)
}

// SearchDBbyID func
func SearchDBbyID(id string) string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	res, _ := getCollection(client, colNameCrawl).Find(ctx, bson.M{"_id": U.ConvertID(id)})
	if err := res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	return U.JSONMarshalString(datas)
}

// DeleteDBbyID func
func DeleteDBbyID(id string) string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	_, err := getCollection(client, colNameCrawl).DeleteOne(ctx, bson.M{"_id": U.ConvertID(id)})
	U.CheckErr(err)

	return "Delete!"
}

//UpdateDBbyID func
func UpdateDBbyID(id, platform, channel, channelID string) string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"_id": U.ConvertID(id)}
	update := bson.D{
		{"$set", bson.D{
			{"platform", platform},
			{"channel", channel},
			{"channelID", channelID},
		},
		},
	}
	_, err := getCollection(client, colNameCrawl).UpdateOne(ctx, filter, update)
	U.CheckErr(err)

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
	_, err := getCollection(client, colNameCrawl).InsertOne(ctx, newData)
	U.CheckErr(err)

	return "create!"
}

//CheckUser func
func CheckUser(googleID, name, email string) bool {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	res, err := getCollection(client, colNameUser).CountDocuments(ctx, bson.M{"googleId": googleID, "name": name})
	U.CheckErr(err)
	if res == 0 {
		createUser(googleID, name, email)
	}
	return true
}

// createUser func
func createUser(googleID, name, email string) {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	_, err := getCollection(client, colNameUser).InsertOne(ctx, bson.M{
		"googleId": googleID,
		"name":     name,
		"email":    email,
	})
	U.CheckErr(err)
}

// UpdateUser func
func UpdateUser(googleID, token string) {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"googleId": googleID}
	update := bson.D{
		{"$set", bson.D{
			{"token", token},
		},
		},
	}
	res, err := getCollection(client, colNameUser).UpdateOne(ctx, filter, update)
	fmt.Println(res)
	U.CheckErr(err)
}
