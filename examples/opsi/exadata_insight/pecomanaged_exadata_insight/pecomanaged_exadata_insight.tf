// Copyright (c) 2017, 2022, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "opsi_private_endpoint_id" {}
variable "dbsystem_database_id" {}
variable "service_name" {}
variable "user_name" {}
variable "password_secret_id" {}
variable "exadata_infra_id" {}
variable "vmcluster_id" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "examples-tag-namespace-all"
  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false
  is_retired       = false
}

variable "exadata_insight_type" {
  default = ["EXACS"]
}

variable "deployment_type" {
  default = "EXACS"
}

variable "credential_details_credential_type" {
  default = "CREDENTIALS_BY_VAULT"
}

variable "credential_details_role" {
  default = "NORMAL"
}

variable "database_resource_type" {
  default = "database"
}

variable "exadata_insight_defined_tags_value" {
  default = "value"
}

variable "exadata_insight_entity_source" {
  default = "PE_COMANAGED_EXADATA"
}

variable "freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "resource_status" {
  default = "ENABLED"
}

// Create PE Comanaged Exadata insight
resource "oci_opsi_exadata_insight" "test_exadata_insight" {
  #Required
  compartment_id                       = var.compartment_ocid
  entity_source                        = var.exadata_insight_entity_source

  #Optional
  exadata_infra_id                     = var.exadata_infra_id
  member_vm_cluster_details {
      vmcluster_id                         = var.vmcluster_id
      compartment_id                       = var.compartment_ocid
      opsi_private_endpoint_id             = var.opsi_private_endpoint_id
      member_database_details {
          entity_source                        = "PE_COMANAGED_DATABASE"
          compartment_id                       = var.compartment_ocid
          database_id                          = var.dbsystem_database_id
          database_resource_type               = var.database_resource_type
          opsi_private_endpoint_id             = var.opsi_private_endpoint_id
          service_name                         = var.service_name
          deployment_type                      = var.deployment_type
          credential_details {
              credential_type                  = var.credential_details_credential_type
              password_secret_id               = var.password_secret_id
              role                             = var.credential_details_role
              user_name                        = var.user_name
          }
      }
  }
  defined_tags                         = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.exadata_insight_defined_tags_value}")}"
  freeform_tags                        = var.freeform_tags
}

variable "exadata_insight_state" {
  default = ["ACTIVE"]
}

variable "exadata_insight_status" {
  default = ["ENABLED"]
}

variable "exadata_type" {
  default = ["EXACS"]
}

// List PE comanaged exadata insights
data "oci_opsi_exadata_insights" "test_exadata_insights" {
  #Optional
  compartment_id               = var.compartment_ocid
  exadata_type                 = var.exadata_insight_type
  state                        = var.exadata_insight_state
  status                       = var.exadata_insight_status
}

// Get a PE comanaged exadata insight
data "oci_opsi_exadata_insight" "test_exadata_insight" {
  exadata_insight_id = oci_opsi_exadata_insight.test_exadata_insight.id
}

