package utils

import (
	"time"
	"fmt"
)

func GetYesterdayDate() string {
	// Mendapatkan tanggal dan waktu saat ini
	now := time.Now()

	// Mengurangi satu hari dari tanggal saat ini
	yesterday := now.AddDate(0, 0, -1)

	// Mengembalikan tanggal dalam format YYYY-MM-DD
	return yesterday.Format("2006-01-02")
}

func GetFormattedTodayDate() string {
	// Mendapatkan tanggal dan waktu saat ini
	now := time.Now()

	// Mendefinisikan nama hari dalam bahasa Indonesia
	days := []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}

	// Mendefinisikan nama bulan dalam bahasa Indonesia
	months := []string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	// Mendapatkan nama hari
	dayName := days[now.Weekday()]

	// Mendapatkan nama bulan
	monthName := months[now.Month()-1]

	// Mengembalikan tanggal dalam format yang diinginkan
	return fmt.Sprintf("%s, %d %s %d", dayName, now.Day(), monthName, now.Year())
}