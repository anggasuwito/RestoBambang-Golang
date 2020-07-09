package master

import (
	"gomux/config"
	"gomux/main/apimaster/controllers"
	"gomux/main/apimaster/repositories"
	"gomux/main/apimaster/usecases"
)

//Init Init
func Init() {
	db := config.InitDB()
	router := config.CreateRouter()

	menuRepo := repositories.InitMenuRepoImpl(db)
	menuUseCase := usecases.InitMenuUseCase(menuRepo)
	controllers.MenuController(router, menuUseCase)

	transaksiRepo := repositories.InitTransaksiRepoImpl(db)
	transaksiUseCase := usecases.InitTransaksiUseCase(transaksiRepo)
	controllers.TransaksiController(router, transaksiUseCase)

	config.RunServer(router)
}
