package services

import (
	"context"
	"github.com/dalmarcogd/gbpl-go/internal/models"
	"gorm.io/gorm"
)

type (
	NoopDatabase    struct{}
	NoopGrpcServer  struct{}
	NoopCache       struct{}
	NoopLogger      struct{}
	NoopHandlers    struct{}
	NoopEnvironment struct{}
)

func NewNoopDatabase() *NoopDatabase {
	return &NoopDatabase{}
}

func (n *NoopDatabase) ServiceManager() ServiceManager {
	return nil
}

func (n *NoopDatabase) Init(_ context.Context) error {
	return nil
}

func (n *NoopDatabase) Close() error {
	return nil
}

func (n *NoopDatabase) WithServiceManager(_ ServiceManager) Database {
	return n
}

func (n *NoopDatabase) DB(_ context.Context) *gorm.DB {
	return nil
}

func NewNoopGrpcServer() *NoopGrpcServer {
	return &NoopGrpcServer{}
}

func (n *NoopGrpcServer) ServiceManager() ServiceManager {
	return nil
}

func (n *NoopGrpcServer) Init(_ context.Context) error {
	return nil
}

func (n *NoopGrpcServer) Close() error {
	return nil
}

func (n *NoopGrpcServer) WithServiceManager(_ ServiceManager) GrpcServer {
	return n
}

func (n *NoopGrpcServer) Run() error {
	return nil
}

func (n *NoopGrpcServer) Address() string {
	return ""
}

func NewNoopCache() *NoopCache {
	return &NoopCache{}
}

func (n *NoopCache) ServiceManager() ServiceManager {
	return nil
}

func (n *NoopCache) Init(_ context.Context) error {
	return nil
}

func (n *NoopCache) Close() error {
	return nil
}

func (n *NoopCache) WithServiceManager(_ ServiceManager) Cache {
	return n
}

func NewNoopLogger() *NoopLogger {
	return &NoopLogger{}
}

func (n *NoopLogger) ServiceManager() ServiceManager {
	return nil
}

func (n *NoopLogger) Init(_ context.Context) error {
	return nil
}

func (n *NoopLogger) Close() error {
	return nil
}

func (n *NoopLogger) WithServiceManager(_ ServiceManager) Logger {
	return n
}

func (n *NoopLogger) Info(_ context.Context, _ string, _ ...map[string]interface{}) {}

func (n *NoopLogger) Warn(_ context.Context, _ string, _ ...map[string]interface{}) {}

func (n *NoopLogger) Error(_ context.Context, _ string, _ ...map[string]interface{}) {}

func (n *NoopLogger) Fatal(_ context.Context, _ string, _ ...map[string]interface{}) {}

func NewNoopHandlers() *NoopHandlers {
	return &NoopHandlers{}
}

func (n *NoopHandlers) ServiceManager() ServiceManager {
	return nil
}

func (n *NoopHandlers) Init(_ context.Context) error {
	return nil
}

func (n *NoopHandlers) Close() error {
	return nil
}

func (n *NoopHandlers) WithServiceManager(_ ServiceManager) Handlers {
	return n
}

func (n *NoopHandlers) CreateUser(_ context.Context, _ *models.User) error {
	return nil
}

func (n *NoopHandlers) UpdateUser(_ context.Context, _ *models.User) error {
	return nil
}

func (n *NoopHandlers) GetUser(_ context.Context, _ *models.User) error {
	return nil
}

func (n *NoopHandlers) GetUsers(_ context.Context, _ *[]models.User) error {
	return nil
}

func (n *NoopHandlers) DeleteUser(_ context.Context, _ *models.User) error {
	return nil
}

func NewNoopEnvironment() *NoopEnvironment {
	return &NoopEnvironment{}
}

func (n *NoopEnvironment) ServiceManager() ServiceManager {
	return nil
}

func (n *NoopEnvironment) Init(_ context.Context) error {
	return nil
}

func (n *NoopEnvironment) Close() error {
	return nil
}

func (n *NoopEnvironment) WithServiceManager(_ ServiceManager) Environment {
	return n
}

func (n *NoopEnvironment) Environment() string {
	return ""
}

func (n *NoopEnvironment) Service() string {
	return ""
}

func (n *NoopEnvironment) Version() string {
	return ""
}

func (n *NoopEnvironment) DebugPprof() bool {
	return false
}

func (n *NoopEnvironment) DatabaseDsn() string {
	return ""
}

func (n *NoopEnvironment) CacheAddress() string {
	return ""
}
