package main

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"

)

func NameCustomer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run(`fail name is empty`, func(t *testing.T) {
		customer := Customer{
			Name : "",
			Email: "john.doe@example.com",
			CustomerID: "L1234567",
		}

		ok,err := govalidator.ValidateStruct(customer)
		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Name is required"))
	})
	
}