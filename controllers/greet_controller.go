package controllers

import (
	"context"
	"fmt"
	"net/http"

	greetv1 "github.com/Inokuchi-Kento/ikapp/gen/greet/v1"
	"github.com/Inokuchi-Kento/ikapp/gen/greet/v1/greetv1connect"
	"github.com/bufbuild/connect-go"
)

type GreetController struct {
}

func (c *GreetController) GreetHandler() (string, http.Handler) {
	return greetv1connect.NewGreetServiceHandler(c)
}

func (c *GreetController) Greet(
	ctx context.Context, req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
