package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/lender", user)
	fmt.Println("web berjalan dengan port 8819")
	http.ListenAndServe(":8819", nil)
}

type tanifund struct {
	nama      string `json:"data_lender"`
	pekerjaan string `json:"pekerjaan_lender"`
}

var dataLender = []tanifund{
	{
		nama:      "febrina",
		pekerjaan: "pegawai swasta",
	},
	{
		nama:      "andayani",
		pekerjaan: "ibu rumah tangga",
	},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", " application/json")

	if r.Method == http.MethodGet {

		result, err := json.Marshal(dataLender)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
}
