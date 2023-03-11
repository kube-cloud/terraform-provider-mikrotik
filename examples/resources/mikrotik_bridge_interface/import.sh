# The ID argument (*17) is a MikroTik's internal id.
# It can be obtained via CLI:
#
# [admin@MikroTik] /interface bridge> :put [ find where name=bridge-name]
# *17
terraform import mikrotik_bridge_interface.bridge <bridge_name>
