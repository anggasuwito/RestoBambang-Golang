package repositories

import (
	"database/sql"
	"gomux/main/master/models"
)

//TransaksiRepoImpl TransaksiRepoImpl
type TransaksiRepoImpl struct {
	db *sql.DB
}

//GetAllTransaksi GetAllTransaksi
func (s TransaksiRepoImpl) GetAllTransaksi() ([]*models.Transaksi, error) {
	dataTransaksi := []*models.Transaksi{}
	query := `select t.id_transaksi,t.tanggal,j.nama_jenis,m.nama_menu,hm.harga,em.nama_ekstra_menu,em.harga_ekstra_menu, hm.harga + em.harga_ekstra_menu as total from transaksi t 
	join list_menu lm on lm.id_list_menu = t.id_list_menu
	join jenis j on j.id_jenis = lm.id_jenis
	join menu m on m.id_menu = lm.id_menu
	join harga_menu hm on hm.id_harga_menu = lm.id_harga_menu
	join ekstra_menu em on em.id_ekstra = t.id_ekstra_menu`
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		transaksi := models.Transaksi{}
		var err = data.Scan(&transaksi.IDTransaksi, &transaksi.TanggalTransaksi, &transaksi.JenisMenu, &transaksi.NamaMenu, &transaksi.HargaMenu, &transaksi.NamaEkstraMenu, &transaksi.HargaEkstraMenu, &transaksi.TotalHarga)
		if err != nil {
			return nil, err
		}
		dataTransaksi = append(dataTransaksi, &transaksi)
	}
	return dataTransaksi, nil
}

// //AddTransaksi InsertTransaksiData
// func (s TransaksiRepoImpl) AddTransaksi(newTransaksi models.Transaksi) error {

// 	tr, err := s.db.Begin()
// 	query := `insert into list_transaksi(id_list_transaksi,tanggal,id_jenis,id_transaksi,id_harga_transaksi,id_stok_transaksi)values(?,?,?,?,?,?)`
// 	_, err = s.db.Query(query, &newTransaksi.IDTransaksi, &newTransaksi.TanggalTransaksi, &newTransaksi.JenisMenu, &newTransaksi.NamaMenu, &newTransaksi.HargaMenu, &newTransaksi.NamaEkstraMenu, &newTransaksi.HargaEkstraMenu, &newTransaksi.TotalHarga)
// 	if err != nil {
// 		tr.Rollback()
// 		log.Fatal(err)
// 	} else {
// 		tr.Commit()
// 	}

// 	return nil
// }

// //GetTransaksiByID GetAllTransaksiById
// func (s TransaksiRepoImpl) GetTransaksiByID(id string) (models.Transaksi, error) {
// 	var transaksi models.Transaksi
// 	query := `select l.id_list_transaksi,l.tanggal,j.nama_jenis,m.nama_transaksi,h.harga,s.stok
// 	from list_transaksi l
// 	join jenis j on l.id_jenis = j.id_jenis
// 	join transaksi m on m.id_transaksi = l.id_transaksi
// 	join harga_transaksi h on h.id_transaksi = l.id_transaksi
// 	join stok_transaksi s on s.id_transaksi = l.id_transaksi
// 	where l.id_list_transaksi = ?`
// 	err := s.db.QueryRow(query, id).Scan(&transaksi.IDTransaksi, &transaksi.TanggalTransaksi, &transaksi.JenisMenu, &transaksi.NamaMenu, &transaksi.HargaMenu, &transaksi.NamaEkstraMenu, &transaksi.HargaEkstraMenu, &transaksi.TotalHarga)
// 	if err != nil {
// 		return transaksi, err
// 	}

// 	return transaksi, nil
// }

// // UpdateTransaksiByID UpdateTransaksiData
// func (s TransaksiRepoImpl) UpdateTransaksiByID(id string, changeTransaksi models.Transaksi) error {
// 	tr, err := s.db.Begin()
// 	_, err = s.GetTransaksiByID(id)

// 	if err != nil {
// 		tr.Rollback()
// 		return err
// 	}

// 	query := "update list_transaksi set id_list_transaksi = ? ,tanggal = ? ,id_jenis = ? ,id_transaksi = ?,id_harga_transaksi = ?,id_stok_transaksi = ? where id_list_transaksi=?"
// 	_, err = s.db.Query(query, &changeTransaksi.IDTransaksi, &changeTransaksi.TanggalTransaksi, &changeTransaksi.JenisMenu, &changeTransaksi.NamaMenu, &changeTransaksi.HargaMenu, &changeTransaksi.NamaEkstraMenu, &changeTransaksi.HargaEkstraMenu, &changeTransaksi.TotalHarga)
// 	if err != nil {
// 		tr.Rollback()
// 		return err
// 	}
// 	tr.Commit()

// 	return nil
// }

// //DeleteDataTransaksiByID DeleteDataTransaksiById
// func (s TransaksiRepoImpl) DeleteDataTransaksiByID(id string) error {
// 	tr, err := s.db.Begin()
// 	_, err = s.GetTransaksiByID(id)

// 	if err != nil {
// 		tr.Rollback()
// 		return err
// 	}
// 	query := "delete from list_transaksi where transaksi_id = ?"
// 	_, err = s.db.Query(query, id)
// 	if err != nil {
// 		tr.Rollback()
// 		log.Fatal(err)
// 	} else {
// 		tr.Commit()
// 	}

// 	return nil
// }

//InitTransaksiRepoImpl InitTransaksiRepoImpl
func InitTransaksiRepoImpl(db *sql.DB) TransaksiRepository {
	return &TransaksiRepoImpl{db}
}
