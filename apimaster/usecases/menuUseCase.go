package usecases

import "restoAPI/apimaster/models"

//MenuUseCase MenuUseCase
type MenuUseCase interface {
	GetAllMenus(keywords string, page string, limit string) ([]*models.Menu, string, error)
	AddMenu(newMenu models.Menu) error
	GetMenuByID(id string) (models.Menu, error)
	UpdateMenusByID(id string, changeMenu models.Menu) error
	DeleteMenusByID(id string) error
}
