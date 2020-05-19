package models

import (
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// CrawlTarget struct
type CrawlTarget struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Platform  string             `bson:"platform,omitempty"`
	Channel   string             `bson:"channel,omitempty"`
	ChannelID string             `bson:"channelID,omitempty"`
}

// TODO : LiveList struct