package helpers

import (
	"fmt"
	"math/rand"
)

func GetFormattedAccountNumber(cust_id int) string {
	bankCode := 1234
	randValue := rand.Intn(10000)
	return fmt.Sprintf("%d-%04d-%06d", bankCode, randValue, cust_id)
}
