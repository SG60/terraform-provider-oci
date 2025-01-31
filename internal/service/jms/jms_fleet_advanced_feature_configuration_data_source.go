// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetAdvancedFeatureConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsFleetAdvancedFeatureConfiguration,
		Schema: map[string]*schema.Schema{
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"advanced_usage_tracking": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"analytic_bucket_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"analytic_namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crypto_event_analysis": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"summarized_events_log": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"log_group_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"jfr_recording": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"lcm": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"post_installation_actions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"disabled_tls_versions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"minimum_key_size_settings": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"certpath": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"key_size": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"jar": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"key_size": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"tls": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"key_size": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"should_replace_certificates_operating_system": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsFleetAdvancedFeatureConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAdvancedFeatureConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetAdvancedFeatureConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetFleetAdvancedFeatureConfigurationResponse
}

func (s *JmsFleetAdvancedFeatureConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetAdvancedFeatureConfigurationDataSourceCrud) Get() error {
	request := oci_jms.GetFleetAdvancedFeatureConfigurationRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetFleetAdvancedFeatureConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetAdvancedFeatureConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetAdvancedFeatureConfigurationDataSource-", JmsFleetAdvancedFeatureConfigurationDataSource(), s.D))

	if s.Res.AdvancedUsageTracking != nil {
		s.D.Set("advanced_usage_tracking", []interface{}{AdvancedUsageTrackingToMap(s.Res.AdvancedUsageTracking)})
	} else {
		s.D.Set("advanced_usage_tracking", nil)
	}

	if s.Res.AnalyticBucketName != nil {
		s.D.Set("analytic_bucket_name", *s.Res.AnalyticBucketName)
	}

	if s.Res.AnalyticNamespace != nil {
		s.D.Set("analytic_namespace", *s.Res.AnalyticNamespace)
	}

	if s.Res.CryptoEventAnalysis != nil {
		s.D.Set("crypto_event_analysis", []interface{}{CryptoEventAnalysisToMap(s.Res.CryptoEventAnalysis)})
	} else {
		s.D.Set("crypto_event_analysis", nil)
	}

	if s.Res.JfrRecording != nil {
		s.D.Set("jfr_recording", []interface{}{JfrRecordingToMap(s.Res.JfrRecording)})
	} else {
		s.D.Set("jfr_recording", nil)
	}

	if s.Res.Lcm != nil {
		s.D.Set("lcm", []interface{}{LcmToMap(s.Res.Lcm)})
	} else {
		s.D.Set("lcm", nil)
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}

func AdvancedUsageTrackingToMap(obj *oci_jms.AdvancedUsageTracking) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func CryptoEventAnalysisToMap(obj *oci_jms.CryptoEventAnalysis) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.SummarizedEventsLog != nil {
		result["summarized_events_log"] = []interface{}{SummarizedEventsLogToMap(obj.SummarizedEventsLog)}
	}

	return result
}

func JfrRecordingToMap(obj *oci_jms.JfrRecording) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func KeySizeAlgorithmToMap(obj oci_jms.KeySizeAlgorithm) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeySize != nil {
		result["key_size"] = int(*obj.KeySize)
	}

	result["name"] = string(obj.Name)

	return result
}

func LcmToMap(obj *oci_jms.Lcm) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.PostInstallationActions != nil {
		result["post_installation_actions"] = []interface{}{PostInstallationActionSettingsToMap(obj.PostInstallationActions)}
	}

	return result
}

func MinimumKeySizeSettingsToMap(obj *oci_jms.MinimumKeySizeSettings) map[string]interface{} {
	result := map[string]interface{}{}

	certpath := []interface{}{}
	for _, item := range obj.Certpath {
		certpath = append(certpath, KeySizeAlgorithmToMap(item))
	}
	result["certpath"] = certpath

	jar := []interface{}{}
	for _, item := range obj.Jar {
		jar = append(jar, KeySizeAlgorithmToMap(item))
	}
	result["jar"] = jar

	tls := []interface{}{}
	for _, item := range obj.Tls {
		tls = append(tls, KeySizeAlgorithmToMap(item))
	}
	result["tls"] = tls

	return result
}

func PostInstallationActionSettingsToMap(obj *oci_jms.PostInstallationActionSettings) map[string]interface{} {
	result := map[string]interface{}{}

	disabledTlsVersions := []interface{}{}
	for _, item := range obj.DisabledTlsVersions {
		disabledTlsVersions = append(disabledTlsVersions, string(item))
	}
	result["disabled_tls_versions"] = disabledTlsVersions

	if obj.MinimumKeySizeSettings != nil {
		result["minimum_key_size_settings"] = []interface{}{MinimumKeySizeSettingsToMap(obj.MinimumKeySizeSettings)}
	}

	if obj.ShouldReplaceCertificatesOperatingSystem != nil {
		result["should_replace_certificates_operating_system"] = bool(*obj.ShouldReplaceCertificatesOperatingSystem)
	}

	return result
}

func SummarizedEventsLogToMap(obj *oci_jms.SummarizedEventsLog) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}
