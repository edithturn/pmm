// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/dbaas/pxc_clusters.proto

package dbaasv1beta1

import (
	fmt "fmt"
	math "math"
	regexp "regexp"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *PXCClusterParams) Validate() error {
	if !(this.ClusterSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.ClusterSize))
	}
	if nil == this.Pxc {
		return github_com_mwitkow_go_proto_validators.FieldError("Pxc", fmt.Errorf("message must exist"))
	}
	if this.Pxc != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pxc); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pxc", err)
		}
	}
	if this.Proxysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proxysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", err)
		}
	}
	if this.Haproxy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Haproxy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Haproxy", err)
		}
	}
	return nil
}

func (this *PXCClusterParams_PXC) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	if !(this.DiskSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("DiskSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.DiskSize))
	}
	return nil
}

func (this *PXCClusterParams_ProxySQL) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	if !(this.DiskSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("DiskSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.DiskSize))
	}
	return nil
}

func (this *PXCClusterParams_HAProxy) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}

func (this *GetPXCClusterCredentialsRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}

func (this *PXCClusterConnectionCredentials) Validate() error {
	return nil
}

func (this *GetPXCClusterCredentialsResponse) Validate() error {
	if this.ConnectionCredentials != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionCredentials); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionCredentials", err)
		}
	}
	return nil
}

var _regex_CreatePXCClusterRequest_Name = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)

func (this *CreatePXCClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if !_regex_CreatePXCClusterRequest_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z]([-a-z0-9]*[a-z0-9])?$"`, this.Name))
	}
	if nil == this.Params {
		return github_com_mwitkow_go_proto_validators.FieldError("Params", fmt.Errorf("message must exist"))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}

func (this *CreatePXCClusterResponse) Validate() error {
	return nil
}

func (this *UpdatePXCClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}

func (this *UpdatePXCClusterRequest_UpdatePXCClusterParams) Validate() error {
	if this.Pxc != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pxc); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pxc", err)
		}
	}
	if this.Proxysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proxysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", err)
		}
	}
	if this.Haproxy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Haproxy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Haproxy", err)
		}
	}
	return nil
}

func (this *UpdatePXCClusterRequest_UpdatePXCClusterParams_PXC) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}

func (this *UpdatePXCClusterRequest_UpdatePXCClusterParams_ProxySQL) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}

func (this *UpdatePXCClusterRequest_UpdatePXCClusterParams_HAProxy) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}

func (this *UpdatePXCClusterResponse) Validate() error {
	return nil
}

func (this *GetPXCClusterResourcesRequest) Validate() error {
	if nil == this.Params {
		return github_com_mwitkow_go_proto_validators.FieldError("Params", fmt.Errorf("message must exist"))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}

func (this *GetPXCClusterResourcesResponse) Validate() error {
	if this.Expected != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Expected); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Expected", err)
		}
	}
	return nil
}
