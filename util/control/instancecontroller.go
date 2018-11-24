package control

import (
	"github.com/astaxie/beego"
)

type InstanceController struct {
	beego.Controller
	ctrlMap *ControllerMap
}

func (ctrl *InstanceController) SetInstanceMap(ctrlMap *ControllerMap) {
	ctrl.ctrlMap = ctrlMap
}

func (ctrl *InstanceController) GetInstanceID() string {
	return ctrl.ctrlMap.GetInstanceID()
}
