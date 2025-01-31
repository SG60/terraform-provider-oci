// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_queue "github.com/oracle/oci-go-sdk/v65/queue"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_queue.QueueAdminClient", &OracleClient{InitClientFn: initQueueQueueAdminClient})
}

func initQueueQueueAdminClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_queue.NewQueueAdminClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) QueueAdminClient() *oci_queue.QueueAdminClient {
	return m.GetClient("oci_queue.QueueAdminClient").(*oci_queue.QueueAdminClient)
}
