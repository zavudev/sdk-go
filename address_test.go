// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/zavudev/sdk-go"
	"github.com/zavudev/sdk-go/internal/testutil"
	"github.com/zavudev/sdk-go/option"
)

func TestAddressNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := zavudev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Addresses.New(context.TODO(), zavudev.AddressNewParams{
		CountryCode:        "DE",
		Locality:           "Berlin",
		PostalCode:         "10115",
		StreetAddress:      "123 Main St",
		AdministrativeArea: zavudev.String("administrativeArea"),
		BusinessName:       zavudev.String("businessName"),
		ExtendedAddress:    zavudev.String("extendedAddress"),
		FirstName:          zavudev.String("John"),
		LastName:           zavudev.String("Doe"),
	})
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAddressGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := zavudev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Addresses.Get(context.TODO(), "addressId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAddressListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := zavudev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Addresses.List(context.TODO(), zavudev.AddressListParams{
		Cursor: zavudev.String("cursor"),
		Limit:  zavudev.Int(100),
	})
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAddressDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := zavudev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Addresses.Delete(context.TODO(), "addressId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
