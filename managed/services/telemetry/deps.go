// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package telemetry

import (
	"context"

	pmmv1 "github.com/percona-platform/saas/gen/telemetry/events/pmm"
	reporter "github.com/percona-platform/saas/gen/telemetry/reporter"
	"github.com/sirupsen/logrus"

	"github.com/percona/pmm/api/serverpb"
)

//go:generate ../../../bin/mockery -name=distributionUtilService -case=snake -inpkg -testonly
//go:generate ../../../bin/mockery -name=sender -case=snake -inpkg -testonly
//go:generate ../../../bin/mockery -name=DataSourceLocator -case=snake -inpkg -testonly
//go:generate ../../../bin/mockery -name=DataSource -case=snake -inpkg -testonly

// distributionUtilService service to get info about OS on which pmm server is running
type distributionUtilService interface {
	getDistributionMethodAndOS(l *logrus.Entry) (serverpb.DistributionMethod, pmmv1.DistributionMethod, string)
	getLinuxDistribution(procVersion string) string
}

// sender is interface which defines method for client which sends report with metrics
type sender interface {
	SendTelemetry(ctx context.Context, report *reporter.ReportRequest) error
}

// DataSourceLocator locates data source by name.
type DataSourceLocator interface {
	LocateTelemetryDataSource(name string) (DataSource, error)
}

// DataSource telemetry data source.
type DataSource interface {
	FetchMetrics(ctx context.Context, config Config) ([][]*pmmv1.ServerMetric_Metric, error)
	Enabled() bool
}
