# Define IPSec Profile
resource "mikrotik_ipsec_profile" "profile" {
  name = "ipsec-profile"
  dh_group ="modp2048"
  dpd_interval = "2m"
  dpd_max_failure = 5
  enc_algorithms = "aes-256,aes-192,aes-128"
  hash_algorithm = "sha1"
  lifetime = "1h30m"
  nat_traversal = true
  proposal_check = "obey"
}

# Define IPSec Policy group (Optional if Default Group exists)
resource "mikrotik_ipsec_policy_group" "group" {
  name = "ipsec-policy-group"
}

# Define IPSec Peer
resource "mikrotik_ipsec_peer" "peer" {
  name = "ipsec-peer"
  address = "192.16.2.14/32"
  profile = mikrotik_ipsec_profile.profile.name
  exchange_mode = "ike2"
  send_initial_contact = false
  passive = true
  local_address = "192.16.3.19"
  port = 0
}

# Define IPSec Identity
resource "mikrotik_ipsec_identity" "identity" {
  peer = mikrotik_ipsec_peer.peer.name
  auth_method ="pre-shared-key"
  secret ="iosxwcf13t6èèu"
  my_id = "auto"
  remote_id = "auto"
  generate_policy = "no"
  policy_template_group = mikrotik_ipsec_policy_group.group.name
  comment = "IPSec Identity"
  disabled = false
}