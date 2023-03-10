// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecases/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	models "github.com/sanyarise/playlist/internal/models"
)

// MockISongUsecase is a mock of ISongUsecase interface.
type MockISongUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockISongUsecaseMockRecorder
}

// MockISongUsecaseMockRecorder is the mock recorder for MockISongUsecase.
type MockISongUsecaseMockRecorder struct {
	mock *MockISongUsecase
}

// NewMockISongUsecase creates a new mock instance.
func NewMockISongUsecase(ctrl *gomock.Controller) *MockISongUsecase {
	mock := &MockISongUsecase{ctrl: ctrl}
	mock.recorder = &MockISongUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISongUsecase) EXPECT() *MockISongUsecaseMockRecorder {
	return m.recorder
}

// CreateSong mocks base method.
func (m *MockISongUsecase) CreateSong(ctx context.Context, song *models.Song) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSong", ctx, song)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSong indicates an expected call of CreateSong.
func (mr *MockISongUsecaseMockRecorder) CreateSong(ctx, song interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSong", reflect.TypeOf((*MockISongUsecase)(nil).CreateSong), ctx, song)
}

// DeleteSong mocks base method.
func (m *MockISongUsecase) DeleteSong(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSong", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSong indicates an expected call of DeleteSong.
func (mr *MockISongUsecaseMockRecorder) DeleteSong(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSong", reflect.TypeOf((*MockISongUsecase)(nil).DeleteSong), ctx, id)
}

// GetSong mocks base method.
func (m *MockISongUsecase) GetSong(ctx context.Context, id uuid.UUID) (*models.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSong", ctx, id)
	ret0, _ := ret[0].(*models.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSong indicates an expected call of GetSong.
func (mr *MockISongUsecaseMockRecorder) GetSong(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSong", reflect.TypeOf((*MockISongUsecase)(nil).GetSong), ctx, id)
}

// UpdateSong mocks base method.
func (m *MockISongUsecase) UpdateSong(ctx context.Context, song *models.Song) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSong", ctx, song)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSong indicates an expected call of UpdateSong.
func (mr *MockISongUsecaseMockRecorder) UpdateSong(ctx, song interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSong", reflect.TypeOf((*MockISongUsecase)(nil).UpdateSong), ctx, song)
}

// MockIPlaylistUsecase is a mock of IPlaylistUsecase interface.
type MockIPlaylistUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIPlaylistUsecaseMockRecorder
}

// MockIPlaylistUsecaseMockRecorder is the mock recorder for MockIPlaylistUsecase.
type MockIPlaylistUsecaseMockRecorder struct {
	mock *MockIPlaylistUsecase
}

// NewMockIPlaylistUsecase creates a new mock instance.
func NewMockIPlaylistUsecase(ctrl *gomock.Controller) *MockIPlaylistUsecase {
	mock := &MockIPlaylistUsecase{ctrl: ctrl}
	mock.recorder = &MockIPlaylistUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPlaylistUsecase) EXPECT() *MockIPlaylistUsecaseMockRecorder {
	return m.recorder
}

// AddSong mocks base method.
func (m *MockIPlaylistUsecase) AddSong(ctx context.Context, song *models.Song) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSong", ctx, song)
}

// AddSong indicates an expected call of AddSong.
func (mr *MockIPlaylistUsecaseMockRecorder) AddSong(ctx, song interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSong", reflect.TypeOf((*MockIPlaylistUsecase)(nil).AddSong), ctx, song)
}

// GetStatus mocks base method.
func (m *MockIPlaylistUsecase) GetStatus(ctx context.Context) (uuid.UUID, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", ctx)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockIPlaylistUsecaseMockRecorder) GetStatus(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockIPlaylistUsecase)(nil).GetStatus), ctx)
}

// Next mocks base method.
func (m *MockIPlaylistUsecase) Next(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockIPlaylistUsecaseMockRecorder) Next(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIPlaylistUsecase)(nil).Next), ctx)
}

// Pause mocks base method.
func (m *MockIPlaylistUsecase) Pause(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pause", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pause indicates an expected call of Pause.
func (mr *MockIPlaylistUsecaseMockRecorder) Pause(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockIPlaylistUsecase)(nil).Pause), ctx)
}

// Play mocks base method.
func (m *MockIPlaylistUsecase) Play(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Play", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Play indicates an expected call of Play.
func (mr *MockIPlaylistUsecaseMockRecorder) Play(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Play", reflect.TypeOf((*MockIPlaylistUsecase)(nil).Play), ctx)
}

// Prev mocks base method.
func (m *MockIPlaylistUsecase) Prev(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prev", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prev indicates an expected call of Prev.
func (mr *MockIPlaylistUsecaseMockRecorder) Prev(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockIPlaylistUsecase)(nil).Prev), ctx)
}
