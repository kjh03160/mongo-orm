package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongo-orm/errorType"
)

type Collection struct {
	*mongo.Collection
}

//func (col *Collection) evaluateSingleResult(result *mongo.SingleResult) error {
//	if result == nil {
//		return errorType.InternalError(col.Name(), filter, errorType.SingleResultErr)
//	}
//	return nil
//}

func (col *Collection) FindAll(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	cursor, err := col.Collection.Find(context.Background(), filter, opts...)
	if err != nil {
		return nil, errorType.InternalError(col.Name(), filter, err)
	}
	return cursor, nil
}

func (col *Collection) FindOne(filter interface{}, opts ...*options.FindOneOptions) (*mongo.SingleResult, error) {
	singleResult := col.Collection.FindOne(context.Background(), filter, opts...)
	if singleResult == nil {
		return nil, errorType.InternalError(col.Name(), filter, errorType.SingleResultErr)
	}
	return singleResult, nil
}

func (col *Collection) FindOneAndModify(filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) (*mongo.SingleResult, error) {
	singleResult := col.Collection.FindOneAndUpdate(context.Background(), filter, update, opts...)
	if singleResult == nil {
		return nil, errorType.InternalError(col.Name(), filter, errorType.SingleResultErr)
	}
	return singleResult, nil
}

func (col *Collection) FindOneAndReplace(filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) (*mongo.SingleResult, error) {
	singleResult := col.Collection.FindOneAndReplace(context.Background(), filter, replacement, opts...)
	if singleResult == nil {
		return nil, errorType.InternalError(col.Name(), filter, errorType.SingleResultErr)
	}
	return singleResult, nil
}

func (col *Collection) FindOneAndDelete(filter interface{}, opts ...*options.FindOneAndDeleteOptions) (*mongo.SingleResult, error) {
	singleResult := col.Collection.FindOneAndDelete(context.Background(), filter, opts...)
	if singleResult == nil {
		return nil, errorType.InternalError(col.Name(), filter, errorType.SingleResultErr)
	}
	return singleResult, nil
}

func (col *Collection) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	updateResult, err := col.Collection.UpdateOne(context.Background(), filter, update, opts...)
	if err != nil {
		return nil, errorType.InternalError(col.Collection.Name(), filter, err, update)
	}
	return updateResult, nil
}

func (col *Collection) UpdateMany(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	updateResult, err := col.Collection.UpdateMany(context.Background(), filter, update, opts...)
	if err != nil {
		return nil, errorType.InternalError(col.Collection.Name(), filter, err, update)
	}
	return updateResult, nil
}

func (col *Collection) ReplaceOne(filter interface{}, document interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, err) {
	result, err := col.Collection.ReplaceOne(context.Background(), filter, document, opts...)
	if err != nil {
		return nil, errorType.InternalError(col.Collection.Name(), filter, err, document)
	}
	return result, nil

}
