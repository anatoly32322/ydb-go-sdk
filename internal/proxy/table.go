package proxy

import (
	"context"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/meta"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

type proxyTable struct {
	client table.Client
	meta   meta.Meta
}

func Table(client table.Client, meta meta.Meta) table.Client {
	return &proxyTable{
		client: client,
		meta:   meta,
	}
}

func (t *proxyTable) CreateSession(ctx context.Context) (s table.ClosableSession, err error) {
	ctx, err = t.meta.Meta(ctx)
	if err != nil {
		return nil, err
	}
	return t.client.CreateSession(ctx)
}

func (t *proxyTable) Do(ctx context.Context, op table.Operation, opts ...table.Option) (err error) {
	ctx, err = t.meta.Meta(ctx)
	if err != nil {
		return err
	}
	return t.client.Do(ctx, op, opts...)
}

func (t *proxyTable) DoTx(ctx context.Context, op table.TxOperation, opts ...table.Option) (err error) {
	ctx, err = t.meta.Meta(ctx)
	if err != nil {
		return err
	}
	return t.client.DoTx(ctx, op, opts...)
}

func (t *proxyTable) Close(ctx context.Context) (err error) {
	// nop
	return nil
}