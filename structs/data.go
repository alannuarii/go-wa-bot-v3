package structs

type IrradianceData struct {
	Irradiance float64 `json:"irradiance"`
	Waktu      string  `json:"waktu"`
}

type Data struct {
	Avg []IrradianceData `json:"avg"`
	Max []IrradianceData `json:"max"`
}

type Response struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Parameter struct {
	ForecastProduksiPV		float64
	ForecastProduksiPVBSS	float64
	ForecastSmoothing		float64
	KebutuhanDoD			float64
	RampRate				float64
	MaxBebanBSS				float64
}