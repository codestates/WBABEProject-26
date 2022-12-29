# 띵동주문이요, 온라인 주문 시스템(Online Ordering System)

## 개요
* Go Backend mini project.  
* 학습하며 배운 Go를 이용해서 간단한 API 서비스를 개발하였습니다.

## 사용 라이브러리
- 설치 및 사용한 라이브러리는 아래와 같습니다.

```
# Gin framework
go get github.com/gin-gonic/gin

# Mongo DB
go get go.mongodb.org/mongo-driver/mongo 
go get go.mongodb.org/mongo-driver/mongo/options
go get go.mongodb.org/mongo-driver/bson

# Swagger
go get -u github.com/swaggo/swag/cmd/swag

# UUID 
go get github.com/gofrs/uuid
```

## 프로젝트 설정 및 실행
- git과 go가 정상적으로 설치되어 있다고 가정합니다.

```
# project clone
git clone git@github.com:codestates/WBABEProject-26.git

# move project directory and cmd run
go mod tidy

# run project
go run main.go
```

## 설정 관련
- 프로젝트 설정의 경우 아래의 파일을 참고하세요
- [/config/config.toml](https://github.com/codestates/WBABEProject-26/blob/main/config/config.toml)
- 상용 서비스 및 오픈 서비스에서는 [.gitignore](https://github.com/codestates/WBABEProject-26/blob/main/.gitignore)에 등록 후 사용합니다.


## API 명세
- API 명세는 다음과 같습니다.

| URL  | Method | Description |
| ------------- |:-------------:|:-------------:|
| /api/v1/store/menu/get      | GET |  메뉴 리스트 조회 |
| /api/v1/store/menu/get/{menu_id}      | GET |  메뉴 id로 대상 메뉴 상세 조회     |
| /api/v1/store/menu/add      | POST | 메뉴 등록    |
| /api/v1/store/menu/update      | PUT | 메뉴 정보 수정     |
| /api/v1/store/menu/delete/{menu_id}   | DELETE |  메뉴 삭제     |
| /api/v1/rating/add    | POST |  사용자 리뷰 및 평점 등록     |
| /api/v1/order_list/order    | GET |  주문 내역 리스트 조회    |
| /api/v1/order_list/order/user/{user_id}    | GET |  사용자별 주문 리스트 조회    |
| /api/v1/order_list/order/add    | POST |  사용자 주문 접수     |
| /api/v1/order_list/order/menu/update    | PUT |  사용자 주문 메뉴 수정     |
| /api/v1/order_list/order/status/update    | PUT |  업주측 주문 상태 변경     |
| /api/v1/account/user/add    | POST | 사용자 등록     |

이미지 넣기

## Swagger 명세
- Swaggo를 사용해 구현되었으며, 접근은 아래와 같이 가능합니다.
- 서버는 구동된 상태여야 합니다.
- http://localhost:8090/swagger/index.html


## Database 구조
- MongoDB를 사용하였으며, 아래의 구조체로 각 컬렉션 엔티티 구조를 대신합니다.

### Menu

```go
type MenuEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"` //MongoDB _id
	Id string `bson:"id"` //메뉴 고유 id
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
```

* 특이사항
    * **Id**
        * 메뉴 고유 id를 의미하며, 다른 곳에서 조회 및 참조 시 이용합니다.
        * ID는 MongoDB _id를 의미하며, 시스템 내부에서 사용 시에만 사용됩니다.
    * **MenuStatus** 
        * enum 값으로 구성됩니다.
    * **Event**
        * 오늘의 메뉴 및 각종 메뉴에 대한 이벤트는 Event의 enum에서 등록 및 배열 처리를 통해 진행하며, Front-End 단에서 해당 enum에 맞게 데이터를 표시합니다.
    * **MenuCategory**
        * enum array로 구성되며 메인에서 각 메뉴별 카테고리에 맞게 표시할 때 사용됩니다.
    * **SubMenu**
        * 각 메뉴의 하위 메뉴로써 대상 메뉴에서 소스 및 부가적 주문 옵션 등록 시 사용합니다.
    * **CreateDate / UpdateDate** 
        * Mongo에서 UTC 시간으로 계산되어 저장되며, 처리 시 Timezone 처리 또는 하드코딩 처리가 필요합니다.

### OrderList

```go
type OrderListEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"` //MongoDB _id
	OrderId string `bson:"orderId"` //고유 id
	OrderUserId string `bson:"orderUserId"` //주문자
	OrderMenu []string `bson:"orderMenu"` //주문 메뉴 리스트
	OrderStatus order_enums.OrderStatus `bson:"orderStatus"` //주문 상태
	CreateDate time.Time `bson:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `bson:"updateDate"` //데이터 수정 시각
}
```

* 특이사항
    * **OrderId**
        * 주문 고유 id를 의미하며, 다른 곳에서 조회 및 참조 시 이용합니다.
        * UUID를 사용합니다.
    * **OrderUserId**
        * 주문 생성한 사용자 id를 가리킵니다.
        * mongodb 특성 상 주문 사용자 데이터를 모두 또는 일부를 설계에 따라 넣을 수 있지만, 샘플 프로젝트라 id만 참조하였습니다.
    * **OrderMenu**
        * 주문한 메뉴를 가리킵니다.
    * **OrderStatus**
        * 메뉴의 상태를 가리킵니다.
        * enum으로 관리됩니다.     
    * **CreateDate / UpdateDate** 
        * Mongo에서 UTC 시간으로 계산되어 저장되며, 처리 시 Timezone 처리 또는 하드코딩 처리가 필요합니다.   

### Rating

```go
type RatingEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"` //MongoDB _id
	UserId string `bson:"userId"` //주문자 (_id)
	OderListId string `bson:"orderListId"` //주문리스트 (_id)
	MenuId string `bson:"menuId"` //메뉴 (_id)
	Rating rating_enum.RatingScore `bson:"rating"`
	ReviewMsg string `bson:"reviewMsg"`
	Recommendation bool `bson:"recommendation"`
	CreateDate time.Time `bson:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `bson:"updateDate"` //데이터 수정 시각
}
```

* 특이사항
    * **UserId**
        * 평점 리뷰 등록한 사용자
        * User entity의 ID 값
    * **OderListId**
        * 리뷰할 주문 id
        * 메뉴 별 등록할 수 있도록 하여 OrderListId의 경우 중복 저장 가능합니다. 
    * **MenuId**
        * 주문 리스트 내의 menu id 
        * 해당 값을 넣은 이유는 향후 메뉴별 상세 리뷰 기능 등을 활용할 때 사용할 수 있게끔 처리되었습니다.
    * **Rating**
        * 평점 enum(1~5)
    * **ReviewMsg**
        * 리뷰 메세지
    * **Recommendation**
        * 해당 메뉴에 대한 추천 여부이며 bool 값을 사용하였습니다.
    * **CreateDate / UpdateDate** 
        * Mongo에서 UTC 시간으로 계산되어 저장되며, 처리 시 Timezone 처리 또는 하드코딩 처리가 필요합니다.   

### User

```go
type UserEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name"` //사용자 이름 
	Phone string `bson:"phone"` //사용자 폰번호
	Addr string `bson:"addr"` //주소
	CreateDate time.Time `bson:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `bson:"updateDate"` //데이터 수정 시각
}
```

* 특이사항
    * 기본적인 샘플용으로 작성되었습니다.


