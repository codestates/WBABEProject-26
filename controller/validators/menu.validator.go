package validators

import (
	menu_enums "wemade_project/enums/menu"

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
            if value, ok := fl.Field().Interface().([]menu_enums.MenuEventType); ok {
                return ValidateRegex(value)
            }
            return true
        }
}

func ValidateRegex(value []menu_enums.MenuEventType) bool {
	for _, val := range value {
		result := menu_enums.MenuEventType.MenuEventStr(val)
		if len(result) == 0 {
			return false
		}
	}
	return true
}
