package client

import "log"

/**
 * Define IPSec Identity Structure
 */
type IpSecIdentity struct {
	Id                  string `mikrotik:".id"`
	Peer                string `mikrotik:"peer"`
	AuthMethod          string `mikrotik:"auth-method"`
	Secret              string `mikrotik:"secret"`
	Username            string `mikrotik:"username"`
	Password            string `mikrotik:"password"`
	EapMethods          string `mikrotik:"eap-methods"`
	Certificate         string `mikrotik:"certificate"`
	RemoteCertificate   string `mikrotik:"remote-certificate"`
	Key                 string `mikrotik:"key"`
	RemoteKey           string `mikrotik:"remote-key"`
	PolicyTemplateGroup string `mikrotik:"policy-template-group"`
	NoTrackChain        string `mikrotik:"notrack-chain"`
	MyId                string `mikrotik:"my-id"`
	RemoteId            string `mikrotik:"remote-id"`
	ModeConfig          string `mikrotik:"mode-config"`
	GeneratePolicy      string `mikrotik:"generate-policy"`
	Comment             string `mikrotik:"comment"`
	Disabled            bool   `mikrotik:"disabled"`
}

/**
 * Function used to ADD IPSec Identity on Mikrotik Router
 */
func (client Mikrotik) AddIpSecIdentity(peer *IpSecIdentity) (*IpSecIdentity, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running ADD IPSec Identity")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/identity/add", peer)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Identity ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Extract the ID From Response
	id := r.Done.Map["ret"]

	// Find and Return peer by ID
	return client.FindIpSecIdentity(id)
}

/**
 * Function used to List IPSec Identity from Mikrotik Router
 */
func (client Mikrotik) ListIpSecIdentity() ([]IpSecIdentity, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running List IPSec Identitys")

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/identity/print"}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Identity List response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Identity Array Pointer
	ipSecIdentitys := []IpSecIdentity{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecIdentitys)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Return result
	return ipSecIdentitys, nil
}

/**
 * Function used to FIND IPSec Identity by ID on Mikrotik Router
 */
func (client Mikrotik) FindIpSecIdentity(id string) (*IpSecIdentity, error) {

	// Log Command to be Run
	log.Printf("[INFO] Running Find IPSec Identity: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/identity/print", "?.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Identity Find response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Instantiate IPSec Identity Pointer
	ipSecIdentity := IpSecIdentity{}

	// Unmarshall Response
	err = Unmarshal(*r, &ipSecIdentity)

	// If There is Error (Unmarshalling)
	if err != nil {

		// Return Error
		return nil, err
	}

	// If MyID is Empty
	if ipSecIdentity.MyId == "" {

		// Initialize with AUTO
		ipSecIdentity.MyId = "auto"
	}

	// If RemoteId is Empty
	if ipSecIdentity.RemoteId == "" {

		// Initialize with AUTO
		ipSecIdentity.RemoteId = "auto"
	}

	// Return result
	return &ipSecIdentity, nil
}

/**
 * Function used to UPDATE IPSec Identity on Mikrotik Router
 */
func (client Mikrotik) UpdateIpSecIdentity(peer *IpSecIdentity) (*IpSecIdentity, error) {

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Generate Mikrotik Command
	cmd := Marshal("/ip/ipsec/identity/set", peer)

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	r, err := c.RunArgs(cmd)

	// Log Command execution result
	log.Printf("[INFO] IPSec Identity ADD response: `%v`", r)

	// If There is Error (Command Processing)
	if err != nil {

		// Return Error
		return nil, err
	}

	// Find and Return peer by ID
	return client.FindIpSecIdentity(peer.Id)
}

/**
 * Function used to DELETE IPSec Identity by ID on Mikrotik Router
 */
func (client Mikrotik) DeleteIpSecIdentity(id string) error {

	// Log Command to be Run
	log.Printf("[INFO] Running Delete IPSec Identity: `ID : %s`", id)

	// Retrieve Mikrotik Client
	c, err := client.getMikrotikClient()

	// If There is Error (Client Retrieving)
	if err != nil {

		// Return Error
		return err
	}

	// Generate Mikrotik Command
	cmd := []string{"/ip/ipsec/identity/remove", "=.id=" + id}

	// Log Command to be Run
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)

	// Run Command
	_, err = c.RunArgs(cmd)

	// Return Error if exists
	return err
}
