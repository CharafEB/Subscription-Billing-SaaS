package castmer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	prt "github.com/microservic/proto/castmerservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	var LoginData prt.LoginData
	if err := json.NewDecoder(r.Body).Decode(&LoginData); err != nil {
		app.HandleError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		app.HandleError(w, http.StatusInternalServerError, "Failed to connect to gRPC server", err)

		return
	}
	defer conn.Close()

	client := prt.NewClientHandlingServiceClient(conn)
	res, err := client.Login(context.Background(), &LoginData) //<-- هنا تباسي ال data 
	if err != nil {
		app.HandleError(w, http.StatusInternalServerError, "Failed to call Login", err)
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res.GetResponse())

}
