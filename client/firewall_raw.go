package client

import (
	"fmt"
	"log"
)

/**
 * Define Firewall Raw Structure
 */
type FirewallRaw struct {
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
	SourceAddressList      string `mikrotik:"src-address-list"`
	DestinationAddressList string `mikrotik:"dst-address-list"`
	SourceMacAddress       string `mikrotik:"src-mac-address"`
	IpSecPolicy            string `mikrotik:"ipsec-policy"`
	Action                 string `mikrotik:"action"`
	Log                    bool   `mikrotik:"log"`
	LogPrefix              string `mikrotik:"log-prefix"`
	Disabled               bool   `mikrotik:"disabled"`
}

/**
 * Function used to ADD Firewall Raw on Mikrotik Router
 */
func (client Mikrotik) AddFirewallRaw(firewallrule *FirewallRaw) (*FirewallRaw, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD Firewall Raw")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/firewall/raw/add", firewallrule)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Raw ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return firewallrule by ID
	return client.FindFirewallRaw(id)
}

/**
 * Function used to List Firewall Raw from Mikrotik Router
 */
func (client Mikrotik) ListFirewallRaw() ([]FirewallRaw, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List Firewall Raws")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/raw/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Raw List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate Firewall Raw Array Pointer
	firewallRaws := []FirewallRaw{}

	// Unmarshall Response
	err = Unmarshal(*r, &firewallRaws)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return firewallRaws, nil
}

/**
 * Function used to FIND Firewall Raw by ID on Mikrotik Router
 */
func (client Mikrotik) FindFirewallRaw(id string) (*FirewallRaw, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find Firewall Raw: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/raw/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Raw Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate Firewall Raw Pointer
	firewallRaw := FirewallRaw{}

	// Unmarshall Response
	err = Unmarshal(*r, &firewallRaw)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If Resource Name Empty (the Not found)
	if firewallRaw.Id == "" {

		// Return Not Found Error
		return nil, NewNotFound(fmt.Sprintf("Firewall Raw `%s` not found", id))
	}

	// Return result
	return &firewallRaw, nil
}

/**
 * Function used to UPDATE Firewall Raw on Mikrotik Router
 */
func (client Mikrotik) UpdateFirewallRaw(firewallrule *FirewallRaw) (*FirewallRaw, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/firewall/raw/set", firewallrule)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] Firewall Raw ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return firewallrule by ID
	return client.FindFirewallRaw(firewallrule.Id)
}

/**
 * Function used to DELETE Firewall Raw by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteFirewallRaw(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete Firewall Raw: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/firewall/raw/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
