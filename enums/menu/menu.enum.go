package menu_enums

/////////////////////////
//	Menu Category
/////////////////////////

type MenuCategoryType int

const(
	/*
	1. MC라는 접두사가 필요한지 잘 모르겠습니다. 오히려 축약어를 써버려서 의미가 직관적이지 않다고 생각합니다. MenuCategoryType 이라는 type이 있으므로, 안에서는 MC_를 제거해도 좋을 것 같습니다.

	2. int형 보다는 string 값을 할당해주는 것이 가독성에 더 좋을 것입니다.
	*/
	MC_KoreaStyle MenuCategoryType = iota +1 //한식
	MC_Sushi //초밥
	MC_FastFood //패스트푸드
	MC_Pizza //피자
	MC_Pasta //파스타
	MC_Bread //빵
)

/*
type의 네이밍만 보더라도 어떤 상수인지 유추할 수 있습니다. 따라서 불필요한 주석이라고 생각됩니다.
*/
/////////////////////////
//	   메뉴 이벤트
/////////////////////////

type MenuEventType int

const (
	FE_TodayMenu MenuEventType = iota + 1
	/*
	2 -> to, 4 -> for 와 같이 풀어쓰는 것이 훨씬 직관적입니다.
	*/
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
	/*
	접두사가 필요한지 잘 모르겠습니다.
	*/
	MSS_OnSeal MenuSellStatusType = iota + 1 //판매중
	MSS_SoldOut //매진
	MSS_Season //시즌 상품
	MSS_SeasonEnd //시즌 상품 종료
	MSS_EventEnd //한시적 이벤트
	MSS_Delete //삭제된 메뉴
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