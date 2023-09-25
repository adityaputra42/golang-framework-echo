package request

type CreatePegawai struct {
	Nama    string `validate:"required" json:"name"`
	Alamat  string `validate:"required" json:"alamat"`
	Telepon string `validate:"required" json:"telepon"`
}
