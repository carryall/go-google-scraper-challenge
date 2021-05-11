package helpers

import "github.com/beego/beego/v2/server/web"

var ActionsWithGetMethod = []string{"List", "New", "Edit"}

// SetControllerAttributes set attributes for controller
func SetControllerAttributes(controller *web.Controller) {
	controllerName, actionName := controller.GetControllerAndAction()

	controller.Data["ControllerName"] = ToKebabCase(controllerName)
	controller.Data["ActionName"] = ToKebabCase(actionName)
}

func IsActionWithGetMethod(actionName string) bool {
	for _, a := range ActionsWithGetMethod {
		if a == actionName {
			return true
		}
	}
	return false
}
