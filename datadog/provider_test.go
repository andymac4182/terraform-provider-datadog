package datadog

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"datadog": testAccProvider,
	}
}

func TestMain(m *testing.M) {
	acctest.UseBinaryDriver("datadog", Provider)
	resource.TestMain(m)
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("DATADOG_API_KEY"); v == "" {
		t.Fatal("DATADOG_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("DATADOG_APP_KEY"); v == "" {
		t.Fatal("DATADOG_APP_KEY must be set for acceptance tests")
	}

	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}
