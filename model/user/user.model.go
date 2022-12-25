package user_model

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/////////////////////////
//		Entity
/////////////////////////

type UserEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name"` //사용자 이름 
	Phone string `bson:"phone"` //사용자 폰번호
	Addr string `bson:"addr"` //주소
	CreateDate time.Time `bson:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `bson:"updateDate"` //데이터 수정 시각
}

type UserCollection struct {
	UserCollection *mongo.Collection
	Ctx context.Context
}


/////////////////////////
//		Init
/////////////////////////

//초기화 함수
func InitWithSelf(col *mongo.Collection, ctx context.Context) UserCollection {
	return UserCollection{col, ctx}
}

/////////////////////////
//		Create 
/////////////////////////

//Add User Entity
func (c *UserCollection) AddEntity(entity UserEntity) (*mongo.InsertOneResult, error) {
	result, inErr := c.UserCollection.InsertOne(c.Ctx, entity)

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"phone": 1}, Options: opt}
	if _, err1 := c.UserCollection.Indexes().CreateOne(c.Ctx, index); err1 != nil {
		return nil, errors.New("could not create index for user Id")
	}
	
	return result, inErr
}

/////////////////////////
//		Read
/////////////////////////

//Id 값으로 조회
func (c *UserCollection) FindByObjectId(objectId interface{}) (*UserEntity, error) {
	var menuItem *UserEntity
	query := bson.M{"_id": objectId}
	if err := c.UserCollection.FindOne(c.Ctx, query).Decode(&menuItem); err != nil {
		return nil, err
	}
	return menuItem, nil
}


//폰번호로 조회
func (c *UserCollection) FindByPhone(phone string) (*UserEntity, error) {
	var menuItem *UserEntity
	query := bson.M{"phone": phone}
	if err := c.UserCollection.FindOne(c.Ctx, query).Decode(&menuItem); err != nil {
		return nil, err
	}
	return menuItem, nil
}


/////////////////////////
//		Update
/////////////////////////


//업데이트 메뉴
func (c *UserCollection) UpdateEntity(_id primitive.ObjectID, updateSet bson.D) (*UserEntity, error) {
	var updateMenu *UserEntity
	query := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$set", Value: updateSet}}

	//업데이트 처리한다.
	result := c.UserCollection.FindOneAndUpdate(c.Ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	if err := result.Decode(&updateMenu); err != nil {
		return nil, errors.New("No exist user id... send id = "+ _id.String())
	}

	return updateMenu, nil;
}