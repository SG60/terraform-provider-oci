// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package globalvar

import (
	"log"
)

const Version = "4.102.0"
const ReleaseDate = "2022-12-14"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}
