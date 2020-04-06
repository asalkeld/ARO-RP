package azureclient

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import "testing"

func TestGetCodeFromType(t *testing.T) {
	tests := []struct {
		_type string
		want  string
	}{
		{
			_type: "Microsoft.Authorization/denyAssignments",
			want:  "authorization-denyassignment",
		},
		{
			_type: "Microsoft.Network/dnsZones/providers/roleAssignments",
			want:  "authorization",
		},
		{
			_type: "Microsoft.Network/privateDnsZones/SRV",
			want:  "privatedns",
		},
		{
			_type: "Microsoft.Authorization/roleDefinitions",
			want:  "authorization-roledefinition",
		},
		{
			_type: "Microsoft.Network/dnsZones",
			want:  "dns",
		},
		{
			_type: "Microsoft.ManagedIdentity/userAssignedIdentities",
			want:  "msi",
		},
		{
			_type: "Microsoft.Network/networkSecurityGroups",
			want:  "network",
		},
		{
			_type: "Microsoft.Compute/virtualMachines",
			want:  "compute",
		},
	}
	for _, tt := range tests {
		t.Run(tt._type, func(t *testing.T) {
			if got := GetCodeFromType(tt._type); got != tt.want {
				t.Errorf("GetCodeFromType() = %v, want %v", got, tt.want)
			}
		})
	}
}
