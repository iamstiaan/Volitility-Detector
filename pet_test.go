// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package whatsauto_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/whatsauto-go"
	"github.com/stainless-sdks/whatsauto-go/internal/testutil"
	"github.com/stainless-sdks/whatsauto-go/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.New(context.TODO(), whatsauto.PetNewParams{
		Pet: whatsauto.PetParam{
			Name:      whatsauto.F("doggie"),
			PhotoURLs: whatsauto.F([]string{"string"}),
			ID:        whatsauto.F(int64(10)),
			Category: whatsauto.F(whatsauto.PetCategoryParam{
				ID:   whatsauto.F(int64(1)),
				Name: whatsauto.F("Dogs"),
			}),
			Status: whatsauto.F(whatsauto.PetStatusAvailable),
			Tags: whatsauto.F([]whatsauto.PetTagParam{{
				ID:   whatsauto.F(int64(0)),
				Name: whatsauto.F("name"),
			}}),
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

func TestPetGet(t *testing.T) {
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
	_, err := client.Pets.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.Update(context.TODO(), whatsauto.PetUpdateParams{
		Pet: whatsauto.PetParam{
			Name:      whatsauto.F("doggie"),
			PhotoURLs: whatsauto.F([]string{"string"}),
			ID:        whatsauto.F(int64(10)),
			Category: whatsauto.F(whatsauto.PetCategoryParam{
				ID:   whatsauto.F(int64(1)),
				Name: whatsauto.F("Dogs"),
			}),
			Status: whatsauto.F(whatsauto.PetStatusAvailable),
			Tags: whatsauto.F([]whatsauto.PetTagParam{{
				ID:   whatsauto.F(int64(0)),
				Name: whatsauto.F("name"),
			}}),
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

func TestPetDelete(t *testing.T) {
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
	err := client.Pets.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByStatus(context.TODO(), whatsauto.PetFindByStatusParams{
		Status: whatsauto.F(whatsauto.PetFindByStatusParamsStatusAvailable),
	})
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByTags(context.TODO(), whatsauto.PetFindByTagsParams{
		Tags: whatsauto.F([]string{"string"}),
	})
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateByIDWithOptionalParams(t *testing.T) {
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
	err := client.Pets.UpdateByID(
		context.TODO(),
		int64(0),
		whatsauto.PetUpdateByIDParams{
			Name:   whatsauto.F("name"),
			Status: whatsauto.F("status"),
		},
	)
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.UploadImage(
		context.TODO(),
		int64(0),
		whatsauto.PetUploadImageParams{
			Image:              io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			AdditionalMetadata: whatsauto.F("additionalMetadata"),
		},
	)
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
