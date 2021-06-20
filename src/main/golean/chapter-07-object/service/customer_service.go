package service

import "golang-learning/src/main/golean/chapter-07-object/model"

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 22, "1312220202", "giami@gimail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}
func (this CustomerService) List() []model.Customer {
	return this.customers

}
