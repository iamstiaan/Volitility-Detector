// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package whatsauto_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/iamstiaan/Volitility-Detector"
	"github.com/iamstiaan/Volitility-Detector/internal/testutil"
	"github.com/iamstiaan/Volitility-Detector/option"
)

func TestUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.New(context.TODO(), whatsauto.UserNewParams{
		User: whatsauto.UserParam{
			ID:         whatsauto.F(int64(10)),
			Email:      whatsauto.F("john@email.com"),
			FirstName:  whatsauto.F("John"),
			LastName:   whatsauto.F("James"),
			Password:   whatsauto.F("12345"),
			Phone:      whatsauto.F("12345"),
			Username:   whatsauto.F("theUser"),
			UserStatus: whatsauto.F(int64(1)),
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

func TestUserGet(t *testing.T) {
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
	_, err := client.User.Get(context.TODO(), "username")
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.User.Update(
		context.TODO(),
		"username",
		whatsauto.UserUpdateParams{
			User: whatsauto.UserParam{
				ID:         whatsauto.F(int64(10)),
				Email:      whatsauto.F("john@email.com"),
				FirstName:  whatsauto.F("John"),
				LastName:   whatsauto.F("James"),
				Password:   whatsauto.F("12345"),
				Phone:      whatsauto.F("12345"),
				Username:   whatsauto.F("theUser"),
				UserStatus: whatsauto.F(int64(1)),
			},
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

func TestUserDelete(t *testing.T) {
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
	err := client.User.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithListWithOptionalParams(t *testing.T) {
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
	_, err := client.User.NewWithList(context.TODO(), whatsauto.UserNewWithListParams{
		Items: []whatsauto.UserParam{{
			ID:         whatsauto.F(int64(10)),
			Email:      whatsauto.F("john@email.com"),
			FirstName:  whatsauto.F("John"),
			LastName:   whatsauto.F("James"),
			Password:   whatsauto.F("12345"),
			Phone:      whatsauto.F("12345"),
			Username:   whatsauto.F("theUser"),
			UserStatus: whatsauto.F(int64(1)),
		}},
	})
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Login(context.TODO(), whatsauto.UserLoginParams{
		Password: whatsauto.F("password"),
		Username: whatsauto.F("username"),
	})
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
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
	err := client.User.Logout(context.TODO())
	if err != nil {
		var apierr *whatsauto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
