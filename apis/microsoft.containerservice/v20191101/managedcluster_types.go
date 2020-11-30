/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20191101

import (
	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	AgentPoolProfile struct {
		Count        *int32 `json:"count,omitempty"`
		VMSize       string `json:"vmSize,omitempty"`
		Name         string `json:"name,omitempty"`
		OsDiskSizeGB *int32 `json:"osDiskSizeGB,omitempty"`
		MaxPods      *int32 `json:"macPods,omitempty"`
		// ... etc
	}

	ServicePrincipalProfile struct {
		ClientId string `json:"clientId,omitempty"`
		Secret   string `json:"secret,omitempty"`
	}

	ManagedClusterProperties struct {
		DnsPrefix               string                   `json:"dnsPrefix,omitempty"`
		KubernetesVersion       string                   `json:"kubernetesVersion,omitempty"`
		ServicePrincipalProfile *ServicePrincipalProfile `json:"servicePrincipalProfile,omitempty"`
		AgentPoolProfiles       []AgentPoolProfile       `json:"agentPoolProfiles,omitempty"`
		// ... etc
	}

	ManagedClusterSpec struct {
		// +kubebuilder:validation:Required
		ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef"`
		// +kubebuilder:validation:Required
		Location   string                   `json:"location"`
		Properties ManagedClusterProperties `json:"properties,omitempty"`
	}

	ManagedClusterStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}
)

// +kubebuilder:object:root=true

// ManagedCluster is the Schema for the managedclusters API
type ManagedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedClusterSpec   `json:"spec,omitempty"`
	Status ManagedClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedClusterList contains a list of ManagedCluster
type ManagedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedCluster{}, &ManagedClusterList{})
}
