// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {
  default = "ocid1.compartment.oc1..aaaaaaaaaq4dqogd2ktatzmuekujkasvwendyhisgfqdky3ojru47w3f634a"
}

variable "dr_plan_defined_tags_value" {
  default = "value"
}

variable "dr_plan_display_name" {
  default = "displayName"
}

variable "dr_plan_dr_plan_type" {
  default = "SWITCHOVER"
}

variable "dr_plan_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "dr_plan_state" {
  default = "ACTIVE"
}

variable "dr_plan_type" {
  default = "SWITCHOVER"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_disaster_recovery_dr_plan" "test_dr_plan" {
  #Required
  display_name           = var.dr_plan_display_name
  dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id
  type                   = var.dr_plan_type

  #Optional
  defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.dr_plan_defined_tags_value}")
  freeform_tags = var.dr_plan_freeform_tags
}

data "oci_disaster_recovery_dr_plans" "test_dr_plans" {
  #Required
  dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id

  #Optional
  display_name = var.dr_plan_display_name
  dr_plan_id   = oci_disaster_recovery_dr_plan.test_dr_plan.id
  dr_plan_type = var.dr_plan_dr_plan_type
  state        = var.dr_plan_state
}

