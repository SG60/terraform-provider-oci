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

// UpdateFleetAdvancedFeatureConfigurationDetails Details object containing advanced feature configurations to be updated
// Ensure that the namespace and bucket storage are created prior to turning on the JfrRecording or CryptoEventAnalysis feature
type UpdateFleetAdvancedFeatureConfigurationDetails struct {

	// Namespace for the fleet advanced feature
	AnalyticNamespace *string `mandatory:"false" json:"analyticNamespace"`

	// Bucket name required to store jfr and related data
	AnalyticBucketName *string `mandatory:"false" json:"analyticBucketName"`

	Lcm *Lcm `mandatory:"false" json:"lcm"`

	CryptoEventAnalysis *CryptoEventAnalysis `mandatory:"false" json:"cryptoEventAnalysis"`

	AdvancedUsageTracking *AdvancedUsageTracking `mandatory:"false" json:"advancedUsageTracking"`

	JfrRecording *JfrRecording `mandatory:"false" json:"jfrRecording"`
}

func (m UpdateFleetAdvancedFeatureConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateFleetAdvancedFeatureConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
