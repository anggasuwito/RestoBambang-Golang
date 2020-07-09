package usecases

import "gomux/main/apimaster/models"

//TransaksiUseCase TransaksiUseCase
type TransaksiUseCase interface {
	GetAllTransaksi() ([]*models.Transaksi, error)
	//AddTransaksi(newTransaksi models.Transaksi) error
	// GetTransaksiByID(id string) (models.Transaksi, error)
	// UpdateTransaksiByID(id string, changeTransaksi models.Transaksi) error
	// DeleteTransaksiByID(id string) error
}
