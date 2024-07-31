package mongodb

import (
	"context"
	models "go-crud/modules/user/model/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	UserRepository interface {
		GetUsers() ([]models.User, error)
		CreateUser(user models.User) error
		GetUserByID(id primitive.ObjectID) (models.User, error)
		UpdateUser(id primitive.ObjectID, update models.User) error
		DeleteUser(id primitive.ObjectID) error
	}

	userRepository struct {
		db *mongo.Client
	}
)

func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) init() *mongo.Collection {
	userCollection := r.db.Database("User").Collection("Users")

	return userCollection
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := r.init()
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, err
}

func (r *userRepository) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	userCollection := r.init()
	_, err := userCollection.InsertOne(ctx, user)

	return err
}

func (r *userRepository) GetUserByID(id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	userCollection := r.init()
	err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	return user, err
}

func (r *userRepository) UpdateUser(id primitive.ObjectID, update models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := r.init()
	update.UpdatedAt = time.Now()
	_, err := userCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})

	return err
}

func (r *userRepository) DeleteUser(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := r.init()
	_, err := userCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err

}
