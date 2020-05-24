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
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Uniq            string             `bson:"_uniq,omitempty"`
	Channel         string             `bson:"channel,omitempty"`
	ChannelID       string             `bson:"channelID,omitempty"`
	Platform        string             `bson:"platform,omitempty"`
	OnLive          bool               `bson:"onLive,omitempty"`
	UpdateDate      string             `bson:"updateDate,omitempty"`
	CreatorDataHref string             `bson:"creatorDataHref,omitempty"`
	creatorDataName string             `bson:"creatorDataName,omitempty"`
	ImgDataSrc      string             `bson:"imgDataSrc,omitempty"`
	LiveAttdc       int32              `bson:"liveAttdc,omitempty"`
	LiveDataHref    string             `bson:"liveDataHref,omitempty"`
	LiveDataTitle   string             `bson:"liveDataTitle,omitempty"`
	Category        string             `bson:"category,omitempty"`
	Detail          string             `bson:"detail,omitempty"`
}

//UserInfo struct
type UserInfo struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	GoogleID string             `bson:"googleId,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Token    string             `bson:"token,omitempty"`
}

//Auth struct
type Auth struct {
	Username string
	Password string
	Hostname string
	Port     string
}
