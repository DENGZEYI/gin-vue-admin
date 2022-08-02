package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/business"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/test"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	TestServiceGroup     test.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
