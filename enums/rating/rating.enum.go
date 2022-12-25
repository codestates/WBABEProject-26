package rating_enum

/////////////////////////
//	Order Status
/////////////////////////

type RatingScore int

const (
	Worst RatingScore = iota +1  //최악
	Bad //나쁨
	Good //보통
	VeryGood //좋음
	Excellent //최고
)