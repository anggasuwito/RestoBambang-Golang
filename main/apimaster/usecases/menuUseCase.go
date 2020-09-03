package usecases

import "gomux/main/apimaster/models"

//MenuUseCase MenuUseCase
type MenuUseCase interface {
	GetAllMenus() ([]*models.Menu, error)
	AddMenu(newMenu models.Menu) error
	GetMenuByID(id string) (models.Menu, error)
	UpdateMenusByID(id string, changeMenu models.Menu) error
	DeleteMenusByID(id string) error
}
