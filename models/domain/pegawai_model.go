package domain

type Pegawai struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}
