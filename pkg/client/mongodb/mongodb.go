package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewClient(ctx context.Context, host, port, username, pass, database, authDB string) (*mongo.Client, error) {
	var mongoDBUrl string
	if username == "" && pass == "" {
		mongoDBUrl = "mongodb://%s:%s"
	} else {
		mongoDBUrl = "mongodb://%s:%s@%s:%s"
	}

}
