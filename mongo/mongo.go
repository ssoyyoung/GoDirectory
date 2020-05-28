package mongo

import (
	m "github.com/ssoyyoung.p/GoDirectory/models"
	crud "github.com/ssoyyoung.p/GoDirectory/mongo/crud"
	U "github.com/ssoyyoung.p/GoDirectory/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Collection name
const colNameLive = "live_list"
const colNameCrawl = "crawl_target"
const colNameUser = "user_info"

// LiveTrueList func
func LiveTrueList() string {

	filter := bson.M{"onLive": true}
	sort := bson.M{"liveAttdc": -1}

	return crud.AllData(colNameLive, filter, sort)
}

// LiveTrueListByFollower  func
func LiveTrueListByFollower(follower []string) string {

	filter := bson.M{"onLive": true}

	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"channel": bson.M{"$in": follower}})

	filter["$or"] = orQuery

	sort := bson.M{"liveAttdc": -1}

	return crud.AllData(colNameLive, filter, sort)
}

// LiveAllList func
func LiveAllList() string {

	filter := bson.M{}
	sort := bson.M{}

	return crud.AllData(colNameCrawl, filter, sort)
}

// SearchDBbyID func
func SearchDBbyID(id string) string {

	filter := bson.M{"_id": U.ConvertID(id)}
	sort := bson.M{}

	return crud.AllData(colNameCrawl, filter, sort)
}

// DeleteDBbyID func
func DeleteDBbyID(id string) string {

	filter := bson.M{"_id": U.ConvertID(id)}

	return crud.Delete(colNameCrawl, filter)
}

//UpdateDBbyID func
func UpdateDBbyID(id, platform, channel, channelID string) string {

	filter := bson.M{"_id": U.ConvertID(id)}
	update := bson.M{
		"$set": bson.M{
			"platform":  platform,
			"channel":   channel,
			"channelID": channelID,
		},
	}

	return crud.Update(colNameCrawl, filter, update)
}

// CreateDB func
func CreateDB(platform, channel, channelID string) string {

	newData := m.CrawlTarget{
		Platform:  platform,
		Channel:   channel,
		ChannelID: channelID,
	}

	return crud.CreateCrawl(colNameCrawl, newData)
}

//CheckUser func
func CheckUser(googleID, name, email string) bool {

	filter := bson.M{"googleId": googleID, "name": name}
	num := crud.Count(colNameUser, filter)
	newData := m.UserInfo{
		GoogleID: googleID,
		Name:     name,
		Email:    email,
	}

	if num == 0 {
		crud.CreateUser(colNameCrawl, newData)
	}
	return true
}

// UpdateUser func
func UpdateUser(googleID, token string) string {

	filter := bson.M{"googleId": googleID}
	update := bson.M{
		"$set": bson.M{
			"token": token,
		},
	}

	return crud.Update(colNameUser, filter, update)
}

// PushFollowing func
func PushFollowing(following, email string) string {

	filter := bson.M{"email": email}
	update := bson.M{
		"$push": bson.M{
			"following": following,
		},
	}

	return crud.Update(colNameUser, filter, update)
}

// PullFollowing func
func PullFollowing(following, email string) string {

	filter := bson.M{"email": email}
	update := bson.M{
		"$pull": bson.M{
			"following": following,
		},
	}

	return crud.Update(colNameUser, filter, update)
}

// SearchDBbyEmail func
func SearchDBbyEmail(email string) []string {

	filter := bson.M{"email": email}

	return crud.GetFollowing(colNameUser, filter)
}
