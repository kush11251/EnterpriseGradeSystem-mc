package main

import (
	"github.com/EnterpriseGradeSystem/pkg/config"
	"github.com/EnterpriseGradeSystem/pkg/controllers"
	"github.com/EnterpriseGradeSystem/pkg/services"
)

func main() {
	services.InitServices()
	controllers.InitControllers()
	config.InitConfig()
}
