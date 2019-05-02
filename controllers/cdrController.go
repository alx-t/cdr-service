package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetCDRJson(w http.ResponseWriter, r *http.Request) {
	// ps, err := project.Read(common.Db, "")
	// j, err := json.Marshal(ProjectsResource{Data: ps})
  //
	// if err != nil {
	// 	common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
	// 	return
	// }

  j := "Get CDR json"

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
