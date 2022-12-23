package receipt_enums

type MenuCategoryType int

const(
	MC_KoreaStyle MenuCategoryType = iota +1 //한식
	MC_Sushi //초밥
	MC_FastFood //패스트푸드
	MC_Pizza //피자
	MC_Pasta //파스타
	MC_Bread //빵
)

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



type FoodSpicyType int

const (
	FS_None FoodSpicyType = iota +1 //맵기 없음
	FS_Level1
	FS_Level2
	FS_Level3
	FS_Level4
	FS_Level5
)