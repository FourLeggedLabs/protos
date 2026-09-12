module github.com/FourLeggedLabs/protos

go 1.26.6

require github.com/FourLeggedLabs/protos/gen/go v0.0.0

require google.golang.org/protobuf v1.36.12

replace github.com/FourLeggedLabs/protos/gen/go => ./gen/go
