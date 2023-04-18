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