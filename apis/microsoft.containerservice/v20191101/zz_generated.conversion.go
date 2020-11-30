// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/
// Code generated by conversion-gen. DO NOT EDIT.

package v20191101

import (
	unsafe "unsafe"

	corev1 "github.com/Azure/k8s-infra/apis/core/v1"
	v1 "github.com/Azure/k8s-infra/apis/microsoft.containerservice/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AgentPoolProfile)(nil), (*v1.AgentPoolProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v20191101_AgentPoolProfile_To_v1_AgentPoolProfile(a.(*AgentPoolProfile), b.(*v1.AgentPoolProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.AgentPoolProfile)(nil), (*AgentPoolProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AgentPoolProfile_To_v20191101_AgentPoolProfile(a.(*v1.AgentPoolProfile), b.(*AgentPoolProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedCluster)(nil), (*v1.ManagedCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v20191101_ManagedCluster_To_v1_ManagedCluster(a.(*ManagedCluster), b.(*v1.ManagedCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ManagedCluster)(nil), (*ManagedCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ManagedCluster_To_v20191101_ManagedCluster(a.(*v1.ManagedCluster), b.(*ManagedCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedClusterList)(nil), (*v1.ManagedClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v20191101_ManagedClusterList_To_v1_ManagedClusterList(a.(*ManagedClusterList), b.(*v1.ManagedClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ManagedClusterList)(nil), (*ManagedClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ManagedClusterList_To_v20191101_ManagedClusterList(a.(*v1.ManagedClusterList), b.(*ManagedClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedClusterProperties)(nil), (*v1.ManagedClusterProperties)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v20191101_ManagedClusterProperties_To_v1_ManagedClusterProperties(a.(*ManagedClusterProperties), b.(*v1.ManagedClusterProperties), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ManagedClusterProperties)(nil), (*ManagedClusterProperties)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ManagedClusterProperties_To_v20191101_ManagedClusterProperties(a.(*v1.ManagedClusterProperties), b.(*ManagedClusterProperties), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedClusterSpec)(nil), (*v1.ManagedClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v20191101_ManagedClusterSpec_To_v1_ManagedClusterSpec(a.(*ManagedClusterSpec), b.(*v1.ManagedClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ManagedClusterSpec)(nil), (*ManagedClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ManagedClusterSpec_To_v20191101_ManagedClusterSpec(a.(*v1.ManagedClusterSpec), b.(*ManagedClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedClusterStatus)(nil), (*v1.ManagedClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v20191101_ManagedClusterStatus_To_v1_ManagedClusterStatus(a.(*ManagedClusterStatus), b.(*v1.ManagedClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ManagedClusterStatus)(nil), (*ManagedClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ManagedClusterStatus_To_v20191101_ManagedClusterStatus(a.(*v1.ManagedClusterStatus), b.(*ManagedClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServicePrincipalProfile)(nil), (*v1.ServicePrincipalProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v20191101_ServicePrincipalProfile_To_v1_ServicePrincipalProfile(a.(*ServicePrincipalProfile), b.(*v1.ServicePrincipalProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1.ServicePrincipalProfile)(nil), (*ServicePrincipalProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServicePrincipalProfile_To_v20191101_ServicePrincipalProfile(a.(*v1.ServicePrincipalProfile), b.(*ServicePrincipalProfile), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v20191101_AgentPoolProfile_To_v1_AgentPoolProfile(in *AgentPoolProfile, out *v1.AgentPoolProfile, s conversion.Scope) error {
	out.Count = (*int32)(unsafe.Pointer(in.Count))
	out.VMSize = in.VMSize
	out.Name = in.Name
	out.OsDiskSizeGB = (*int32)(unsafe.Pointer(in.OsDiskSizeGB))
	out.MaxPods = (*int32)(unsafe.Pointer(in.MaxPods))
	return nil
}

// Convert_v20191101_AgentPoolProfile_To_v1_AgentPoolProfile is an autogenerated conversion function.
func Convert_v20191101_AgentPoolProfile_To_v1_AgentPoolProfile(in *AgentPoolProfile, out *v1.AgentPoolProfile, s conversion.Scope) error {
	return autoConvert_v20191101_AgentPoolProfile_To_v1_AgentPoolProfile(in, out, s)
}

func autoConvert_v1_AgentPoolProfile_To_v20191101_AgentPoolProfile(in *v1.AgentPoolProfile, out *AgentPoolProfile, s conversion.Scope) error {
	out.Count = (*int32)(unsafe.Pointer(in.Count))
	out.VMSize = in.VMSize
	out.Name = in.Name
	out.OsDiskSizeGB = (*int32)(unsafe.Pointer(in.OsDiskSizeGB))
	out.MaxPods = (*int32)(unsafe.Pointer(in.MaxPods))
	return nil
}

// Convert_v1_AgentPoolProfile_To_v20191101_AgentPoolProfile is an autogenerated conversion function.
func Convert_v1_AgentPoolProfile_To_v20191101_AgentPoolProfile(in *v1.AgentPoolProfile, out *AgentPoolProfile, s conversion.Scope) error {
	return autoConvert_v1_AgentPoolProfile_To_v20191101_AgentPoolProfile(in, out, s)
}

func autoConvert_v20191101_ManagedCluster_To_v1_ManagedCluster(in *ManagedCluster, out *v1.ManagedCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v20191101_ManagedClusterSpec_To_v1_ManagedClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v20191101_ManagedClusterStatus_To_v1_ManagedClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v20191101_ManagedCluster_To_v1_ManagedCluster is an autogenerated conversion function.
func Convert_v20191101_ManagedCluster_To_v1_ManagedCluster(in *ManagedCluster, out *v1.ManagedCluster, s conversion.Scope) error {
	return autoConvert_v20191101_ManagedCluster_To_v1_ManagedCluster(in, out, s)
}

func autoConvert_v1_ManagedCluster_To_v20191101_ManagedCluster(in *v1.ManagedCluster, out *ManagedCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ManagedClusterSpec_To_v20191101_ManagedClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ManagedClusterStatus_To_v20191101_ManagedClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ManagedCluster_To_v20191101_ManagedCluster is an autogenerated conversion function.
func Convert_v1_ManagedCluster_To_v20191101_ManagedCluster(in *v1.ManagedCluster, out *ManagedCluster, s conversion.Scope) error {
	return autoConvert_v1_ManagedCluster_To_v20191101_ManagedCluster(in, out, s)
}

func autoConvert_v20191101_ManagedClusterList_To_v1_ManagedClusterList(in *ManagedClusterList, out *v1.ManagedClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ManagedCluster, len(*in))
		for i := range *in {
			if err := Convert_v20191101_ManagedCluster_To_v1_ManagedCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v20191101_ManagedClusterList_To_v1_ManagedClusterList is an autogenerated conversion function.
func Convert_v20191101_ManagedClusterList_To_v1_ManagedClusterList(in *ManagedClusterList, out *v1.ManagedClusterList, s conversion.Scope) error {
	return autoConvert_v20191101_ManagedClusterList_To_v1_ManagedClusterList(in, out, s)
}

func autoConvert_v1_ManagedClusterList_To_v20191101_ManagedClusterList(in *v1.ManagedClusterList, out *ManagedClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedCluster, len(*in))
		for i := range *in {
			if err := Convert_v1_ManagedCluster_To_v20191101_ManagedCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_ManagedClusterList_To_v20191101_ManagedClusterList is an autogenerated conversion function.
func Convert_v1_ManagedClusterList_To_v20191101_ManagedClusterList(in *v1.ManagedClusterList, out *ManagedClusterList, s conversion.Scope) error {
	return autoConvert_v1_ManagedClusterList_To_v20191101_ManagedClusterList(in, out, s)
}

func autoConvert_v20191101_ManagedClusterProperties_To_v1_ManagedClusterProperties(in *ManagedClusterProperties, out *v1.ManagedClusterProperties, s conversion.Scope) error {
	out.DnsPrefix = in.DnsPrefix
	out.KubernetesVersion = in.KubernetesVersion
	out.ServicePrincipalProfile = (*v1.ServicePrincipalProfile)(unsafe.Pointer(in.ServicePrincipalProfile))
	out.AgentPoolProfiles = *(*[]v1.AgentPoolProfile)(unsafe.Pointer(&in.AgentPoolProfiles))
	return nil
}

// Convert_v20191101_ManagedClusterProperties_To_v1_ManagedClusterProperties is an autogenerated conversion function.
func Convert_v20191101_ManagedClusterProperties_To_v1_ManagedClusterProperties(in *ManagedClusterProperties, out *v1.ManagedClusterProperties, s conversion.Scope) error {
	return autoConvert_v20191101_ManagedClusterProperties_To_v1_ManagedClusterProperties(in, out, s)
}

func autoConvert_v1_ManagedClusterProperties_To_v20191101_ManagedClusterProperties(in *v1.ManagedClusterProperties, out *ManagedClusterProperties, s conversion.Scope) error {
	out.DnsPrefix = in.DnsPrefix
	out.KubernetesVersion = in.KubernetesVersion
	out.ServicePrincipalProfile = (*ServicePrincipalProfile)(unsafe.Pointer(in.ServicePrincipalProfile))
	out.AgentPoolProfiles = *(*[]AgentPoolProfile)(unsafe.Pointer(&in.AgentPoolProfiles))
	return nil
}

// Convert_v1_ManagedClusterProperties_To_v20191101_ManagedClusterProperties is an autogenerated conversion function.
func Convert_v1_ManagedClusterProperties_To_v20191101_ManagedClusterProperties(in *v1.ManagedClusterProperties, out *ManagedClusterProperties, s conversion.Scope) error {
	return autoConvert_v1_ManagedClusterProperties_To_v20191101_ManagedClusterProperties(in, out, s)
}

func autoConvert_v20191101_ManagedClusterSpec_To_v1_ManagedClusterSpec(in *ManagedClusterSpec, out *v1.ManagedClusterSpec, s conversion.Scope) error {
	out.ResourceGroupRef = (*corev1.KnownTypeReference)(unsafe.Pointer(in.ResourceGroupRef))
	out.Location = in.Location
	if err := Convert_v20191101_ManagedClusterProperties_To_v1_ManagedClusterProperties(&in.Properties, &out.Properties, s); err != nil {
		return err
	}
	return nil
}

// Convert_v20191101_ManagedClusterSpec_To_v1_ManagedClusterSpec is an autogenerated conversion function.
func Convert_v20191101_ManagedClusterSpec_To_v1_ManagedClusterSpec(in *ManagedClusterSpec, out *v1.ManagedClusterSpec, s conversion.Scope) error {
	return autoConvert_v20191101_ManagedClusterSpec_To_v1_ManagedClusterSpec(in, out, s)
}

func autoConvert_v1_ManagedClusterSpec_To_v20191101_ManagedClusterSpec(in *v1.ManagedClusterSpec, out *ManagedClusterSpec, s conversion.Scope) error {
	// INFO: in.APIVersion opted out of conversion generation
	out.ResourceGroupRef = (*corev1.KnownTypeReference)(unsafe.Pointer(in.ResourceGroupRef))
	out.Location = in.Location
	if err := Convert_v1_ManagedClusterProperties_To_v20191101_ManagedClusterProperties(&in.Properties, &out.Properties, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ManagedClusterSpec_To_v20191101_ManagedClusterSpec is an autogenerated conversion function.
func Convert_v1_ManagedClusterSpec_To_v20191101_ManagedClusterSpec(in *v1.ManagedClusterSpec, out *ManagedClusterSpec, s conversion.Scope) error {
	return autoConvert_v1_ManagedClusterSpec_To_v20191101_ManagedClusterSpec(in, out, s)
}

func autoConvert_v20191101_ManagedClusterStatus_To_v1_ManagedClusterStatus(in *ManagedClusterStatus, out *v1.ManagedClusterStatus, s conversion.Scope) error {
	out.ID = in.ID
	out.ProvisioningState = in.ProvisioningState
	return nil
}

// Convert_v20191101_ManagedClusterStatus_To_v1_ManagedClusterStatus is an autogenerated conversion function.
func Convert_v20191101_ManagedClusterStatus_To_v1_ManagedClusterStatus(in *ManagedClusterStatus, out *v1.ManagedClusterStatus, s conversion.Scope) error {
	return autoConvert_v20191101_ManagedClusterStatus_To_v1_ManagedClusterStatus(in, out, s)
}

func autoConvert_v1_ManagedClusterStatus_To_v20191101_ManagedClusterStatus(in *v1.ManagedClusterStatus, out *ManagedClusterStatus, s conversion.Scope) error {
	out.ID = in.ID
	// INFO: in.DeploymentID opted out of conversion generation
	out.ProvisioningState = in.ProvisioningState
	return nil
}

// Convert_v1_ManagedClusterStatus_To_v20191101_ManagedClusterStatus is an autogenerated conversion function.
func Convert_v1_ManagedClusterStatus_To_v20191101_ManagedClusterStatus(in *v1.ManagedClusterStatus, out *ManagedClusterStatus, s conversion.Scope) error {
	return autoConvert_v1_ManagedClusterStatus_To_v20191101_ManagedClusterStatus(in, out, s)
}

func autoConvert_v20191101_ServicePrincipalProfile_To_v1_ServicePrincipalProfile(in *ServicePrincipalProfile, out *v1.ServicePrincipalProfile, s conversion.Scope) error {
	out.ClientId = in.ClientId
	out.Secret = in.Secret
	return nil
}

// Convert_v20191101_ServicePrincipalProfile_To_v1_ServicePrincipalProfile is an autogenerated conversion function.
func Convert_v20191101_ServicePrincipalProfile_To_v1_ServicePrincipalProfile(in *ServicePrincipalProfile, out *v1.ServicePrincipalProfile, s conversion.Scope) error {
	return autoConvert_v20191101_ServicePrincipalProfile_To_v1_ServicePrincipalProfile(in, out, s)
}

func autoConvert_v1_ServicePrincipalProfile_To_v20191101_ServicePrincipalProfile(in *v1.ServicePrincipalProfile, out *ServicePrincipalProfile, s conversion.Scope) error {
	out.ClientId = in.ClientId
	out.Secret = in.Secret
	return nil
}

// Convert_v1_ServicePrincipalProfile_To_v20191101_ServicePrincipalProfile is an autogenerated conversion function.
func Convert_v1_ServicePrincipalProfile_To_v20191101_ServicePrincipalProfile(in *v1.ServicePrincipalProfile, out *ServicePrincipalProfile, s conversion.Scope) error {
	return autoConvert_v1_ServicePrincipalProfile_To_v20191101_ServicePrincipalProfile(in, out, s)
}