package commodity

type Commodity struct {
	Komoditas    string  `json:"komoditas"`
	AreaProvinsi string  `json:"area_provinsi"`
	AreaKota     string  `json:"area_kota"`
	Size         string  `json:"size"`
	Price        string  `json:"price"`
	PriceUSD     float64 `json:"price_usd"`
	TglParsed    string  `json:"tgl_parsed"`
	Uuid         string  `json:"uuid"`
}
