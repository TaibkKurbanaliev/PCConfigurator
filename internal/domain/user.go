package domain

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type User struct {
	ID           bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string        `json:"name" bson:"name"`
	Email        string        `json:"email,omitempty" bson:"email,omitempty"`
	Phone        string        `json:"phone,omitempty" bson:"phone,omitempty"`
	Password     string        `json:"password" bson:"password,omitempty"`
	RegisteredAt time.Time     `json:"registeredAt,omitempty" bson:"registeredAt,omitempty"`
	LastVisitAt  time.Time     `json:"lastVisitAt,omitempty" bson:"lastVisitAt,omitempty"`
}
