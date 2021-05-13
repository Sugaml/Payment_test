package models

import (
	"01cloud-payment/internal/util"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomPaymentSetting(t *testing.T) PaymentSetting {
	arg := PaymentSetting{
		UserID:      uint(util.RandomUserID(1, 100)),
		Country:     util.RandomCountry(),
		State:       util.RandomState(),
		City:        util.RandomCity(),
		Street:      util.RandomStreet(),
		Postal_Code: strconv.Itoa((util.RandomUserID(1000, 9900))),
	}
	paymentsetting, err := arg.Save(TestDB)
	require.NoError(t, err)
	require.NotEmpty(t, paymentsetting)
	require.Equal(t, arg.UserID, paymentsetting.UserID)
	require.Equal(t, arg.Country, paymentsetting.Country)
	require.Equal(t, arg.State, paymentsetting.State)
	require.Equal(t, arg.City, paymentsetting.City)
	require.Equal(t, arg.Street, paymentsetting.Street)
	require.Equal(t, arg.Postal_Code, paymentsetting.Postal_Code)
	require.NotZero(t, paymentsetting.ID)
	require.NotZero(t, paymentsetting.CreatedAt)
	return *paymentsetting
}
func TestSave(t *testing.T) {
	createRandomPaymentSetting(t)
}

func TestFind(t *testing.T) {
	paymentsetting1 := createRandomPaymentSetting(t)

	paymentsetting2, err := paymentsetting1.Find(TestDB, uint64(paymentsetting1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, paymentsetting2)
	require.Equal(t, paymentsetting1.ID, paymentsetting2.ID)
	require.Equal(t, paymentsetting1.Country, paymentsetting2.Country)
	require.Equal(t, paymentsetting1.State, paymentsetting2.State)
	require.Equal(t, paymentsetting1.City, paymentsetting2.City)
	require.Equal(t, paymentsetting1.Street, paymentsetting2.Street)
	require.Equal(t, paymentsetting1.Postal_Code, paymentsetting2.Postal_Code)
	require.WithinDuration(t, paymentsetting1.CreatedAt, paymentsetting2.CreatedAt, time.Second)
}

func TestUpdate(t *testing.T) {
	paymentsetting1 := createRandomPaymentSetting(t)
	paymentsetting0, err := paymentsetting1.Find(TestDB, uint64(paymentsetting1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, paymentsetting0)
	fmt.Print(paymentsetting0.Country)
	paymentsetting0.Country = util.RandomCountry()
	fmt.Print(paymentsetting0.Country)
	paymentsetting2, err := paymentsetting0.Update(TestDB)
	require.NoError(t, err)
	require.NotEmpty(t, paymentsetting2)
	require.Equal(t, paymentsetting0.Country, paymentsetting2.Country)
	require.Equal(t, paymentsetting0.State, paymentsetting2.State)
	require.Equal(t, paymentsetting0.City, paymentsetting2.City)
	require.Equal(t, paymentsetting0.Street, paymentsetting2.Street)
	require.Equal(t, paymentsetting0.Postal_Code, paymentsetting2.Postal_Code)
	require.WithinDuration(t, paymentsetting1.CreatedAt, paymentsetting2.CreatedAt, time.Second)
}
