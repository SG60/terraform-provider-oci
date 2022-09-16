// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TimeClusterCommandDescriptor Command descriptor for querylanguage TIMECLUSTER command.
type TimeClusterCommandDescriptor struct {

	// Command fragment display string from user specified query string formatted by query builder.
	DisplayQueryString *string `mandatory:"true" json:"displayQueryString"`

	// Command fragment internal string from user specified query string formatted by query builder.
	InternalQueryString *string `mandatory:"true" json:"internalQueryString"`

	// querylanguage command designation for example; reporting vs filtering
	Category *string `mandatory:"false" json:"category"`

	// Fields referenced in command fragment from user specified query string.
	ReferencedFields []AbstractField `mandatory:"false" json:"referencedFields"`

	// Fields declared in command fragment from user specified query string.
	DeclaredFields []AbstractField `mandatory:"false" json:"declaredFields"`

	// Field denoting if this is a hidden command that is not shown in the query string.
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// Optional timestamp datatype field if specified. Default field is 'Start Time'.
	Time AbstractField `mandatory:"false" json:"time"`

	// Option to control the size of buckets in the timeseries. Will be adjusted to a larger span if time range is very large.
	Span *string `mandatory:"false" json:"span"`

	// Group by fields specified in the query string.
	GroupByFields []AbstractField `mandatory:"false" json:"groupByFields"`

	// Statistical functions specified in the query string. Atleast 1 is required for a TIMECLUSTER command.
	Functions []FunctionField `mandatory:"false" json:"functions"`
}

//GetDisplayQueryString returns DisplayQueryString
func (m TimeClusterCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m TimeClusterCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m TimeClusterCommandDescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m TimeClusterCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m TimeClusterCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

//GetIsHidden returns IsHidden
func (m TimeClusterCommandDescriptor) GetIsHidden() *bool {
	return m.IsHidden
}

func (m TimeClusterCommandDescriptor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TimeClusterCommandDescriptor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TimeClusterCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTimeClusterCommandDescriptor TimeClusterCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeTimeClusterCommandDescriptor
	}{
		"TIME_CLUSTER",
		(MarshalTypeTimeClusterCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TimeClusterCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string         `json:"category"`
		ReferencedFields    []abstractfield `json:"referencedFields"`
		DeclaredFields      []abstractfield `json:"declaredFields"`
		IsHidden            *bool           `json:"isHidden"`
		Time                abstractfield   `json:"time"`
		Span                *string         `json:"span"`
		GroupByFields       []abstractfield `json:"groupByFields"`
		Functions           []FunctionField `json:"functions"`
		DisplayQueryString  *string         `json:"displayQueryString"`
		InternalQueryString *string         `json:"internalQueryString"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Category = model.Category

	m.ReferencedFields = make([]AbstractField, len(model.ReferencedFields))
	for i, n := range model.ReferencedFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ReferencedFields[i] = nn.(AbstractField)
		} else {
			m.ReferencedFields[i] = nil
		}
	}

	m.DeclaredFields = make([]AbstractField, len(model.DeclaredFields))
	for i, n := range model.DeclaredFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DeclaredFields[i] = nn.(AbstractField)
		} else {
			m.DeclaredFields[i] = nil
		}
	}

	m.IsHidden = model.IsHidden

	nn, e = model.Time.UnmarshalPolymorphicJSON(model.Time.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Time = nn.(AbstractField)
	} else {
		m.Time = nil
	}

	m.Span = model.Span

	m.GroupByFields = make([]AbstractField, len(model.GroupByFields))
	for i, n := range model.GroupByFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.GroupByFields[i] = nn.(AbstractField)
		} else {
			m.GroupByFields[i] = nil
		}
	}

	m.Functions = make([]FunctionField, len(model.Functions))
	for i, n := range model.Functions {
		m.Functions[i] = n
	}

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}
