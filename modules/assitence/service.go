package assitence

import "github.com/Bravoezz/infra_srv_tic/models"

type AssistanceService struct {
	assRepo *AssistanceRepository
}

func NewAssistanceService(asR *AssistanceRepository) *AssistanceService {
	return &AssistanceService{asR}
}

func (srv *AssistanceService) List() (*[]models.Assistance, error) {
	return srv.assRepo.List()
}

func (srv AssistanceService) FindById(id int) (*models.Assistance, error) {
	return srv.assRepo.FindById(id)
}

func (srv AssistanceService) Insert(assistance *models.Assistance) (*models.Assistance, error) {
	return srv.assRepo.Insert(assistance)
}
