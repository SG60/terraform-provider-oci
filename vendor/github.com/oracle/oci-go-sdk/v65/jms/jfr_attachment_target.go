// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JfrAttachmentTarget The target to collect JFR data. A target is a managed instance, with options to further limit to specific application and/or Java runtime.
// When the applicationKey isn't specified, then all applications are selected.
// When the jreKey isn't specified, then all supported Java runtime versions are selected.
type JfrAttachmentTarget struct {

	// OCID of the Managed Instance to collect JFR data.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// Unique key that identify the application for JFR data collection.
	ApplicationKey *string `mandatory:"false" json:"applicationKey"`

	// Unique key that identify the JVM for JFR data collection.
	JreKey *string `mandatory:"false" json:"jreKey"`
}

func (m JfrAttachmentTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JfrAttachmentTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
