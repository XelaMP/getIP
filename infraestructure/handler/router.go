package handler

import (
	"github.com/XelaMP/getIP.git/infraestructure/handler/ip"
	"github.com/XelaMP/getIP.git/model"
)

func InitRoutes(specification model.RouterSpecification) {
	ip.NewRouter(specification)
}
