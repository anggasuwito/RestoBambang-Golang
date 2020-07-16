package models

//Menu Menu
type Menu struct {
	IDMenu    string `json:"id_menu"`
	JenisMenu string `json:"jenis_menu"`
	NamaMenu  string `json:"nama_menu"`
	HargaMenu string `json:"harga_menu"`
	StokMenu  string `json:"stok_menu"`
}
