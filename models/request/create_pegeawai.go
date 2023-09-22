package request

type CreatePegawai struct {
	Name    string `validate:"required" json:"name"`
	Alamat  string `validate:"required" json:"alamat"`
	Telepon string `validate:"required,max=15" json:"telepon"`
}
