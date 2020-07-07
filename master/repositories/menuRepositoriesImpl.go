package repositories

import (
	"database/sql"
	"gomux/main/master/models"
	"log"
)

//MenuRepoImpl MenuRepoImpl
type MenuRepoImpl struct {
	db *sql.DB
}

//GetAllMenus GetAllMenus
func (s MenuRepoImpl) GetAllMenus() ([]*models.Menu, error) {
	dataMenus := []*models.Menu{}
	query := `select l.id_list_menu,l.tanggal,j.nama_jenis,m.nama_menu,h.harga,s.stok
	from list_menu l
	join jenis j on l.id_jenis = j.id_jenis
	join menu m on m.id_menu = l.id_menu
	join harga_menu h on h.id_menu = l.id_menu
	join stok_menu s on s.id_menu = l.id_menu
	order by l.id_list_menu`
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		menu := models.Menu{}
		var err = data.Scan(&menu.IDMenu, &menu.TanggalMenu, &menu.JenisMenu, &menu.NamaMenu, &menu.HargaMenu, &menu.StokMenu)
		if err != nil {
			return nil, err
		}
		dataMenus = append(dataMenus, &menu)
	}
	return dataMenus, nil
}

//AddMenu InsertMenuData
func (s MenuRepoImpl) AddMenu(newMenu models.Menu) error {

	tr, err := s.db.Begin()
	query := `insert into list_menu(id_list_menu,tanggal,id_jenis,id_menu,id_harga_menu,id_stok_menu)values(?,?,?,?,?,?)`
	_, err = s.db.Query(query, &newMenu.IDMenu, &newMenu.TanggalMenu, &newMenu.JenisMenu, &newMenu.NamaMenu, &newMenu.HargaMenu, &newMenu.StokMenu)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}

	return nil
}

//GetMenuByID GetAllMenuById
func (s MenuRepoImpl) GetMenuByID(id string) (models.Menu, error) {
	var menu models.Menu
	query := `select l.id_list_menu,l.tanggal,j.nama_jenis,m.nama_menu,h.harga,s.stok
	from list_menu l
	join jenis j on l.id_jenis = j.id_jenis
	join menu m on m.id_menu = l.id_menu
	join harga_menu h on h.id_menu = l.id_menu
	join stok_menu s on s.id_menu = l.id_menu
	where l.id_list_menu = ?`
	err := s.db.QueryRow(query, id).Scan(&menu.IDMenu, &menu.TanggalMenu, &menu.JenisMenu, &menu.NamaMenu, &menu.HargaMenu, &menu.StokMenu)
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

	query := "update list_menu set id_list_menu = ? ,tanggal = ? ,id_jenis = ? ,id_menu = ?,id_harga_menu = ?,id_stok_menu = ? where id_list_menu=?"
	_, err = s.db.Query(query, &changeMenu.IDMenu, &changeMenu.TanggalMenu, &changeMenu.JenisMenu, &changeMenu.NamaMenu, &changeMenu.HargaMenu, &changeMenu.StokMenu, id)
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
	query := "delete from list_menu where menu_id = ?"
	_, err = s.db.Query(query, id)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}

	return nil
}

//GetAllJenis GetAllJenis
func (s MenuRepoImpl) GetAllJenis() ([]*models.Jenis, error) {
	dataJenis := []*models.Jenis{}
	query := "select * from jenis"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		jenis := models.Jenis{}
		var err = data.Scan(&jenis.IDJenis, &jenis.NamaJenis)
		if err != nil {
			return nil, err
		}
		dataJenis = append(dataJenis, &jenis)
	}
	return dataJenis, nil
}

//InitMenuRepoImpl InitMenuRepoImpl
func InitMenuRepoImpl(db *sql.DB) MenuRepository {
	return &MenuRepoImpl{db}
}
