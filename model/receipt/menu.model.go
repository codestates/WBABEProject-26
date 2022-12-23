package receipt_model

import (
	"context"
	"errors"
	"time"
	receipt_enums "wemade_project/enums/receipt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//음식 기타 정보
type FoodEtcInfo struct {
	OriginInfo string //원산지
	SpicyInfo receipt_enums.FoodSpicyType //맵기 정보
}


//음식 서브 메뉴
type SubMenu struct {
	SubMenuName string //서브 메뉴 타이틀
	Name string //서브 메뉴 이름
	Price int //가격
}


//판매음식 메뉴
type MenuEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Id string `bson:"id"` //고유 id
	Name string `bson:"name"` //메뉴 이름
	IsCanOrder bool `bson:"isCanOrder"`	//주문 가능 여부	
	Price int `bson:"price"` //가격
	Event []receipt_enums.MenuEventType `bson:"event"` //이벤트
	MenuCategory []receipt_enums.MenuCategoryType `bson:"menuCategory"` //매뉴 카테고리
	SubMenu []SubMenu `bson:"subMenu"` //서브메뉴
	FoodEtcInfo FoodEtcInfo `bson:"foodEtcInfo"` //기타 정보
	CreateDate time.Time `bson:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `bson:"updateDate"` //데이터 수정 시각
}

// type MenuModel interface {
// 	CreateMenu(menu *MenuEntity) (*MenuEntity, error)
// 	FindAllMenu() ([]MenuEntity)
// }


type MenuCollection struct {
	MenuCollection *mongo.Collection
	Ctx context.Context
}

/////////////////////////
//		Init
/////////////////////////

func InitWithSelf(menuCol *mongo.Collection, ctx context.Context) MenuCollection {
	
	return MenuCollection{menuCol, ctx}
}

/////////////////////////
//		Create 
/////////////////////////

//Add Entity
func (m *MenuCollection) AddEntity(entity MenuEntity) (*mongo.InsertOneResult, error) {
	result, inErr := m.MenuCollection.InsertOne(m.Ctx, entity)

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"id": 1}, Options: opt}
	if _, err1 := m.MenuCollection.Indexes().CreateOne(m.Ctx, index); err1 != nil {
		return nil, errors.New("could not create index for menu Id")
	}
	
	return result, inErr
}


/////////////////////////
//		Read
/////////////////////////

//Id 값으로 조회
func (m *MenuCollection) FindByInnerId(innerId interface{}) (*MenuEntity, error) {
	var menuItem *MenuEntity
	query := bson.M{"_id": innerId}
	if err := m.MenuCollection.FindOne(m.Ctx, query).Decode(&menuItem); err != nil {
		return nil, err
	}
	return menuItem, nil
}




/*
var newPost *models.DBPost
	query := bson.M{"_id": res.InsertedID}
	if err = p.postCollection.FindOne(p.ctx, query).Decode(&newPost); err != nil {
		return nil, err
	}
*/


/*
type Dog struct {
	ID     primitive.ObjectID
*/