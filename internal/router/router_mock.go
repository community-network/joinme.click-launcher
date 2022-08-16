// Code generated by MockGen. DO NOT EDIT.
// Source: router.go

package router

import (
	url "net/url"
	reflect "reflect"

	game_launcher "github.com/cetteup/joinme.click-launcher/pkg/game_launcher"
	software_finder "github.com/cetteup/joinme.click-launcher/pkg/software_finder"
	gomock "github.com/golang/mock/gomock"
	registry "golang.org/x/sys/windows/registry"
)

// MockRegistryRepository is a mock of RegistryRepository interface.
type MockRegistryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryRepositoryMockRecorder
}

// MockRegistryRepositoryMockRecorder is the mock recorder for MockRegistryRepository.
type MockRegistryRepositoryMockRecorder struct {
	mock *MockRegistryRepository
}

// NewMockRegistryRepository creates a new mock instance.
func NewMockRegistryRepository(ctrl *gomock.Controller) *MockRegistryRepository {
	mock := &MockRegistryRepository{ctrl: ctrl}
	mock.recorder = &MockRegistryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistryRepository) EXPECT() *MockRegistryRepositoryMockRecorder {
	return m.recorder
}

// CreateKey mocks base method.
func (m *MockRegistryRepository) CreateKey(k registry.Key, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", k, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateKey indicates an expected call of CreateKey.
func (mr *MockRegistryRepositoryMockRecorder) CreateKey(k, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockRegistryRepository)(nil).CreateKey), k, path)
}

// GetStringValue mocks base method.
func (m *MockRegistryRepository) GetStringValue(k registry.Key, path, valueName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringValue", k, path, valueName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStringValue indicates an expected call of GetStringValue.
func (mr *MockRegistryRepositoryMockRecorder) GetStringValue(k, path, valueName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringValue", reflect.TypeOf((*MockRegistryRepository)(nil).GetStringValue), k, path, valueName)
}

// SetStringValue mocks base method.
func (m *MockRegistryRepository) SetStringValue(k registry.Key, path, valueName, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStringValue", k, path, valueName, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStringValue indicates an expected call of SetStringValue.
func (mr *MockRegistryRepositoryMockRecorder) SetStringValue(k, path, valueName, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStringValue", reflect.TypeOf((*MockRegistryRepository)(nil).SetStringValue), k, path, valueName, value)
}

// MockGameFinder is a mock of GameFinder interface.
type MockGameFinder struct {
	ctrl     *gomock.Controller
	recorder *MockGameFinderMockRecorder
}

// MockGameFinderMockRecorder is the mock recorder for MockGameFinder.
type MockGameFinderMockRecorder struct {
	mock *MockGameFinder
}

// NewMockGameFinder creates a new mock instance.
func NewMockGameFinder(ctrl *gomock.Controller) *MockGameFinder {
	mock := &MockGameFinder{ctrl: ctrl}
	mock.recorder = &MockGameFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameFinder) EXPECT() *MockGameFinderMockRecorder {
	return m.recorder
}

// GetInstallDir mocks base method.
func (m *MockGameFinder) GetInstallDir(config software_finder.Config) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallDir", config)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallDir indicates an expected call of GetInstallDir.
func (mr *MockGameFinderMockRecorder) GetInstallDir(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallDir", reflect.TypeOf((*MockGameFinder)(nil).GetInstallDir), config)
}

// GetInstallDirFromSomewhere mocks base method.
func (m *MockGameFinder) GetInstallDirFromSomewhere(configs []software_finder.Config) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallDirFromSomewhere", configs)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallDirFromSomewhere indicates an expected call of GetInstallDirFromSomewhere.
func (mr *MockGameFinderMockRecorder) GetInstallDirFromSomewhere(configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallDirFromSomewhere", reflect.TypeOf((*MockGameFinder)(nil).GetInstallDirFromSomewhere), configs)
}

// IsInstalled mocks base method.
func (m *MockGameFinder) IsInstalled(config software_finder.Config) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstalled", config)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInstalled indicates an expected call of IsInstalled.
func (mr *MockGameFinderMockRecorder) IsInstalled(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstalled", reflect.TypeOf((*MockGameFinder)(nil).IsInstalled), config)
}

// IsInstalledAnywhere mocks base method.
func (m *MockGameFinder) IsInstalledAnywhere(configs []software_finder.Config) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstalledAnywhere", configs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInstalledAnywhere indicates an expected call of IsInstalledAnywhere.
func (mr *MockGameFinderMockRecorder) IsInstalledAnywhere(configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstalledAnywhere", reflect.TypeOf((*MockGameFinder)(nil).IsInstalledAnywhere), configs)
}

// MockGameLauncher is a mock of GameLauncher interface.
type MockGameLauncher struct {
	ctrl     *gomock.Controller
	recorder *MockGameLauncherMockRecorder
}

// MockGameLauncherMockRecorder is the mock recorder for MockGameLauncher.
type MockGameLauncherMockRecorder struct {
	mock *MockGameLauncher
}

// NewMockGameLauncher creates a new mock instance.
func NewMockGameLauncher(ctrl *gomock.Controller) *MockGameLauncher {
	mock := &MockGameLauncher{ctrl: ctrl}
	mock.recorder = &MockGameLauncherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameLauncher) EXPECT() *MockGameLauncherMockRecorder {
	return m.recorder
}

// PrepareLaunch mocks base method.
func (m *MockGameLauncher) PrepareLaunch(config game_launcher.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareLaunch", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareLaunch indicates an expected call of PrepareLaunch.
func (mr *MockGameLauncherMockRecorder) PrepareLaunch(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareLaunch", reflect.TypeOf((*MockGameLauncher)(nil).PrepareLaunch), config)
}

// StartGame mocks base method.
func (m *MockGameLauncher) StartGame(u *url.URL, config game_launcher.Config, launchType game_launcher.LaunchType, cmdBuilder game_launcher.CommandBuilder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartGame", u, config, launchType, cmdBuilder)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartGame indicates an expected call of StartGame.
func (mr *MockGameLauncherMockRecorder) StartGame(u, config, launchType, cmdBuilder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartGame", reflect.TypeOf((*MockGameLauncher)(nil).StartGame), u, config, launchType, cmdBuilder)
}
