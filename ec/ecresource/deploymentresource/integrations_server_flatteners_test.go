// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package deploymentresource

import (
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/stretchr/testify/assert"
)

func Test_flattenIntegrationsServerResource(t *testing.T) {
	type args struct {
		in   []*models.IntegrationsServerResourceInfo
		name string
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "empty resource list returns empty list",
			args: args{in: []*models.IntegrationsServerResourceInfo{}},
			want: []interface{}{},
		},
		{
			name: "empty current plan returns empty list",
			args: args{in: []*models.IntegrationsServerResourceInfo{
				{
					Info: &models.IntegrationsServerInfo{
						PlanInfo: &models.IntegrationsServerPlansInfo{
							Pending: &models.IntegrationsServerPlanInfo{},
						},
					},
				},
			}},
			want: []interface{}{},
		},
		{
			name: "parses the integrations_server resource",
			args: args{in: []*models.IntegrationsServerResourceInfo{
				{
					Region:                    ec.String("some-region"),
					RefID:                     ec.String("main-integrations_server"),
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Info: &models.IntegrationsServerInfo{
						ID:     &mock.ValidClusterID,
						Name:   ec.String("some-integrations_server-name"),
						Region: "some-region",
						Status: ec.String("started"),
						Metadata: &models.ClusterMetadataInfo{
							Endpoint: "integrations_serverresource.cloud.elastic.co",
							Ports: &models.ClusterMetadataPortInfo{
								HTTP:  ec.Int32(9200),
								HTTPS: ec.Int32(9243),
							},
						},
						PlanInfo: &models.IntegrationsServerPlansInfo{Current: &models.IntegrationsServerPlanInfo{
							Plan: &models.IntegrationsServerPlan{
								IntegrationsServer: &models.IntegrationsServerConfiguration{
									Version: "7.7.0",
								},
								ClusterTopology: []*models.IntegrationsServerTopologyElement{
									{
										ZoneCount:               1,
										InstanceConfigurationID: "aws.integrations_server.r4",
										Size: &models.TopologySize{
											Resource: ec.String("memory"),
											Value:    ec.Int32(1024),
										},
									},
								},
							},
						}},
					},
				},
			}},
			want: []interface{}{
				map[string]interface{}{
					"elasticsearch_cluster_ref_id": "main-elasticsearch",
					"ref_id":                       "main-integrations_server",
					"resource_id":                  mock.ValidClusterID,
					"region":                       "some-region",
					"http_endpoint":                "http://integrations_serverresource.cloud.elastic.co:9200",
					"https_endpoint":               "https://integrations_serverresource.cloud.elastic.co:9243",
					"topology": []interface{}{
						map[string]interface{}{
							"instance_configuration_id": "aws.integrations_server.r4",
							"size":                      "1g",
							"size_resource":             "memory",
							"zone_count":                int32(1),
						},
					},
				},
			},
		},
		{
			name: "parses the integrations_server resource with config overrides, ignoring a stopped resource",
			args: args{in: []*models.IntegrationsServerResourceInfo{
				{
					Region:                    ec.String("some-region"),
					RefID:                     ec.String("main-integrations_server"),
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Info: &models.IntegrationsServerInfo{
						ID:     &mock.ValidClusterID,
						Name:   ec.String("some-integrations_server-name"),
						Region: "some-region",
						Status: ec.String("started"),
						Metadata: &models.ClusterMetadataInfo{
							Endpoint: "integrations_serverresource.cloud.elastic.co",
							Ports: &models.ClusterMetadataPortInfo{
								HTTP:  ec.Int32(9200),
								HTTPS: ec.Int32(9243),
							},
						},
						PlanInfo: &models.IntegrationsServerPlansInfo{Current: &models.IntegrationsServerPlanInfo{
							Plan: &models.IntegrationsServerPlan{
								IntegrationsServer: &models.IntegrationsServerConfiguration{
									Version:                  "7.8.0",
									UserSettingsYaml:         `some.setting: value`,
									UserSettingsOverrideYaml: `some.setting: value2`,
									UserSettingsJSON: map[string]interface{}{
										"some.setting": "value",
									},
									UserSettingsOverrideJSON: map[string]interface{}{
										"some.setting": "value2",
									},
									SystemSettings: &models.IntegrationsServerSystemSettings{},
								},
								ClusterTopology: []*models.IntegrationsServerTopologyElement{
									{
										ZoneCount:               1,
										InstanceConfigurationID: "aws.integrations_server.r4",
										Size: &models.TopologySize{
											Resource: ec.String("memory"),
											Value:    ec.Int32(1024),
										},
									},
								},
							},
						}},
					},
				},
				{
					Region:                    ec.String("some-region"),
					RefID:                     ec.String("main-integrations_server"),
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Info: &models.IntegrationsServerInfo{
						ID:     &mock.ValidClusterID,
						Name:   ec.String("some-integrations_server-name"),
						Region: "some-region",
						Status: ec.String("stopped"),
						Metadata: &models.ClusterMetadataInfo{
							Endpoint: "integrations_serverresource.cloud.elastic.co",
							Ports: &models.ClusterMetadataPortInfo{
								HTTP:  ec.Int32(9200),
								HTTPS: ec.Int32(9243),
							},
						},
						PlanInfo: &models.IntegrationsServerPlansInfo{Current: &models.IntegrationsServerPlanInfo{
							Plan: &models.IntegrationsServerPlan{
								IntegrationsServer: &models.IntegrationsServerConfiguration{
									Version:                  "7.8.0",
									UserSettingsYaml:         `some.setting: value`,
									UserSettingsOverrideYaml: `some.setting: value2`,
									UserSettingsJSON: map[string]interface{}{
										"some.setting": "value",
									},
									UserSettingsOverrideJSON: map[string]interface{}{
										"some.setting": "value2",
									},
									SystemSettings: &models.IntegrationsServerSystemSettings{},
								},
								ClusterTopology: []*models.IntegrationsServerTopologyElement{
									{
										ZoneCount:               1,
										InstanceConfigurationID: "aws.integrations_server.r4",
										Size: &models.TopologySize{
											Resource: ec.String("memory"),
											Value:    ec.Int32(1024),
										},
									},
								},
							},
						}},
					},
				},
			}},
			want: []interface{}{map[string]interface{}{
				"elasticsearch_cluster_ref_id": "main-elasticsearch",
				"ref_id":                       "main-integrations_server",
				"resource_id":                  mock.ValidClusterID,
				"region":                       "some-region",
				"http_endpoint":                "http://integrations_serverresource.cloud.elastic.co:9200",
				"https_endpoint":               "https://integrations_serverresource.cloud.elastic.co:9243",
				"topology": []interface{}{map[string]interface{}{
					"instance_configuration_id": "aws.integrations_server.r4",
					"size":                      "1g",
					"size_resource":             "memory",
					"zone_count":                int32(1),
				}},
				"config": []interface{}{map[string]interface{}{
					"user_settings_yaml":          "some.setting: value",
					"user_settings_override_yaml": "some.setting: value2",
					"user_settings_json":          "{\"some.setting\":\"value\"}",
					"user_settings_override_json": "{\"some.setting\":\"value2\"}",
				}},
			}},
		},
		{
			name: "parses the integrations_server resource with config overrides and system settings",
			args: args{in: []*models.IntegrationsServerResourceInfo{
				{
					Region:                    ec.String("some-region"),
					RefID:                     ec.String("main-integrations_server"),
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Info: &models.IntegrationsServerInfo{
						ID:     &mock.ValidClusterID,
						Name:   ec.String("some-integrations_server-name"),
						Region: "some-region",
						Status: ec.String("started"),
						Metadata: &models.ClusterMetadataInfo{
							Endpoint: "integrations_serverresource.cloud.elastic.co",
							Ports: &models.ClusterMetadataPortInfo{
								HTTP:  ec.Int32(9200),
								HTTPS: ec.Int32(9243),
							},
						},
						PlanInfo: &models.IntegrationsServerPlansInfo{Current: &models.IntegrationsServerPlanInfo{
							Plan: &models.IntegrationsServerPlan{
								IntegrationsServer: &models.IntegrationsServerConfiguration{
									Version:                  "7.8.0",
									UserSettingsYaml:         `some.setting: value`,
									UserSettingsOverrideYaml: `some.setting: value2`,
									UserSettingsJSON: map[string]interface{}{
										"some.setting": "value",
									},
									UserSettingsOverrideJSON: map[string]interface{}{
										"some.setting": "value2",
									},
									SystemSettings: &models.IntegrationsServerSystemSettings{
										DebugEnabled: ec.Bool(true),
									},
								},
								ClusterTopology: []*models.IntegrationsServerTopologyElement{
									{
										ZoneCount:               1,
										InstanceConfigurationID: "aws.integrations_server.r4",
										Size: &models.TopologySize{
											Resource: ec.String("memory"),
											Value:    ec.Int32(1024),
										},
									},
								},
							},
						}},
					},
				},
			}},
			want: []interface{}{map[string]interface{}{
				"elasticsearch_cluster_ref_id": "main-elasticsearch",
				"ref_id":                       "main-integrations_server",
				"resource_id":                  mock.ValidClusterID,
				"region":                       "some-region",
				"http_endpoint":                "http://integrations_serverresource.cloud.elastic.co:9200",
				"https_endpoint":               "https://integrations_serverresource.cloud.elastic.co:9243",
				"topology": []interface{}{map[string]interface{}{
					"instance_configuration_id": "aws.integrations_server.r4",
					"size":                      "1g",
					"size_resource":             "memory",
					"zone_count":                int32(1),
				}},
				"config": []interface{}{map[string]interface{}{
					"user_settings_yaml":          "some.setting: value",
					"user_settings_override_yaml": "some.setting: value2",
					"user_settings_json":          "{\"some.setting\":\"value\"}",
					"user_settings_override_json": "{\"some.setting\":\"value2\"}",

					"debug_enabled": true,
				}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flattenIntegrationsServerResources(tt.args.in, tt.args.name)
			assert.Equal(t, tt.want, got)
		})
	}
}
