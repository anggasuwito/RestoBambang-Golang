package models

//Transaksi Transaksi
type Transaksi struct {
	IDTransaksi      string `json:"id_transaksi"`
	TanggalTransaksi string `json:"tanggal_transaksi"`
	IDMenu           string `json:"id_menu"`
	JenisMenu        string `json:"jenis_menu"`
	NamaMenu         string `json:"nama_menu"`
	HargaMenu        string `json:"harga_menu"`
	Quantity         string `json:"quantity"`
	TotalHarga       string `json:"total_harga"`
}
