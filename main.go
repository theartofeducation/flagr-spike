package main

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/davecgh/go-spew/spew"
	"github.com/theantichris/goflagr"
)

const (
	flagrURL = "http://localhost:18000/api/v1"
)

func main() {
	config := goflagr.NewConfiguration()
	config.BasePath = flagrURL
	client := goflagr.NewAPIClient(config)

	ctx := context.Background()
	auth := goflagr.BasicAuth{
		UserName: "flagr_username",
		Password: "flagr_password",
	}
	ctx = context.WithValue(ctx, goflagr.ContextBasicAuth, auth)

	preload := optional.NewBool(true)
	params := goflagr.FindFlagsOpts{
		Preload: preload,
	}
	flags, response, err := client.FlagApi.FindFlags(ctx, &params)
	if err != nil {
		fmt.Printf("error getting flag: %s\n", err.Error())
	}

	fmt.Printf("http status: %d\n\n", response.StatusCode)

	fmt.Print("Flag:\n\n")
	spew.Dump(flags)

	// fmt.Println("Flag 2:")
	// spew.Dump(flags[0])

	// fmt.Println("Flag 3:")
	// spew.Dump(flags[1])

	// fmt.Println("Flag 3 Variants:")
	// spew.Dump(flags[1].Variants)
}
