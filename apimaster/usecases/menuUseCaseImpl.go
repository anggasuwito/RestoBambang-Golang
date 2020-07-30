package usecases

import (
	"restoAPI/apimaster/models"
	"restoAPI/apimaster/repositories"
	"restoAPI/utils"
)

//MenuUsecaseImpl MenuUsecaseImpl
type MenuUsecaseImpl struct {
	menuRepo repositories.MenuRepository
}

//GetAllMenus GetMenus
func (s MenuUsecaseImpl) GetAllMenus(keywords string, page string, limit string) ([]*models.Menu, string, error) {
	menu, totalData, err := s.menuRepo.GetAllMenus(keywords, page, limit)
	if err != nil {
		return nil, "", err
	}
	return menu, totalData, nil
}

//AddMenu InsertMenus
func (s MenuUsecaseImpl) AddMenu(newMenu models.Menu) error {
	err := utils.ValidateInputNotNil(newMenu.JenisMenu, newMenu.NamaMenu, newMenu.HargaMenu, newMenu.StokMenu)

	if err != nil {
		return err
	}
	err = s.menuRepo.AddMenu(newMenu)
	if err != nil {
		return err
	}
	return err
}

//GetMenuByID GetMenusByID
func (s MenuUsecaseImpl) GetMenuByID(id string) (models.Menu, error) {
	menu, err := s.menuRepo.GetMenuByID(id)
	if err != nil {
		return menu, err
	}
	return menu, nil
}

// UpdateMenusByID UpdateMenusByID
func (s MenuUsecaseImpl) UpdateMenusByID(id string, changeMenu models.Menu) error {
	err := s.menuRepo.UpdateMenusByID(id, changeMenu)

	if err != nil {
		return err
	}
	return nil
}

//DeleteMenusByID DeleteMenusByID
func (s MenuUsecaseImpl) DeleteMenusByID(id string) error {
	err := s.menuRepo.DeleteDataMenuByID(id)

	if err != nil {
		return err
	}
	return nil
}

//InitMenuUseCase InitMenuUseCase
func InitMenuUseCase(menuRepo repositories.MenuRepository) MenuUseCase {
	return &MenuUsecaseImpl{menuRepo}
}
