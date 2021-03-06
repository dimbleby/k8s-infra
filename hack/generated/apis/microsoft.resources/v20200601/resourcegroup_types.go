// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v20200601

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupSpec   `json:"spec,omitempty"`
	Status            ResourceGroupStatus `json:"status,omitempty"`
}

var _ genruntime.KubernetesResource = &ResourceGroup{}

// AzureName returns the Azure name of the resource
func (rg *ResourceGroup) AzureName() string {
	return rg.Spec.AzureName
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rg *ResourceGroup) Owner() *genruntime.ResourceReference {
	return nil
}

var _ genruntime.LocatableResource = &ResourceGroup{}

func (rg *ResourceGroup) Location() string {
	return rg.Spec.Location
}

// +kubebuilder:object:root=true
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

type ResourceGroupStatus struct {
	ID string `json:"id,omitempty"`

	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy string `json:"managedBy,omitempty"`

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`

	Properties *ResourceGroupStatusProperties `json:"properties,omitempty"` // TODO: Is this required or optional?
}

var _ genruntime.ArmTransformer = &ResourceGroupStatus{}

func (status *ResourceGroupStatus) CreateEmptyArmValue() interface{} {
	return ResourceGroupStatusArm{}
}

// ConvertToArm converts from a Kubernetes CRD object to an ARM object
func (status *ResourceGroupStatus) ConvertToArm(name string) (interface{}, error) {
	if status == nil {
		return nil, nil
	}
	result := ResourceGroupStatusArm{}
	result.ID = status.ID
	result.Location = status.Location
	result.ManagedBy = status.ManagedBy
	result.Name = status.Name
	result.Tags = status.Tags
	if status.Properties != nil {
		properties, err := status.Properties.ConvertToArm(name)
		if err != nil {
			return nil, err
		}
		propertiesTyped := properties.(ResourceGroupStatusPropertiesArm)
		result.Properties = &propertiesTyped
	}
	return result, nil
}

// PopulateFromArm populates a Kubernetes CRD object from an Azure ARM object
func (status *ResourceGroupStatus) PopulateFromArm(owner genruntime.KnownResourceReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroupStatusArm)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromArm() function. Expected ResourceGroupStatusArm, got %T", armInput)
	}
	status.ID = typedInput.ID
	status.Location = typedInput.Location
	status.ManagedBy = typedInput.ManagedBy
	status.Name = typedInput.Name
	status.Tags = typedInput.Tags
	var err error
	if typedInput.Properties != nil {
		properties := ResourceGroupStatusProperties{}
		err = properties.PopulateFromArm(owner, *typedInput.Properties)
		if err != nil {
			return err
		}
		status.Properties = &properties
	}
	return nil
}

type ResourceGroupStatusProperties struct {
	ProvisioningState string `json:"provisioningState,omitempty"` // TODO: Wrong, needs to be in properties
}

var _ genruntime.ArmTransformer = &ResourceGroupStatusProperties{}

func (p *ResourceGroupStatusProperties) CreateEmptyArmValue() interface{} {
	return ResourceGroupStatusPropertiesArm{}
}

// ConvertToArm converts from a Kubernetes CRD object to an ARM object
func (p *ResourceGroupStatusProperties) ConvertToArm(name string) (interface{}, error) {
	if p == nil {
		return nil, nil
	}
	result := ResourceGroupStatusPropertiesArm{}
	result.ProvisioningState = p.ProvisioningState
	return result, nil
}

// PopulateFromArm populates a Kubernetes CRD object from an Azure ARM object
func (p *ResourceGroupStatusProperties) PopulateFromArm(owner genruntime.KnownResourceReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroupStatusPropertiesArm)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromArm() function. Expected ResourceGroupStatusPropertiesArm, got %T", armInput)
	}
	p.ProvisioningState = typedInput.ProvisioningState
	return nil
}

type ResourceGroupSpec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	// +kubebuilder:validation:Required
	// Location is the Azure location for the group (eg westus2, southcentralus, etc...)
	Location string `json:"location"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy string `json:"managedBy,omitempty"`

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ArmTransformer = &ResourceGroupSpec{}

func (spec *ResourceGroupSpec) CreateEmptyArmValue() interface{} {
	return ResourceGroupSpecArm{}
}

// ConvertToArm converts from a Kubernetes CRD object to an ARM object
func (spec *ResourceGroupSpec) ConvertToArm(name string) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	result := ResourceGroupSpecArm{}
	result.ApiVersion = "2020-06-01" // TODO: Update this to match what the codegenerated resources do with APIVersion eventually
	result.Location = spec.Location
	result.Name = name
	result.ManagedBy = spec.ManagedBy
	result.Tags = spec.Tags
	result.Type = ResourceGroupTypeResourceGroup
	return result, nil
}

// PopulateFromArm populates a Kubernetes CRD object from an Azure ARM object
func (spec *ResourceGroupSpec) PopulateFromArm(owner genruntime.KnownResourceReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroupSpecArm)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromArm() function. Expected ResourceGroupSpecArm, got %T", armInput)
	}
	// spec.ApiVersion = typedInput.ApiVersion
	spec.AzureName = genruntime.ExtractKubernetesResourceNameFromArmName(typedInput.Name)
	spec.Location = typedInput.Location
	spec.ManagedBy = typedInput.ManagedBy
	spec.Tags = typedInput.Tags
	return nil
}

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
