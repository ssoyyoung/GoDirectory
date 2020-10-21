package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ScheduleList strict
type ScheduleList struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Uniq        string             `json:"_uniq" bson:"_uniq,omitempty"`
	Channel     string             `json:"channel" bson:"channel,omitempty"`
	ChannelID   string             `json:"channelID" bson:"channelID,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Image       string             `json:"image" bson:"image,omitempty"`
	URL         string             `json:"url" bson:"url,omitempty"`
	Regular     bool               `json:"regular" bson:"regular,omitempty"`
	Category    string             `json:"category" bson:"category,omitempty"`
	Platform    string             `json:"platform" bson:"platform,omitempty"`
	Days        []string           `json:"days" bson:"days,omitempty"`
	Year        int                `json:"year" bson:"year,omitempty"`
	Month       int                `json:"month" bson:"month,omitempty"`
	Day         int                `json:"day" bson:"day,omitempty"`
	Hours       string             `json:"hours" bson:"hours,omitempty"`
}

// CrawlTarget struct
type CrawlTarget struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Platform  string             `bson:"platform,omitempty"`
	Channel   string             `bson:"channel,omitempty"`
	ChannelID string             `bson:"channelID,omitempty"`
}

// LiveList struct
type LiveList struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Uniq            string             `json:"_uniq" bson:"_uniq,omitempty"`
	Channel         string             `json:"channel" bson:"channel,omitempty"`
	ChannelID       string             `json:"channelID" bson:"channelID,omitempty"`
	Platform        string             `json:"platform" bson:"platform,omitempty"`
	OnLive          bool               `json:"onLive" bson:"onLive,omitempty"`
	UpdateDate      string             `json:"updateDate" bson:"updateDate,omitempty"`
	CreatorDataHref string             `json:"creatorDataHref" bson:"creatorDataHref,omitempty"`
	CreatorDataName string             `json:"creatorDataName" bson:"creatorDataName,omitempty"`
	CreatorDataLogo string             `json:"creatorDataLogo" bson:"creatorDataLogo"`
	ImgDataSrc      string             `json:"imgDataSrc" bson:"imgDataSrc,omitempty"`
	LiveAttdc       int32              `json:"liveAttdc" bson:"liveAttdc,omitempty"`
	LiveDataHref    string             `json:"liveDataHref" bson:"liveDataHref,omitempty"`
	LiveDataTitle   string             `json:"liveDataTitle" bson:"liveDataTitle,omitempty"`
	Category        string             `json:"category" bson:"category,omitempty"`
	Detail          string             `json:"detail" bson:"detail,omitempty"`
}

// LiveListForAdmin struct
type LiveListForAdmin struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Channel   string             `json:"channel" bson:"channel,omitempty"`
	ChannelID string             `json:"channelID" bson:"channelID,omitempty"`
	Category  string             `json:"category" bson:"category,omitempty"`
	Platform  string             `json:"platform" bson:"platform,omitempty"`
}

//UserInfo struct
type UserInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	GoogleID  string             `bson:"googleId,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Token     string             `bson:"token,omitempty"`
	Following []string           `bson:"following"`
	Block     []string           `bson:"block"`
}

//Auth struct
type Auth struct {
	Username string
	Password string
	Hostname string
	Port     string
}

// AuthElastic struct
type AuthElastic struct {
	HOST_IP string
	PORT    string
}

// Feedback struct
type Feedback struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Cookie  string             `bson:"cookie,omitempty"`
	Email   string             `bson:"email,omitempty"`
	Title   string             `bson:"title,omitempty"`
	Message string             `bson:"message,omitempty"`
}

// UserHistory struct
type UserHistory struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	LoginType     string             `bson:"loginType,omitempty"`
	Username      string             `bson:"username,omitempty"`
	Pathname      string             `bson:"pathname,omitempty"`
	Residencetime float64            `bson:"residencetime,omitempty"`
	UpdateDate    string             `bson:"updateDate,omitempty"`
}

// UserViewHistory struct
type UserViewHistory struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	LoginType  string             `bson:"loginType,omitempty"`
	Username   string             `bson:"username,omitempty"`
	Streaming  string             `bson:"streaming,omitempty"`
	Platform   string             `bson:"platform,omitempty"`
	Uniq       string             `bson:"_uniq,omitempty"`
	UpdateDate string             `bson:"updateDate,omitempty"`
}

// SignUp struct
type SignUp struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userID,omitempty"`
	SerialNo  string             `bson:"serialNo,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Nickname  string             `bson:"nickname,omitempty"`
	Birthday  string             `bson:"birthday,omitempty"`
	Gender    string             `bson:"gender,omitempty"`
	Ctags     []string           `bson:"ctags,omitempty"`
	Tags      []string           `bson:"tags,omitempty"`
	Token     string             `bson:"token,omitempty"`
	Following []string           `bson:"following"`
	Block     []string           `bson:"block"`
}

// Survey struct
type Survey struct {
	Platform   []string `bson:"platform,omitempty"`
	Category   []string `bson:"category,omitempty"`
	Streamers  []string `bson:"streamers,omitempty"`
	Userinfo   string   `bson:"userinfo,omitempty"`
	UpdateDate string   `bson:"updateDate,omitempty"`
}
