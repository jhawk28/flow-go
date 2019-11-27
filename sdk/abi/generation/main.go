package main

import (
	"os"

	"github.com/dapperlabs/flow-go/sdk/abi/generation/code"
	"github.com/dapperlabs/flow-go/sdk/abi/types"
)

func main() {

	if len(os.Args) != 2 {
		panic("use filename as one and only argument")
	}

	//abiFilename := os.Args[1]

	types := map[string]*types.Composite{
		"Car": {
			Fields: map[string]*types.Field{
				"fullName": {
					Identifier: "fullName",
					Type:       types.String{},
				},
			},
			Identifier: "Car",
			Initializers: [][]*types.Parameter{
				{
					&types.Parameter{
						Field: types.Field{
							Identifier: "model",
							Type:       types.String{},
						},
						Label: "",
					},
					&types.Parameter{
						Field: types.Field{
							Identifier: "make",
							Type:       types.String{},
						},
						Label: "",
					},
				},
			},
		},
	}

	code.GenerateGo("examples", types)

}
