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
	aadb2cv1alpha1 "kubeform.dev/provider-azurerm-api/apis/aadb2c/v1alpha1"
	activev1alpha1 "kubeform.dev/provider-azurerm-api/apis/active/v1alpha1"
	advancedv1alpha1 "kubeform.dev/provider-azurerm-api/apis/advanced/v1alpha1"
	analysisv1alpha1 "kubeform.dev/provider-azurerm-api/apis/analysis/v1alpha1"
	apimanagementv1alpha1 "kubeform.dev/provider-azurerm-api/apis/apimanagement/v1alpha1"
	appv1alpha1 "kubeform.dev/provider-azurerm-api/apis/app/v1alpha1"
	applicationv1alpha1 "kubeform.dev/provider-azurerm-api/apis/application/v1alpha1"
	attestationv1alpha1 "kubeform.dev/provider-azurerm-api/apis/attestation/v1alpha1"
	automationv1alpha1 "kubeform.dev/provider-azurerm-api/apis/automation/v1alpha1"
	availabilityv1alpha1 "kubeform.dev/provider-azurerm-api/apis/availability/v1alpha1"
	backupv1alpha1 "kubeform.dev/provider-azurerm-api/apis/backup/v1alpha1"
	bastionv1alpha1 "kubeform.dev/provider-azurerm-api/apis/bastion/v1alpha1"
	batchv1alpha1 "kubeform.dev/provider-azurerm-api/apis/batch/v1alpha1"
	blueprintv1alpha1 "kubeform.dev/provider-azurerm-api/apis/blueprint/v1alpha1"
	botv1alpha1 "kubeform.dev/provider-azurerm-api/apis/bot/v1alpha1"
	cdnv1alpha1 "kubeform.dev/provider-azurerm-api/apis/cdn/v1alpha1"
	cognitivev1alpha1 "kubeform.dev/provider-azurerm-api/apis/cognitive/v1alpha1"
	communicationv1alpha1 "kubeform.dev/provider-azurerm-api/apis/communication/v1alpha1"
	consumptionv1alpha1 "kubeform.dev/provider-azurerm-api/apis/consumption/v1alpha1"
	containerv1alpha1 "kubeform.dev/provider-azurerm-api/apis/container/v1alpha1"
	cosmosdbv1alpha1 "kubeform.dev/provider-azurerm-api/apis/cosmosdb/v1alpha1"
	costv1alpha1 "kubeform.dev/provider-azurerm-api/apis/cost/v1alpha1"
	customv1alpha1 "kubeform.dev/provider-azurerm-api/apis/custom/v1alpha1"
	dashboardv1alpha1 "kubeform.dev/provider-azurerm-api/apis/dashboard/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"
	databasev1alpha1 "kubeform.dev/provider-azurerm-api/apis/database/v1alpha1"
	databoxv1alpha1 "kubeform.dev/provider-azurerm-api/apis/databox/v1alpha1"
	databricksv1alpha1 "kubeform.dev/provider-azurerm-api/apis/databricks/v1alpha1"
	dedicatedhardwarev1alpha1 "kubeform.dev/provider-azurerm-api/apis/dedicatedhardware/v1alpha1"
	dedicatedhostv1alpha1 "kubeform.dev/provider-azurerm-api/apis/dedicatedhost/v1alpha1"
	devspacev1alpha1 "kubeform.dev/provider-azurerm-api/apis/devspace/v1alpha1"
	devtestv1alpha1 "kubeform.dev/provider-azurerm-api/apis/devtest/v1alpha1"
	digitalv1alpha1 "kubeform.dev/provider-azurerm-api/apis/digital/v1alpha1"
	diskv1alpha1 "kubeform.dev/provider-azurerm-api/apis/disk/v1alpha1"
	dnsv1alpha1 "kubeform.dev/provider-azurerm-api/apis/dns/v1alpha1"
	eventgridv1alpha1 "kubeform.dev/provider-azurerm-api/apis/eventgrid/v1alpha1"
	eventhubv1alpha1 "kubeform.dev/provider-azurerm-api/apis/eventhub/v1alpha1"
	expressroutev1alpha1 "kubeform.dev/provider-azurerm-api/apis/expressroute/v1alpha1"
	firewallv1alpha1 "kubeform.dev/provider-azurerm-api/apis/firewall/v1alpha1"
	frontdoorv1alpha1 "kubeform.dev/provider-azurerm-api/apis/frontdoor/v1alpha1"
	functionv1alpha1 "kubeform.dev/provider-azurerm-api/apis/function/v1alpha1"
	hdinsightv1alpha1 "kubeform.dev/provider-azurerm-api/apis/hdinsight/v1alpha1"
	healthbotv1alpha1 "kubeform.dev/provider-azurerm-api/apis/healthbot/v1alpha1"
	healthcarev1alpha1 "kubeform.dev/provider-azurerm-api/apis/healthcare/v1alpha1"
	hpcv1alpha1 "kubeform.dev/provider-azurerm-api/apis/hpc/v1alpha1"
	imagev1alpha1 "kubeform.dev/provider-azurerm-api/apis/image/v1alpha1"
	integrationv1alpha1 "kubeform.dev/provider-azurerm-api/apis/integration/v1alpha1"
	iotcentralv1alpha1 "kubeform.dev/provider-azurerm-api/apis/iotcentral/v1alpha1"
	iothubv1alpha1 "kubeform.dev/provider-azurerm-api/apis/iothub/v1alpha1"
	iotsecurityv1alpha1 "kubeform.dev/provider-azurerm-api/apis/iotsecurity/v1alpha1"
	iottimev1alpha1 "kubeform.dev/provider-azurerm-api/apis/iottime/v1alpha1"
	ipv1alpha1 "kubeform.dev/provider-azurerm-api/apis/ip/v1alpha1"
	keyvaultv1alpha1 "kubeform.dev/provider-azurerm-api/apis/keyvault/v1alpha1"
	kubernetesclusterv1alpha1 "kubeform.dev/provider-azurerm-api/apis/kubernetescluster/v1alpha1"
	kustov1alpha1 "kubeform.dev/provider-azurerm-api/apis/kusto/v1alpha1"
	lbv1alpha1 "kubeform.dev/provider-azurerm-api/apis/lb/v1alpha1"
	lighthousev1alpha1 "kubeform.dev/provider-azurerm-api/apis/lighthouse/v1alpha1"
	linuxv1alpha1 "kubeform.dev/provider-azurerm-api/apis/linux/v1alpha1"
	loadv1alpha1 "kubeform.dev/provider-azurerm-api/apis/load/v1alpha1"
	localv1alpha1 "kubeform.dev/provider-azurerm-api/apis/local/v1alpha1"
	loganalyticsv1alpha1 "kubeform.dev/provider-azurerm-api/apis/loganalytics/v1alpha1"
	logicappv1alpha1 "kubeform.dev/provider-azurerm-api/apis/logicapp/v1alpha1"
	logzv1alpha1 "kubeform.dev/provider-azurerm-api/apis/logz/v1alpha1"
	machinev1alpha1 "kubeform.dev/provider-azurerm-api/apis/machine/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-azurerm-api/apis/maintenance/v1alpha1"
	managedv1alpha1 "kubeform.dev/provider-azurerm-api/apis/managed/v1alpha1"
	managementv1alpha1 "kubeform.dev/provider-azurerm-api/apis/management/v1alpha1"
	mapsv1alpha1 "kubeform.dev/provider-azurerm-api/apis/maps/v1alpha1"
	mariadbv1alpha1 "kubeform.dev/provider-azurerm-api/apis/mariadb/v1alpha1"
	marketplacev1alpha1 "kubeform.dev/provider-azurerm-api/apis/marketplace/v1alpha1"
	mediav1alpha1 "kubeform.dev/provider-azurerm-api/apis/media/v1alpha1"
	monitorv1alpha1 "kubeform.dev/provider-azurerm-api/apis/monitor/v1alpha1"
	mssqlv1alpha1 "kubeform.dev/provider-azurerm-api/apis/mssql/v1alpha1"
	mysqlv1alpha1 "kubeform.dev/provider-azurerm-api/apis/mysql/v1alpha1"
	natv1alpha1 "kubeform.dev/provider-azurerm-api/apis/nat/v1alpha1"
	netappv1alpha1 "kubeform.dev/provider-azurerm-api/apis/netapp/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-azurerm-api/apis/network/v1alpha1"
	notificationhubv1alpha1 "kubeform.dev/provider-azurerm-api/apis/notificationhub/v1alpha1"
	orchestratedv1alpha1 "kubeform.dev/provider-azurerm-api/apis/orchestrated/v1alpha1"
	packetv1alpha1 "kubeform.dev/provider-azurerm-api/apis/packet/v1alpha1"
	pointv1alpha1 "kubeform.dev/provider-azurerm-api/apis/point/v1alpha1"
	policyv1alpha1 "kubeform.dev/provider-azurerm-api/apis/policy/v1alpha1"
	portalv1alpha1 "kubeform.dev/provider-azurerm-api/apis/portal/v1alpha1"
	postgresqlv1alpha1 "kubeform.dev/provider-azurerm-api/apis/postgresql/v1alpha1"
	powerbiv1alpha1 "kubeform.dev/provider-azurerm-api/apis/powerbi/v1alpha1"
	privatev1alpha1 "kubeform.dev/provider-azurerm-api/apis/private/v1alpha1"
	proximityv1alpha1 "kubeform.dev/provider-azurerm-api/apis/proximity/v1alpha1"
	publicipv1alpha1 "kubeform.dev/provider-azurerm-api/apis/publicip/v1alpha1"
	purviewv1alpha1 "kubeform.dev/provider-azurerm-api/apis/purview/v1alpha1"
	recoveryv1alpha1 "kubeform.dev/provider-azurerm-api/apis/recovery/v1alpha1"
	redisv1alpha1 "kubeform.dev/provider-azurerm-api/apis/redis/v1alpha1"
	relayv1alpha1 "kubeform.dev/provider-azurerm-api/apis/relay/v1alpha1"
	resourcev1alpha1 "kubeform.dev/provider-azurerm-api/apis/resource/v1alpha1"
	rolev1alpha1 "kubeform.dev/provider-azurerm-api/apis/role/v1alpha1"
	routev1alpha1 "kubeform.dev/provider-azurerm-api/apis/route/v1alpha1"
	searchv1alpha1 "kubeform.dev/provider-azurerm-api/apis/search/v1alpha1"
	securityv1alpha1 "kubeform.dev/provider-azurerm-api/apis/security/v1alpha1"
	sentinelv1alpha1 "kubeform.dev/provider-azurerm-api/apis/sentinel/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-azurerm-api/apis/service/v1alpha1"
	servicebusv1alpha1 "kubeform.dev/provider-azurerm-api/apis/servicebus/v1alpha1"
	sharedimagev1alpha1 "kubeform.dev/provider-azurerm-api/apis/sharedimage/v1alpha1"
	signalrv1alpha1 "kubeform.dev/provider-azurerm-api/apis/signalr/v1alpha1"
	siterecoveryv1alpha1 "kubeform.dev/provider-azurerm-api/apis/siterecovery/v1alpha1"
	snapshotv1alpha1 "kubeform.dev/provider-azurerm-api/apis/snapshot/v1alpha1"
	spatialv1alpha1 "kubeform.dev/provider-azurerm-api/apis/spatial/v1alpha1"
	springv1alpha1 "kubeform.dev/provider-azurerm-api/apis/spring/v1alpha1"
	sqlv1alpha1 "kubeform.dev/provider-azurerm-api/apis/sql/v1alpha1"
	sshv1alpha1 "kubeform.dev/provider-azurerm-api/apis/ssh/v1alpha1"
	stackv1alpha1 "kubeform.dev/provider-azurerm-api/apis/stack/v1alpha1"
	staticv1alpha1 "kubeform.dev/provider-azurerm-api/apis/static/v1alpha1"
	storagev1alpha1 "kubeform.dev/provider-azurerm-api/apis/storage/v1alpha1"
	streamv1alpha1 "kubeform.dev/provider-azurerm-api/apis/stream/v1alpha1"
	subnetv1alpha1 "kubeform.dev/provider-azurerm-api/apis/subnet/v1alpha1"
	subscriptionv1alpha1 "kubeform.dev/provider-azurerm-api/apis/subscription/v1alpha1"
	synapsev1alpha1 "kubeform.dev/provider-azurerm-api/apis/synapse/v1alpha1"
	templatev1alpha1 "kubeform.dev/provider-azurerm-api/apis/template/v1alpha1"
	tenantv1alpha1 "kubeform.dev/provider-azurerm-api/apis/tenant/v1alpha1"
	trafficmanagerv1alpha1 "kubeform.dev/provider-azurerm-api/apis/trafficmanager/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-azurerm-api/apis/user/v1alpha1"
	videov1alpha1 "kubeform.dev/provider-azurerm-api/apis/video/v1alpha1"
	virtualv1alpha1 "kubeform.dev/provider-azurerm-api/apis/virtual/v1alpha1"
	vmwarev1alpha1 "kubeform.dev/provider-azurerm-api/apis/vmware/v1alpha1"
	vpnv1alpha1 "kubeform.dev/provider-azurerm-api/apis/vpn/v1alpha1"
	webv1alpha1 "kubeform.dev/provider-azurerm-api/apis/web/v1alpha1"
	windowsv1alpha1 "kubeform.dev/provider-azurerm-api/apis/windows/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)

var localSchemeBuilder = runtime.SchemeBuilder{
	aadb2cv1alpha1.AddToScheme,
	activev1alpha1.AddToScheme,
	advancedv1alpha1.AddToScheme,
	analysisv1alpha1.AddToScheme,
	apimanagementv1alpha1.AddToScheme,
	appv1alpha1.AddToScheme,
	applicationv1alpha1.AddToScheme,
	attestationv1alpha1.AddToScheme,
	automationv1alpha1.AddToScheme,
	availabilityv1alpha1.AddToScheme,
	backupv1alpha1.AddToScheme,
	bastionv1alpha1.AddToScheme,
	batchv1alpha1.AddToScheme,
	blueprintv1alpha1.AddToScheme,
	botv1alpha1.AddToScheme,
	cdnv1alpha1.AddToScheme,
	cognitivev1alpha1.AddToScheme,
	communicationv1alpha1.AddToScheme,
	consumptionv1alpha1.AddToScheme,
	containerv1alpha1.AddToScheme,
	cosmosdbv1alpha1.AddToScheme,
	costv1alpha1.AddToScheme,
	customv1alpha1.AddToScheme,
	dashboardv1alpha1.AddToScheme,
	datav1alpha1.AddToScheme,
	databasev1alpha1.AddToScheme,
	databoxv1alpha1.AddToScheme,
	databricksv1alpha1.AddToScheme,
	dedicatedhardwarev1alpha1.AddToScheme,
	dedicatedhostv1alpha1.AddToScheme,
	devspacev1alpha1.AddToScheme,
	devtestv1alpha1.AddToScheme,
	digitalv1alpha1.AddToScheme,
	diskv1alpha1.AddToScheme,
	dnsv1alpha1.AddToScheme,
	eventgridv1alpha1.AddToScheme,
	eventhubv1alpha1.AddToScheme,
	expressroutev1alpha1.AddToScheme,
	firewallv1alpha1.AddToScheme,
	frontdoorv1alpha1.AddToScheme,
	functionv1alpha1.AddToScheme,
	hdinsightv1alpha1.AddToScheme,
	healthbotv1alpha1.AddToScheme,
	healthcarev1alpha1.AddToScheme,
	hpcv1alpha1.AddToScheme,
	imagev1alpha1.AddToScheme,
	integrationv1alpha1.AddToScheme,
	iotcentralv1alpha1.AddToScheme,
	iothubv1alpha1.AddToScheme,
	iotsecurityv1alpha1.AddToScheme,
	iottimev1alpha1.AddToScheme,
	ipv1alpha1.AddToScheme,
	keyvaultv1alpha1.AddToScheme,
	kubernetesclusterv1alpha1.AddToScheme,
	kustov1alpha1.AddToScheme,
	lbv1alpha1.AddToScheme,
	lighthousev1alpha1.AddToScheme,
	linuxv1alpha1.AddToScheme,
	loadv1alpha1.AddToScheme,
	localv1alpha1.AddToScheme,
	loganalyticsv1alpha1.AddToScheme,
	logicappv1alpha1.AddToScheme,
	logzv1alpha1.AddToScheme,
	machinev1alpha1.AddToScheme,
	maintenancev1alpha1.AddToScheme,
	managedv1alpha1.AddToScheme,
	managementv1alpha1.AddToScheme,
	mapsv1alpha1.AddToScheme,
	mariadbv1alpha1.AddToScheme,
	marketplacev1alpha1.AddToScheme,
	mediav1alpha1.AddToScheme,
	monitorv1alpha1.AddToScheme,
	mssqlv1alpha1.AddToScheme,
	mysqlv1alpha1.AddToScheme,
	natv1alpha1.AddToScheme,
	netappv1alpha1.AddToScheme,
	networkv1alpha1.AddToScheme,
	notificationhubv1alpha1.AddToScheme,
	orchestratedv1alpha1.AddToScheme,
	packetv1alpha1.AddToScheme,
	pointv1alpha1.AddToScheme,
	policyv1alpha1.AddToScheme,
	portalv1alpha1.AddToScheme,
	postgresqlv1alpha1.AddToScheme,
	powerbiv1alpha1.AddToScheme,
	privatev1alpha1.AddToScheme,
	proximityv1alpha1.AddToScheme,
	publicipv1alpha1.AddToScheme,
	purviewv1alpha1.AddToScheme,
	recoveryv1alpha1.AddToScheme,
	redisv1alpha1.AddToScheme,
	relayv1alpha1.AddToScheme,
	resourcev1alpha1.AddToScheme,
	rolev1alpha1.AddToScheme,
	routev1alpha1.AddToScheme,
	searchv1alpha1.AddToScheme,
	securityv1alpha1.AddToScheme,
	sentinelv1alpha1.AddToScheme,
	servicev1alpha1.AddToScheme,
	servicebusv1alpha1.AddToScheme,
	sharedimagev1alpha1.AddToScheme,
	signalrv1alpha1.AddToScheme,
	siterecoveryv1alpha1.AddToScheme,
	snapshotv1alpha1.AddToScheme,
	spatialv1alpha1.AddToScheme,
	springv1alpha1.AddToScheme,
	sqlv1alpha1.AddToScheme,
	sshv1alpha1.AddToScheme,
	stackv1alpha1.AddToScheme,
	staticv1alpha1.AddToScheme,
	storagev1alpha1.AddToScheme,
	streamv1alpha1.AddToScheme,
	subnetv1alpha1.AddToScheme,
	subscriptionv1alpha1.AddToScheme,
	synapsev1alpha1.AddToScheme,
	templatev1alpha1.AddToScheme,
	tenantv1alpha1.AddToScheme,
	trafficmanagerv1alpha1.AddToScheme,
	userv1alpha1.AddToScheme,
	videov1alpha1.AddToScheme,
	virtualv1alpha1.AddToScheme,
	vmwarev1alpha1.AddToScheme,
	vpnv1alpha1.AddToScheme,
	webv1alpha1.AddToScheme,
	windowsv1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
