package db

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"rest/internal/user"
	"rest/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewStorage(database *mongo.Database, collection string, logger logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     &logger,
	}
}

func (d *db) Create(ctx context.Context, user *user.CreateUserDto) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("error inserting user: %w", err)
	}
	oid, ok := result.InsertedID.(bson.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Errorf("failed to convert to hex: %v", result)
	return "", fmt.Errorf("failed to convert to hex: %v", result)
}

func (d *db) FindOne(ctx context.Context, id string) (u *user.User, err error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to object_id. hex: %w, err: %s", err)
	}
	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		// TODO 404
		return u, fmt.Errorf("failed to find user by id. id: %s, err: %s", id, result.Err())
	}
	if err := result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user by id. id: %s, err: %s", id, err)
	}
}

func (d *db) Update(ctx context.Context, user *user.User) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
