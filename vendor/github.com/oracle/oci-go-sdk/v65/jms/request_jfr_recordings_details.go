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

// RequestJfrRecordingsDetails Details of the request to start JFR recordings.
// When the targets aren't specified, then all managed instances currently in the fleet are selected.
type RequestJfrRecordingsDetails struct {

	// The profile used for JFR events selection. If the name isn't recognized, the settings from jfcV1 or jfcV2
	// will be used depending on the JVM version.
	// Both jfcV2 and jfcV1 should be provided to ensure JFR collection on different JVM versions.
	JfcProfileName *string `mandatory:"true" json:"jfcProfileName"`

	// The attachment targets to start JFR.
	Targets []JfrAttachmentTarget `mandatory:"false" json:"targets"`

	// The BASE64 encoded string of JFR settings XML with schema used by JDK 8.
	JfcV1 *string `mandatory:"false" json:"jfcV1"`

	// The BASE64 encoded string of JFR settings XML with schema used by JDK 9 and after (https://raw.githubusercontent.com/openjdk/jdk/master/src/jdk.jfr/share/classes/jdk/jfr/internal/jfc/jfc.xsd).
	JfcV2 *string `mandatory:"false" json:"jfcV2"`

	// Duration of the JFR recording in minutes.
	RecordingDurationInMinutes *int `mandatory:"false" json:"recordingDurationInMinutes"`

	// The maximum size limit for the JFR file collected.
	RecordingSizeInMb *int `mandatory:"false" json:"recordingSizeInMb"`
}

func (m RequestJfrRecordingsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestJfrRecordingsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
