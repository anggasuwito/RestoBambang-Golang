package repositories

import (
	"database/sql"
	"fmt"
	"gomux/main/apimaster/models"

	"github.com/google/uuid"
)

//MenuRepoImpl MenuRepoImpl
type MenuRepoImpl struct {
	db *sql.DB
}

//GetAllMenus GetAllMenus
func (s MenuRepoImpl) GetAllMenus(keywords string, page string, limit string) ([]*models.Menu, string, error) {
	dataMenus := []*models.Menu{}
	query := fmt.Sprintf(`select * from menu WHERE (jenis_menu LIKE '%v%v%v' or nama_menu LIKE '%v%v%v') LIMIT %v,%v`, "%", keywords, "%", "%", keywords, "%", page, limit)
	data, err := s.db.Query(query)
	if err != nil {
		return nil, "", err
	}
	for data.Next() {
		menu := models.Menu{}
		var err = data.Scan(&menu.IDMenu, &menu.JenisMenu, &menu.NamaMenu, &menu.HargaMenu, &menu.StokMenu)
		if err != nil {
			return nil, "", err
		}
		dataMenus = append(dataMenus, &menu)
	}
	var getTotalData string
	totalData := fmt.Sprintf(`select count(*) from (select * from menu WHERE (jenis_menu LIKE '%v%v%v' or nama_menu LIKE '%v%v%v')) as menu`, "%", keywords, "%", "%", keywords, "%")
	s.db.QueryRow(totalData).Scan(&getTotalData)
	return dataMenus, getTotalData, nil
}

//AddMenu InsertMenuData
func (s MenuRepoImpl) AddMenu(newMenu models.Menu) error {
	newMenuID := uuid.New().String()
	tr, err := s.db.Begin()
	query1 := `insert into menu(menu_id,jenis_menu,nama_menu,harga_menu,stok_menu)values(?,?,?,?,?)`
	_, err = s.db.Query(query1, newMenuID, &newMenu.JenisMenu, &newMenu.NamaMenu, &newMenu.HargaMenu, &newMenu.StokMenu)
	if err != nil {
		tr.Rollback()
		return err
	}

	tr.Commit()
	return nil
}

//GetMenuByID GetAllMenuById
func (s MenuRepoImpl) GetMenuByID(id string) (models.Menu, error) {
	var menu models.Menu
	query := `select * from menu where menu_id = ?`
	err := s.db.QueryRow(query, id).Scan(&menu.IDMenu, &menu.JenisMenu, &menu.NamaMenu, &menu.HargaMenu, &menu.StokMenu)
	if err != nil {
		return menu, err
	}

	return menu, nil
}

// UpdateMenusByID UpdateMenuData
func (s MenuRepoImpl) UpdateMenusByID(id string, changeMenu models.Menu) error {
	tr, err := s.db.Begin()
	_, err = s.GetMenuByID(id)

	if err != nil {
		tr.Rollback()
		return err
	}

	query := "update menu set jenis_menu = ? ,nama_menu = ?,harga_menu = ?,stok_menu = ? where menu_id=?"
	_, err = s.db.Query(query, &changeMenu.JenisMenu, &changeMenu.NamaMenu, &changeMenu.HargaMenu, &changeMenu.StokMenu, id)
	if err != nil {
		tr.Rollback()
		return err
	}
	tr.Commit()

	return nil
}

//DeleteDataMenuByID DeleteDataMenuById
func (s MenuRepoImpl) DeleteDataMenuByID(id string) error {
	tr, err := s.db.Begin()
	_, err = s.GetMenuByID(id)

	if err != nil {
		tr.Rollback()
		return err
	}
	query1 := "delete from transaksi where menu_id = ?"
	_, err = s.db.Query(query1, id)
	if err != nil {
		tr.Rollback()
		return err
	}

	query2 := "delete from menu where menu_id = ?"
	_, err = s.db.Query(query2, id)
	if err != nil {
		tr.Rollback()
		return err
	}
	tr.Commit()

	return nil
}

//InitMenuRepoImpl InitMenuRepoImpl
func InitMenuRepoImpl(db *sql.DB) MenuRepository {
	return &MenuRepoImpl{db}
}
