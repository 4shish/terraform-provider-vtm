// Copyright (C) 2018-2019, Pulse Secure, LLC.
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	vtm "github.com/pulse-vadc/go-vtm/6.0"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VTM_BASE_URL", nil),
				Description: "Base URL: 'https://vtm:9070/api' or 'https://sd:8100/api/tmcm/<ver>/instance/<vtm>",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "admin",
				Description: "vTM admin user",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VTM_PASSWORD", nil),
				Description: "vTM admin password",
			},
			"verify_ssl_cert": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VTM_VERIFY_SSL_CERT", true),
				Description: "Check that vTM REST interface SSL certificate is trusted",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"vtm_backups_full":          resourceSystemBackupsFull(),
			"vtm_action":                resourceAction(),
			"vtm_action_program":        resourceActionProgram(),
			"vtm_appliance_nat":         resourceApplianceNat(),
			"vtm_aptimizer_profile":     resourceAptimizerProfile(),
			"vtm_aptimizer_scope":       resourceAptimizerScope(),
			"vtm_bandwidth":             resourceBandwidth(),
			"vtm_bgpneighbor":           resourceBgpneighbor(),
			"vtm_cloud_api_credential":  resourceCloudApiCredential(),
			"vtm_custom":                resourceCustom(),
			"vtm_dns_server_zone":       resourceDnsServerZone(),
			"vtm_dns_server_zone_file":  resourceDnsServerZoneFile(),
			"vtm_event_type":            resourceEventType(),
			"vtm_extra_file":            resourceExtraFile(),
			"vtm_glb_service":           resourceGlbService(),
			"vtm_global_settings":       resourceGlobalSettings(),
			"vtm_kerberos_keytab":       resourceKerberosKeytab(),
			"vtm_kerberos_krb5conf":     resourceKerberosKrb5Conf(),
			"vtm_kerberos_principal":    resourceKerberosPrincipal(),
			"vtm_license_key":           resourceLicenseKey(),
			"vtm_location":              resourceLocation(),
			"vtm_log_export":            resourceLogExport(),
			"vtm_monitor":               resourceMonitor(),
			"vtm_monitor_script":        resourceMonitorScript(),
			"vtm_persistence":           resourcePersistence(),
			"vtm_pool":                  resourcePool(),
			"vtm_protection":            resourceProtection(),
			"vtm_rate":                  resourceRate(),
			"vtm_rule":                  resourceRule(),
			"vtm_rule_authenticator":    resourceRuleAuthenticator(),
			"vtm_saml_trustedidp":       resourceSamlTrustedidp(),
			"vtm_security":              resourceSecurity(),
			"vtm_service_level_monitor": resourceServiceLevelMonitor(),
			"vtm_servicediscovery":      resourceServicediscovery(),
			"vtm_ssl_ca":                resourceSslCa(),
			"vtm_ssl_client_key":        resourceSslClientKey(),
			"vtm_ssl_server_key":        resourceSslServerKey(),
			"vtm_ssl_ticket_key":        resourceSslTicketKey(),
			"vtm_traffic_ip_group":      resourceTrafficIpGroup(),
			"vtm_traffic_manager":       resourceTrafficManager(),
			"vtm_user_authenticator":    resourceUserAuthenticator(),
			"vtm_user_group":            resourceUserGroup(),
			"vtm_virtual_server":        resourceVirtualServer(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"vtm_backups_full":                                     dataSourceSystemBackupsFull(),
			"vtm_backups_full_list":                                dataSourceSystemBackupsFullList(),
			"vtm_action":                                           dataSourceAction(),
			"vtm_action_arguments_table":                           dataSourceActionArgumentsTable(),
			"vtm_action_list":                                      dataSourceActionList(),
			"vtm_action_program":                                   dataSourceActionProgram(),
			"vtm_action_program_list":                              dataSourceActionProgramList(),
			"vtm_action_stats":                                     dataSourceActionStatistics(),
			"vtm_appliance_nat":                                    dataSourceApplianceNat(),
			"vtm_appliance_nat_many_to_one_all_ports_table":        dataSourceApplianceNatManyToOneAllPortsTable(),
			"vtm_appliance_nat_many_to_one_port_locked_table":      dataSourceApplianceNatManyToOnePortLockedTable(),
			"vtm_appliance_nat_one_to_one_table":                   dataSourceApplianceNatOneToOneTable(),
			"vtm_appliance_nat_port_mapping_table":                 dataSourceApplianceNatPortMappingTable(),
			"vtm_aptimizer_profile":                                dataSourceAptimizerProfile(),
			"vtm_aptimizer_profile_list":                           dataSourceAptimizerProfileList(),
			"vtm_aptimizer_scope":                                  dataSourceAptimizerScope(),
			"vtm_aptimizer_scope_list":                             dataSourceAptimizerScopeList(),
			"vtm_bandwidth":                                        dataSourceBandwidth(),
			"vtm_bandwidth_list":                                   dataSourceBandwidthList(),
			"vtm_bandwidth_stats":                                  dataSourceBandwidthStatistics(),
			"vtm_bgpneighbor":                                      dataSourceBgpneighbor(),
			"vtm_bgpneighbor_list":                                 dataSourceBgpneighborList(),
			"vtm_cache_asp_session_cache_stats":                    dataSourceCacheAspSessionCacheStatistics(),
			"vtm_cache_ip_session_cache_stats":                     dataSourceCacheIpSessionCacheStatistics(),
			"vtm_cache_j2ee_session_cache_stats":                   dataSourceCacheJ2EeSessionCacheStatistics(),
			"vtm_cache_ssl_cache_stats":                            dataSourceCacheSslCacheStatistics(),
			"vtm_cache_ssl_session_cache_stats":                    dataSourceCacheSslSessionCacheStatistics(),
			"vtm_cache_uni_session_cache_stats":                    dataSourceCacheUniSessionCacheStatistics(),
			"vtm_cache_web_cache_stats":                            dataSourceCacheWebCacheStatistics(),
			"vtm_cloud_api_credential":                             dataSourceCloudApiCredential(),
			"vtm_cloud_api_credential_list":                        dataSourceCloudApiCredentialList(),
			"vtm_cloud_api_credential_stats":                       dataSourceCloudApiCredentialStatistics(),
			"vtm_connection_rate_limit_stats":                      dataSourceConnectionRateLimitStatistics(),
			"vtm_custom":                                           dataSourceCustom(),
			"vtm_custom_list":                                      dataSourceCustomList(),
			"vtm_custom_string_lists_table":                        dataSourceCustomStringListsTable(),
			"vtm_dns_server_zone":                                  dataSourceDnsServerZone(),
			"vtm_dns_server_zone_file":                             dataSourceDnsServerZoneFile(),
			"vtm_dns_server_zone_file_list":                        dataSourceDnsServerZoneFileList(),
			"vtm_dns_server_zone_list":                             dataSourceDnsServerZoneList(),
			"vtm_event_stats":                                      dataSourceEventStatistics(),
			"vtm_event_type":                                       dataSourceEventType(),
			"vtm_event_type_list":                                  dataSourceEventTypeList(),
			"vtm_extra_file":                                       dataSourceExtraFile(),
			"vtm_extra_file_list":                                  dataSourceExtraFileList(),
			"vtm_extras_user_counters_32_stats":                    dataSourceExtrasUserCounters32Statistics(),
			"vtm_extras_user_counters_64_stats":                    dataSourceExtrasUserCounters64Statistics(),
			"vtm_glb_service":                                      dataSourceGlbService(),
			"vtm_glb_service_dnssec_keys_table":                    dataSourceGlbServiceDnssecKeysTable(),
			"vtm_glb_service_list":                                 dataSourceGlbServiceList(),
			"vtm_glb_service_location_settings_table":              dataSourceGlbServiceLocationSettingsTable(),
			"vtm_glb_service_stats":                                dataSourceGlbServiceStatistics(),
			"vtm_global_settings":                                  dataSourceGlobalSettings(),
			"vtm_global_settings_appliance_returnpath_table":       dataSourceGlobalSettingsApplianceReturnpathTable(),
			"vtm_globals_stats":                                    dataSourceGlobalsStatistics(),
			"vtm_information":                                      dataSourceSystemInformation(),
			"vtm_kerberos_keytab":                                  dataSourceKerberosKeytab(),
			"vtm_kerberos_keytab_list":                             dataSourceKerberosKeytabList(),
			"vtm_kerberos_krb5conf":                                dataSourceKerberosKrb5Conf(),
			"vtm_kerberos_krb5conf_list":                           dataSourceKerberosKrb5ConfList(),
			"vtm_kerberos_principal":                               dataSourceKerberosPrincipal(),
			"vtm_kerberos_principal_list":                          dataSourceKerberosPrincipalList(),
			"vtm_license_key":                                      dataSourceLicenseKey(),
			"vtm_license_key_list":                                 dataSourceLicenseKeyList(),
			"vtm_listen_ip_stats":                                  dataSourceListenIpStatistics(),
			"vtm_location":                                         dataSourceLocation(),
			"vtm_location_list":                                    dataSourceLocationList(),
			"vtm_location_stats":                                   dataSourceLocationStatistics(),
			"vtm_log_export":                                       dataSourceLogExport(),
			"vtm_log_export_list":                                  dataSourceLogExportList(),
			"vtm_log_export_metadata_table":                        dataSourceLogExportMetadataTable(),
			"vtm_monitor":                                          dataSourceMonitor(),
			"vtm_monitor_arguments_table":                          dataSourceMonitorArgumentsTable(),
			"vtm_monitor_list":                                     dataSourceMonitorList(),
			"vtm_monitor_script":                                   dataSourceMonitorScript(),
			"vtm_monitor_script_list":                              dataSourceMonitorScriptList(),
			"vtm_network_interface_stats":                          dataSourceNetworkInterfaceStatistics(),
			"vtm_nodes_node_inet46_stats":                          dataSourceNodesNodeInet46Statistics(),
			"vtm_nodes_node_stats":                                 dataSourceNodesNodeStatistics(),
			"vtm_nodes_per_pool_node_stats":                        dataSourceNodesPerPoolNodeStatistics(),
			"vtm_per_location_service_stats":                       dataSourcePerLocationServiceStatistics(),
			"vtm_per_node_slm_per_node_service_level_inet46_stats": dataSourcePerNodeSlmPerNodeServiceLevelInet46Statistics(),
			"vtm_per_node_slm_per_node_service_level_stats":        dataSourcePerNodeSlmPerNodeServiceLevelStatistics(),
			"vtm_persistence":                                      dataSourcePersistence(),
			"vtm_persistence_list":                                 dataSourcePersistenceList(),
			"vtm_pool":                                             dataSourcePool(),
			"vtm_pool_list":                                        dataSourcePoolList(),
			"vtm_pool_nodes_table_table":                           dataSourcePoolNodesTableTable(),
			"vtm_pool_stats":                                       dataSourcePoolStatistics(),
			"vtm_protection":                                       dataSourceProtection(),
			"vtm_protection_list":                                  dataSourceProtectionList(),
			"vtm_rate":                                             dataSourceRate(),
			"vtm_rate_list":                                        dataSourceRateList(),
			"vtm_rule":                                             dataSourceRule(),
			"vtm_rule_authenticator":                               dataSourceRuleAuthenticator(),
			"vtm_rule_authenticator_list":                          dataSourceRuleAuthenticatorList(),
			"vtm_rule_authenticator_stats":                         dataSourceRuleAuthenticatorStatistics(),
			"vtm_rule_list":                                        dataSourceRuleList(),
			"vtm_rule_stats":                                       dataSourceRuleStatistics(),
			"vtm_saml_trustedidp":                                  dataSourceSamlTrustedidp(),
			"vtm_saml_trustedidp_list":                             dataSourceSamlTrustedidpList(),
			"vtm_security":                                         dataSourceSecurity(),
			"vtm_service_level_monitor":                            dataSourceServiceLevelMonitor(),
			"vtm_service_level_monitor_list":                       dataSourceServiceLevelMonitorList(),
			"vtm_service_level_monitor_stats":                      dataSourceServiceLevelMonitorStatistics(),
			"vtm_service_protection_stats":                         dataSourceServiceProtectionStatistics(),
			"vtm_servicediscovery":                                 dataSourceServicediscovery(),
			"vtm_servicediscovery_list":                            dataSourceServicediscoveryList(),
			"vtm_ssl_ca":                                           dataSourceSslCa(),
			"vtm_ssl_ca_list":                                      dataSourceSslCaList(),
			"vtm_ssl_client_key":                                   dataSourceSslClientKey(),
			"vtm_ssl_client_key_list":                              dataSourceSslClientKeyList(),
			"vtm_ssl_ocsp_stapling_stats":                          dataSourceSslOcspStaplingStatistics(),
			"vtm_ssl_server_key":                                   dataSourceSslServerKey(),
			"vtm_ssl_server_key_list":                              dataSourceSslServerKeyList(),
			"vtm_ssl_ticket_key":                                   dataSourceSslTicketKey(),
			"vtm_ssl_ticket_key_list":                              dataSourceSslTicketKeyList(),
			"vtm_state":                                            dataSourceSystemState(),
			"vtm_traffic_ip_group":                                 dataSourceTrafficIpGroup(),
			"vtm_traffic_ip_group_ip_mapping_table":                dataSourceTrafficIpGroupIpMappingTable(),
			"vtm_traffic_ip_group_list":                            dataSourceTrafficIpGroupList(),
			"vtm_traffic_ips_ip_gateway_stats":                     dataSourceTrafficIpsIpGatewayStatistics(),
			"vtm_traffic_ips_traffic_ip_inet46_stats":              dataSourceTrafficIpsTrafficIpInet46Statistics(),
			"vtm_traffic_ips_traffic_ip_stats":                     dataSourceTrafficIpsTrafficIpStatistics(),
			"vtm_traffic_manager":                                  dataSourceTrafficManager(),
			"vtm_traffic_manager_appliance_card_table":             dataSourceTrafficManagerApplianceCardTable(),
			"vtm_traffic_manager_appliance_sysctl_table":           dataSourceTrafficManagerApplianceSysctlTable(),
			"vtm_traffic_manager_hosts_table":                      dataSourceTrafficManagerHostsTable(),
			"vtm_traffic_manager_if_table":                         dataSourceTrafficManagerIfTable(),
			"vtm_traffic_manager_ip_table":                         dataSourceTrafficManagerIpTable(),
			"vtm_traffic_manager_list":                             dataSourceTrafficManagerList(),
			"vtm_traffic_manager_routes_table":                     dataSourceTrafficManagerRoutesTable(),
			"vtm_traffic_manager_trafficip_table":                  dataSourceTrafficManagerTrafficipTable(),
			"vtm_user_authenticator":                               dataSourceUserAuthenticator(),
			"vtm_user_authenticator_list":                          dataSourceUserAuthenticatorList(),
			"vtm_user_group":                                       dataSourceUserGroup(),
			"vtm_user_group_list":                                  dataSourceUserGroupList(),
			"vtm_user_group_permissions_table":                     dataSourceUserGroupPermissionsTable(),
			"vtm_virtual_server":                                   dataSourceVirtualServer(),
			"vtm_virtual_server_list":                              dataSourceVirtualServerList(),
			"vtm_virtual_server_ocsp_issuers_table":                dataSourceVirtualServerOcspIssuersTable(),
			"vtm_virtual_server_profile_table":                     dataSourceVirtualServerProfileTable(),
			"vtm_virtual_server_server_cert_host_mapping_table":    dataSourceVirtualServerServerCertHostMappingTable(),
			"vtm_virtual_server_stats":                             dataSourceVirtualServerStatistics(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	baseUrl := d.Get("base_url").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	verifySslCert := d.Get("verify_ssl_cert").(bool)

	tm, contactable, contactErr := vtm.NewVirtualTrafficManager(baseUrl, username, password, verifySslCert, true)
	if contactable != true {
		return nil, fmt.Errorf("Failed to connect to Virtual Traffic Manager at '%v': %v", baseUrl, contactErr.ErrorText)
	}
	return tm, nil
}
