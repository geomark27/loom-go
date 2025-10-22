package helpers

import (
	"context"
	"time"
)

// ContextKey tipo para keys de contexto
type ContextKey string

const (
	// UserIDKey clave para userID en contexto
	UserIDKey ContextKey = "userID"
	// RequestIDKey clave para requestID en contexto
	RequestIDKey ContextKey = "requestID"
	// TenantIDKey clave para tenantID en contexto
	TenantIDKey ContextKey = "tenantID"
)

// GetUserID obtiene userID del contexto
func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDKey).(string)
	return userID, ok
}

// SetUserID establece userID en contexto
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// GetRequestID obtiene requestID del contexto
func GetRequestID(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(RequestIDKey).(string)
	return requestID, ok
}

// SetRequestID establece requestID en contexto
func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

// GetTenantID obtiene tenantID del contexto
func GetTenantID(ctx context.Context) (string, bool) {
	tenantID, ok := ctx.Value(TenantIDKey).(string)
	return tenantID, ok
}

// SetTenantID establece tenantID en contexto
func SetTenantID(ctx context.Context, tenantID string) context.Context {
	return context.WithValue(ctx, TenantIDKey, tenantID)
}

// WithTimeout crea contexto con timeout
func WithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, timeout)
}

// WithDeadline crea contexto con deadline
func WithDeadline(ctx context.Context, deadline time.Time) (context.Context, context.CancelFunc) {
	return context.WithDeadline(ctx, deadline)
}
