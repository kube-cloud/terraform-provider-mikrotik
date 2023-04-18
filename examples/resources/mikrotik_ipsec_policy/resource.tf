# Define Proposal
resource "mikrotik_ipsec_proposal" "proposal" {
  name = "ipsec-proposal"
  auth_algorithms = "sha512,sha256,sha1"
  enc_algorithms = "aes-256-cbc,aes-192-cbc,aes-128-cbc"
  lifetime = "30m"
  pfs_group = "modp2048"
  disabled = false
}

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

# Define IPSec Policy
resource "mikrotik_ipsec_policy" "policy" {
  peer = mikrotik_ipsec_peer.peer.name
  tunnel = true
  source_address = "172.20.0.0/16"
  source_port = 0
  destination_address = "10.20.0.0/16"
  destination_port = 0
  protocol = "all"
  template = false
  action = "encrypt"
  level = "require"
  ipsec_protocol = "esp"
  proposal = mikrotik_ipsec_proposal.proposal.name
  disabled = false
}