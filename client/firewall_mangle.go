package client

import (
	"fmt"
	"log"
)

/**
 * Define Firewall Mangle Structure
 */
type FirewallMangle struct {
	Id                     string `mikrotik:".id"`
	Chain                  string `mikrotik:"chain"`
	SourceAddress          string `mikrotik:"src-address"`
	DestinationAddress     string `mikrotik:"dst-address"`
	SourcePort             int    `mikrotik:"src-port"`
	DestinationPort        int    `mikrotik:"dst-port"`
	AnyPort                int    `mikrotik:"port"`
	Protocol               string `mikrotik:"protocol"`
	InInterface            string `mikrotik:"in-interface"`
	OutInterface           string `mikrotik:"out-interface"`
	InInterfaceList        string `mikrotik:"in-interface-list"`
	OutInterfaceList       string `mikrotik:"out-interface-list"`
	PacketMark             string `mikrotik:"packet-mark"`
	ConnectionMark         string `mikrotik:"connection-mark"`
	RoutingMark            string `mikrotik:"routing-mark"`
	RoutingTable           string `mikrotik:"routing-table"`
	ConnectionType         string `mikrotik:"connection-type"`
	SourceAddressList      string `mikrotik:"src-address-list"`
	DestinationAddressList string `mikrotik:"dst-address-list"`
	Layer7Protocol         string `mikrotik:"layer7-protocol"`
	SourceMacAddress       string `mikrotik:"src-mac-address"`
	IpSecPolicy            string `mikrotik:"ipsec-policy"`
	InBridgePort           string `mikrotik:"in-bridge-port"`
	OutBridgePort          string `mikrotik:"out-bridge-port"`
	InBridgePortList       string `mikrotik:"in-bridge-port-list"`
	OutBridgePortList      string `mikrotik:"out-bridge-port-list"`
	Action                 string `mikrotik:"action"`
	Log                    bool   `mikrotik:"log"`
	LogPrefix              string `mikrotik:"log-prefix"`
	Disabled               bool   `mikrotik:"disabled"`
	ConnectionState        string `mikrotik:"connection-state"`
	ConnectionNatState     string `mikrotik:"connection-nat-state"`
	TcpFlags               string `mikrotik:"tcp-flags"`
}

/**
 * Function used to ADD Firewall Mangle on Mikrotik Router
 */
func (client Mikrotik) AddFirewallMangle(firewallrule *FirewallMangle) (*FirewallMangle, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD Firewall Mangle")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/firewall/mangle/add", firewallrule)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Mangle ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return firewallrule by ID
	return client.FindFirewallMangle(id)
}

/**
 * Function used to List Firewall Mangle from Mikrotik Router
 */
func (client Mikrotik) ListFirewallMangle() ([]FirewallMangle, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List Firewall Mangles")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/mangle/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Mangle List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate Firewall Mangle Array Pointer
	firewallMangles := []FirewallMangle{}

	// Unmarshall Response
	err = Unmarshal(*r, &firewallMangles)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return firewallMangles, nil
}

/**
 * Function used to FIND Firewall Mangle by ID on Mikrotik Router
 */
func (client Mikrotik) FindFirewallMangle(id string) (*FirewallMangle, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find Firewall Mangle: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/mangle/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Mangle Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate Firewall Mangle Pointer
	firewallMangle := FirewallMangle{}

	// Unmarshall Response
	err = Unmarshal(*r, &firewallMangle)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If Resource Name Empty (the Not found)
	if firewallMangle.Id == "" {

		// Return Not Found Error
		return nil, NewNotFound(fmt.Sprintf("Firewall Mangle `%s` not found", id))
	}

	// Return result
	return &firewallMangle, nil
}

/**
 * Function used to UPDATE Firewall Mangle on Mikrotik Router
 */
func (client Mikrotik) UpdateFirewallMangle(firewallrule *FirewallMangle) (*FirewallMangle, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/firewall/mangle/set", firewallrule)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Mangle ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return firewallrule by ID
	return client.FindFirewallMangle(firewallrule.Id)
}

/**
 * Function used to DELETE Firewall Mangle by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteFirewallMangle(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete Firewall Mangle: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/mangle/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
