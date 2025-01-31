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

// DrPlanExecutionTypeEnum Enum with underlying type: string
type DrPlanExecutionTypeEnum string

// Set of constants representing the allowable values for DrPlanExecutionTypeEnum
const (
	DrPlanExecutionTypeSwitchover         DrPlanExecutionTypeEnum = "SWITCHOVER"
	DrPlanExecutionTypeSwitchoverPrecheck DrPlanExecutionTypeEnum = "SWITCHOVER_PRECHECK"
	DrPlanExecutionTypeFailover           DrPlanExecutionTypeEnum = "FAILOVER"
	DrPlanExecutionTypeFailoverPrecheck   DrPlanExecutionTypeEnum = "FAILOVER_PRECHECK"
)

var mappingDrPlanExecutionTypeEnum = map[string]DrPlanExecutionTypeEnum{
	"SWITCHOVER":          DrPlanExecutionTypeSwitchover,
	"SWITCHOVER_PRECHECK": DrPlanExecutionTypeSwitchoverPrecheck,
	"FAILOVER":            DrPlanExecutionTypeFailover,
	"FAILOVER_PRECHECK":   DrPlanExecutionTypeFailoverPrecheck,
}

var mappingDrPlanExecutionTypeEnumLowerCase = map[string]DrPlanExecutionTypeEnum{
	"switchover":          DrPlanExecutionTypeSwitchover,
	"switchover_precheck": DrPlanExecutionTypeSwitchoverPrecheck,
	"failover":            DrPlanExecutionTypeFailover,
	"failover_precheck":   DrPlanExecutionTypeFailoverPrecheck,
}

// GetDrPlanExecutionTypeEnumValues Enumerates the set of values for DrPlanExecutionTypeEnum
func GetDrPlanExecutionTypeEnumValues() []DrPlanExecutionTypeEnum {
	values := make([]DrPlanExecutionTypeEnum, 0)
	for _, v := range mappingDrPlanExecutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanExecutionTypeEnumStringValues Enumerates the set of values in String for DrPlanExecutionTypeEnum
func GetDrPlanExecutionTypeEnumStringValues() []string {
	return []string{
		"SWITCHOVER",
		"SWITCHOVER_PRECHECK",
		"FAILOVER",
		"FAILOVER_PRECHECK",
	}
}

// GetMappingDrPlanExecutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanExecutionTypeEnum(val string) (DrPlanExecutionTypeEnum, bool) {
	enum, ok := mappingDrPlanExecutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
