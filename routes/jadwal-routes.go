package routes

import (
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Jadwal_Minggu/controllers"
)

var RegisterJadwalsRoutes = func(router *mux.Router) {
	// router.HandleFunc("/", controllers.Index).Methods("GET")
	// router.HandleFunc("/jemaatBaru", controllers.Create).Methods("POST")
	router.HandleFunc("/jadwal/", controllers.CreateJadwal).Methods("POST")
	router.HandleFunc("/jadwal/", controllers.GetJadwal).Methods("GET")
	router.HandleFunc("/jadwal/{jadwalId}", controllers.GetJadwalById).Methods("GET")
	router.HandleFunc("/jadwal/{jadwalId}", controllers.UpdateJadwal).Methods("PUT")
	router.HandleFunc("/jadwal/{jadwalId}", controllers.DeleteJadwal).Methods("DELETE")
}
