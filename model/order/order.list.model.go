package order_list_model

import (
	"context"
	"errors"
	"time"
	order_enums "wemade_project/enums/order"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/////////////////////////
//		Entity
/////////////////////////

type OrderListEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	OrderId string `bson:"orderId"` //고유 id
	OrderUserId string `bson:"orderUserId"` //주문자
	OrderMenu []string `bson:"orderMenu"` //주문 메뉴 리스트
	OrderStatus order_enums.OrderStatus `bosn:"orderStatus"` //주문 상태
	// TotalPrice int `bson:"totalPrice"` //총 가격
	CreateDate time.Time `bson:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `bson:"updateDate"` //데이터 수정 시각
}

type OrderListCollection struct {
	OrderListCollection *mongo.Collection
	Ctx context.Context
}


/////////////////////////
//		Init
/////////////////////////

//초기화 함수
func InitWithSelf(col *mongo.Collection, ctx context.Context) OrderListCollection {
	return OrderListCollection{col, ctx}
}



/////////////////////////
//		Create 
/////////////////////////

//Add Entity
func (c *OrderListCollection) AddEntity(entity OrderListEntity) (*mongo.InsertOneResult, error) {
	result, inErr := c.OrderListCollection.InsertOne(c.Ctx, entity)

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"orderId": 1}, Options: opt}
	if _, err1 := c.OrderListCollection.Indexes().CreateOne(c.Ctx, index); err1 != nil {
		return nil, errors.New("could not create index for OrderList Id")
	}
	
	return result, inErr
}


/////////////////////////
//		Read
/////////////////////////

//Id 값으로 조회
func (c *OrderListCollection) FindByObjectId(objectId interface{}) (*OrderListEntity, error) {
	var menuItem *OrderListEntity
	query := bson.M{"_id": objectId}
	if err := c.OrderListCollection.FindOne(c.Ctx, query).Decode(&menuItem); err != nil {
		return nil, err
	}
	return menuItem, nil
}

//Order Id 값으로 조회
func (c *OrderListCollection) FindByMenuId(orderId string) (*OrderListEntity, error) {
	var orderListItem *OrderListEntity
	query := bson.M{"orderId": orderId}
	if err := c.OrderListCollection.FindOne(c.Ctx, query).Decode(&orderListItem); err != nil {
		return nil, err
	}
	return orderListItem, nil
} 


/////////////////////////
//		Update
/////////////////////////


//업데이트 메뉴
func (c *OrderListCollection) UpdateEntity(_id primitive.ObjectID, updateSet bson.D) (*OrderListEntity, error) {
	var updateOrderListItem *OrderListEntity
	query := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$set", Value: updateSet}}

	//업데이트 처리한다.
	result := c.OrderListCollection.FindOneAndUpdate(c.Ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	if err := result.Decode(&updateOrderListItem); err != nil {
		return nil, errors.New("No exist order list id... send id = "+ _id.String())
	}

	return updateOrderListItem, nil;
}


/////////////////////////
//		Delete
/////////////////////////

//엔티티를 물리적으로 삭제하는 함수
//가급적 미사용 처리 (메서드 및 외부 호출 불가 상태)
func (c *OrderListCollection) deleteEntity(_id primitive.ObjectID) error {
	query := bson.D{{Key: "_id", Value: _id}}

	res, err := c.OrderListCollection.DeleteOne(c.Ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("There is not exist order list _id... Send _id = "+ _id.String())
	}
	return nil
}

