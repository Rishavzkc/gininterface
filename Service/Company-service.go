package Service

import (
	"gininterface/Config"
	"gininterface/Model"
)

type CompanyService interface {
	GetCompany() []Model.Company
	CreateCompany(Model.Company) Model.Company
	UpdateCompany(Model.Company) Model.Company
	DeleteCompany(Model.Company) Model.Company
}
type companyService struct {
	companies []Model.Company
}

func (service *companyService) GetCompany() []Model.Company {
	return service.companies
}

func (service *companyService) CreateCompany(company Model.Company) Model.Company {

	service.companies = append(service.companies, company)

	Config.Database.Create(&company)
	return company
}
func (service *companyService) UpdateCompany(company Model.Company) Model.Company {
	service.companies = append(service.companies, company)

	return company
}

func (service *companyService) DeleteCompany(company Model.Company) Model.Company {

	service.companies = append(service.companies, company)
	return company

}

func New() CompanyService {
	return &companyService{}
}
