# Define Proposal
resource "mikrotik_ipsec_proposal" "proposal" {
  name = "ipsec-proposal"
  auth_algorithms = "sha512,sha256,sha1"
  enc_algorithms = "aes-256-cbc,aes-192-cbc,aes-128-cbc"
  lifetime = "30m"
  pfs_group = "modp2048"
  disabled = false
}