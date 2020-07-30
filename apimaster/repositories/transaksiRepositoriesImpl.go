package repositories

import (
	"database/sql"
	"restoAPI/apimaster/models"

	"github.com/google/uuid"
)

//TransaksiRepoImpl TransaksiRepoImpl
type TransaksiRepoImpl struct {
	db *sql.DB
}

//GetAllTransaksi GetAllTransaksi
func (s TransaksiRepoImpl) GetAllTransaksi() ([]*models.Transaksi, error) {
	dataTransaksi := []*models.Transaksi{}
	query := `select t.transaksi_id, t.tanggal,t.menu_id, m.jenis_menu, m.nama_menu,m.harga_menu,t.quantity,m.harga_menu*t.quantity as total
	from transaksi t join menu m on m.menu_id = t.menu_id`
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		transaksi := models.Transaksi{}
		var err = data.Scan(&transaksi.IDTransaksi, &transaksi.TanggalTransaksi, &transaksi.IDMenu, &transaksi.JenisMenu, &transaksi.NamaMenu, &transaksi.HargaMenu, &transaksi.Quantity, &transaksi.TotalHarga)
		if err != nil {
			return nil, err
		}
		dataTransaksi = append(dataTransaksi, &transaksi)
	}
	return dataTransaksi, nil
}

//AddTransaksi InsertTransaksiData
func (s TransaksiRepoImpl) AddTransaksi(newTransaksi models.Transaksi) error {
	newTransaksiID := uuid.New().String()
	tr, err := s.db.Begin()
	query := `insert into transaksi(transaksi_id,tanggal,menu_id,quantity)values(?,?,?,?)`
	_, err = s.db.Query(query, newTransaksiID, &newTransaksi.TanggalTransaksi, &newTransaksi.IDMenu, &newTransaksi.Quantity)
	if err != nil {
		tr.Rollback()
		return err
	}
	tr.Commit()

	return nil
}

//GetTransaksiByID GetAllTransaksiById
func (s TransaksiRepoImpl) GetTransaksiByID(id string) (models.Transaksi, error) {
	var transaksi models.Transaksi

	query := `select t.transaksi_id, t.tanggal,t.menu_id ,m.jenis_menu, m.nama_menu,m.harga_menu,t.quantity,m.harga_menu*t.quantity as total
	from transaksi t join menu m on m.menu_id = t.menu_id
	where t.transaksi_id = ?`
	err := s.db.QueryRow(query, id).Scan(&transaksi.IDTransaksi, &transaksi.TanggalTransaksi, &transaksi.IDMenu, &transaksi.JenisMenu, &transaksi.NamaMenu, &transaksi.HargaMenu, &transaksi.Quantity, &transaksi.TotalHarga)

	if err != nil {
		return transaksi, err
	}

	return transaksi, nil
}

// UpdateTransaksiByID UpdateTransaksiData
func (s TransaksiRepoImpl) UpdateTransaksiByID(id string, changeTransaksi models.Transaksi) error {
	tr, err := s.db.Begin()
	_, err = s.GetTransaksiByID(id)

	if err != nil {
		tr.Rollback()
		return err
	}

	query := "update transaksi set tanggal = ? ,menu_id = ? ,quantity = ? where transaksi_id=?"
	_, err = s.db.Query(query, &changeTransaksi.TanggalTransaksi, &changeTransaksi.IDMenu, &changeTransaksi.Quantity, id)
	if err != nil {
		tr.Rollback()
		return err
	}
	tr.Commit()

	return nil
}

//DeleteDataTransaksiByID DeleteDataTransaksiById
func (s TransaksiRepoImpl) DeleteDataTransaksiByID(id string) error {
	tr, err := s.db.Begin()
	_, err = s.GetTransaksiByID(id)

	if err != nil {
		tr.Rollback()
		return err
	}
	query := "delete from transaksi where transaksi_id = ?"
	_, err = s.db.Query(query, id)
	if err != nil {
		tr.Rollback()
		return err
	}
	tr.Commit()

	return nil
}

//InitTransaksiRepoImpl InitTransaksiRepoImpl
func InitTransaksiRepoImpl(db *sql.DB) TransaksiRepository {
	return &TransaksiRepoImpl{db}
}
