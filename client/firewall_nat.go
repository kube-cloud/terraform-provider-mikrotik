package client

import (
	"fmt"
	"log"
)

/**
 * Define Firewall Nat Structure
 */
type FirewallNat struct {
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
}

/**
 * Function used to ADD Firewall Nat on Mikrotik Router
 */
func (client Mikrotik) AddFirewallNat(firewallrule *FirewallNat) (*FirewallNat, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD Firewall Nat")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/firewall/nat/add", firewallrule)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Nat ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return firewallrule by ID
	return client.FindFirewallNat(id)
}

/**
 * Function used to List Firewall Nat from Mikrotik Router
 */
func (client Mikrotik) ListFirewallNat() ([]FirewallNat, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List Firewall Nats")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/nat/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Nat List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate Firewall Nat Array Pointer
	firewallNats := []FirewallNat{}

	// Unmarshall Response
	err = Unmarshal(*r, &firewallNats)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return firewallNats, nil
}

/**
 * Function used to FIND Firewall Nat by ID on Mikrotik Router
 */
func (client Mikrotik) FindFirewallNat(id string) (*FirewallNat, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find Firewall Nat: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/nat/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Nat Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate Firewall Nat Pointer
	firewallNat := FirewallNat{}

	// Unmarshall Response
	err = Unmarshal(*r, &firewallNat)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If Resource Name Empty (the Not found)
	if firewallNat.Id == "" {

		// Return Not Found Error
		return nil, NewNotFound(fmt.Sprintf("Firewall Nat `%s` not found", id))
	}

	// Return result
	return &firewallNat, nil
}

/**
 * Function used to UPDATE Firewall Nat on Mikrotik Router
 */
func (client Mikrotik) UpdateFirewallNat(firewallrule *FirewallNat) (*FirewallNat, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/firewall/nat/set", firewallrule)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Nat ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return firewallrule by ID
	return client.FindFirewallNat(firewallrule.Id)
}

/**
 * Function used to DELETE Firewall Nat by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteFirewallNat(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete Firewall Nat: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/nat/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
