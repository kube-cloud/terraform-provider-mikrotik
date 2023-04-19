# Define Firewall Nat
resource "mikrotik_firewall_nat" "nat" {

  # Filter Chain : srcnat, dstnat, etc...
  chain = "srcnat"

  # Rule Disabled Flag
  disabled = false
}