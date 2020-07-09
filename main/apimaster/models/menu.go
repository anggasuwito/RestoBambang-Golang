package models

//Menu Menu
type Menu struct {
	IDMenu      string `json:"id_menu"`
	TanggalMenu string `json:"tanggal_menu"`
	JenisMenu   string `json:"jenis_menu"`
	NamaMenu    string `json:"nama_menu"`
	HargaMenu   string `json:"harga_menu"`
	StokMenu    string `json:"stok_menu"`
}

//Jenis Jenis
type Jenis struct {
	IDJenis   string `json:"id_jenis"`
	NamaJenis string `json:"nama_jenis"`
}
