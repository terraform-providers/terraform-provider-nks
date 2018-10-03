package stackpoint

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceKeysets_lookup(t *testing.T) {
	orgID := os.Getenv("SPC_ORG_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDataSourceStackPointKeysets_lookup, orgID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.stackpoint_keysets.keyset-default", "org_id", orgID),
				),
			},
		},
	})
}

const testAccDataSourceStackPointKeysets_lookup = `
data "stackpoint_keysets" "keyset-default" {
  org_id = "%s"
}
`
