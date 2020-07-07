package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
	"gomux/utils"
	"log"
)

//MenuUsecaseImpl MenuUsecaseImpl
type MenuUsecaseImpl struct {
	menuRepo repositories.MenuRepository
}

//GetAllMenus GetMenus
func (s MenuUsecaseImpl) GetAllMenus() ([]*models.Menu, error) {
	menu, err := s.menuRepo.GetAllMenus()
	if err != nil {
		return nil, err
	}
	return menu, nil
}

//AddMenu InsertMenus
func (s MenuUsecaseImpl) AddMenu(newMenu models.Menu) error {
	err := utils.ValidateInputNotNil(newMenu.IDMenu, newMenu.TanggalMenu, newMenu.JenisMenu, newMenu.NamaMenu, newMenu.HargaMenu, newMenu.StokMenu)

	if err != nil {
		return err
	}
	err = s.menuRepo.AddMenu(newMenu)
	if err != nil {
		log.Fatal(err)
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

//GetAllJenis GetMenus
func (s MenuUsecaseImpl) GetAllJenis() ([]*models.Jenis, error) {
	menu, err := s.menuRepo.GetAllJenis()
	if err != nil {
		return nil, err
	}
	return menu, nil
}

//InitMenuUseCase InitMenuUseCase
func InitMenuUseCase(menuRepo repositories.MenuRepository) MenuUseCase {
	return &MenuUsecaseImpl{menuRepo}
}
