package dto

type IrsTransactionReq struct {
	ID         string `json:"id" query:"id"`
	PIN        string `json:"pin" query:"pin"`
	User       string `json:"user" query:"user"`
	Pass       string `json:"pass" query:"pass"`
	KodeProduk string `json:"kodeproduk" query:"kodeproduk"`
	Tujuan     string `json:"tujuan" query:"tujuan"`
	Counter    int    `json:"counter" query:"counter"`
	IDTrx      string `json:"idtrx" query:"idtrx"`
	Jenis      int    `json:"jenis" query:"jenis"`
}

type IrsTransactionRes struct {
	Success bool   `json:"success"`
	Produk  string `json:"produk"`
	Tujuan  string `json:"tujuan"`
	ReffID  string `json:"reffid"`
	RC      string `json:"rc"`
	Msg     string `json:"msg"`
}
