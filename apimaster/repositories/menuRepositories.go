package repositories

import "restoAPI/apimaster/models"

//MenuRepository MenuRepository
type MenuRepository interface {
	AddMenu(newMenu models.Menu) error
	GetMenuByID(id string) (models.Menu, error)
	UpdateMenusByID(id string, changeMenu models.Menu) error
	DeleteDataMenuByID(id string) error
	GetAllMenus(keywords string, page string, limit string) ([]*models.Menu, string, error)
}
