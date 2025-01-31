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

// RequestCryptoAnalysesDetails Details of the request to start a JFR analysis.
// When the targets aren't specified, then all managed instances currently in the fleet are selected.
type RequestCryptoAnalysesDetails struct {

	// The attachment targets to start JFR.
	Targets []JfrAttachmentTarget `mandatory:"false" json:"targets"`

	// Duration of the JFR recording in minutes.
	RecordingDurationInMinutes *int `mandatory:"false" json:"recordingDurationInMinutes"`
}

func (m RequestCryptoAnalysesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestCryptoAnalysesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
