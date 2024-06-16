package utils

import (
	"fmt"
	"strings"
	"go-wa-bot-v3/structs"
)

func FormatUsers(parameter structs.Parameter) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(
		"*SETTING PARAMETER BSS*\n"+
			"*%s*\n\n"+
			"Forecast Total Produksi PV: *%.2f kWh*\n"+
			"Forecast Kebutuhan Smoothing BSS: *%.2f kWh*\n"+
			"Forecast Total Produksi PV + BSS: *%.2f kWh*\n"+
			"Kebutuhan %%DoD yang Dibutuhkan: *%.2f %%*\n"+
			"Setting C-rate: *0.2 C*\n"+
			"Setting Ramp Rate: *%.2f kW/s*\n"+
			"Setting Max Beban BSS: *%.2f kWh*\n",
		GetFormattedTodayDate(),
		parameter.ForecastProduksiPV, 
		parameter.ForecastSmoothing, 
		parameter.ForecastProduksiPVBSS, 
		parameter.KebutuhanDoD, 
		parameter.RampRate, 
		parameter.MaxBebanBSS,
	))

	return sb.String()
}
