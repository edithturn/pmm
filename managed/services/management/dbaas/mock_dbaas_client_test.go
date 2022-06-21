// Code generated by mockery v1.0.0. DO NOT EDIT.

package dbaas

import (
	context "context"

	controllerv1beta1 "github.com/percona-platform/dbaas-api/gen/controller"
	mock "github.com/stretchr/testify/mock"
	grpc "google.golang.org/grpc"
)

// mockDbaasClient is an autogenerated mock type for the dbaasClient type
type mockDbaasClient struct {
	mock.Mock
}

// CheckKubernetesClusterConnection provides a mock function with given fields: ctx, kubeConfig
func (_m *mockDbaasClient) CheckKubernetesClusterConnection(ctx context.Context, kubeConfig string) (*controllerv1beta1.CheckKubernetesClusterConnectionResponse, error) {
	ret := _m.Called(ctx, kubeConfig)

	var r0 *controllerv1beta1.CheckKubernetesClusterConnectionResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) *controllerv1beta1.CheckKubernetesClusterConnectionResponse); ok {
		r0 = rf(ctx, kubeConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.CheckKubernetesClusterConnectionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, kubeConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePSMDBCluster provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) CreatePSMDBCluster(ctx context.Context, in *controllerv1beta1.CreatePSMDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.CreatePSMDBClusterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.CreatePSMDBClusterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.CreatePSMDBClusterRequest, ...grpc.CallOption) *controllerv1beta1.CreatePSMDBClusterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.CreatePSMDBClusterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.CreatePSMDBClusterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePXCCluster provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) CreatePXCCluster(ctx context.Context, in *controllerv1beta1.CreatePXCClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.CreatePXCClusterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.CreatePXCClusterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.CreatePXCClusterRequest, ...grpc.CallOption) *controllerv1beta1.CreatePXCClusterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.CreatePXCClusterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.CreatePXCClusterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePSMDBCluster provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) DeletePSMDBCluster(ctx context.Context, in *controllerv1beta1.DeletePSMDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.DeletePSMDBClusterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.DeletePSMDBClusterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.DeletePSMDBClusterRequest, ...grpc.CallOption) *controllerv1beta1.DeletePSMDBClusterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.DeletePSMDBClusterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.DeletePSMDBClusterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePXCCluster provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) DeletePXCCluster(ctx context.Context, in *controllerv1beta1.DeletePXCClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.DeletePXCClusterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.DeletePXCClusterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.DeletePXCClusterRequest, ...grpc.CallOption) *controllerv1beta1.DeletePXCClusterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.DeletePXCClusterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.DeletePXCClusterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) GetLogs(ctx context.Context, in *controllerv1beta1.GetLogsRequest, opts ...grpc.CallOption) (*controllerv1beta1.GetLogsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.GetLogsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.GetLogsRequest, ...grpc.CallOption) *controllerv1beta1.GetLogsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.GetLogsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.GetLogsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPSMDBClusterCredentials provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) GetPSMDBClusterCredentials(ctx context.Context, in *controllerv1beta1.GetPSMDBClusterCredentialsRequest, opts ...grpc.CallOption) (*controllerv1beta1.GetPSMDBClusterCredentialsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.GetPSMDBClusterCredentialsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.GetPSMDBClusterCredentialsRequest, ...grpc.CallOption) *controllerv1beta1.GetPSMDBClusterCredentialsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.GetPSMDBClusterCredentialsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.GetPSMDBClusterCredentialsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPXCClusterCredentials provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) GetPXCClusterCredentials(ctx context.Context, in *controllerv1beta1.GetPXCClusterCredentialsRequest, opts ...grpc.CallOption) (*controllerv1beta1.GetPXCClusterCredentialsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.GetPXCClusterCredentialsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.GetPXCClusterCredentialsRequest, ...grpc.CallOption) *controllerv1beta1.GetPXCClusterCredentialsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.GetPXCClusterCredentialsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.GetPXCClusterCredentialsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResources provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) GetResources(ctx context.Context, in *controllerv1beta1.GetResourcesRequest, opts ...grpc.CallOption) (*controllerv1beta1.GetResourcesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.GetResourcesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.GetResourcesRequest, ...grpc.CallOption) *controllerv1beta1.GetResourcesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.GetResourcesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.GetResourcesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallPSMDBOperator provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) InstallPSMDBOperator(ctx context.Context, in *controllerv1beta1.InstallPSMDBOperatorRequest, opts ...grpc.CallOption) (*controllerv1beta1.InstallPSMDBOperatorResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.InstallPSMDBOperatorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.InstallPSMDBOperatorRequest, ...grpc.CallOption) *controllerv1beta1.InstallPSMDBOperatorResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.InstallPSMDBOperatorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.InstallPSMDBOperatorRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallPXCOperator provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) InstallPXCOperator(ctx context.Context, in *controllerv1beta1.InstallPXCOperatorRequest, opts ...grpc.CallOption) (*controllerv1beta1.InstallPXCOperatorResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.InstallPXCOperatorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.InstallPXCOperatorRequest, ...grpc.CallOption) *controllerv1beta1.InstallPXCOperatorResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.InstallPXCOperatorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.InstallPXCOperatorRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPSMDBClusters provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) ListPSMDBClusters(ctx context.Context, in *controllerv1beta1.ListPSMDBClustersRequest, opts ...grpc.CallOption) (*controllerv1beta1.ListPSMDBClustersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.ListPSMDBClustersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.ListPSMDBClustersRequest, ...grpc.CallOption) *controllerv1beta1.ListPSMDBClustersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.ListPSMDBClustersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.ListPSMDBClustersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPXCClusters provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) ListPXCClusters(ctx context.Context, in *controllerv1beta1.ListPXCClustersRequest, opts ...grpc.CallOption) (*controllerv1beta1.ListPXCClustersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.ListPXCClustersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.ListPXCClustersRequest, ...grpc.CallOption) *controllerv1beta1.ListPXCClustersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.ListPXCClustersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.ListPXCClustersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestartPSMDBCluster provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) RestartPSMDBCluster(ctx context.Context, in *controllerv1beta1.RestartPSMDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.RestartPSMDBClusterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.RestartPSMDBClusterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.RestartPSMDBClusterRequest, ...grpc.CallOption) *controllerv1beta1.RestartPSMDBClusterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.RestartPSMDBClusterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.RestartPSMDBClusterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestartPXCCluster provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) RestartPXCCluster(ctx context.Context, in *controllerv1beta1.RestartPXCClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.RestartPXCClusterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.RestartPXCClusterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.RestartPXCClusterRequest, ...grpc.CallOption) *controllerv1beta1.RestartPXCClusterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.RestartPXCClusterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.RestartPXCClusterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartMonitoring provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) StartMonitoring(ctx context.Context, in *controllerv1beta1.StartMonitoringRequest, opts ...grpc.CallOption) (*controllerv1beta1.StartMonitoringResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.StartMonitoringResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.StartMonitoringRequest, ...grpc.CallOption) *controllerv1beta1.StartMonitoringResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.StartMonitoringResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.StartMonitoringRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePSMDBCluster provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) UpdatePSMDBCluster(ctx context.Context, in *controllerv1beta1.UpdatePSMDBClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.UpdatePSMDBClusterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.UpdatePSMDBClusterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.UpdatePSMDBClusterRequest, ...grpc.CallOption) *controllerv1beta1.UpdatePSMDBClusterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.UpdatePSMDBClusterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.UpdatePSMDBClusterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePXCCluster provides a mock function with given fields: ctx, in, opts
func (_m *mockDbaasClient) UpdatePXCCluster(ctx context.Context, in *controllerv1beta1.UpdatePXCClusterRequest, opts ...grpc.CallOption) (*controllerv1beta1.UpdatePXCClusterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *controllerv1beta1.UpdatePXCClusterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *controllerv1beta1.UpdatePXCClusterRequest, ...grpc.CallOption) *controllerv1beta1.UpdatePXCClusterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*controllerv1beta1.UpdatePXCClusterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *controllerv1beta1.UpdatePXCClusterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
