package crud

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

// GetCollection func
func GetCollection(client *mongo.Client, colName string) *mongo.Collection {
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

// ConnectDB to MongoDB
func ConnectDB() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
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

// AllData func
func AllData(collection string, filter bson.M, sort bson.M) string {
	// define bson.M type data
	var datas []bson.M

	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(sort)

	res, err := GetCollection(client, collection).Find(ctx, filter, findOptions)
	U.CheckErr(err)

	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	return U.JSONMarshalString(datas)
}

// AllDataReturnJson func
func AllDataReturnJson(collection string, filter bson.M, sort bson.M) string {
	// define bson.M type data
	var datas []m.LiveList

	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(sort)

	res, err := GetCollection(client, collection).Find(ctx, filter, findOptions)
	U.CheckErr(err)

	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	jsonBytes, err := json.Marshal(datas)
	jsonString := string(jsonBytes)

	return jsonString

	//return U.ConvertData(datas)
}

// AllDataForAdmin func
func AllDataForAdmin(collection string, filter bson.M, sort bson.M) string {

	var admin []m.LiveListForAdmin

	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(sort)

	res, err := GetCollection(client, collection).Find(ctx, filter, findOptions)
	U.CheckErr(err)

	if err = res.All(ctx, &admin); err != nil {
		fmt.Println(err)
	}

	jsonBytes, err := json.Marshal(admin)
	jsonString := string(jsonBytes)

	return jsonString
}

// AllScheduleList func
func AllScheduleList(collection string, filter bson.M, sort bson.M) string {

	var schedule []m.ScheduleList

	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(sort)

	res, err := GetCollection(client, collection).Find(ctx, filter, findOptions)
	U.CheckErr(err)

	if err = res.All(ctx, &schedule); err != nil {
		fmt.Println(err)
	}

	jsonBytes, err := json.Marshal(schedule)
	jsonString := string(jsonBytes)

	return jsonString
}

// GetFollowing func
func GetFollowing(collection string, filter bson.M) []string {
	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	var result []m.UserInfo

	res, err := GetCollection(client, collection).Find(ctx, filter)
	U.CheckErr(err)

	if err = res.All(ctx, &result); err != nil {
		fmt.Println(err)
	}
	if len(result) == 0 {
		return nil
	}
	return result[0].Following
}

// GetBlocking func
func GetBlocking(collection string, filter bson.M) []string {
	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	var result []m.UserInfo

	res, err := GetCollection(client, collection).Find(ctx, filter)
	U.CheckErr(err)

	if err = res.All(ctx, &result); err != nil {
		fmt.Println(err)
	}

	return result[0].Block
}

// Delete func
func Delete(collection string, filter bson.M) string {
	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	_, err := GetCollection(client, collection).DeleteOne(ctx, filter)
	U.CheckErr(err)

	return "Delete!"
}

// Update func
func Update(collection string, filter bson.M, update bson.M) string {
	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	_, err := GetCollection(client, collection).UpdateOne(ctx, filter, update)
	U.CheckErr(err)

	return "Update!"
}

// CreateCrawl func
func CreateCrawl(collection string, newData m.LiveList) string {
	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	_, err := GetCollection(client, collection).InsertOne(ctx, newData)
	U.CheckErr(err)

	return "create!"
}

// CreateUser func
func CreateUser(collection string, newData m.UserInfo) string {
	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	_, err := GetCollection(client, collection).InsertOne(ctx, newData)
	U.CheckErr(err)

	return "create!"
}

// Count func
func Count(collection string, filter bson.M) int64 {
	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	res, err := GetCollection(client, collection).CountDocuments(ctx, filter)
	U.CheckErr(err)

	return res
}

// InsertFeedback func
func InsertFeedback(collection string, newData m.Feedback) string {
	client, ctx, cancel := ConnectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	_, err := GetCollection(client, collection).InsertOne(ctx, newData)
	U.CheckErr(err)

	return "Done!"
}
