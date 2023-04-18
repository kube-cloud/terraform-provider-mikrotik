# Define IPSec Policy group
resource "mikrotik_ipsec_policy_group" "group" {
  name = "ipsec-policy-group"
}