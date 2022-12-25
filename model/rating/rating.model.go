package rating_model

import (
	"context"
	"errors"
	"time"
	rating_enum "wemade_project/enums/rating"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/////////////////////////
//		Entity
/////////////////////////

type RatingEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserId string `bson:"userId"` //주문자 (_id)
	OderListId string `bson:"orderListId"` //주문리스트 (_id)
	MenuId string `bson:"menuId"` //메뉴 (_id)
	Rating rating_enum.RatingScore `bson:"rating"`
	ReviewMsg string `bson:"reviewMsg"`
	Recommendation bool `bson:"recommendation"`
	CreateDate time.Time `bson:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `bson:"updateDate"` //데이터 수정 시각
}

type RatingCollection struct {
	RatingCollection *mongo.Collection
	Ctx context.Context
}

/////////////////////////
//		Init
/////////////////////////

//초기화 함수
func InitWithSelf(col *mongo.Collection, ctx context.Context) RatingCollection {
	return RatingCollection{col, ctx}
}


/////////////////////////
//		Create 
/////////////////////////

//Add Entity
func (c *RatingCollection) AddEntity(entity RatingEntity) (*mongo.InsertOneResult, error) {
	result, inErr := c.RatingCollection.InsertOne(c.Ctx, entity)

	opt := options.Index()
	// opt.SetUnique(true)

	//메뉴 Id만 인덱싱을 한다.
	index := mongo.IndexModel{Keys: bson.M{"menuId": 1}, Options: opt}
	if _, err1 := c.RatingCollection.Indexes().CreateOne(c.Ctx, index); err1 != nil {
		return nil, errors.New("could not create index for OrderList Id")
	}
	
	return result, inErr
}


/////////////////////////
//		Read
/////////////////////////

//Id 값으로 조회
func (c *RatingCollection) FindByObjectId(objectId interface{}) (*RatingEntity, error) {
	var menuItem *RatingEntity
	query := bson.M{"_id": objectId}
	if err := c.RatingCollection.FindOne(c.Ctx, query).Decode(&menuItem); err != nil {
		return nil, err
	}
	return menuItem, nil
}


//Menu Id
func (c *RatingCollection) FindListByMenuId (menuId interface{})  ([]*RatingEntity, error) {
	// var ratingItem *RatingEntity
	// if err := c.RatingCollection.FindOne(c.Ctx, query).Decode(&ratingItem); err != nil {
	// 	return nil, err
	// }

	query := bson.M{"menuId": menuId}

	//향후 페이징 처리 필요
	opt := options.FindOptions{}
	opt.SetSort(bson.M{"createDate": -1})

	cursor, err := c.RatingCollection.Find(c.Ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(c.Ctx)

	var ratingEntityList []*RatingEntity

	for cursor.Next(c.Ctx) {
		item := &RatingEntity{}
		err := cursor.Decode(item)

		if err != nil {
			return nil, err
		}

		ratingEntityList = append(ratingEntityList, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(ratingEntityList) == 0 {
		return []*RatingEntity{}, nil
	}

	return ratingEntityList, nil
}

/////////////////////////
//		Update
/////////////////////////


//업데이트 메뉴
func (c *RatingCollection) UpdateEntity(_id primitive.ObjectID, updateSet bson.D) (*RatingEntity, error) {
	var updateOrderListItem *RatingEntity
	query := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$set", Value: updateSet}}

	//업데이트 처리한다.
	result := c.RatingCollection.FindOneAndUpdate(c.Ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	if err := result.Decode(&updateOrderListItem); err != nil {
		return nil, errors.New("No exist rating id... send id = "+ _id.String())
	}

	return updateOrderListItem, nil;
}


/////////////////////////
//		Delete
/////////////////////////

//엔티티를 물리적으로 삭제하는 함수
//가급적 미사용 처리 (메서드 및 외부 호출 불가 상태)
func (c *RatingCollection) deleteEntity(_id primitive.ObjectID) error {
	query := bson.D{{Key: "_id", Value: _id}}

	res, err := c.RatingCollection.DeleteOne(c.Ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("There is not exist rating _id... Send _id = "+ _id.String())
	}
	return nil
}