package v1

import (
	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

func (mc *ManagedCluster) GetResourceGroupObjectRef() *azcorev1.KnownTypeReference {
	return mc.Spec.ResourceGroupRef
}

func (*ManagedCluster) ResourceType() string {
	return "Microsoft.ContainerService/managedClusters"
}
