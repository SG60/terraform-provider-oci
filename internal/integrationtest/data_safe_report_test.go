// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafereportSingularDataSourceRepresentation = map[string]interface{}{
		"report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.rep_identifier}`},
	}

	DataSafereportDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"report_definition_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_report_definition.test_report_definition.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"type":                      acctest.Representation{RepType: acctest.Optional, Create: `GENERATED`},
	}

	DataSafeReportResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_report_definition", "test_report_definition", acctest.Required, acctest.Create, reportDefinitionRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	//compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	reportDefId := utils.GetEnvSettingWithBlankDefault("report_ocid")
	reportDefIdVariableStr := fmt.Sprintf("variable \"report_ocid\" { default = \"%s\" }\n", reportDefId)

	reportIdentifier := utils.GetEnvSettingWithBlankDefault("rep_identifier")
	reportIdentifierStr := fmt.Sprintf("variable \"rep_identifier\" { default = \"%s\" }\n", reportIdentifier)

	datasourceName := "data.oci_data_safe_reports.test_reports"
	singularDatasourceName := "data.oci_data_safe_report.test_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_reports", "test_reports", acctest.Required, acctest.Create, DataSafereportDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeReportResourceConfig + reportDefIdVariableStr + reportIdentifierStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "report_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_report", "test_report", acctest.Required, acctest.Create, DataSafereportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeReportResourceConfig + reportDefIdVariableStr + reportIdentifierStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "report_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mime_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_generated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
	})
}
