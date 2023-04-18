package client

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func SkipLegacyBgpIfUnsupported(t *testing.T) {
	if !IsLegacyBgpSupported() {
		t.Skip()
	}
}

func IsLegacyBgpSupported() bool {

	// Load .env File if exists
	envFile, _ := godotenv.Read("../.env")

	// Get Legacy BGP Support
	legacyBgpSupported := os.Getenv("LEGACY_BGP_SUPPORT")

	// If OS ENV Not Set
	if legacyBgpSupported == "" {

		// Try .env
		legacyBgpSupported = envFile["LEGACY_BGP_SUPPORT"]
	}

	// If Lecacy BGP Supported
	if legacyBgpSupported == "true" {

		// Return true
		return true
	}

	// Return false
	return false
}

func SkipIpAddressV6IfUnsupported(t *testing.T) {
	if !IsIpAddressV6Supported() {
		t.Skip()
	}
}

func IsIpAddressV6Supported() bool {

	// Load .env File if exists
	envFile, _ := godotenv.Read("../.env")

	// Get IPV6 Supported
	ipV6Supported := os.Getenv("IP_ADDRESS_V6_SUPPORT")

	// If OS ENV Not Set
	if ipV6Supported == "" {

		// Try .env
		ipV6Supported = envFile["IP_ADDRESS_V6_SUPPORT"]
	}

	// If IP V6 Supported
	if ipV6Supported == "true" {

		// Return true
		return true
	}

	// Return false
	return false
}
