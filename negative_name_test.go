package main

import (
	"testing"

	"github.com/asaskevich/govalidator"
	 . "github.com/onsi/gomega"

)

func TestNameCustomer(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`fail name is empty`, func(t *testing.T) {
		customer := Customer{
			Name : "",
			Email: "john.doe@example.com",
			CustomerID: "L1234567",
		}

		ok,err := govalidator.ValidateStruct(customer)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name is required"))
	})
	
}