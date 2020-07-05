package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	Following []string           `bson:"following,omitempty"`
	Block     []string           `bson:"block,omitempty"`
}

//Auth struct
type Auth struct {
	Username string
	Password string
	Hostname string
	Port     string
}
