# Backplane-CLI Detailed Code Coverage Report

## 📊 Overall Coverage
- **Total Coverage**: `45.72%`

## 📉 Files with Low Test Coverage (Sorted in Ascending Order)
| File | Total Lines | Covered Lines | Missed Lines | Coverage % |
|------|------------|--------------|-------------|------------|
| `cmd/ocm-backplane/main.go` | 2 | 0 | 2 | **0.0%** |
| `cmd/ocm-backplane/status/status.go` | 20 | 0 | 20 | **0.0%** |
| `cmd/ocm-backplane/version/version.go` | 3 | 0 | 3 | **0.0%** |
| `internal/upgrade/options.go` | 14 | 0 | 14 | **0.0%** |
| `internal/upgrade/upgrade.go` | 88 | 0 | 88 | **0.0%** |
| `internal/upgrade/writer.go` | 36 | 0 | 36 | **0.0%** |
| `pkg/backplaneapi/clientUtils.go` | 62 | 0 | 62 | **0.0%** |
| `pkg/cli/session/mocks/sessionMock.go` | 38 | 0 | 38 | **0.0%** |
| `pkg/info/mocks/infoMock.go` | 14 | 0 | 14 | **0.0%** |
| `pkg/jira/issueService.go` | 47 | 0 | 47 | **0.0%** |
| `pkg/ocm/ocmWrapper.go` | 289 | 0 | 289 | **0.0%** |
| `pkg/utils/clientUtils.go` | 73 | 0 | 73 | **0.0%** |
| `pkg/utils/shell.go` | 3 | 0 | 3 | **0.0%** |
| `pkg/utils/mocks/shellCheckerMock.go` | 14 | 0 | 14 | **0.0%** |
| `pkg/client/mocks/ClientWithResponsesMock.go` | 435 | 10 | 425 | **2.29%** |
| `pkg/healthcheck/connectivity_checks.go` | 83 | 7 | 73 | **8.43%** |
| `pkg/utils/mocks/ClusterMock.go` | 46 | 4 | 42 | **8.69%** |
| `cmd/ocm-backplane/config/set.go` | 67 | 9 | 58 | **13.43%** |
| `cmd/ocm-backplane/managedJob/logsManagedJob.go` | 57 | 8 | 48 | **14.03%** |
| `pkg/credentials/aws.go` | 24 | 4 | 20 | **16.66%** |
| `cmd/ocm-backplane/testJob/getTestJobLogs.go` | 50 | 10 | 40 | **20.0%** |
| `cmd/ocm-backplane/cloud/common.go` | 173 | 36 | 136 | **20.8%** |
| `cmd/ocm-backplane/cloud/console.go` | 100 | 22 | 78 | **22.0%** |
| `cmd/ocm-backplane/config/get.go` | 33 | 9 | 24 | **27.27%** |
| `pkg/client/mocks/ClientMock.go` | 435 | 127 | 297 | **29.19%** |
| `pkg/remediation/remediation.go` | 42 | 13 | 24 | **30.95%** |
| `cmd/ocm-backplane/accessrequest/expireAccessRequest.go` | 27 | 9 | 18 | **33.33%** |
| `cmd/ocm-backplane/accessrequest/getAccessRequest.go` | 26 | 9 | 17 | **34.61%** |
| `cmd/ocm-backplane/upgrade/upgrade.go` | 13 | 5 | 8 | **38.46%** |
| `pkg/monitoring/monitoring.go` | 190 | 75 | 86 | **39.47%** |
| `cmd/ocm-backplane/console/console.go` | 767 | 312 | 398 | **40.67%** |
| `cmd/ocm-backplane/accessrequest/createAccessRequest.go` | 70 | 30 | 40 | **42.85%** |
| `cmd/ocm-backplane/managedJob/deleteManagedJob.go` | 54 | 26 | 18 | **48.14%** |
| `pkg/utils/util.go` | 139 | 67 | 62 | **48.2%** |
| `cmd/ocm-backplane/healthcheck/health_check.go` | 6 | 3 | 3 | **50.0%** |
| `pkg/credentials/gcp.go` | 4 | 2 | 2 | **50.0%** |
| `pkg/cli/session/session.go` | 226 | 114 | 77 | **50.44%** |
| `pkg/pagerduty/mocks/clientMock.go` | 45 | 24 | 21 | **53.33%** |
| `pkg/awsutil/sts.go` | 118 | 63 | 47 | **53.38%** |
| `pkg/pagerduty/client.go` | 13 | 7 | 5 | **53.84%** |
| `pkg/utils/renderingutils.go` | 67 | 39 | 24 | **58.2%** |
| `cmd/ocm-backplane/testJob/getTestJob.go` | 48 | 28 | 10 | **58.33%** |
| `cmd/ocm-backplane/cloud/cloud.go` | 5 | 3 | 2 | **60.0%** |
| `cmd/ocm-backplane/cloud/credentials.go` | 72 | 45 | 25 | **62.5%** |
| `cmd/ocm-backplane/config/troubleshoot.go` | 79 | 50 | 23 | **63.29%** |
| `pkg/accessrequest/accessRequest.go` | 149 | 99 | 41 | **66.44%** |
| `cmd/ocm-backplane/elevate/elevate.go` | 12 | 8 | 4 | **66.66%** |
| `cmd/ocm-backplane/logout/logout.go` | 43 | 29 | 10 | **67.44%** |
| `cmd/ocm-backplane/managedJob/createManagedJob.go` | 170 | 115 | 37 | **67.64%** |
| `pkg/elevate/elevate.go` | 47 | 32 | 8 | **68.08%** |
| `cmd/ocm-backplane/managedJob/getManagedJob.go` | 63 | 43 | 10 | **68.25%** |
| `internal/github/github.go` | 98 | 67 | 25 | **68.36%** |
| `cmd/ocm-backplane/login/login.go` | 368 | 252 | 79 | **68.47%** |
| `cmd/ocm-backplane/script/listScripts.go` | 51 | 35 | 8 | **68.62%** |
| `pkg/login/kubeConfig.go` | 118 | 81 | 25 | **68.64%** |
| `cmd/ocm-backplane/remediation/remediation.go` | 134 | 93 | 29 | **69.4%** |
| `pkg/backplaneapi/mocks/clientUtilsMock.go` | 59 | 41 | 18 | **69.49%** |
| `pkg/jira/ohssService.go` | 31 | 22 | 5 | **70.96%** |
| `pkg/healthcheck/check_proxy.go` | 47 | 34 | 9 | **72.34%** |
| `pkg/healthcheck/check_vpn.go` | 37 | 27 | 7 | **72.97%** |
| `cmd/ocm-backplane/testJob/createTestJob.go` | 177 | 130 | 24 | **73.44%** |
| `cmd/ocm-backplane/session/session.go` | 36 | 27 | 8 | **75.0%** |
| `pkg/cli/config/config.go` | 165 | 126 | 30 | **76.36%** |
| `pkg/login/printClusterInfo.go` | 43 | 33 | 6 | **76.74%** |
| `cmd/ocm-backplane/script/describeScript.go` | 77 | 61 | 8 | **79.22%** |
| `pkg/pagerduty/pagerduty.go` | 67 | 55 | 6 | **82.08%** |
| `cmd/ocm-backplane/root.go` | 23 | 19 | 4 | **82.6%** |
| `pkg/ocm/mocks/ocmWrapperMock.go` | 163 | 136 | 27 | **83.43%** |
| `pkg/utils/cluster.go` | 51 | 43 | 4 | **84.31%** |
| `pkg/utils/jwt.go` | 34 | 30 | 2 | **88.23%** |
| `cmd/ocm-backplane/monitoring/monitoring.go` | 48 | 43 | 5 | **89.58%** |
| `cmd/ocm-backplane/accessrequest/accessRequest.go` | 12 | 12 | 0 | **100.0%** |
| `cmd/ocm-backplane/config/config.go` | 17 | 17 | 0 | **100.0%** |
| `cmd/ocm-backplane/managedJob/managedJob.go` | 16 | 16 | 0 | **100.0%** |
| `cmd/ocm-backplane/script/script.go` | 20 | 20 | 0 | **100.0%** |
| `cmd/ocm-backplane/testJob/testJob.go` | 21 | 21 | 0 | **100.0%** |
| `internal/github/options.go` | 4 | 4 | 0 | **100.0%** |
| `pkg/backplaneapi/deprecation.go` | 4 | 4 | 0 | **100.0%** |
| `pkg/cli/globalflags/globalflags.go` | 25 | 25 | 0 | **100.0%** |
| `pkg/cli/globalflags/logs.go` | 19 | 19 | 0 | **100.0%** |
| `pkg/healthcheck/mocks/httpClientMock.go` | 15 | 15 | 0 | **100.0%** |
| `pkg/healthcheck/mocks/networkMock.go` | 15 | 15 | 0 | **100.0%** |
| `pkg/info/buildInfo.go` | 2 | 2 | 0 | **100.0%** |
| `pkg/info/info.go` | 7 | 7 | 0 | **100.0%** |
| `pkg/info/mocks/buildInfoMock.go` | 15 | 15 | 0 | **100.0%** |
| `pkg/jira/mocks/jiraMock.go` | 55 | 55 | 0 | **100.0%** |

## 🔥 Critical Areas That Need Improvement
- `cmd/ocm-backplane/main.go` → **0.0% coverage** (Needs urgent attention!)
- `cmd/ocm-backplane/status/status.go` → **0.0% coverage** (Needs urgent attention!)
- `cmd/ocm-backplane/version/version.go` → **0.0% coverage** (Needs urgent attention!)
- `internal/upgrade/options.go` → **0.0% coverage** (Needs urgent attention!)
- `internal/upgrade/upgrade.go` → **0.0% coverage** (Needs urgent attention!)
- `internal/upgrade/writer.go` → **0.0% coverage** (Needs urgent attention!)
- `pkg/backplaneapi/clientUtils.go` → **0.0% coverage** (Needs urgent attention!)
- `pkg/cli/session/mocks/sessionMock.go` → **0.0% coverage** (Needs urgent attention!)
- `pkg/info/mocks/infoMock.go` → **0.0% coverage** (Needs urgent attention!)
- `pkg/jira/issueService.go` → **0.0% coverage** (Needs urgent attention!)
- `pkg/ocm/ocmWrapper.go` → **0.0% coverage** (Needs urgent attention!)
- `pkg/utils/clientUtils.go` → **0.0% coverage** (Needs urgent attention!)
- `pkg/utils/shell.go` → **0.0% coverage** (Needs urgent attention!)
- `pkg/utils/mocks/shellCheckerMock.go` → **0.0% coverage** (Needs urgent attention!)
- `pkg/client/mocks/ClientWithResponsesMock.go` → **2.29% coverage** (Needs urgent attention!)
- `pkg/healthcheck/connectivity_checks.go` → **8.43% coverage** (Needs urgent attention!)
- `pkg/utils/mocks/ClusterMock.go` → **8.69% coverage** (Needs urgent attention!)
- `cmd/ocm-backplane/config/set.go` → **13.43% coverage** (Needs urgent attention!)
- `cmd/ocm-backplane/managedJob/logsManagedJob.go` → **14.03% coverage** (Needs urgent attention!)
- `pkg/credentials/aws.go` → **16.66% coverage** (Needs urgent attention!)
- `cmd/ocm-backplane/testJob/getTestJobLogs.go` → **20.0% coverage** (Needs urgent attention!)
- `cmd/ocm-backplane/cloud/common.go` → **20.8% coverage** (Needs urgent attention!)
- `cmd/ocm-backplane/cloud/console.go` → **22.0% coverage** (Needs urgent attention!)
- `cmd/ocm-backplane/config/get.go` → **27.27% coverage** (Needs urgent attention!)
- `pkg/client/mocks/ClientMock.go` → **29.19% coverage** (Needs urgent attention!)

## 🚀 Test Improvement Suggestions
- Focus on adding unit tests for files with **less than 50% coverage**.
- Prioritize core functionality files such as authentication, session management, and API handling.
- Leverage the `osde2e` test harness for integration testing.
