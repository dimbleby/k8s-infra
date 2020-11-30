package v20191101

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/Azure/k8s-infra/apis/microsoft.containerservice/v1"
)

const (
	apiVersion = "2019-11-01"
)

func (src *ManagedCluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.ManagedCluster)

	if err := Convert_v20191101_ManagedCluster_To_v1_ManagedCluster(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *ManagedCluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.ManagedCluster)

	if err := Convert_v1_ManagedCluster_To_v20191101_ManagedCluster(src, dst, nil); err != nil {
		return err
	}

	return nil
}
