resource "mikrotik_bridge_interface_port" "bridge_port" {
  interface = "interface_name"
  bridge = "bridge_name"
  horizon = "horizon"
  learn = true
  unknown_multicast_flood = true
  unknown_unicast_flood = true
  broadcast_flood = true
  trusted = true
  hardware_offload = true
  auto_isolate = true
  restricted_role = true
  restricted_tcn = true
  bpdu_guard = true
  path_cost = 10
  internal_path_cost = 10
  edge = "auto|yes|no|yes-discover|no-discover"
  point_to_point = "auto|yes|NO"
  disabled = false
  comment = "Bridge Description"
}