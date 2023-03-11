# The ID argument (*17) is a MikroTik's internal id.
# It can be obtained via CLI:
#
# [admin@MikroTik] /interface bridge port> :put
# *19
terraform import mikrotik_bridge_interface_port.bridge_port <bridge_port_id>
