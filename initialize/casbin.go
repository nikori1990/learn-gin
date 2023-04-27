package initialize

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"learn-gin/global"
	"learn-gin/model/system"
)

func InitCasbin() {
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(global.DB, &system.CasbinRule{}, "casbin_permission") // Your driver and data source.

	text := `
	[request_definition]
	r = sub, dom, obj, act

	[policy_definition]
	p = sub, dom, obj, act

	[role_definition]
	g = _, _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act || r.sub == "root"
	`

	m, err := model.NewModelFromString(text)
	if err != nil {
		global.Logger.Errorf("error: model: %s", err)
	}

	enforcer, _ := casbin.NewEnforcer(m, adapter)
	err = enforcer.LoadPolicy()
	if err != nil {
		global.Logger.Errorf("error: enforcer: %s", err)
	}

	global.CasbinEnforcer = enforcer
}
