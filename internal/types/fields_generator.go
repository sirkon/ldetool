package types

//go:generate go install ./internal/gen-builtin
//go:generate  gen-builtin --type Int --name int
//go:generate  gen-builtin --type Int8 --name int8
//go:generate  gen-builtin --type Int16 --name int16
//go:generate  gen-builtin --type Int32 --name int32
//go:generate  gen-builtin --type Int64 --name int64
//go:generate  gen-builtin --type Uint --name uint
//go:generate  gen-builtin --type Uint8 --name uint8
//go:generate  gen-builtin --type Uint16 --name uint16
//go:generate  gen-builtin --type Uint32 --name uint32
//go:generate  gen-builtin --type Uint64 --name uint64
//go:generate  gen-builtin --type Hex --name hex     --handler AddUint
//go:generate  gen-builtin --type Hex8 --name hex8   --handler AddUint8
//go:generate  gen-builtin --type Hex16 --name hex16 --handler AddUint16
//go:generate  gen-builtin --type Hex32 --name hex32 --handler AddUint32
//go:generate  gen-builtin --type Hex64 --name hex64 --handler AddUint64
//go:generate  gen-builtin --type Oct --name oct     --handler AddUint
//go:generate  gen-builtin --type Oct8 --name oct8   --handler AddUint8
//go:generate  gen-builtin --type Oct16 --name oct16 --handler AddUint16
//go:generate  gen-builtin --type Oct32 --name oct32 --handler AddUint32
//go:generate  gen-builtin --type Oct64 --name oct64 --handler AddUint64
//go:generate  gen-builtin --type Dec32 --name dec32 --handler AddInt32
//go:generate  gen-builtin --type Dec64 --name dec64 --handler AddInt64
//go:generate  gen-builtin --type Dec128 --name dec128 --native "struct { Lo uint64; Hi uint64 }"
//go:generate  gen-builtin --type Float32 --name float32
//go:generate  gen-builtin --type Float64 --name float64
//go:generate  gen-builtin --type String --name string
//go:generate  gen-builtin --type Str --name str
