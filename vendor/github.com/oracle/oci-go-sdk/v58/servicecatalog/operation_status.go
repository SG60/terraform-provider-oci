// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"strings"
)

// OperationStatusEnum Enum with underlying type: string
type OperationStatusEnum string

// Set of constants representing the allowable values for OperationStatusEnum
const (
	OperationStatusAccepted  OperationStatusEnum = "ACCEPTED"
	OperationStatusFailed    OperationStatusEnum = "FAILED"
	OperationStatusSucceeded OperationStatusEnum = "SUCCEEDED"
)

var mappingOperationStatusEnum = map[string]OperationStatusEnum{
	"ACCEPTED":  OperationStatusAccepted,
	"FAILED":    OperationStatusFailed,
	"SUCCEEDED": OperationStatusSucceeded,
}

// GetOperationStatusEnumValues Enumerates the set of values for OperationStatusEnum
func GetOperationStatusEnumValues() []OperationStatusEnum {
	values := make([]OperationStatusEnum, 0)
	for _, v := range mappingOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationStatusEnumStringValues Enumerates the set of values in String for OperationStatusEnum
func GetOperationStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationStatusEnum(val string) (OperationStatusEnum, bool) {
	mappingOperationStatusEnumIgnoreCase := make(map[string]OperationStatusEnum)
	for k, v := range mappingOperationStatusEnum {
		mappingOperationStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperationStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
