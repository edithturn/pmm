// Code generated by mockery v1.0.0. DO NOT EDIT.

package telemetry

import (
	logrus "github.com/sirupsen/logrus"
	mock "github.com/stretchr/testify/mock"

	pmmv1 "github.com/percona-platform/saas/gen/telemetry/events/pmm"

	serverpb "github.com/percona/pmm/api/serverpb"
)

// mockDistributionUtilService is an autogenerated mock type for the distributionUtilService type
type mockDistributionUtilService struct {
	mock.Mock
}

// getDistributionMethodAndOS provides a mock function with given fields: l
func (_m *mockDistributionUtilService) getDistributionMethodAndOS(l *logrus.Entry) (serverpb.DistributionMethod, pmmv1.DistributionMethod, string) {
	ret := _m.Called(l)

	var r0 serverpb.DistributionMethod
	if rf, ok := ret.Get(0).(func(*logrus.Entry) serverpb.DistributionMethod); ok {
		r0 = rf(l)
	} else {
		r0 = ret.Get(0).(serverpb.DistributionMethod)
	}

	var r1 pmmv1.DistributionMethod
	if rf, ok := ret.Get(1).(func(*logrus.Entry) pmmv1.DistributionMethod); ok {
		r1 = rf(l)
	} else {
		r1 = ret.Get(1).(pmmv1.DistributionMethod)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(*logrus.Entry) string); ok {
		r2 = rf(l)
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// getLinuxDistribution provides a mock function with given fields: procVersion
func (_m *mockDistributionUtilService) getLinuxDistribution(procVersion string) string {
	ret := _m.Called(procVersion)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(procVersion)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
