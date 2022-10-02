package client

import (
	"fmt"
	"log"
)

// BridgeInterface represents Bridge Interface Resource
type BridgeInterface struct {
	Id       string `mikrotik:".id"`
	Mtu      int    `mikrotik:"mtu`
	Name     string `mikrotik:"name"`
	Disabled bool   `mikrotik:"disabled"`
	AutoMac  bool   `mikrotik:"auto-mac"`
	AdminMac string `mikrotik:"admin-mac"`
	Comment  string `mikrotik:"comment"`
}

func (client Mikrotik) FindBridgeInterface(name string) (*BridgeInterface, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/interface/bridge/print", "?name=" + name}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	fmt.Println("============RUN BRIDGE REQUEST=============")
	fmt.Println(cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] Found record: %v", r)
	fmt.Println("=========================")
	fmt.Println("=====> RAW RESULT")
	fmt.Println(r)
	record := BridgeInterface{}
	err = Unmarshal(*r, &record)

	if err != nil {
		return nil, err
	}

	if record.Name == "" {
		return nil, NewNotFound(fmt.Sprintf("Bridge Interface `%s` not found", name))
	}
	fmt.Println("=========================")
	fmt.Println("=====> BRIDGE RESULT")
	fmt.Println(record)
	fmt.Println("=========================")
	fmt.Println("=========================")
	return &record, nil
}

func (client Mikrotik) AddBridgeInterface(d *BridgeInterface) (*BridgeInterface, error) {
	c, err := client.getMikrotikClient()
	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/bridge/add", d)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] command returned: %v", r)

	return client.FindBridgeInterface(d.Name)
}

func (client Mikrotik) UpdateBridgeInterface(d *BridgeInterface) (*BridgeInterface, error) {
	c, err := client.getMikrotikClient()
	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/bridge/set", d)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] command returned: %v", r)

	return client.FindBridgeInterface(d.Name)
}

func (client Mikrotik) DeleteBridgeInterface(name string) error {
	c, err := client.getMikrotikClient()
	if err != nil {
		return err
	}

	cmd := []string{"/interface/bridge/remove", "=numbers=" + name}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Command returned: %v", r)

	return nil
}
