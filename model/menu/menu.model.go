package menu_model

import (
	"context"
	"errors"
	"time"
	menu_enums "wemade_project/enums/menu"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//음식 기타 정보
type FoodEtcInfo struct {
	OriginInfo string //원산지
	SpicyInfo menu_enums.FoodSpicyType //맵기 정보
}


//음식 서브 메뉴
type SubMenu struct {
	SubMenuName string //서브 메뉴 타이틀
	Name string //서브 메뉴 이름
	Price int //가격
}


//판매음식 메뉴
type MenuEntity struct {
	/*
	메뉴 삭제에 대한 플래그는 Status에 포함시키는 것보다는 따로 필드를 두는 것은 어떠할까요?
	실제 서비스 운용시에도 data를 삭제하면 IsDeleted와 같은 필드에 true, false로 값을 업데이트하여 사용하는 것이 일반적입니다.
	*/
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Id string `bson:"id"` //고유 id
	Name string `bson:"name"` //메뉴 이름
	MenuStatus menu_enums.MenuSellStatusType `bson:"menuStatus"`	//주문 가능 여부	
	Price int `bson:"price"` //가격
	Event []menu_enums.MenuEventType `bson:"event"` //이벤트
	MenuCategory []menu_enums.MenuCategoryType `bson:"menuCategory"` //매뉴 카테고리
	SubMenu []SubMenu `bson:"subMenu"` //서브메뉴
	FoodEtcInfo FoodEtcInfo `bson:"foodEtcInfo"` //기타 정보
	CreateDate time.Time `bson:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `bson:"updateDate"` //데이터 수정 시각
}

/*
사용하지 않는 것은 지워주시는 것이 좋아 보입니다.
*/
//recommendation_count [] user_id

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
func (m *MenuCollection) FindByObjectId(objectId interface{}) (*MenuEntity, error) {
	var menuItem *MenuEntity
	query := bson.M{"_id": objectId}
	if err := m.MenuCollection.FindOne(m.Ctx, query).Decode(&menuItem); err != nil {
		return nil, err
	}
	return menuItem, nil
}

//Menu Id 값으로 조회
func (m *MenuCollection) FindEntity2MenuId(menuId string) (*MenuEntity, error) {
	var menuItem *MenuEntity
	query := bson.M{"id": menuId}
	if err := m.MenuCollection.FindOne(m.Ctx, query).Decode(&menuItem); err != nil {
		return nil, err
	}
	return menuItem, nil
} 

/**
* 메뉴 리스트를 조회하는 함수
*/
func (m *MenuCollection) FindEntityList2All() ( []*MenuEntity, error) {
	//검색 대상
	query := bson.M{}

	//향후 페이징 처리 필요
	opt := options.FindOptions{}
	opt.SetSort(bson.M{"createDate": -1})

	/*
	*/
	//* 각 카테고리별 sort 리스트 출력 order by 추천, 평점, 최신 

	cursor, err := m.MenuCollection.Find(m.Ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(m.Ctx)

	//순회하며 데이터를 추가해준다.
	var menuEntityList []*MenuEntity
	for cursor.Next(m.Ctx) {
		item := &MenuEntity{}
		err := cursor.Decode(item)
		if err != nil {
			return nil, err
		}
		menuEntityList = append(menuEntityList, item)
	}
	
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return menuEntityList, nil
}

/////////////////////////
//		Update
/////////////////////////


//업데이트 메뉴
func (m *MenuCollection) UpdateEntity(_id primitive.ObjectID, updateSet bson.D) (*MenuEntity, error) {
	var updateMenu *MenuEntity
	query := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$set", Value: updateSet}}

	//업데이트 처리한다.
	result := m.MenuCollection.FindOneAndUpdate(m.Ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	if err := result.Decode(&updateMenu); err != nil {
		return nil, errors.New("No exist menu id... send id = "+ _id.String())
	}

	return updateMenu, nil;
}


/////////////////////////
//		Delete
/////////////////////////

/*
Delete라는 메서드를 호출 하되, 실제로 값은 Update를 통해 IsDeleted 값만 변경하는 것은 어떨까요?
update를 통해서 status를 통해 삭제된 메뉴를 표시하는 것 보다는 삭제하는 메서드를 호출하면 업데이트와 삭제 행위에 대한 구분이 가능해집니다.
*/

//엔티티를 물리적으로 삭제하는 함수
//가급적 미사용 처리 (메서드 및 외부 호출 불가 상태)
func (m *MenuCollection) deleteEntity(_id primitive.ObjectID) error {
	query := bson.D{{Key: "_id", Value: _id}}

	res, err := m.MenuCollection.DeleteOne(m.Ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("There is not exist menu _id... Send _id = "+ _id.String())
	}
	return nil
}