package repositories

import "gomux/main/master/models"

//MenuRepository MenuRepository
type MenuRepository interface {
	GetAllJenis() ([]*models.Jenis, error)
	AddMenu(newMenu models.Menu) error
	GetMenuByID(id string) (models.Menu, error)
	UpdateMenusByID(id string, changeMenu models.Menu) error
	DeleteDataMenuByID(id string) error
	GetAllMenus() ([]*models.Menu, error)
	// GetAllMenuByID(id string) (models.Menu, error)
	// DeleteDataMenuByID(id string) error
	// UpdateMenuData(id string, changeMenu models.Menu) error
	// InsertMenuData(models.Menu) error
}
