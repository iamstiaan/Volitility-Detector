// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package whatsauto_test

import (
	"context"
	"os"
	"testing"

	"github.com/iamstiaan/Volitility-Detector"
	"github.com/iamstiaan/Volitility-Detector/internal/testutil"
	"github.com/iamstiaan/Volitility-Detector/option"
)

func TestUsage(t *testing.T) {
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
	order, err := client.Store.NewOrder(context.TODO(), whatsauto.StoreNewOrderParams{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", order.ID)
}
