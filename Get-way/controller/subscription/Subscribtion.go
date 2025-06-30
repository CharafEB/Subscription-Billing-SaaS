package subscription

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	prt "github.com/microservic/proto/subscriptionservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (app *Application) Subscribtion(w http.ResponseWriter, r *http.Request) {
	var SubscribtionData prt.SubscribtionData
	if err := json.NewDecoder(r.Body).Decode(&SubscribtionData); err != nil {
		app.HandleError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	conn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		app.HandleError(w, http.StatusInternalServerError, "Failed to connect to gRPC server", err)

		return
	}
	defer conn.Close()

	client := prt.NewClientHandlingServiceClient(conn)
	res, err := client.Subscribe(context.Background(), &SubscribtionData)
	if err != nil {
		app.HandleError(w, http.StatusInternalServerError, "Failed to Subscribe", err)
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res.GetResponse())

}
