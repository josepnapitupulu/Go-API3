package models

import (
	"github.com/jinzhu/gorm"
	"github.com/josepnapitupulu/API_Jadwal_Minggu/config"
	)

var db *gorm.DB

type Jadwal struct {
	gorm.Model
	NamaMinggu string `json:"nama_minggu"`
	TanggalIbadah string `json:"tanggal_ibadah"`
	WaktuIbadah string `json:"waktu_ibadah"`
	Penkotbah string `json:"penkotbah"`
	Shift string `json:"shift"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Jadwal{})
}

func (b *Jadwal) CreateJadwal() *Jadwal {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllJadwals() []Jadwal {
	var Jadwals []Jadwal
	db.Find(&Jadwals)
	return Jadwals
}

func GetJadwalbyId(nik_anak int64) (*Jadwal, *gorm.DB) {
	var getJadwal Jadwal
	db := db.Where("ID=?", nik_anak).Find(&getJadwal)
	return &getJadwal, db
}

func DeleteJadwal(nik_anak int64) Jadwal {
	var jadwal Jadwal
	db.Where("ID=?", nik_anak).Delete(jadwal)
	return jadwal 
}