package service

import (
	model2 "golang-learning/chapter-07-object/model"
)

type CustomerService struct {
	customers   []model2.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model2.NewCustomer(1, "张三", "男", 22, "1312220202", "giami@gimail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}
func (this CustomerService) List() []model2.Customer {
	return this.customers

}
