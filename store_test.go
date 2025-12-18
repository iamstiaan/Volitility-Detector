// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package whatsauto_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/whatsauto-go"
	"github.com/stainless-sdks/whatsauto-go/internal/testutil"
	"github.com/stainless-sdks/whatsauto-go/option"
	"github.com/stainless-sdks/whatsauto-go/shared"
)

func TestStoreNewOrderWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := whatsauto.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.NewOrder(context.TODO(), whatsauto.StoreNewOrderParams{
		Order: shared.OrderParam{
			ID:       whatsauto.F(int64(10)),
			Complete: whatsauto.F(true),
			PetID:    whatsauto.F(int64(198772)),
			Quantity: whatsauto.F(int64(7)),
			ShipDate: whatsauto.F(time.Now()),
			Status:   whatsauto.F(shared.OrderStatusApproved),
		},
	})
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreInventory(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := whatsauto.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.Inventory(context.TODO())
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
