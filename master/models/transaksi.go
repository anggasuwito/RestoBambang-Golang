package models

//Transaksi Transaksi
type Transaksi struct {
	IDTransaksi      string `json:"id_menu"`
	TanggalTransaksi string `json:"tanggal_menu"`
	JenisMenu        string `json:"jenis_menu"`
	NamaMenu         string `json:"nama_menu"`
	HargaMenu        string `json:"harga_menu"`
	NamaEkstraMenu   string `json:"nama_ekstra_menu"`
	HargaEkstraMenu  string `json:"harga_ekstra_menu"`
	TotalHarga       string `json:"total_harga"`
}
