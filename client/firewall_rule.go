package client

import (
	"fmt"
	"log"
)

/**
 * Define Firewall Rule Structure
 */
type FirewallRule struct {
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
 * Function used to ADD Firewall Rule on Mikrotik Router
 */
func (client Mikrotik) AddFirewallRule(firewallrule *FirewallRule) (*FirewallRule, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD Firewall Rule")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/firewall/filter/add", firewallrule)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Rule ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return firewallrule by ID
	return client.FindFirewallRule(id)
}

/**
 * Function used to List Firewall Rule from Mikrotik Router
 */
func (client Mikrotik) ListFirewallRule() ([]FirewallRule, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List Firewall Rules")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/filter/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Rule List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate Firewall Rule Array Pointer
	firewallRules := []FirewallRule{}

	// Unmarshall Response
	err = Unmarshal(*r, &firewallRules)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return firewallRules, nil
}

/**
 * Function used to FIND Firewall Rule by ID on Mikrotik Router
 */
func (client Mikrotik) FindFirewallRule(id string) (*FirewallRule, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find Firewall Rule: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/filter/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Rule Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate Firewall Rule Pointer
	firewallRule := FirewallRule{}

	// Unmarshall Response
	err = Unmarshal(*r, &firewallRule)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If Resource Name Empty (the Not found)
	if firewallRule.Id == "" {

		// Return Not Found Error
		return nil, NewNotFound(fmt.Sprintf("Firewall Rule `%s` not found", id))
	}

	// Return result
	return &firewallRule, nil
}

/**
 * Function used to UPDATE Firewall Rule on Mikrotik Router
 */
func (client Mikrotik) UpdateFirewallRule(firewallrule *FirewallRule) (*FirewallRule, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/firewall/filter/set", firewallrule)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Rule ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return firewallrule by ID
	return client.FindFirewallRule(firewallrule.Id)
}

/**
 * Function used to DELETE Firewall Rule by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteFirewallRule(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete Firewall Rule: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/filter/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
