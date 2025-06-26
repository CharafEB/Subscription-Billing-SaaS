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

func (app *Application) SingUp(w http.ResponseWriter, r *http.Request) {
	var signUpData prt.SingUpData
	if err := json.NewDecoder(r.Body).Decode(&signUpData); err != nil {
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
	res, err := client.SignUp(context.Background(), &signUpData) //<-- هنا تباسي ال data 
	if err != nil {
		app.HandleError(w, http.StatusInternalServerError, "Failed to call SingUp", err)
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res.GetResponse())

}
