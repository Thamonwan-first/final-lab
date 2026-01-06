package main

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestCustomerInvalid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run(`fail customer is invalid`, func(t *testing.T) {
		customer := Customer{
			Name:       "john doe",
			Email:      "john.doe@example.com",
			CustomerID: "X1234567",
		}
		ok, err := govalidator.ValidateStruct(customer)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal(fmt.Sprintf("CustomerID: %s does not validate as matches(^[LMH]\\d{7}$)", customer.CustomerID)))

	})
}
