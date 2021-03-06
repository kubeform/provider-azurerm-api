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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"fmt"
	"strings"

	base "kubeform.dev/apimachinery/api/v1alpha1"
	"kubeform.dev/apimachinery/pkg/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *HadoopCluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-hdinsight-azurerm-kubeform-com-v1alpha1-hadoopcluster,mutating=false,failurePolicy=fail,groups=hdinsight.azurerm.kubeform.com,resources=hadoopclusters,versions=v1alpha1,name=hadoopcluster.hdinsight.azurerm.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &HadoopCluster{}

var hadoopclusterForceNewList = map[string]bool{
	"/cluster_version":                                     true,
	"/component_version/*/hadoop":                          true,
	"/gateway/*/username":                                  true,
	"/location":                                            true,
	"/metastores/*/ambari/*/database_name":                 true,
	"/metastores/*/ambari/*/server":                        true,
	"/metastores/*/ambari/*/username":                      true,
	"/metastores/*/hive/*/database_name":                   true,
	"/metastores/*/hive/*/server":                          true,
	"/metastores/*/hive/*/username":                        true,
	"/metastores/*/oozie/*/database_name":                  true,
	"/metastores/*/oozie/*/server":                         true,
	"/metastores/*/oozie/*/username":                       true,
	"/name":                                                true,
	"/network/*/connection_direction":                      true,
	"/network/*/private_link_enabled":                      true,
	"/resource_group_name":                                 true,
	"/roles/*/head_node/*/ssh_keys":                        true,
	"/roles/*/head_node/*/subnet_id":                       true,
	"/roles/*/head_node/*/username":                        true,
	"/roles/*/head_node/*/virtual_network_id":              true,
	"/roles/*/head_node/*/vm_size":                         true,
	"/roles/*/worker_node/*/min_instance_count":            true,
	"/roles/*/worker_node/*/ssh_keys":                      true,
	"/roles/*/worker_node/*/subnet_id":                     true,
	"/roles/*/worker_node/*/username":                      true,
	"/roles/*/worker_node/*/virtual_network_id":            true,
	"/roles/*/worker_node/*/vm_size":                       true,
	"/roles/*/zookeeper_node/*/ssh_keys":                   true,
	"/roles/*/zookeeper_node/*/subnet_id":                  true,
	"/roles/*/zookeeper_node/*/username":                   true,
	"/roles/*/zookeeper_node/*/virtual_network_id":         true,
	"/roles/*/zookeeper_node/*/vm_size":                    true,
	"/security_profile/*/aadds_resource_id":                true,
	"/security_profile/*/cluster_users_group_dns":          true,
	"/security_profile/*/domain_name":                      true,
	"/security_profile/*/domain_username":                  true,
	"/security_profile/*/ldaps_urls":                       true,
	"/security_profile/*/msi_resource_id":                  true,
	"/storage_account/*/is_default":                        true,
	"/storage_account/*/storage_container_id":              true,
	"/storage_account/*/storage_resource_id":               true,
	"/storage_account_gen2/*/filesystem_id":                true,
	"/storage_account_gen2/*/is_default":                   true,
	"/storage_account_gen2/*/managed_identity_resource_id": true,
	"/storage_account_gen2/*/storage_resource_id":          true,
	"/tier":            true,
	"/tls_min_version": true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *HadoopCluster) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *HadoopCluster) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*HadoopCluster)
	oldObj := res.Spec.Resource

	jsnitr := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		TagKey:                 "tf",
		ValidateJsonRawMessage: true,
		TypeEncoders:           GetEncoder(),
		TypeDecoders:           GetDecoder(),
	}.Froze()

	byteNew, err := jsnitr.Marshal(newObj)
	if err != nil {
		return err
	}
	tempNew := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteNew, &tempNew)
	if err != nil {
		return err
	}

	byteOld, err := jsnitr.Marshal(oldObj)
	if err != nil {
		return err
	}
	tempOld := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteOld, &tempOld)
	if err != nil {
		return err
	}

	for key, _ := range hadoopclusterForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`hadoopcluster "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *HadoopCluster) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`hadoopcluster "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
