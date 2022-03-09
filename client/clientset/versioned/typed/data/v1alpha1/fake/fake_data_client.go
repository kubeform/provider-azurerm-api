/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/data/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDataV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDataV1alpha1) Factories(namespace string) v1alpha1.FactoryInterface {
	return &FakeFactories{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryCustomDatasets(namespace string) v1alpha1.FactoryCustomDatasetInterface {
	return &FakeFactoryCustomDatasets{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDataFlows(namespace string) v1alpha1.FactoryDataFlowInterface {
	return &FakeFactoryDataFlows{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetAzureBlobs(namespace string) v1alpha1.FactoryDatasetAzureBlobInterface {
	return &FakeFactoryDatasetAzureBlobs{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetBinaries(namespace string) v1alpha1.FactoryDatasetBinaryInterface {
	return &FakeFactoryDatasetBinaries{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetCosmosdbSqlapis(namespace string) v1alpha1.FactoryDatasetCosmosdbSqlapiInterface {
	return &FakeFactoryDatasetCosmosdbSqlapis{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetDelimitedTexts(namespace string) v1alpha1.FactoryDatasetDelimitedTextInterface {
	return &FakeFactoryDatasetDelimitedTexts{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetHTTPs(namespace string) v1alpha1.FactoryDatasetHTTPInterface {
	return &FakeFactoryDatasetHTTPs{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetJSONs(namespace string) v1alpha1.FactoryDatasetJSONInterface {
	return &FakeFactoryDatasetJSONs{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetMysqls(namespace string) v1alpha1.FactoryDatasetMysqlInterface {
	return &FakeFactoryDatasetMysqls{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetParquets(namespace string) v1alpha1.FactoryDatasetParquetInterface {
	return &FakeFactoryDatasetParquets{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetPostgresqls(namespace string) v1alpha1.FactoryDatasetPostgresqlInterface {
	return &FakeFactoryDatasetPostgresqls{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetSQLServerTables(namespace string) v1alpha1.FactoryDatasetSQLServerTableInterface {
	return &FakeFactoryDatasetSQLServerTables{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryDatasetSnowflakes(namespace string) v1alpha1.FactoryDatasetSnowflakeInterface {
	return &FakeFactoryDatasetSnowflakes{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryIntegrationRuntimeAzures(namespace string) v1alpha1.FactoryIntegrationRuntimeAzureInterface {
	return &FakeFactoryIntegrationRuntimeAzures{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryIntegrationRuntimeAzureSsises(namespace string) v1alpha1.FactoryIntegrationRuntimeAzureSsisInterface {
	return &FakeFactoryIntegrationRuntimeAzureSsises{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryIntegrationRuntimeManageds(namespace string) v1alpha1.FactoryIntegrationRuntimeManagedInterface {
	return &FakeFactoryIntegrationRuntimeManageds{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryIntegrationRuntimeSelfHosteds(namespace string) v1alpha1.FactoryIntegrationRuntimeSelfHostedInterface {
	return &FakeFactoryIntegrationRuntimeSelfHosteds{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedCustomServices(namespace string) v1alpha1.FactoryLinkedCustomServiceInterface {
	return &FakeFactoryLinkedCustomServices{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceAzureBlobStorages(namespace string) v1alpha1.FactoryLinkedServiceAzureBlobStorageInterface {
	return &FakeFactoryLinkedServiceAzureBlobStorages{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceAzureDatabrickses(namespace string) v1alpha1.FactoryLinkedServiceAzureDatabricksInterface {
	return &FakeFactoryLinkedServiceAzureDatabrickses{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceAzureFileStorages(namespace string) v1alpha1.FactoryLinkedServiceAzureFileStorageInterface {
	return &FakeFactoryLinkedServiceAzureFileStorages{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceAzureFunctions(namespace string) v1alpha1.FactoryLinkedServiceAzureFunctionInterface {
	return &FakeFactoryLinkedServiceAzureFunctions{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceAzureSQLDatabases(namespace string) v1alpha1.FactoryLinkedServiceAzureSQLDatabaseInterface {
	return &FakeFactoryLinkedServiceAzureSQLDatabases{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceAzureSearches(namespace string) v1alpha1.FactoryLinkedServiceAzureSearchInterface {
	return &FakeFactoryLinkedServiceAzureSearches{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceAzureTableStorages(namespace string) v1alpha1.FactoryLinkedServiceAzureTableStorageInterface {
	return &FakeFactoryLinkedServiceAzureTableStorages{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceCosmosdbs(namespace string) v1alpha1.FactoryLinkedServiceCosmosdbInterface {
	return &FakeFactoryLinkedServiceCosmosdbs{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceCosmosdbMongoapis(namespace string) v1alpha1.FactoryLinkedServiceCosmosdbMongoapiInterface {
	return &FakeFactoryLinkedServiceCosmosdbMongoapis{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceDataLakeStorageGen2s(namespace string) v1alpha1.FactoryLinkedServiceDataLakeStorageGen2Interface {
	return &FakeFactoryLinkedServiceDataLakeStorageGen2s{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceKeyVaults(namespace string) v1alpha1.FactoryLinkedServiceKeyVaultInterface {
	return &FakeFactoryLinkedServiceKeyVaults{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceKustos(namespace string) v1alpha1.FactoryLinkedServiceKustoInterface {
	return &FakeFactoryLinkedServiceKustos{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceMysqls(namespace string) v1alpha1.FactoryLinkedServiceMysqlInterface {
	return &FakeFactoryLinkedServiceMysqls{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceOdatas(namespace string) v1alpha1.FactoryLinkedServiceOdataInterface {
	return &FakeFactoryLinkedServiceOdatas{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceOdbcs(namespace string) v1alpha1.FactoryLinkedServiceOdbcInterface {
	return &FakeFactoryLinkedServiceOdbcs{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServicePostgresqls(namespace string) v1alpha1.FactoryLinkedServicePostgresqlInterface {
	return &FakeFactoryLinkedServicePostgresqls{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceSQLServers(namespace string) v1alpha1.FactoryLinkedServiceSQLServerInterface {
	return &FakeFactoryLinkedServiceSQLServers{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceSftps(namespace string) v1alpha1.FactoryLinkedServiceSftpInterface {
	return &FakeFactoryLinkedServiceSftps{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceSnowflakes(namespace string) v1alpha1.FactoryLinkedServiceSnowflakeInterface {
	return &FakeFactoryLinkedServiceSnowflakes{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceSynapses(namespace string) v1alpha1.FactoryLinkedServiceSynapseInterface {
	return &FakeFactoryLinkedServiceSynapses{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryLinkedServiceWebs(namespace string) v1alpha1.FactoryLinkedServiceWebInterface {
	return &FakeFactoryLinkedServiceWebs{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryManagedPrivateEndpoints(namespace string) v1alpha1.FactoryManagedPrivateEndpointInterface {
	return &FakeFactoryManagedPrivateEndpoints{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryPipelines(namespace string) v1alpha1.FactoryPipelineInterface {
	return &FakeFactoryPipelines{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryTriggerBlobEvents(namespace string) v1alpha1.FactoryTriggerBlobEventInterface {
	return &FakeFactoryTriggerBlobEvents{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryTriggerCustomEvents(namespace string) v1alpha1.FactoryTriggerCustomEventInterface {
	return &FakeFactoryTriggerCustomEvents{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryTriggerSchedules(namespace string) v1alpha1.FactoryTriggerScheduleInterface {
	return &FakeFactoryTriggerSchedules{c, namespace}
}

func (c *FakeDataV1alpha1) FactoryTriggerTumblingWindows(namespace string) v1alpha1.FactoryTriggerTumblingWindowInterface {
	return &FakeFactoryTriggerTumblingWindows{c, namespace}
}

func (c *FakeDataV1alpha1) LakeAnalyticsAccounts(namespace string) v1alpha1.LakeAnalyticsAccountInterface {
	return &FakeLakeAnalyticsAccounts{c, namespace}
}

func (c *FakeDataV1alpha1) LakeAnalyticsFirewallRules(namespace string) v1alpha1.LakeAnalyticsFirewallRuleInterface {
	return &FakeLakeAnalyticsFirewallRules{c, namespace}
}

func (c *FakeDataV1alpha1) LakeStores(namespace string) v1alpha1.LakeStoreInterface {
	return &FakeLakeStores{c, namespace}
}

func (c *FakeDataV1alpha1) LakeStoreFiles(namespace string) v1alpha1.LakeStoreFileInterface {
	return &FakeLakeStoreFiles{c, namespace}
}

func (c *FakeDataV1alpha1) LakeStoreFirewallRules(namespace string) v1alpha1.LakeStoreFirewallRuleInterface {
	return &FakeLakeStoreFirewallRules{c, namespace}
}

func (c *FakeDataV1alpha1) LakeStoreVirtualNetworkRules(namespace string) v1alpha1.LakeStoreVirtualNetworkRuleInterface {
	return &FakeLakeStoreVirtualNetworkRules{c, namespace}
}

func (c *FakeDataV1alpha1) ProtectionBackupInstanceBlobStorages(namespace string) v1alpha1.ProtectionBackupInstanceBlobStorageInterface {
	return &FakeProtectionBackupInstanceBlobStorages{c, namespace}
}

func (c *FakeDataV1alpha1) ProtectionBackupInstanceDisks(namespace string) v1alpha1.ProtectionBackupInstanceDiskInterface {
	return &FakeProtectionBackupInstanceDisks{c, namespace}
}

func (c *FakeDataV1alpha1) ProtectionBackupInstancePostgresqls(namespace string) v1alpha1.ProtectionBackupInstancePostgresqlInterface {
	return &FakeProtectionBackupInstancePostgresqls{c, namespace}
}

func (c *FakeDataV1alpha1) ProtectionBackupPolicyBlobStorages(namespace string) v1alpha1.ProtectionBackupPolicyBlobStorageInterface {
	return &FakeProtectionBackupPolicyBlobStorages{c, namespace}
}

func (c *FakeDataV1alpha1) ProtectionBackupPolicyDisks(namespace string) v1alpha1.ProtectionBackupPolicyDiskInterface {
	return &FakeProtectionBackupPolicyDisks{c, namespace}
}

func (c *FakeDataV1alpha1) ProtectionBackupPolicyPostgresqls(namespace string) v1alpha1.ProtectionBackupPolicyPostgresqlInterface {
	return &FakeProtectionBackupPolicyPostgresqls{c, namespace}
}

func (c *FakeDataV1alpha1) ProtectionBackupVaults(namespace string) v1alpha1.ProtectionBackupVaultInterface {
	return &FakeProtectionBackupVaults{c, namespace}
}

func (c *FakeDataV1alpha1) Shares(namespace string) v1alpha1.ShareInterface {
	return &FakeShares{c, namespace}
}

func (c *FakeDataV1alpha1) ShareAccounts(namespace string) v1alpha1.ShareAccountInterface {
	return &FakeShareAccounts{c, namespace}
}

func (c *FakeDataV1alpha1) ShareDatasetBlobStorages(namespace string) v1alpha1.ShareDatasetBlobStorageInterface {
	return &FakeShareDatasetBlobStorages{c, namespace}
}

func (c *FakeDataV1alpha1) ShareDatasetDataLakeGen1s(namespace string) v1alpha1.ShareDatasetDataLakeGen1Interface {
	return &FakeShareDatasetDataLakeGen1s{c, namespace}
}

func (c *FakeDataV1alpha1) ShareDatasetDataLakeGen2s(namespace string) v1alpha1.ShareDatasetDataLakeGen2Interface {
	return &FakeShareDatasetDataLakeGen2s{c, namespace}
}

func (c *FakeDataV1alpha1) ShareDatasetKustoClusters(namespace string) v1alpha1.ShareDatasetKustoClusterInterface {
	return &FakeShareDatasetKustoClusters{c, namespace}
}

func (c *FakeDataV1alpha1) ShareDatasetKustoDatabases(namespace string) v1alpha1.ShareDatasetKustoDatabaseInterface {
	return &FakeShareDatasetKustoDatabases{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDataV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
