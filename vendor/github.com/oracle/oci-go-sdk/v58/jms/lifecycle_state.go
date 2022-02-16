// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateActive   LifecycleStateEnum = "ACTIVE"
	LifecycleStateCreating LifecycleStateEnum = "CREATING"
	LifecycleStateDeleted  LifecycleStateEnum = "DELETED"
	LifecycleStateDeleting LifecycleStateEnum = "DELETING"
	LifecycleStateFailed   LifecycleStateEnum = "FAILED"
	LifecycleStateUpdating LifecycleStateEnum = "UPDATING"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"ACTIVE":   LifecycleStateActive,
	"CREATING": LifecycleStateCreating,
	"DELETED":  LifecycleStateDeleted,
	"DELETING": LifecycleStateDeleting,
	"FAILED":   LifecycleStateFailed,
	"UPDATING": LifecycleStateUpdating,
}

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
	values := make([]LifecycleStateEnum, 0)
	for _, v := range mappingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleStateEnumStringValues Enumerates the set of values in String for LifecycleStateEnum
func GetLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStateEnum(val string) (LifecycleStateEnum, bool) {
	mappingLifecycleStateEnumIgnoreCase := make(map[string]LifecycleStateEnum)
	for k, v := range mappingLifecycleStateEnum {
		mappingLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
