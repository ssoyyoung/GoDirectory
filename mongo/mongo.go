package mongo

import (
	"time"

	m "github.com/ssoyyoung.p/GoDirectory/models"
	crud "github.com/ssoyyoung.p/GoDirectory/mongo/crud"
	U "github.com/ssoyyoung.p/GoDirectory/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Collection name
const colNameLive = "live_list"
const colNameUser = "user_info"
const colNameSchedule = "schedule_list"
const colNameFeedback = "feedback"
const colNameUserHistory = "userHistory"
const colNameUserViewHistory = "userViewHistory"
const colNameSignUp = "userList"

// LiveTrueList func
func LiveTrueList() string {

	language := []string{"ko", "kr"} // Adding language filter

	filter := bson.M{"onLive": true}
	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"language": bson.M{"$in": language}})

	filter["$or"] = orQuery

	sort := bson.M{"liveAttdc": -1}

	return crud.AllData(colNameLive, filter, sort)
}

// LiveTrueListByFollower func
func LiveTrueListByFollower(follower []string) string {

	filter := bson.M{"onLive": true}

	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"_uniq": bson.M{"$in": follower}})

	filter["$or"] = orQuery

	sort := bson.M{"liveAttdc": -1}

	return crud.AllDataReturnJson(colNameLive, filter, sort)
}

// LiveAllListByBlocking func
func LiveAllListByBlocking(blocking []string) string {
	filter := bson.M{}

	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"_uniq": bson.M{"$in": blocking}})

	filter["$or"] = orQuery

	sort := bson.M{}

	return crud.AllData(colNameLive, filter, sort)
}

// LiveAllList func
func LiveAllList() string {

	filter := bson.M{}
	sort := bson.M{}

	return crud.AllDataForAdmin(colNameLive, filter, sort) // change!
}

// AllList func
func AllList() string {

	filter := bson.M{}
	sort := bson.M{}

	return crud.AllData(colNameLive, filter, sort) // change!
}

// AllScheduleList func
func AllScheduleList() string {

	currentTime := time.Now() //get current time
	month := currentTime.Month()
	day := currentTime.Day() - 7

	if day < 1 {
		month = month - 1
		day = 31 + day
	}

	filter := bson.M{
		"month": bson.M{
			"$gte": month,
		},
		"day": bson.M{
			"$gte": day,
		},
	}
	sort := bson.M{}

	return crud.AllScheduleList(colNameSchedule, filter, sort)
}

// GetScheduleListByDay func
func GetScheduleListByDay(month, day int, hours string) string {

	if hours == "" {
		filter := bson.M{"month": month, "day": day}
		sort := bson.M{}
		return crud.AllScheduleList(colNameSchedule, filter, sort)
	}

	hours += ":00"
	filter := bson.M{"month": month, "day": day, "hours": hours}
	sort := bson.M{}
	return crud.AllScheduleList(colNameSchedule, filter, sort)

}

// GetCategoryList func
func GetCategoryList(cate string) string {

	language := []string{"ko", "kr"} // Adding language filter

	filter := bson.M{"onLive": true, "category": cate}
	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"language": bson.M{"$in": language}})
	filter["$or"] = orQuery

	sort := bson.M{"liveAttdc": -1}

	return crud.AllData(colNameLive, filter, sort)
}

// GetPlatformList func
func GetPlatformList(platform string) string {

	filter := bson.M{"platform": platform}
	sort := bson.M{}

	return crud.AllData(colNameLive, filter, sort)
}

// GetCategoryCount func
func GetCategoryCount(cate string) int {

	language := []string{"ko", "kr"} // Adding language filter

	filter := bson.M{"onLive": true, "category": cate}
	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"language": bson.M{"$in": language}})
	filter["$or"] = orQuery

	sort := bson.M{"liveAttdc": -1}

	return crud.AllDataCount(colNameLive, filter, sort)
}

// SearchDBbyID func
func SearchDBbyID(id string) string {

	filter := bson.M{"_id": U.ConvertID(id)}
	sort := bson.M{}

	return crud.AllDataForAdmin(colNameLive, filter, sort) // change!
}

// DeleteDBbyID func
func DeleteDBbyID(id string) string {

	filter := bson.M{"_id": U.ConvertID(id)}

	return crud.Delete(colNameLive, filter) // change!
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

	return crud.Update(colNameLive, filter, update) // change!
}

// CreateDB func
func CreateDB(platform, channel, channelID, category string) string {

	filter := bson.M{"_uniq": platform + channelID}
	num := crud.Count(colNameLive, filter)

	if num == 0 {
		newData := m.LiveList{
			Platform:  platform,
			Uniq:      platform + channelID,
			Channel:   channel,
			ChannelID: channelID,
			Category:  category,
		}

		return crud.CreateCrawl(colNameLive, newData)
	}

	return "exist!"
}

// CheckPW func
func CheckPW(userID, passWD string) (bool, string) {
	filter := bson.M{"userID": userID}
	sort := bson.M{}

	dbPassWD, nickname := crud.GetPassWD(colNameSignUp, filter, sort)

	if dbPassWD == passWD {
		return true, nickname
	}
	return false, nickname
}

// CheckDB func
func CheckDB(platform, channelID string) string {
	filter := bson.M{"_uniq": platform + channelID}
	num := crud.Count(colNameLive, filter)

	if num == 0 {
		return "true"
	}

	return "false"
}

//CheckUser func
func CheckUser(googleID, name, email string) bool {
	following := []string{}
	block := []string{}

	filter := bson.M{"googleId": googleID, "name": name}
	num := crud.Count(colNameUser, filter)
	newData := m.UserInfo{
		GoogleID:  googleID,
		Name:      name,
		Email:     email,
		Following: following,
		Block:     block,
	}

	if num == 0 {
		crud.CreateUser(colNameUser, newData)
	}
	return true
}

// CheckID func
func CheckID(userID string) (string, string) {

	filter := bson.M{"userID": userID}
	sort := bson.M{}

	num := crud.Count(colNameSignUp, filter)

	if num == 0 {
		return "false", ""
	}

	res := crud.GetSerialNo(colNameSignUp, filter, sort)

	return "true", res

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

// UpdateUserInfo func
func UpdateUserInfo(userID, token string) string {

	filter := bson.M{"userID": userID}
	update := bson.M{
		"$set": bson.M{
			"token": token,
		},
	}

	return crud.Update(colNameSignUp, filter, update)
}

// UpdateUserToken func
func UpdateUserToken(userID, token string) string {

	filter := bson.M{"userID": userID}
	update := bson.M{
		"$set": bson.M{
			"token": token,
		},
	}

	return crud.Update(colNameSignUp, filter, update)
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

// GetFollowing func
func GetFollowing(email string) []string {

	filter := bson.M{"email": email}

	return crud.GetFollowing(colNameUser, filter)
}

// PushBlocking func
func PushBlocking(following, email string) string {

	filter := bson.M{"email": email}
	update := bson.M{
		"$push": bson.M{
			"block": following,
		},
	}

	return crud.Update(colNameUser, filter, update)
}

// PullBlocking func
func PullBlocking(following, email string) string {

	filter := bson.M{"email": email}
	update := bson.M{
		"$pull": bson.M{
			"block": following,
		},
	}

	return crud.Update(colNameUser, filter, update)
}

// GetBlocking func
func GetBlocking(email string) []string {

	filter := bson.M{"email": email}

	return crud.GetBlocking(colNameUser, filter)
}

// SearchBar func
func SearchBar(query string) string {
	filter := bson.M{}
	orQuery := []bson.M{}
	orQuery = append(orQuery, bson.M{"category": bson.M{"$regex": query}})
	orQuery = append(orQuery, bson.M{"creatorDataName": bson.M{"$regex": query}})
	orQuery = append(orQuery, bson.M{"liveDataTitle": bson.M{"$regex": query}})

	filter["$or"] = orQuery

	sort := bson.M{"liveAttdc": -1}

	return crud.AllDataReturnJson(colNameLive, filter, sort)
}

// InsertFeedback func
func InsertFeedback(cookie, title, email, message string) string {
	newData := m.Feedback{
		Cookie:  cookie,
		Email:   email,
		Title:   title,
		Message: message,
	}

	return crud.InsertFeedback(colNameFeedback, newData)
}

// InsertUserHistory func
func InsertUserHistory(cookie, username, pathname string, residencetime int) string {
	newData := m.UserHistory{
		Cookie:        cookie,
		Username:      username,
		Pathname:      pathname,
		Residencetime: residencetime,
	}

	return crud.InsertUserHistory(colNameUserHistory, newData)
}

// InsertViewHistory func
func InsertViewHistory(cookie, username, streaming, platform, _uniq string) string {
	newData := m.UserViewHistory{
		Cookie:    cookie,
		Username:  username,
		Streaming: streaming,
		Platform:  platform,
		Uniq:      _uniq,
	}

	return crud.InsertViewHistory(colNameUserViewHistory, newData)
}

// SignUp func
func SignUp(id, serialNo, password, nickname, birthday, t string, tags, ctags []string) string {
	following := []string{}
	block := []string{}

	signUpData := m.SignUp{
		UserID:    id,
		SerialNo:  serialNo,
		Password:  password,
		Nickname:  nickname,
		Birthday:  birthday,
		Ctags:     ctags,
		Tags:      tags,
		Token:     t,
		Following: following,
		Block:     block,
	}

	return crud.CreateSignUpUser(colNameSignUp, signUpData)
}
