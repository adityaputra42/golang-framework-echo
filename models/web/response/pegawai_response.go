package response

type PegawaiResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}
