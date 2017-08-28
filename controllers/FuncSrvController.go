package controllers

import (
	"encoding/json"

	"github.com/asofdate/auth-core/groupcache"
	"github.com/asofdate/auth-core/models"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/logger"
	"github.com/hzwy23/utils/router"
)

type FuncSrvController struct {
	funcRoute models.FuncRoute
	router.Controller
}

func GetServiceManagePage(ctx router.Context) {
	rst, err := groupcache.GetStaticFile("ServiceManage")
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.Get(ctx.Request, "as_of_date_page_not_exist"))
		return
	}
	ctx.ResponseWriter.Write(rst)
}

// 查询功能服务
// @param themeId 主题编码
// @param resId   资源编码
func (this *FuncSrvController) Get() {
	this.Ctx.Request.ParseForm()
	resId := this.Ctx.Request.FormValue("resId")
	themeId := this.Ctx.Request.FormValue("themeId")

	rst, err := this.funcRoute.Get(resId, themeId)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, "查询失败，请联系管理员")
		return
	}
	hret.Json(this.Ctx.ResponseWriter, rst)
}

// 删除功能服务配置信息
func (this *FuncSrvController) Delete() {
	this.Ctx.Request.ParseForm()

	js := this.Ctx.Request.FormValue("JSON")
	var rows []models.FuncRoute
	err := json.Unmarshal([]byte(js), &rows)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, "解析json数据失败，请联系管理员")
		return
	}

	err = this.funcRoute.Delete(rows)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 422, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, "success")
}

// 更新功能服务配置信息
func (this *FuncSrvController) Put() {
	this.Ctx.Request.ParseForm()

	form := this.Ctx.Request.Form
	var row models.FuncRoute
	row.ResId = form.Get("res_id")
	row.ResName = form.Get("res_name")
	row.ResUpId = form.Get("res_up_id")
	row.ResUrl = form.Get("res_url")
	row.ServiceCd = form.Get("service_cd")
	row.ResOpenType = form.Get("res_type")
	row.NewIframe = form.Get("new_iframe")
	row.ThemeId = form.Get("theme_id")
	row.Uuid = form.Get("uuid")
	var err error
	if this.funcRoute.IsExists(row.ResId, row.ThemeId) {
		err = this.funcRoute.Update(row)
	} else {
		err = this.funcRoute.AddTheme(row)
	}
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, "success")
}

// 新建功能服务配置信息
func (this *FuncSrvController) Post() {
	this.Ctx.Request.ParseForm()

	form := this.Ctx.Request.Form
	var row models.FuncRoute
	row.ResId = form.Get("res_id")
	row.ResName = form.Get("res_name")
	row.ResUpId = form.Get("res_up_id")
	row.ResUrl = form.Get("res_url")
	row.ServiceCd = form.Get("service_cd")
	row.ResOpenType = form.Get("res_type")
	row.NewIframe = form.Get("new_iframe")
	row.ThemeId = form.Get("theme_id")

	err := this.funcRoute.Post(row)
	if err != nil {
		logger.Error(err)
		hret.Error(this.Ctx.ResponseWriter, 421, err.Error())
		return
	}
	hret.Success(this.Ctx.ResponseWriter, "success")
}

func init() {
	groupcache.RegisterStaticFile("ServiceManage", "./views/hauth/service.tpl")
}
