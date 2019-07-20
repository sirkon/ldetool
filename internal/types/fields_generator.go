package types

//go:generate go install ./internal/gen-builtin
//go:generate  gen-builtin --declarable --native --type Int --name int
//go:generate  gen-builtin --declarable --native --type Int8 --name int8
//go:generate  gen-builtin --declarable --native --type Int16 --name int16
//go:generate  gen-builtin --declarable --native --type Int32 --name int32
//go:generate  gen-builtin --declarable --native --type Int64 --name int64
//go:generate  gen-builtin --declarable --native --type Uint --name uint
//go:generate  gen-builtin --declarable --native --type Uint8 --name uint8
//go:generate  gen-builtin --declarable --native --type Uint16 --name uint16
//go:generate  gen-builtin --declarable --native --type Uint32 --name uint32
//go:generate  gen-builtin --declarable --native --type Uint64 --name uint64
//go:generate  gen-builtin --declarable --type Hex --name hex     --handler AddUint   --go-name uint
//go:generate  gen-builtin --declarable --type Hex8 --name hex8   --handler AddUint8  --go-name uint8
//go:generate  gen-builtin --declarable --type Hex16 --name hex16 --handler AddUint16 --go-name uint16
//go:generate  gen-builtin --declarable --type Hex32 --name hex32 --handler AddUint32 --go-name uint32
//go:generate  gen-builtin --declarable --type Hex64 --name hex64 --handler AddUint64 --go-name uint64
//go:generate  gen-builtin --declarable --type Oct --name oct     --handler AddUint   --go-name uint
//go:generate  gen-builtin --declarable --type Oct8 --name oct8   --handler AddUint8  --go-name uint8
//go:generate  gen-builtin --declarable --type Oct16 --name oct16 --handler AddUint16 --go-name uint16
//go:generate  gen-builtin --declarable --type Oct32 --name oct32 --handler AddUint32 --go-name uint32
//go:generate  gen-builtin --declarable --type Oct64 --name oct64 --handler AddUint64 --go-name uint64
//go:generate  gen-builtin --decimal --type Dec32 --name dec32 --handler AddInt32 --go-name int32
//go:generate  gen-builtin --decimal --type Dec64 --name dec64 --handler AddInt64 --go-name int64
//go:generate  gen-builtin --decimal --type Dec128 --name dec128 --go-name "struct { Lo uint64; Hi uint64 }"
//go:generate  gen-builtin --declarable --native --type Float32 --name float32
//go:generate  gen-builtin --declarable --native --type Float64 --name float64
//go:generate  gen-builtin --declarable --native --type String --name string
//go:generate  gen-builtin --declarable --native --type Str --name str
