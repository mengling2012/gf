// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

import (
	"context"

	"github.com/gogf/gf/v2/os/gctx"
)

// internalCtxData stores data in ctx for internal usage purpose.
type internalCtxData struct {
	// Operation DB.
	DB DB

	// The first column in result response from database server.
	// This attribute is used for Value/Count selection statement purpose,
	// which is to avoid HOOK handler that might modify the result columns
	// that can confuse the Value/Count selection statement logic.
	FirstResultColumn string
}

const (
	internalCtxDataKeyInCtx gctx.StrKey = "InternalCtxData"

	// IgnoreResultInCtx
	// This option is only available in ClickHouse.
	// Because ClickHouse does not support fetching insert/update results and returns errors when executed
	// So need to ignore the results to avoid triggering errors
	// Rather than ignoring errors after they are triggered
	IgnoreResultInCtx gctx.StrKey = "IgnoreResult"
)

func (c *Core) injectInternalCtxData(ctx context.Context) context.Context {
	// If the internal data is already injected, it does nothing.
	if ctx.Value(internalCtxDataKeyInCtx) != nil {
		return ctx
	}
	return context.WithValue(ctx, internalCtxDataKeyInCtx, &internalCtxData{
		DB: c.db,
	})
}

func (c *Core) InjectIgnoreResult(ctx context.Context) context.Context {
	if ctx.Value(IgnoreResultInCtx) != nil {
		return ctx
	}
	return context.WithValue(ctx, IgnoreResultInCtx, &internalCtxData{
		DB: c.db,
	})
}

func (c *Core) GetIgnoreResultFromCtx(ctx context.Context) *internalCtxData {
	if v := ctx.Value(internalCtxDataKeyInCtx); v != nil {
		return v.(*internalCtxData)
	}
	return nil
}

func (c *Core) getInternalCtxDataFromCtx(ctx context.Context) *internalCtxData {
	if v := ctx.Value(internalCtxDataKeyInCtx); v != nil {
		return v.(*internalCtxData)
	}
	return nil
}
