package main

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestCustomerValid(t *testing.T){
	g := gomega.NewGomegaWithT(t)

	t.Run(`pass valid customer`, func(t *testing.T) {
		customer := Customer{
			Name:       "John Doe",
			Email:      "john.doe@example.com",
			CustomerID: "L1234567",
		}
		ok, err := govalidator.ValidateStruct(customer)
		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
	})
}