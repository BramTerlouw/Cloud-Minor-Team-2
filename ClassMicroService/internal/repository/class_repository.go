package repository

import (
	"Class/graph/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// IClassRepository GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Class.
type IClassRepository interface {
	CreateClass(newClass *model.Class) (*model.Class, error)
	UpdateClass(id string, updatedClass model.Class) (*model.Class, error)
	DeleteClass(id string, existingClass model.Class) error
	GetClassByID(id string) (*model.Class, error)
	ListClasses(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.ClassInfo, error)
}

// ClassRepository GOLANG STRUCT
// Contains a model.Class list and a mongo.Collection.
type ClassRepository struct {
	Classes    []*model.Class
	collection *mongo.Collection
}

// NewClassRepository GOLANG FACTORY
// Returns a ClassRepository implementing IClassRepository.
func NewClassRepository(collection *mongo.Collection) IClassRepository {
	return &ClassRepository{
		Classes:    []*model.Class{},
		collection: collection,
	}
}

func (r *ClassRepository) CreateClass(newClass *model.Class) (*model.Class, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.collection.InsertOne(ctx, newClass)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": newClass.ID}
	var fetchedClass model.Class

	err = r.collection.FindOne(ctx, filter).Decode(&fetchedClass)
	if err != nil {
		return nil, err
	}

	return &fetchedClass, nil
}

func (r *ClassRepository) UpdateClass(id string, updatedClass model.Class) (*model.Class, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedClass}
	var result model.Class

	err := r.collection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)
	if err != nil {
		return nil, err // Return any MongoDB-related errors.
	}

	return &result, nil
}

func (r *ClassRepository) DeleteClass(id string, existingClass model.Class) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{"$set": existingClass}
	var result model.Class

	err := r.collection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

func (r *ClassRepository) GetClassByID(id string) (*model.Class, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	var result model.Class

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ClassRepository) ListClasses(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.ClassInfo, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	var classes []*model.ClassInfo

	cursor, err := r.collection.Find(ctx, bsonFilter, paginateOptions)
	if err != nil {
		return nil, err // Return any MongoDB-related errors.
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	for cursor.Next(ctx) {
		var Class model.ClassInfo
		if err := cursor.Decode(&Class); err != nil {
			return nil, err
		}
		classes = append(classes, &Class)
	}

	return classes, nil
}
