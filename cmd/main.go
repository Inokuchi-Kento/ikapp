package main

import (
	"net/http"

	"github.com/Inokuchi-Kento/ikapp/controllers"
	"github.com/Inokuchi-Kento/ikapp/gen/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	greeter := &controllers.GreetController{}
	mux := http.NewServeMux()

	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	http.ListenAndServe("localhost:8010", h2c.NewHandler(mux, &http2.Server{}))
}

// func main() {
// 	if err := run(context.Background()); err != nil {
// 		log.Printf("failed to terminate server: %v", err)
// 		os.Exit(1)
// 	}
// }

// func run(ctx context.Context) error {
// 	cfg, err := config.New()
// 	if err != nil {
// 		return err
// 	}

// 	db := repositories.OpenDB(context.Background(), cfg)
// 	fmt.Println(db)

// 	return nil
// }
