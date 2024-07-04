package assitence

import (
	"github.com/Bravoezz/infra_srv_tic/models"
	"github.com/Bravoezz/infra_srv_tic/utils"
	"net/http"
	"strconv"
)

type AssistanceController struct {
	asSrv *AssistanceService
}

func NewAssistanceController(asSrv *AssistanceService) *AssistanceController {
	return &AssistanceController{asSrv}
}

func (ctr *AssistanceController) ListAstc(w http.ResponseWriter, r *http.Request) {
	asstList, err := ctr.asSrv.List()
	if err != nil {
		utils.WriteErr(w, 500, err)
		return
	}

	utils.WriteJSON(w, 200, asstList)
	return
}

func (ctr *AssistanceController) FindById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)
	foundAss, err := ctr.asSrv.FindById(idInt)

	if err != nil {
		utils.WriteErr(w, 500, err)
		return
	}

	utils.WriteJSON(w, 200, foundAss)
	return
}

func (ctr AssistanceController) Insert(w http.ResponseWriter, r *http.Request) {
	assiBody := new(models.Assistance)
	err := utils.ParseJSON(r, assiBody)
	if err != nil {
		utils.WriteErr(w, http.StatusForbidden, err)
		return
	}

	found, err := ctr.asSrv.Insert(assiBody)
	if err != nil {
		utils.WriteErr(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, 200, found)
	return
}
