package client

import "log"

/**
 * Define IPSec Proposal Structure
 */
type IpSecProposal struct {
	Id             string `mikrotik:".id"`
	Name           string `mikrotik:"name"`
	AuthAlgorithms string `mikrotik:"auth-algorithms"`
	EncAlgorithms  string `mikrotik:"enc-algorithms"`
	Lifetime       string `mikrotik:"lifetime"`
	PfsGroup       string `mikrotik:"pfs-group"`
	Disabled       bool   `mikrotik:"disabled"`
}

/**
 * Function used to ADD IPSec Proposal on Mikrotik Router
 */
func (client Mikrotik) AddIpSecProposal(proposal *IpSecProposal) (*IpSecProposal, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD IPSec Proposal")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/proposal/add", proposal)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Proposal ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return Proposal by ID
	return client.FindIpSecProposal(id)
}

/**
 * Function used to List IPSec Proposal from Mikrotik Router
 */
func (client Mikrotik) ListIpSecProposal() ([]IpSecProposal, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List IPSec Proposals")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/proposal/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Proposal List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Proposal Array Pointer
	ipSecProposals := []IpSecProposal{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecProposals)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return ipSecProposals, nil
}

/**
 * Function used to FIND IPSec Proposal by ID on Mikrotik Router
 */
func (client Mikrotik) FindIpSecProposal(id string) (*IpSecProposal, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find IPSec Proposal: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/proposal/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Proposal Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Proposal Pointer
	ipSecProposal := IpSecProposal{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecProposal)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return &ipSecProposal, nil
}

/**
 * Function used to UPDATE IPSec Proposal on Mikrotik Router
 */
func (client Mikrotik) UpdateIpSecProposal(proposal *IpSecProposal) (*IpSecProposal, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/proposal/set", proposal)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Proposal ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return Proposal by ID
	return client.FindIpSecProposal(proposal.Id)
}

/**
 * Function used to DELETE IPSec Proposal by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteIpSecProposal(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete IPSec Proposal: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/proposal/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
