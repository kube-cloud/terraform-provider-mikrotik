package client

import (
	"fmt"
	"log"
)

// Bridge Interface Port Binding Resource
type BridgeInterfacePort struct {
	Id                    string `mikrotik:".id"`
	Bridge                string `mikrotik:"bridge"`
	Interface             string `mikrotik:"interface"`
	Horizon               string `mikrotik:"horizon"`
	Learn                 string `mikrotik:"learn"`
	UnknownMulticastFlood bool   `mikrotik:"unknown-multicast-flood"`
	UnknownUnicastFlood   bool   `mikrotik:"unknown-unicast-flood"`
	BroadcastFlood        bool   `mikrotik:"broadcast-flood"`
	Trusted               bool   `mikrotik:"trusted"`
	HardwareOffload       bool   `mikrotik:"hw"`
	AutoIsolate           bool   `mikrotik:"auto-isolate"`
	RestrictedRole        bool   `mikrotik:"restricted-role"`
	RestrictedTcn         bool   `mikrotik:"restricted-tcn"`
	BpduGuard             bool   `mikrotik:"bpdu-guard"`
	Priority              int    `mikrotik:"priority"`
	PathCost              int    `mikrotik:"path-cost"`
	InternalPathCost      int    `mikrotik:"internal-path-cost"`
	Edge                  string `mikrotik:"edge"`
	PointToPoint          string `mikrotik:"point-to-point"`
	Disabled              bool   `mikrotik:"disabled"`
	Comment               string `mikrotik:"comment"`
}

func (client Mikrotik) FindBridgeInterfacePort(iface string) (*BridgeInterfacePort, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/interface/bridge/port/print", "?interface=" + iface}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] Found record: %v", r)
	record := BridgeInterfacePort{}
	err = Unmarshal(*r, &record)

	if err != nil {
		return nil, err
	}

	if record.Interface == "" {
		return nil, NewNotFound(fmt.Sprintf("Bridge Interface Port `%s` not found", iface))
	}
	return &record, nil
}

func (client Mikrotik) AddBridgeInterfacePort(d *BridgeInterfacePort) (*BridgeInterfacePort, error) {
	c, err := client.getMikrotikClient()
	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/bridge/port/add", d)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] command returned: %v", r)

	return client.FindBridgeInterfacePort(d.Interface)
}

func (client Mikrotik) UpdateBridgeInterfacePort(d *BridgeInterfacePort) (*BridgeInterfacePort, error) {
	c, err := client.getMikrotikClient()
	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/bridge/port/set", d)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] command returned: %v", r)

	return client.FindBridgeInterfacePort(d.Interface)
}

func (client Mikrotik) DeleteBridgeInterfacePort(iface string) error {
	c, err := client.getMikrotikClient()
	if err != nil {
		return err
	}

	existingBridgeInterfaceport, err0 := client.FindBridgeInterfacePort(iface)
	if err0 != nil {
		return err0
	}

	cmd := []string{"/interface/bridge/port/remove", "=numbers=" + existingBridgeInterfaceport.Id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Command returned: %v", r)

	return nil
}
