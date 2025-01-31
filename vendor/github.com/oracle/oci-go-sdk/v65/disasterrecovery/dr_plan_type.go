// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"strings"
)

// DrPlanTypeEnum Enum with underlying type: string
type DrPlanTypeEnum string

// Set of constants representing the allowable values for DrPlanTypeEnum
const (
	DrPlanTypeSwitchover DrPlanTypeEnum = "SWITCHOVER"
	DrPlanTypeFailover   DrPlanTypeEnum = "FAILOVER"
)

var mappingDrPlanTypeEnum = map[string]DrPlanTypeEnum{
	"SWITCHOVER": DrPlanTypeSwitchover,
	"FAILOVER":   DrPlanTypeFailover,
}

var mappingDrPlanTypeEnumLowerCase = map[string]DrPlanTypeEnum{
	"switchover": DrPlanTypeSwitchover,
	"failover":   DrPlanTypeFailover,
}

// GetDrPlanTypeEnumValues Enumerates the set of values for DrPlanTypeEnum
func GetDrPlanTypeEnumValues() []DrPlanTypeEnum {
	values := make([]DrPlanTypeEnum, 0)
	for _, v := range mappingDrPlanTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanTypeEnumStringValues Enumerates the set of values in String for DrPlanTypeEnum
func GetDrPlanTypeEnumStringValues() []string {
	return []string{
		"SWITCHOVER",
		"FAILOVER",
	}
}

// GetMappingDrPlanTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanTypeEnum(val string) (DrPlanTypeEnum, bool) {
	enum, ok := mappingDrPlanTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
