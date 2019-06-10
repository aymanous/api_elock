package apiService

import (
	"net/http"

	helperhttp "../../Helper/Http"
	"github.com/gorilla/mux"
)

func newHandler(model dao) *handler {
	var out handler
	out.model = model
	return &out
}

type handler struct {
	model dao
}

func (obj *handler) openLock(w http.ResponseWriter, r *http.Request) error {
	badgeID := mux.Vars(r)["id"]

	if !obj.model.OpenLock(badgeID) {
		return helperhttp.RespondContent(w, 500, false)
	}

	return helperhttp.RespondContent(w, 200, true)

}

func (obj *handler) addBadge(w http.ResponseWriter, r *http.Request) error {
	badgeID := mux.Vars(r)["id"]
	nom := r.URL.Query().Get("nom")
	prenom := r.URL.Query().Get("prenom")

	if !obj.model.AddBadge(badgeID, nom, prenom) {
		return helperhttp.RespondContent(w, 500, false)
	}

	return helperhttp.RespondContent(w, 200, true)
}

func (obj *handler) deleteBadge(w http.ResponseWriter, r *http.Request) error {
	badgeID := mux.Vars(r)["id"]
	return helperhttp.RespondContent(w, 200, obj.model.DeleteBadge(badgeID))
}

func (obj *handler) getServerAddress(w http.ResponseWriter, r *http.Request) error {
	return helperhttp.RespondContent(w, 200, obj.model.GetServerAddress())
}

func (obj *handler) changeMode(w http.ResponseWriter, r *http.Request) error {
	mode := mux.Vars(r)["mode"]

	if !obj.model.ChangeMode(mode) {
		return helperhttp.RespondContent(w, 500, "KO")
	}

	return helperhttp.RespondContent(w, 200, "OK")
}

func (obj *handler) getCurrentMode(w http.ResponseWriter, r *http.Request) error {
	var code int
	mode := obj.model.GetCurrentMode()

	switch mode {
	case "ADD":
		code = 201
	case "DELETE":
		code = 202
	default:
		code = 200 //READ
	}

	return helperhttp.RespondContent(w, code, mode)
}

func (obj *handler) getLastLog(w http.ResponseWriter, r *http.Request) error {
	return helperhttp.RespondContent(w, 200, obj.model.GetLastLog())
}

func (obj *handler) getBadgesList(w http.ResponseWriter, r *http.Request) error {
	return helperhttp.RespondContent(w, 200, obj.model.GetBadgesList())
}

func (obj *handler) setNomPrenom(w http.ResponseWriter, r *http.Request) error {
	code := mux.Vars(r)["code"]
	return helperhttp.RespondContent(w, 200, obj.model.SetNomPrenom(code, r.URL.Query().Get("prenom"), r.URL.Query().Get("nom")))
}
