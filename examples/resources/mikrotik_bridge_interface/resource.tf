resource "mikrotik_bridge_interface" "bridge_resurce_name" {
  mtu = 1500
  name ="bridge_name"
  disabled = false
  auto_mac = true
  admin_mac = "74:4D:28:F3:A7:14"
  comment = "bridge description"
}
