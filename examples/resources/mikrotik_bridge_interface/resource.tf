resource "mikrotik_bridge_interface" "bridge_resurce_name" {
  mtu = 1500
  name ="bridge_name"
  disabled = false
  auto_mac = true
  admin_mac = false
  comment = "bridge description"
}
