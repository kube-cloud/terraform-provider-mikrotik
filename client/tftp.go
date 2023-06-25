package client

import (
	"fmt"
	"log"
)

/**
 * Define TFTP Structure
 */
type Tftp struct {
	Id              string `mikrotik:".id"`
	IpAddresses     string `mikrotik:"ip-addresses"`
	RequestFileName string `mikrotik:"req-filename"`
	RealFileName    string `mikrotik:"real-filename"`
	Allow           bool   `mikrotik:"allow"`
	ReadOnly        bool   `mikrotik:"read-only"`
	Disabled        bool   `mikrotik:"disabled"`
	Comment         string `mikrotik:"comment"`
}

/**
 * Function used to ADD TFTP Server on Mikrotik Router
 */
func (client Mikrotik) AddTftp(tftp *Tftp) (*Tftp, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD TFTP Server")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/tftp/add", tftp)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] TFTP Server ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return tftp by ID
	return client.FindTftp(id)
}

/**
 * Function used to List TFTP Server from Mikrotik Router
 */
func (client Mikrotik) ListTftp() ([]Tftp, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List TFTP Servers")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/tftp/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] TFTP Server List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate TFTP Server Array Pointer
	ipSecPolicies := []Tftp{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecPolicies)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return ipSecPolicies, nil
}

/**
 * Function used to FIND TFTP Server by ID on Mikrotik Router
 */
func (client Mikrotik) FindTftp(id string) (*Tftp, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find TFTP Server: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/tftp/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] TFTP Server Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate TFTP Server Pointer
	tftp := Tftp{}

	// Unmarshall Response
	err = Unmarshal(*r, &tftp)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If Resource Name Empty (the Not found)
	if tftp.Id == "" {

		// Return Not Found Error
		return nil, NewNotFound(fmt.Sprintf("TFTP Server `%s` not found", id))
	}

	// Return result
	return &tftp, nil
}

/**
 * Function used to UPDATE TFTP Server on Mikrotik Router
 */
func (client Mikrotik) UpdateTftp(tftp *Tftp) (*Tftp, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/tftp/set", tftp)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] TFTP Server ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return tftp by ID
	return client.FindTftp(tftp.Id)
}

/**
 * Function used to DELETE TFTP Server by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteTftp(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete TFTP Server: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/tftp/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
