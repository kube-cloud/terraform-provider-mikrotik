# Define Firewall Rule
resource "mikrotik_firewall_rule" "rule" {

  # Filter Chain : input, output, forward, etc...
  chain = "input"

  # Rule Disabled Flag
  disabled = false
}