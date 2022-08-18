package keeper

import (
	"github.com/psangwoo/parsdk/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
