package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User      interface{}        `json:"user" bson:"user"`
	Path      string             `json:"path" bson:"path"`
	Method    string             `json:"method" bson:"method"`
	Status    int                `json:"status" bson:"status"`
	Query     string             `json:"query" bson:"query"`
	Body      interface{}        `json:"body" bson:"body"`
	Ip        string             `json:"ip" bson:"ip"`
	UserAgent string             `json:"userAgent" bson:"userAgent"`
	Errors    string             `json:"errors" bson:"errors"`
	Cost      string             `json:"cost" bson:"cost"`
}
