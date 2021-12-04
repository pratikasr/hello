package keeper

import (
	"github.com/pratikasr/hello/x/hello/types"
)

var _ types.QueryServer = Keeper{}
