package conn

import (
	"context"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
)

type (
	ctxTransactionControlKey struct{}
	ctxDataQueryOptionsKey   struct{}
	ctxScanQueryOptionsKey   struct{}
	ctxModeTypeKey           struct{}
	ctxTxControlHookKey      struct{}

	txControlHook func(txControl *table.TransactionControl)
)

func WithTxControlHook(ctx context.Context, hook txControlHook) context.Context {
	return context.WithValue(ctx, ctxTxControlHookKey{}, hook)
}

// WithQueryMode returns a copy of context with given QueryMode
func WithQueryMode(ctx context.Context, m QueryMode) context.Context {
	return context.WithValue(ctx, ctxModeTypeKey{}, m)
}

// QueryModeFromContext returns defined QueryMode or DefaultQueryMode
func QueryModeFromContext(ctx context.Context, defaultQueryMode QueryMode) QueryMode {
	if m, ok := ctx.Value(ctxModeTypeKey{}).(QueryMode); ok {
		return m
	}

	return defaultQueryMode
}

func WithTxControl(ctx context.Context, txc *table.TransactionControl) context.Context {
	return context.WithValue(ctx, ctxTransactionControlKey{}, txc)
}

func txControl(ctx context.Context, defaultTxControl *table.TransactionControl) (txControl *table.TransactionControl) {
	defer func() {
		if hook, has := ctx.Value(ctxTxControlHookKey{}).(txControlHook); has && hook != nil {
			hook(txControl)
		}
	}()
	if txc, ok := ctx.Value(ctxTransactionControlKey{}).(*table.TransactionControl); ok {
		return txc
	}

	return defaultTxControl
}

func (c *Conn) WithScanQueryOptions(ctx context.Context, opts ...options.ExecuteScanQueryOption) context.Context {
	return context.WithValue(ctx,
		ctxScanQueryOptionsKey{},
		append(
			append([]options.ExecuteScanQueryOption{}, c.scanQueryOptions(ctx)...),
			opts...,
		),
	)
}

func (c *Conn) scanQueryOptions(ctx context.Context) []options.ExecuteScanQueryOption {
	if opts, ok := ctx.Value(ctxScanQueryOptionsKey{}).([]options.ExecuteScanQueryOption); ok {
		return append(c.scanOpts, opts...)
	}

	return c.scanOpts
}

func (c *Conn) WithDataQueryOptions(ctx context.Context, opts ...options.ExecuteDataQueryOption) context.Context {
	return context.WithValue(ctx,
		ctxDataQueryOptionsKey{},
		append(
			append([]options.ExecuteDataQueryOption{}, c.dataQueryOptions(ctx)...),
			opts...,
		),
	)
}

func (c *Conn) dataQueryOptions(ctx context.Context) []options.ExecuteDataQueryOption {
	if opts, ok := ctx.Value(ctxDataQueryOptionsKey{}).([]options.ExecuteDataQueryOption); ok {
		return append(c.dataOpts, opts...)
	}

	return c.dataOpts
}

func (c *Conn) withKeepInCache(ctx context.Context) context.Context {
	return c.WithDataQueryOptions(ctx, options.WithKeepInCache(true))
}
