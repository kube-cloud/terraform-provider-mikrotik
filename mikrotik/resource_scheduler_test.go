package mikrotik

import (
	"fmt"
	"testing"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var origOnEvent string = "testing"
var origInterval int = 0
var updatedOnEvent string = "updated"
var updatedInterval int = 300

func TestAccMikrotikScheduler_create(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-acc-create")

	resourceName := "mikrotik_scheduler.bar"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckMikrotikSchedulerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccScheduler(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSchedulerExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "on_event"),
					resource.TestCheckResourceAttrSet(resourceName, "start_date"),
					resource.TestCheckResourceAttrSet(resourceName, "start_time"),
					resource.TestCheckResourceAttrSet(resourceName, "interval")),
			},
		},
	})
}

func TestAccMikrotikScheduler_updateInterval(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-acc-update")

	resourceName := "mikrotik_scheduler.bar"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckMikrotikSchedulerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccScheduler(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSchedulerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "interval", "0")),
			},
			{
				Config: testAccSchedulerUpdatedInterval(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSchedulerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "interval", "300")),
			},
		},
	})
}

func TestAccMikrotikScheduler_updatedOnEvent(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-acc-update-event")

	resourceName := "mikrotik_scheduler.bar"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckMikrotikSchedulerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccScheduler(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSchedulerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "on_event", origOnEvent)),
			},
			{
				Config: testAccSchedulerUpdatedOnEvent(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSchedulerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "on_event", updatedOnEvent)),
			},
		},
	})
}

func TestAccMikrotikScheduler_import(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-acc-import")

	resourceName := "mikrotik_scheduler.bar"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckMikrotikSchedulerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccScheduler(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSchedulerExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id")),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccScheduler(name string) string {
	return fmt.Sprintf(`
resource "mikrotik_scheduler" "bar" {
    name = "%s"
    on_event = "%s"
}
`, name, origOnEvent)
}

func testAccSchedulerUpdatedInterval(name string) string {
	return fmt.Sprintf(`
resource "mikrotik_scheduler" "bar" {
    name = "%s"
    on_event = "%s"
    interval = "%d"
}
`, name, origOnEvent, updatedInterval)
}

func testAccSchedulerUpdatedOnEvent(name string) string {
	return fmt.Sprintf(`
resource "mikrotik_scheduler" "bar" {
    name = "%s"
    on_event = "%s"
    interval = "%d"
}
`, name, updatedOnEvent, origInterval)
}

func testAccCheckMikrotikSchedulerDestroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_scheduler" {
			continue
		}

		scheduler, err := c.FindScheduler(rs.Primary.ID)

		_, ok := err.(*client.NotFound)
		if !ok && err != nil {
			return err
		}

		if scheduler != nil {
			return fmt.Errorf("scheduler (%s) still exists", scheduler.Name)
		}
	}
	return nil
}

func testAccSchedulerExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("mikrotik_scheduler does not exist in the statefile")
		}

		c := client.NewClient(client.GetConfigFromEnv())

		scheduler, err := c.FindScheduler(rs.Primary.ID)

		_, ok = err.(*client.NotFound)
		if !ok && err != nil {
			return fmt.Errorf("Unable to get the scheduler with error: %v", err)
		}

		if scheduler == nil {
			return fmt.Errorf("Unable to get the scheduler with name: %s", rs.Primary.ID)
		}

		if scheduler.Name == rs.Primary.ID {
			return nil
		}
		return nil
	}
}
