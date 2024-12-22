package mongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, pass, database, authDB string) (*mongo.Database, error) {
	var mongoDBUrl string
	var isAuth bool
	if username == "" && pass == "" {
		mongoDBUrl = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongoDBUrl = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, pass, host, port)
	}

	clientOptions := options.Client().ApplyURI(mongoDBUrl)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   pass,
		})
	}

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, errors.New("Failed to connect to MongoDB: " + err.Error())
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.New("Failed to ping MongoDB: " + err.Error())
	}

	return client.Database(database), nil
}
