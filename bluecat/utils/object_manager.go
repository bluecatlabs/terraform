// Copyright 2020 BlueCat Networks. All rights reserved

package utils

import (
	"encoding/json"
	"fmt"
	"terraform-provider-bluecat/bluecat/entities"
	"terraform-provider-bluecat/bluecat/models"
)

// ObjectManager The BlueCat object manager
type ObjectManager struct {
	Connector BCConnector
}

// Host record

// CreateHostRecord Create the Host record
func (objMgr *ObjectManager) CreateHostRecord(configuration string, view string, zone string, absoluteName string, ip4Address string, ttl int, properties string) (*entities.HostRecord, error) {

	hostRecord := models.NewHostRecord(entities.HostRecord{
		Configuration: configuration,
		View:          view,
		Zone:          zone,
		IP4Address:    ip4Address,
		AbsoluteName:  absoluteName,
		TTL:           ttl,
		Properties:    properties,
	})

	_, err := objMgr.Connector.CreateObject(hostRecord)
	return hostRecord, err
}

// GetHostRecord Get the Host record
func (objMgr *ObjectManager) GetHostRecord(configuration string, view string, absoluteName string) (*entities.HostRecord, error) {

	hostRecord := models.HostRecord(entities.HostRecord{
		Configuration: configuration,
		View:          view,
		AbsoluteName:  absoluteName,
	})

	err := objMgr.Connector.GetObject(hostRecord, &hostRecord)
	return hostRecord, err
}

// UpdateHostRecord Update the Host record
func (objMgr *ObjectManager) UpdateHostRecord(configuration string, view string, zone string, absoluteName string, ip4Address string, ttl int, properties string) (*entities.HostRecord, error) {

	hostRecord := models.HostRecord(entities.HostRecord{
		Configuration: configuration,
		View:          view,
		Zone:          zone,
		IP4Address:    ip4Address,
		AbsoluteName:  absoluteName,
		TTL:           ttl,
		Properties:    properties,
	})

	err := objMgr.Connector.UpdateObject(hostRecord, &hostRecord)
	return hostRecord, err
}

// DeleteHostRecord Delete the Host record
func (objMgr *ObjectManager) DeleteHostRecord(configuration string, view string, absoluteName string) (string, error) {

	hostRecord := models.HostRecord(entities.HostRecord{
		Configuration: configuration,
		View:          view,
		AbsoluteName:  absoluteName,
	})

	return objMgr.Connector.DeleteObject(hostRecord)
}

// CNAME record

// CreateCNAMERecord Create the CNAME record
func (objMgr *ObjectManager) CreateCNAMERecord(configuration string, view string, zone string, absoluteName string, linkedRecord string, ttl int, properties string) (*entities.CNAMERecord, error) {

	cnameRecord := models.NewCNAMERecord(entities.CNAMERecord{
		Configuration: configuration,
		View:          view,
		Zone:          zone,
		LinkedRecord:  linkedRecord,
		AbsoluteName:  absoluteName,
		TTL:           ttl,
		Properties:    properties,
	})

	_, err := objMgr.Connector.CreateObject(cnameRecord)
	return cnameRecord, err
}

// GetCNAMERecord Get the CNAME record
func (objMgr *ObjectManager) GetCNAMERecord(configuration string, view string, absoluteName string) (*entities.CNAMERecord, error) {

	cnameRecord := models.CNAMERecord(entities.CNAMERecord{
		Configuration: configuration,
		View:          view,
		AbsoluteName:  absoluteName,
	})

	err := objMgr.Connector.GetObject(cnameRecord, &cnameRecord)
	return cnameRecord, err
}

// UpdateCNAMERecord Update the CNAME record
func (objMgr *ObjectManager) UpdateCNAMERecord(configuration string, view string, zone string, absoluteName string, linkedRecord string, ttl int, properties string) (*entities.CNAMERecord, error) {

	cnameRecord := models.CNAMERecord(entities.CNAMERecord{
		Configuration: configuration,
		View:          view,
		Zone:          zone,
		LinkedRecord:  linkedRecord,
		AbsoluteName:  absoluteName,
		TTL:           ttl,
		Properties:    properties,
	})

	err := objMgr.Connector.UpdateObject(cnameRecord, &cnameRecord)
	return cnameRecord, err
}

// DeleteCNAMERecord Delete the CNAME record
func (objMgr *ObjectManager) DeleteCNAMERecord(configuration string, view string, absoluteName string) (string, error) {

	cnameRecord := models.CNAMERecord(entities.CNAMERecord{
		Configuration: configuration,
		View:          view,
		AbsoluteName:  absoluteName,
	})

	return objMgr.Connector.DeleteObject(cnameRecord)
}

// Configuration

// CreateConfiguration Create a new Configuration
func (objMgr *ObjectManager) CreateConfiguration(name string, properties string) (*entities.Configuration, error) {

	configuration := models.NewConfiguration(entities.Configuration{
		Name:       name,
		Properties: properties,
	})

	_, err := objMgr.Connector.CreateObject(configuration)
	return configuration, err
}

// GetConfiguration Get the Configuration info
func (objMgr *ObjectManager) GetConfiguration(name string) (*entities.Configuration, error) {

	configuration := models.Configuration(entities.Configuration{
		Name: name,
	})

	err := objMgr.Connector.GetObject(configuration, &configuration)
	return configuration, err
}

// UpdateConfiguration Update the Configuration info
func (objMgr *ObjectManager) UpdateConfiguration(name string, properties string) (*entities.Configuration, error) {

	configuration := models.Configuration(entities.Configuration{
		Name:       name,
		Properties: properties,
	})

	err := objMgr.Connector.UpdateObject(configuration, &configuration)
	return configuration, err
}

// DeleteConfiguration Delete the configuration
func (objMgr *ObjectManager) DeleteConfiguration(name string) (string, error) {

	configuration := models.Configuration(entities.Configuration{
		Name: name,
	})

	return objMgr.Connector.DeleteObject(configuration)
}

// Block

// CreateBlock Create a new Block
func (objMgr *ObjectManager) CreateBlock(configuration string, name string, address string, cidr string, parentBlock string, properties string) (*entities.Block, error) {

	block := models.NewBlock(entities.Block{
		Configuration: configuration,
		Name:          name,
		Address:       address,
		CIDR:          cidr,
		ParentBlock:   parentBlock,
		Properties:    properties,
	})

	_, err := objMgr.Connector.CreateObject(block)
	return block, err
}

// GetBlock Get the Block info
func (objMgr *ObjectManager) GetBlock(configuration string, address string, cidr string) (*entities.Block, error) {

	block := models.Block(entities.Block{
		Configuration: configuration,
		Address:       address,
		CIDR:          cidr,
	})

	err := objMgr.Connector.GetObject(block, &block)
	return block, err
}

// UpdateBlock Update the Block info
func (objMgr *ObjectManager) UpdateBlock(configuration string, name string, address string, cidr string, parentBlock string, properties string) (*entities.Block, error) {

	block := models.Block(entities.Block{
		Configuration: configuration,
		Name:          name,
		Address:       address,
		CIDR:          cidr,
		ParentBlock:   parentBlock,
		Properties:    properties,
	})

	err := objMgr.Connector.UpdateObject(block, &block)
	return block, err
}

// DeleteBlock Delete the Block
func (objMgr *ObjectManager) DeleteBlock(configuration string, address string, cidr string) (string, error) {

	block := models.Block(entities.Block{
		Configuration: configuration,
		Address:       address,
		CIDR:          cidr,
	})

	return objMgr.Connector.DeleteObject(block)
}

// Network

func generateNetworkProperties(props string, gateway string) string {
	result := props
	if len(gateway) > 0 {
		result = fmt.Sprintf("%s|gateway=%s", result, gateway)
	}
	return result
}

// CreateNetwork Create a new Network
func (objMgr *ObjectManager) CreateNetwork(configuration string, block string, name string, cidr string, gateway string, properties string) (*entities.Network, error) {

	network := models.NewNetwork(entities.Network{
		Configuration: configuration,
		BlockAddr:     block,
		Name:          name,
		CIDR:          cidr,
		Gateway:       gateway,
		Properties:    generateNetworkProperties(properties, gateway),
	})
	_, err := objMgr.Connector.CreateObject(network)
	return network, err
}

// GetNetwork Get the Network info
func (objMgr *ObjectManager) GetNetwork(configuration string, cidr string) (*entities.Network, error) {

	network := models.Network(entities.Network{
		Configuration: configuration,
		CIDR:          cidr,
	})

	err := objMgr.Connector.GetObject(network, &network)
	return network, err
}

// UpdateNetwork Update the Network info
func (objMgr *ObjectManager) UpdateNetwork(configuration string, name string, cidr string, gateway string, properties string) (*entities.Network, error) {

	network := models.Network(entities.Network{
		Configuration: configuration,
		Name:          name,
		CIDR:          cidr,
		Properties:    generateNetworkProperties(properties, gateway),
	})

	err := objMgr.Connector.UpdateObject(network, &network)
	return network, err
}

// DeleteNetwork Delete the Network
func (objMgr *ObjectManager) DeleteNetwork(configuration string, cidr string) (string, error) {

	network := models.Network(entities.Network{
		Configuration: configuration,
		CIDR:          cidr,
	})

	return objMgr.Connector.DeleteObject(network)
}

// IP

// ReserveIPAddress Create the new IP address for later use
func (objMgr *ObjectManager) ReserveIPAddress(configuration string, network string) (*entities.IPAddress, error) {
	return objMgr.createIPAddress(configuration, network, "", "", "", models.AllocateReserved, "")
}

// CreateStaticIP Create the new static IP address
func (objMgr *ObjectManager) CreateStaticIP(configuration string, network string, address string, macAddress string, name string, properties string) (*entities.IPAddress, error) {
	return objMgr.createIPAddress(configuration, network, address, macAddress, name, models.AllocateStatic, properties)
}

// createIPAddress Create the new IP address. Allocate the next available on the network if IP address is not provided
func (objMgr *ObjectManager) createIPAddress(configuration string, cidr string, address string, macAddress string, name string, addrType string, properties string) (*entities.IPAddress, error) {
	addrEntity := entities.IPAddress{
		Configuration: configuration,
		CIDR:          cidr,
		Name:          name,
		Address:       address,
		Mac:           macAddress,
		Action:        addrType,
		Properties:    properties,
	}

	ipAddr := new(entities.IPAddress)
	if len(address) > 0 {
		ipAddr = models.IPAddress(addrEntity)
	} else {
		ipAddr = models.GetNextIPAddress(addrEntity)
		log.Debugf("Requesting the new IP address in the network %s", cidr)
	}
	res, err := objMgr.Connector.CreateObject(ipAddr)
	if err == nil {
		err = json.Unmarshal([]byte(res), &ipAddr)
		if err == nil {
			log.Debugf("Failed to decode the IP object %s", err)
		}
	}
	return ipAddr, err
}

// GetIPAddress Get the IP Address info
func (objMgr *ObjectManager) GetIPAddress(configuration string, address string) (*entities.IPAddress, error) {

	ipAddr := models.IPAddress(entities.IPAddress{
		Configuration: configuration,
		Address:       address,
	})

	err := objMgr.Connector.GetObject(ipAddr, &ipAddr)
	return ipAddr, err
}

// SetMACAddress Update the MAC address for the existing IP address
func (objMgr *ObjectManager) SetMACAddress(configuration string, address string, macAddress string) (*entities.IPAddress, error) {
	ipAddr := models.IPAddress(entities.IPAddress{
		Configuration: configuration,
		Address:       address,
		Mac:           macAddress,
	})
	err := objMgr.Connector.UpdateObject(ipAddr, &ipAddr)
	return ipAddr, err
}

// UpdateIPAddress Update the IP address info
func (objMgr *ObjectManager) UpdateIPAddress(configuration string, address string, macAddress string, name string, addrType string, properties string) (*entities.IPAddress, error) {
	ipAddr := models.IPAddress(entities.IPAddress{
		Configuration: configuration,
		Name:          name,
		Address:       address,
		Mac:           macAddress,
		Action:        addrType,
		Properties:    properties,
	})
	err := objMgr.Connector.UpdateObject(ipAddr, &ipAddr)
	return ipAddr, err
}

// DeleteIPAddress Delete the existing IP address
func (objMgr *ObjectManager) DeleteIPAddress(configuration string, address string) (string, error) {
	ipAddr := models.IPAddress(entities.IPAddress{
		Configuration: configuration,
		Address:       address,
	})
	return objMgr.Connector.DeleteObject(ipAddr)
}
