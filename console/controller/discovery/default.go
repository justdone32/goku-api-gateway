package discovery

import (
	"fmt"
	"net/http"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/service"
)

func setDefault(w http.ResponseWriter, r *http.Request) {

	_, err := controller.CheckLogin(w, r, controller.OperationLoadBalance, controller.OperationEDIT)
	if err != nil {
		return
	}
	if err != r.ParseForm() {
		controller.WriteError(w, "260000", "data", "[param_check] Parse form body error | 解析form表单参数错误", err)
		return
	}
	name := r.FormValue("name")
	if !service.ValidateName(name) {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[param_check] invalid  [name]"), nil)
		return
	}

	err = service.SetDefaut(name)
	if err != nil {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[error] %s", err.Error()), err)
		return
	}

	controller.WriteResultInfo(w,
		"data",
		"data",
		nil)
}
