package main

import (
	stdctx "context"
	"fmt"
	"github.com/drywaters/lenslocked/context"
	"github.com/drywaters/lenslocked/models"
)

func main() {
	ctx := stdctx.Background()

	user := models.User{
		Email: "daniel@bitofbytes.io",
	}
	ctx = context.WithUser(ctx, &user)

	retrievedUser := context.User(ctx)
	fmt.Println(retrievedUser.Email)
}
