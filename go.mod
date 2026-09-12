module github.com/FourLeggedLabs/protos

go 1.27.0

require (
	github.com/FourLeggedLabs/protos/gen/go v0.0.0
	google.golang.org/protobuf v1.36.12
)

replace github.com/FourLeggedLabs/protos/gen/go => ./gen/go
