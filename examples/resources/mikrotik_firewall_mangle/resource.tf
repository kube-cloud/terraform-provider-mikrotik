# Define Firewall Mangle
resource "mikrotik_firewall_mangle" "mangle" {

  # Filter Chain : forward, input, output, postrouting, prerouting, etc...
  chain = "input"

  # Rule Disabled Flag
  disabled = false
}