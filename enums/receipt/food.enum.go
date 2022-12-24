package receipt_enums

/////////////////////////
//	Menu Category
/////////////////////////

type MenuCategoryType int

const(
	MC_KoreaStyle MenuCategoryType = iota +1 //한식
	MC_Sushi //초밥
	MC_FastFood //패스트푸드
	MC_Pizza //피자
	MC_Pasta //파스타
	MC_Bread //빵
)


/////////////////////////
//	   메뉴 이벤트
/////////////////////////

type MenuEventType int

const (
	FE_TodayMenu MenuEventType = iota + 1
	FE_Sale4CustomThanks
)

func (m MenuEventType) MenuEventStr() string {
	switch m {
	case 1:
		return "TodayMenu"
	case 2:
		return "Sale4CustomThanks"
	}
	return ""
}

/////////////////////////
//   메뉴 판매 상태
/////////////////////////

type MenuSellStatusType int

const (
	MSS_OnSeal MenuSellStatusType = iota + 1 //판매중
	MSS_SoldOut //매진
	MSS_SeasonEnd //시즌 상품
	MSS_EventEnd //한시적 이벤트
)


/////////////////////////
//	음식 메뉴 중 매운 강도
/////////////////////////

type FoodSpicyType int

const (
	FS_None FoodSpicyType = iota +1 //맵기 없음
	FS_Level1
	FS_Level2
	FS_Level3
	FS_Level4
	FS_Level5
)