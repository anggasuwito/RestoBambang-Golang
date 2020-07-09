package repositories

import "gomux/main/apimaster/models"

//TransaksiRepository TransaksiRepository
type TransaksiRepository interface {
	//AddTransaksi(newTransaksi models.Transaksi) error
	// GetTransaksiByID(id string) (models.Transaksi, error)
	// UpdateTransaksiByID(id string, changeTransaksi models.Transaksi) error
	// DeleteDataTransaksiByID(id string) error
	GetAllTransaksi() ([]*models.Transaksi, error)
}
