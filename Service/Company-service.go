package Service

import (
	"gininterface/Model"
)

type CompanyService interface {
	GetCompany(Model.Company)
}
type companyService struct {
	companies []Model.Company
}

func New() CompanyService {
	return &companyService{}
}

func (service *companyService) GetCompany(company Model.Company) {

}
