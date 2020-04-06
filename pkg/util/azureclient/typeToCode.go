package azureclient

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import "strings"

// GetCodeFromType gets the code (that keys APIVersions) from a resource type
func GetCodeFromType(_type string) string {
	switch _type {
	case "Microsoft.Authorization/roleDefinitions":
		return "authorization-roledefinition"
	case "Microsoft.Network/dnsZones":
		return "dns"
	default:
		if strings.Contains(_type, "roleAssignments") {
			return "authorization"
		}
		if strings.Contains(_type, "denyAssignments") {
			return "authorization-denyassignment"
		}
		if strings.Contains(_type, "Microsoft.Network/privateDnsZones") {
			return "privatedns"
		}
	}
	provider := strings.Split(_type, "/")[0]
	switch provider {
	case "Microsoft.ManagedIdentity":
		return "msi"
	default:
		return strings.ToLower(strings.Replace(provider, "Microsoft.", "", 1))
	}
}
