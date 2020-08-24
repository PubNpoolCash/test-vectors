package main

import (
	"github.com/filecoin-project/test-vectors/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("../chaos/cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CreateActorArgs{},
	); err != nil {
		panic(err)
	}
}