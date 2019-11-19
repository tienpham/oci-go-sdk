// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDatabaseWallet The Autonomous Database wallet details.
type AutonomousDatabaseWallet struct {

	// The current lifecycle state of the Autonomous Database wallet.
	LifecycleState AutonomousDatabaseWalletLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the wallet was last rotated.
	TimeRotated *common.SDKTime `mandatory:"false" json:"timeRotated"`
}

func (m AutonomousDatabaseWallet) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseWalletLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseWalletLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseWalletLifecycleStateEnum
const (
	AutonomousDatabaseWalletLifecycleStateActive   AutonomousDatabaseWalletLifecycleStateEnum = "ACTIVE"
	AutonomousDatabaseWalletLifecycleStateUpdating AutonomousDatabaseWalletLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousDatabaseWalletLifecycleState = map[string]AutonomousDatabaseWalletLifecycleStateEnum{
	"ACTIVE":   AutonomousDatabaseWalletLifecycleStateActive,
	"UPDATING": AutonomousDatabaseWalletLifecycleStateUpdating,
}

// GetAutonomousDatabaseWalletLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseWalletLifecycleStateEnum
func GetAutonomousDatabaseWalletLifecycleStateEnumValues() []AutonomousDatabaseWalletLifecycleStateEnum {
	values := make([]AutonomousDatabaseWalletLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseWalletLifecycleState {
		values = append(values, v)
	}
	return values
}