// Code generated by MockGen. DO NOT EDIT.
// Source: router.go

// Package router is a generated GoMock package.
package router

import (
	url "net/url"
	reflect "reflect"

	game_launcher "github.com/cetteup/joinme.click-launcher/pkg/game_launcher"
	software_finder "github.com/cetteup/joinme.click-launcher/pkg/software_finder"
	gomock "github.com/golang/mock/gomock"
	registry "golang.org/x/sys/windows/registry"
)

// MockregistryRepository is a mock of registryRepository interface.
type MockregistryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockregistryRepositoryMockRecorder
}

// MockregistryRepositoryMockRecorder is the mock recorder for MockregistryRepository.
type MockregistryRepositoryMockRecorder struct {
	mock *MockregistryRepository
}

// NewMockregistryRepository creates a new mock instance.
func NewMockregistryRepository(ctrl *gomock.Controller) *MockregistryRepository {
	mock := &MockregistryRepository{ctrl: ctrl}
	mock.recorder = &MockregistryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockregistryRepository) EXPECT() *MockregistryRepositoryMockRecorder {
	return m.recorder
}

// CreateKey mocks base method.
func (m *MockregistryRepository) CreateKey(k registry.Key, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", k, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateKey indicates an expected call of CreateKey.
func (mr *MockregistryRepositoryMockRecorder) CreateKey(k, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockregistryRepository)(nil).CreateKey), k, path)
}

// GetStringValue mocks base method.
func (m *MockregistryRepository) GetStringValue(k registry.Key, path, valueName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringValue", k, path, valueName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStringValue indicates an expected call of GetStringValue.
func (mr *MockregistryRepositoryMockRecorder) GetStringValue(k, path, valueName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringValue", reflect.TypeOf((*MockregistryRepository)(nil).GetStringValue), k, path, valueName)
}

// SetStringValue mocks base method.
func (m *MockregistryRepository) SetStringValue(k registry.Key, path, valueName, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStringValue", k, path, valueName, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStringValue indicates an expected call of SetStringValue.
func (mr *MockregistryRepositoryMockRecorder) SetStringValue(k, path, valueName, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStringValue", reflect.TypeOf((*MockregistryRepository)(nil).SetStringValue), k, path, valueName, value)
}

// MockgameFinder is a mock of gameFinder interface.
type MockgameFinder struct {
	ctrl     *gomock.Controller
	recorder *MockgameFinderMockRecorder
}

// MockgameFinderMockRecorder is the mock recorder for MockgameFinder.
type MockgameFinderMockRecorder struct {
	mock *MockgameFinder
}

// NewMockgameFinder creates a new mock instance.
func NewMockgameFinder(ctrl *gomock.Controller) *MockgameFinder {
	mock := &MockgameFinder{ctrl: ctrl}
	mock.recorder = &MockgameFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockgameFinder) EXPECT() *MockgameFinderMockRecorder {
	return m.recorder
}

// GetInstallDir mocks base method.
func (m *MockgameFinder) GetInstallDir(config software_finder.Config) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallDir", config)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallDir indicates an expected call of GetInstallDir.
func (mr *MockgameFinderMockRecorder) GetInstallDir(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallDir", reflect.TypeOf((*MockgameFinder)(nil).GetInstallDir), config)
}

// GetInstallDirFromSomewhere mocks base method.
func (m *MockgameFinder) GetInstallDirFromSomewhere(configs []software_finder.Config) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallDirFromSomewhere", configs)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallDirFromSomewhere indicates an expected call of GetInstallDirFromSomewhere.
func (mr *MockgameFinderMockRecorder) GetInstallDirFromSomewhere(configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallDirFromSomewhere", reflect.TypeOf((*MockgameFinder)(nil).GetInstallDirFromSomewhere), configs)
}

// IsInstalled mocks base method.
func (m *MockgameFinder) IsInstalled(config software_finder.Config) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstalled", config)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInstalled indicates an expected call of IsInstalled.
func (mr *MockgameFinderMockRecorder) IsInstalled(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstalled", reflect.TypeOf((*MockgameFinder)(nil).IsInstalled), config)
}

// IsInstalledAnywhere mocks base method.
func (m *MockgameFinder) IsInstalledAnywhere(configs []software_finder.Config) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstalledAnywhere", configs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInstalledAnywhere indicates an expected call of IsInstalledAnywhere.
func (mr *MockgameFinderMockRecorder) IsInstalledAnywhere(configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstalledAnywhere", reflect.TypeOf((*MockgameFinder)(nil).IsInstalledAnywhere), configs)
}

// MockgameLauncher is a mock of gameLauncher interface.
type MockgameLauncher struct {
	ctrl     *gomock.Controller
	recorder *MockgameLauncherMockRecorder
}

// MockgameLauncherMockRecorder is the mock recorder for MockgameLauncher.
type MockgameLauncherMockRecorder struct {
	mock *MockgameLauncher
}

// NewMockgameLauncher creates a new mock instance.
func NewMockgameLauncher(ctrl *gomock.Controller) *MockgameLauncher {
	mock := &MockgameLauncher{ctrl: ctrl}
	mock.recorder = &MockgameLauncherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockgameLauncher) EXPECT() *MockgameLauncherMockRecorder {
	return m.recorder
}

// PrepareLaunch mocks base method.
func (m *MockgameLauncher) PrepareLaunch(config game_launcher.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareLaunch", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareLaunch indicates an expected call of PrepareLaunch.
func (mr *MockgameLauncherMockRecorder) PrepareLaunch(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareLaunch", reflect.TypeOf((*MockgameLauncher)(nil).PrepareLaunch), config)
}

// StartGame mocks base method.
func (m *MockgameLauncher) StartGame(u *url.URL, config game_launcher.Config, launchType game_launcher.LaunchType, cmdBuilder game_launcher.CommandBuilder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartGame", u, config, launchType, cmdBuilder)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartGame indicates an expected call of StartGame.
func (mr *MockgameLauncherMockRecorder) StartGame(u, config, launchType, cmdBuilder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartGame", reflect.TypeOf((*MockgameLauncher)(nil).StartGame), u, config, launchType, cmdBuilder)
}