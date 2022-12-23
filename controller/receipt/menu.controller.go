package receipt_controller

import (
	"fmt"
	"net/http"

	"wemade_project/dto"
	receipt_dto "wemade_project/dto"
	receipt_service "wemade_project/service/receipt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/////////////////////////
//	  Struct
/////////////////////////

type MenuController struct {
	menuService receipt_service.MenuService
}


/////////////////////////
//	  Init func
/////////////////////////

//생성자 역할 함수
func InitWithSelf(menuService receipt_service.MenuService) MenuController {
	return MenuController{menuService: menuService}
}


/////////////////////////
//	  Get (Read)
/////////////////////////

//메뉴 조회
func (mc *MenuController) GetMenu() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ginCtx.JSON(http.StatusOK, gin.H{"Message" : "Get Menu"})
	}
}

/////////////////////////
//	  Add (Create)
/////////////////////////

//메뉴 등록 함수
func (mc *MenuController) AddMenu() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//요청 데이터를 확인한다.
		var addMenuReq receipt_dto.CreateMenuRequest;

		if err := ginCtx.ShouldBindJSON(&addMenuReq); err != nil {
			fmt.Println("err = ", err )
			handleBindError(ginCtx, addMenuReq, "checkerMenuEvent", err)
			return 
		}

		//전달된 아이템 등록 처리
		result, mongoErr := mc.menuService.AddMenuItem(addMenuReq)

		//조회에 에러가 난 경우
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
		}
		
		// res := dto.ResponseBody{Result: true, Data: result}
		// fmt.Println("res = ", res)
		//fmt.Println("re = ", utils.ChangeStruct2JsonStr(addMenuReq))
	
		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}


/////Error bind
func handleBindError(c *gin.Context, obj interface{}, tag string, err error) {
	var errs []gin.H
	// /switch err.(type) {
	switch err.(type) {
		case validator.ValidationErrors:	
			vErrs := err.(validator.ValidationErrors)
			for _, vErr := range vErrs {
				var message string
				value := vErr.Value()
				switch vErr.ActualTag() {
					case "checkerMenuEvent":
						message = fmt.Sprintln("Menu Event Code is wrong. Check Value. ", value)
					default:
            			message = err.Error()    
				}
				errs = append(errs, gin.H{
                	"field":   vErr.Field(),
                	"value":   value,
                	"message": message,
            	})
			}
		default: 
			errs = append(errs, gin.H{
                	"message": err.Error(),
            })
	}

	
	// fmt.Println("tag = ",vErr.ActualTag(), " value = ", vErr.Value(), " fil = ", vErr.Field())	
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"errors": errs,
	})	 
}

/*
switch err.(type) {
    case validator.ValidationErrors:
        var errs []gin.H
        vErrs := err.(validator.ValidationErrors)
        e := reflect.TypeOf(obj).Elem()
        for _, vErr := range vErrs {
            field, _ := e.FieldByName(vErr.Field())
            tagName, _ := field.Tag.Lookup(tag)
            value := vErr.Value()
            var message string
            switch vErr.ActualTag() {
            case "required":
                message = fmt.Sprintf("required %s", tagName)
            case "hexadecimal":
                message = fmt.Sprintf("required hexadecimal format")
            case "gte":
                message = fmt.Sprintf("greater than or quauls to %s", vErr.Param())
            case "cgte":
                message = fmt.Sprintf("greater than or quauls to %s", vErr.Param())
            case "numeric":
                message = fmt.Sprintf("%s must be numeric", tagName)
			case "checkerMenuEvent":
				message = fmt.Sprintln("Menu Event Code is wrong. Check Value. ", value)
            default:
                message = err.Error()
            }
            errs = append(errs, gin.H{
                "field":   tagName,
                "value":   value,
                "message": message,
            })
        }
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "errors": errs,
        })
        return
    }
    c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())

*/





/*

   "age": 23,
    "pnum": "01012352"
var newPerson *CreatePerson2Request

		if err := ctx.ShouldBindJSON(&newPerson); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

func (p *Controller) NewPersonInsert(c *gin.Context) {
	name := c.PostForm("name")
	sAge := c.PostForm("age")
	spnum := c.PostForm("pnum")

	if len(name) <= 0 || len(spnum) <= 0 {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", nil)
		return
	}
*/



