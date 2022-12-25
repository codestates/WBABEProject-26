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
	OrderStatus order_enums.OrderStatus `bson:"orderStatus"` //주문 상태
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
func (c *OrderListCollection) FindByOrderId(orderId string) (*OrderListEntity, error) {
	var orderListItem *OrderListEntity
	query := bson.M{"orderId": orderId}
	if err := c.OrderListCollection.FindOne(c.Ctx, query).Decode(&orderListItem); err != nil {
		return nil, err
	}
	return orderListItem, nil
} 

//User Id로 조회
func (c *OrderListCollection) Find4OrderUserIdAndStatus(orderUserId string, sortOpt string)  ([]*OrderListEntity, error)  {
	//검색 대상
	query := bson.M{"orderStatus": bson.M{"$ne": 4}}
	
	//만약 주문 완료 내역만 보는 경우
	if (sortOpt == "on") {
		query = bson.M{"orderStatus": bson.M{"$eq": 4}}
	}


	//향후 페이징 처리 필요
	opt := options.FindOptions{}
	opt.SetSort(bson.M{"createDate": -1})

	cursor, err := c.OrderListCollection.Find(c.Ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(c.Ctx)

	var orderListItem []*OrderListEntity

	for cursor.Next(c.Ctx) {
		item := &OrderListEntity{}
		err := cursor.Decode(item)

		if err != nil {
			return nil, err
		}

		orderListItem = append(orderListItem, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(orderListItem) == 0 {
		return []*OrderListEntity{}, nil
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


