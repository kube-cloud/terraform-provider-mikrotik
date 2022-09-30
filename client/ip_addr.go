package client

import (
	"fmt"
	"log"
)

type IpAddress struct {
	Id        string `mikrotik:".id"`
	Address   string `mikrotik:"address"`
	Comment   string `mikrotik:"comment"`
	Disabled  bool   `mikrotik:"disabled"`
	Interface string `mikrotik:"interface"`
	Network   string `mikrotik:"network"`
}

func (client Mikrotik) AddIpAddress(addr *IpAddress) (*IpAddress, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/ip/address/add", addr)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip address creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	id := r.Done.Map["ret"]

	return client.FindIpAddress(id)
}

func (client Mikrotik) ListIpAddress() ([]IpAddress, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/ip/address/print"}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] found ip address: %v", r)

	ipaddr := []IpAddress{}

	err = Unmarshal(*r, &ipaddr)

	if err != nil {
		return nil, err
	}

	return ipaddr, nil
}

func (client Mikrotik) FindIpAddress(id string) (*IpAddress, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/ip/address/print", "?.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip address response: %v", r)

	if err != nil {
		return nil, err
	}

	ipaddr := IpAddress{}
	err = Unmarshal(*r, &ipaddr)

	if err != nil {
		return nil, err
	}

	if ipaddr.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("ip address `%s` not found", id))
	}

	return &ipaddr, nil
}

func (client Mikrotik) UpdateIpAddress(addr *IpAddress) (*IpAddress, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/ip/address/set", addr)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return client.FindIpAddress(addr.Id)
}

func (client Mikrotik) DeleteIpAddress(id string) error {
	c, err := client.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{"/ip/address/remove", "=.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)
	return err
}
