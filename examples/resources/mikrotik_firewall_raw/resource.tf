# Define Firewall Raw
resource "mikrotik_firewall_raw" "raw" {

  # Filter Chain : output, prerouting, etc...
  chain = "output"

  # Rule Disabled Flag
  disabled = false
}