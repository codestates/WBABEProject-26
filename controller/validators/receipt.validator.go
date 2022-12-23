package validators

import (
	receipt_enums "wemade_project/enums/receipt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

/////////////////////////
//		Reg Validator
/////////////////////////

func RegValidator4MenuEvent() {
	   if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("checkerMenuEvent", validator4MenuEvent())
    }
}


/////////////////////////
//		Valid
/////////////////////////

func validator4MenuEvent() validator.Func {
	return func(fl validator.FieldLevel) bool {
            if value, ok := fl.Field().Interface().([]receipt_enums.MenuEventType); ok {
                return ValidateRegex(value)
            }
            return true
        }
}

func ValidateRegex(value []receipt_enums.MenuEventType) bool {
	for _, val := range value {
		result := receipt_enums.MenuEventType.MenuEventStr(val)
		if len(result) == 0 {
			return false
		}
	}
	return true
}
