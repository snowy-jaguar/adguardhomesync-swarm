// Copyright 2025 snowy-jaguar
// Contact: @snowyjaguar (Discord)
// Contact: contact@snowyjaguar.xyz (Email)
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.



// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bakito/adguardhome-sync/pkg/client (interfaces: Client)
//
// Generated by this command:
//
//	mockgen -package client -destination pkg/mocks/client/mock.go github.com/bakito/adguardhome-sync/pkg/client Client
//

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"

	model "github.com/snowy-jaguar/adguardhomesync-swarm/pkg/client/model"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AccessList mocks base method.
func (m *MockClient) AccessList() (*model.AccessList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessList")
	ret0, _ := ret[0].(*model.AccessList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessList indicates an expected call of AccessList.
func (mr *MockClientMockRecorder) AccessList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessList", reflect.TypeOf((*MockClient)(nil).AccessList))
}

// AddClient mocks base method.
func (m *MockClient) AddClient(client *model.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddClient", client)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddClient indicates an expected call of AddClient.
func (mr *MockClientMockRecorder) AddClient(client any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClient", reflect.TypeOf((*MockClient)(nil).AddClient), client)
}

// AddDHCPStaticLease mocks base method.
func (m *MockClient) AddDHCPStaticLease(lease model.DhcpStaticLease) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDHCPStaticLease", lease)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDHCPStaticLease indicates an expected call of AddDHCPStaticLease.
func (mr *MockClientMockRecorder) AddDHCPStaticLease(lease any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDHCPStaticLease", reflect.TypeOf((*MockClient)(nil).AddDHCPStaticLease), lease)
}

// AddFilter mocks base method.
func (m *MockClient) AddFilter(whitelist bool, f model.Filter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilter", whitelist, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilter indicates an expected call of AddFilter.
func (mr *MockClientMockRecorder) AddFilter(whitelist, f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilter", reflect.TypeOf((*MockClient)(nil).AddFilter), whitelist, f)
}

// AddRewriteEntries mocks base method.
func (m *MockClient) AddRewriteEntries(e ...model.RewriteEntry) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range e {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddRewriteEntries", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRewriteEntries indicates an expected call of AddRewriteEntries.
func (mr *MockClientMockRecorder) AddRewriteEntries(e ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRewriteEntries", reflect.TypeOf((*MockClient)(nil).AddRewriteEntries), e...)
}

// BlockedServicesSchedule mocks base method.
func (m *MockClient) BlockedServicesSchedule() (*model.BlockedServicesSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockedServicesSchedule")
	ret0, _ := ret[0].(*model.BlockedServicesSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockedServicesSchedule indicates an expected call of BlockedServicesSchedule.
func (mr *MockClientMockRecorder) BlockedServicesSchedule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockedServicesSchedule", reflect.TypeOf((*MockClient)(nil).BlockedServicesSchedule))
}

// Clients mocks base method.
func (m *MockClient) Clients() (*model.Clients, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clients")
	ret0, _ := ret[0].(*model.Clients)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Clients indicates an expected call of Clients.
func (mr *MockClientMockRecorder) Clients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clients", reflect.TypeOf((*MockClient)(nil).Clients))
}

// DNSConfig mocks base method.
func (m *MockClient) DNSConfig() (*model.DNSConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DNSConfig")
	ret0, _ := ret[0].(*model.DNSConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DNSConfig indicates an expected call of DNSConfig.
func (mr *MockClientMockRecorder) DNSConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DNSConfig", reflect.TypeOf((*MockClient)(nil).DNSConfig))
}

// DeleteClient mocks base method.
func (m *MockClient) DeleteClient(client *model.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClient", client)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClient indicates an expected call of DeleteClient.
func (mr *MockClientMockRecorder) DeleteClient(client any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClient", reflect.TypeOf((*MockClient)(nil).DeleteClient), client)
}

// DeleteDHCPStaticLease mocks base method.
func (m *MockClient) DeleteDHCPStaticLease(lease model.DhcpStaticLease) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDHCPStaticLease", lease)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDHCPStaticLease indicates an expected call of DeleteDHCPStaticLease.
func (mr *MockClientMockRecorder) DeleteDHCPStaticLease(lease any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDHCPStaticLease", reflect.TypeOf((*MockClient)(nil).DeleteDHCPStaticLease), lease)
}

// DeleteFilter mocks base method.
func (m *MockClient) DeleteFilter(whitelist bool, f model.Filter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFilter", whitelist, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFilter indicates an expected call of DeleteFilter.
func (mr *MockClientMockRecorder) DeleteFilter(whitelist, f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFilter", reflect.TypeOf((*MockClient)(nil).DeleteFilter), whitelist, f)
}

// DeleteRewriteEntries mocks base method.
func (m *MockClient) DeleteRewriteEntries(e ...model.RewriteEntry) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range e {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRewriteEntries", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRewriteEntries indicates an expected call of DeleteRewriteEntries.
func (mr *MockClientMockRecorder) DeleteRewriteEntries(e ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRewriteEntries", reflect.TypeOf((*MockClient)(nil).DeleteRewriteEntries), e...)
}

// DhcpConfig mocks base method.
func (m *MockClient) DhcpConfig() (*model.DhcpStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DhcpConfig")
	ret0, _ := ret[0].(*model.DhcpStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DhcpConfig indicates an expected call of DhcpConfig.
func (mr *MockClientMockRecorder) DhcpConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DhcpConfig", reflect.TypeOf((*MockClient)(nil).DhcpConfig))
}

// Filtering mocks base method.
func (m *MockClient) Filtering() (*model.FilterStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filtering")
	ret0, _ := ret[0].(*model.FilterStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filtering indicates an expected call of Filtering.
func (mr *MockClientMockRecorder) Filtering() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filtering", reflect.TypeOf((*MockClient)(nil).Filtering))
}

// Host mocks base method.
func (m *MockClient) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host.
func (mr *MockClientMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockClient)(nil).Host))
}

// Parental mocks base method.
func (m *MockClient) Parental() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parental")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parental indicates an expected call of Parental.
func (mr *MockClientMockRecorder) Parental() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parental", reflect.TypeOf((*MockClient)(nil).Parental))
}

// ProfileInfo mocks base method.
func (m *MockClient) ProfileInfo() (*model.ProfileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProfileInfo")
	ret0, _ := ret[0].(*model.ProfileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProfileInfo indicates an expected call of ProfileInfo.
func (mr *MockClientMockRecorder) ProfileInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProfileInfo", reflect.TypeOf((*MockClient)(nil).ProfileInfo))
}

// QueryLog mocks base method.
func (m *MockClient) QueryLog(limit int) (*model.QueryLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLog", limit)
	ret0, _ := ret[0].(*model.QueryLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLog indicates an expected call of QueryLog.
func (mr *MockClientMockRecorder) QueryLog(limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLog", reflect.TypeOf((*MockClient)(nil).QueryLog), limit)
}

// QueryLogConfig mocks base method.
func (m *MockClient) QueryLogConfig() (*model.QueryLogConfigWithIgnored, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLogConfig")
	ret0, _ := ret[0].(*model.QueryLogConfigWithIgnored)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLogConfig indicates an expected call of QueryLogConfig.
func (mr *MockClientMockRecorder) QueryLogConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLogConfig", reflect.TypeOf((*MockClient)(nil).QueryLogConfig))
}

// RefreshFilters mocks base method.
func (m *MockClient) RefreshFilters(whitelist bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshFilters", whitelist)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshFilters indicates an expected call of RefreshFilters.
func (mr *MockClientMockRecorder) RefreshFilters(whitelist any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshFilters", reflect.TypeOf((*MockClient)(nil).RefreshFilters), whitelist)
}

// RewriteList mocks base method.
func (m *MockClient) RewriteList() (*model.RewriteEntries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RewriteList")
	ret0, _ := ret[0].(*model.RewriteEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RewriteList indicates an expected call of RewriteList.
func (mr *MockClientMockRecorder) RewriteList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RewriteList", reflect.TypeOf((*MockClient)(nil).RewriteList))
}

// SafeBrowsing mocks base method.
func (m *MockClient) SafeBrowsing() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SafeBrowsing")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SafeBrowsing indicates an expected call of SafeBrowsing.
func (mr *MockClientMockRecorder) SafeBrowsing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeBrowsing", reflect.TypeOf((*MockClient)(nil).SafeBrowsing))
}

// SafeSearchConfig mocks base method.
func (m *MockClient) SafeSearchConfig() (*model.SafeSearchConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SafeSearchConfig")
	ret0, _ := ret[0].(*model.SafeSearchConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SafeSearchConfig indicates an expected call of SafeSearchConfig.
func (mr *MockClientMockRecorder) SafeSearchConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeSearchConfig", reflect.TypeOf((*MockClient)(nil).SafeSearchConfig))
}

// SetAccessList mocks base method.
func (m *MockClient) SetAccessList(accessList *model.AccessList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessList", accessList)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAccessList indicates an expected call of SetAccessList.
func (mr *MockClientMockRecorder) SetAccessList(accessList any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessList", reflect.TypeOf((*MockClient)(nil).SetAccessList), accessList)
}

// SetBlockedServicesSchedule mocks base method.
func (m *MockClient) SetBlockedServicesSchedule(schedule *model.BlockedServicesSchedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBlockedServicesSchedule", schedule)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBlockedServicesSchedule indicates an expected call of SetBlockedServicesSchedule.
func (mr *MockClientMockRecorder) SetBlockedServicesSchedule(schedule any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBlockedServicesSchedule", reflect.TypeOf((*MockClient)(nil).SetBlockedServicesSchedule), schedule)
}

// SetCustomRules mocks base method.
func (m *MockClient) SetCustomRules(rules *[]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCustomRules", rules)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCustomRules indicates an expected call of SetCustomRules.
func (mr *MockClientMockRecorder) SetCustomRules(rules any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCustomRules", reflect.TypeOf((*MockClient)(nil).SetCustomRules), rules)
}

// SetDNSConfig mocks base method.
func (m *MockClient) SetDNSConfig(config *model.DNSConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDNSConfig", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDNSConfig indicates an expected call of SetDNSConfig.
func (mr *MockClientMockRecorder) SetDNSConfig(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDNSConfig", reflect.TypeOf((*MockClient)(nil).SetDNSConfig), config)
}

// SetDhcpConfig mocks base method.
func (m *MockClient) SetDhcpConfig(status *model.DhcpStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDhcpConfig", status)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDhcpConfig indicates an expected call of SetDhcpConfig.
func (mr *MockClientMockRecorder) SetDhcpConfig(status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDhcpConfig", reflect.TypeOf((*MockClient)(nil).SetDhcpConfig), status)
}

// SetProfileInfo mocks base method.
func (m *MockClient) SetProfileInfo(settings *model.ProfileInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProfileInfo", settings)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProfileInfo indicates an expected call of SetProfileInfo.
func (mr *MockClientMockRecorder) SetProfileInfo(settings any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProfileInfo", reflect.TypeOf((*MockClient)(nil).SetProfileInfo), settings)
}

// SetQueryLogConfig mocks base method.
func (m *MockClient) SetQueryLogConfig(ql *model.QueryLogConfigWithIgnored) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetQueryLogConfig", ql)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetQueryLogConfig indicates an expected call of SetQueryLogConfig.
func (mr *MockClientMockRecorder) SetQueryLogConfig(ql any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetQueryLogConfig", reflect.TypeOf((*MockClient)(nil).SetQueryLogConfig), ql)
}

// SetSafeSearchConfig mocks base method.
func (m *MockClient) SetSafeSearchConfig(settings *model.SafeSearchConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSafeSearchConfig", settings)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSafeSearchConfig indicates an expected call of SetSafeSearchConfig.
func (mr *MockClientMockRecorder) SetSafeSearchConfig(settings any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSafeSearchConfig", reflect.TypeOf((*MockClient)(nil).SetSafeSearchConfig), settings)
}

// SetStatsConfig mocks base method.
func (m *MockClient) SetStatsConfig(sc *model.PutStatsConfigUpdateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatsConfig", sc)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatsConfig indicates an expected call of SetStatsConfig.
func (mr *MockClientMockRecorder) SetStatsConfig(sc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatsConfig", reflect.TypeOf((*MockClient)(nil).SetStatsConfig), sc)
}

// Setup mocks base method.
func (m *MockClient) Setup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup.
func (mr *MockClientMockRecorder) Setup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockClient)(nil).Setup))
}

// Stats mocks base method.
func (m *MockClient) Stats() (*model.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(*model.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats.
func (mr *MockClientMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockClient)(nil).Stats))
}

// StatsConfig mocks base method.
func (m *MockClient) StatsConfig() (*model.GetStatsConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatsConfig")
	ret0, _ := ret[0].(*model.GetStatsConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatsConfig indicates an expected call of StatsConfig.
func (mr *MockClientMockRecorder) StatsConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatsConfig", reflect.TypeOf((*MockClient)(nil).StatsConfig))
}

// Status mocks base method.
func (m *MockClient) Status() (*model.ServerStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(*model.ServerStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockClientMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockClient)(nil).Status))
}

// ToggleFiltering mocks base method.
func (m *MockClient) ToggleFiltering(enabled bool, interval int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleFiltering", enabled, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleFiltering indicates an expected call of ToggleFiltering.
func (mr *MockClientMockRecorder) ToggleFiltering(enabled, interval any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleFiltering", reflect.TypeOf((*MockClient)(nil).ToggleFiltering), enabled, interval)
}

// ToggleParental mocks base method.
func (m *MockClient) ToggleParental(enable bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleParental", enable)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleParental indicates an expected call of ToggleParental.
func (mr *MockClientMockRecorder) ToggleParental(enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleParental", reflect.TypeOf((*MockClient)(nil).ToggleParental), enable)
}

// ToggleProtection mocks base method.
func (m *MockClient) ToggleProtection(enable bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleProtection", enable)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleProtection indicates an expected call of ToggleProtection.
func (mr *MockClientMockRecorder) ToggleProtection(enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleProtection", reflect.TypeOf((*MockClient)(nil).ToggleProtection), enable)
}

// ToggleSafeBrowsing mocks base method.
func (m *MockClient) ToggleSafeBrowsing(enable bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleSafeBrowsing", enable)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleSafeBrowsing indicates an expected call of ToggleSafeBrowsing.
func (mr *MockClientMockRecorder) ToggleSafeBrowsing(enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleSafeBrowsing", reflect.TypeOf((*MockClient)(nil).ToggleSafeBrowsing), enable)
}

// UpdateClient mocks base method.
func (m *MockClient) UpdateClient(client *model.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClient", client)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClient indicates an expected call of UpdateClient.
func (mr *MockClientMockRecorder) UpdateClient(client any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClient", reflect.TypeOf((*MockClient)(nil).UpdateClient), client)
}

// UpdateFilter mocks base method.
func (m *MockClient) UpdateFilter(whitelist bool, f model.Filter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFilter", whitelist, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFilter indicates an expected call of UpdateFilter.
func (mr *MockClientMockRecorder) UpdateFilter(whitelist, f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilter", reflect.TypeOf((*MockClient)(nil).UpdateFilter), whitelist, f)
}
