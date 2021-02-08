package helpers

import "github.com/beego/beego/v2/server/web"

// SetControllerAttributes set attributes for controller
func SetControllerAttributes(controller *web.Controller) {
	controllerName, actionName := controller.GetControllerAndAction()

	controller.Data["ControllerName"] = ToKebabCase(controllerName)
	controller.Data["ActionName"] = ToKebabCase(actionName)
}
