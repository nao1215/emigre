//go:build tools
// +build tools

package tools

import (
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/k1LoW/tbls"
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
	_ "github.com/swaggo/swag/cmd/swag"
)
