package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Jadwal_Minggu/models"
	"github.com/josepnapitupulu/API_Jadwal_Minggu/utils"
)

var NewJadwal models.Jadwal
// var tmpl * template.Template
// func init(){
// 	tmpl = template.Must(template.ParseFiles("jemaat.html"))
// }

// func Index(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/jemaat.html")
// 	temp.Execute(w, nil)
// }

// func Create(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/createJemaat.html")
// 	temp.Execute(w, nil)
// }

func GetJadwal(w http.ResponseWriter, r *http.Request) {
	newJadwals := models.GetAllJadwals()
	res, _ := json.Marshal(newJadwals)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetJadwalById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jadwalId := vars["jadwalId"]
	ID, err := strconv.ParseInt(jadwalId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jadwalDetails, _ := models.GetJadwalbyId(ID)
	res, _ := json.Marshal(jadwalDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateJadwal(w http.ResponseWriter, r *http.Request) {
	CreateJadwal := &models.Jadwal{}
	utils.ParseBody(r, CreateJadwal)
	b := CreateJadwal.CreateJadwal()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteJadwal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jadwalId := vars["jadwalId"]
	ID, err := strconv.ParseInt(jadwalId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	jadwal := models.DeleteJadwal(ID)
	res, _ := json.Marshal(jadwal)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateJadwal(w http.ResponseWriter, r *http.Request) {
    var updateJadwal = &models.Jadwal{}
    utils.ParseBody(r, updateJadwal)
    vars := mux.Vars(r)
    jadwalId := vars["jadwalId"]
    ID, err := strconv.ParseInt(jadwalId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    jadwalDetails, db := models.GetJadwalbyId(ID)
    if updateJadwal.NamaMinggu != "" {
        jadwalDetails.NamaMinggu = updateJadwal.NamaMinggu
    }
    if updateJadwal.TanggalIbadah != "" {
        jadwalDetails.TanggalIbadah = updateJadwal.TanggalIbadah
    }
    if updateJadwal.WaktuIbadah != "" {
        jadwalDetails.WaktuIbadah = updateJadwal.WaktuIbadah
    }
    if updateJadwal.Penkotbah != "" {
        jadwalDetails.Penkotbah = updateJadwal.Penkotbah
    }
    if updateJadwal.Shift != "" {
        jadwalDetails.Shift = updateJadwal.Shift
    }
    db.Save(&jadwalDetails)
    res, _ := json.Marshal(jadwalDetails)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}