package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

const (
	dbServer   = "localhost\\DB2019"
	dbUser     = "sa"
	dbPassword = "Tel@12345"
	dbName     = "gotest"
)

type Subscription struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type SubscriberService struct {
	db *sql.DB
}

func (s *SubscriberService) GetSubscriberDetails(sub_id int) (*Subscription, error) {
	var content Subscription
	err := s.db.QueryRow("SELECT sub_id,Sub_name,descrptn from [dbo].[subscription_dtls] where sub_id = @sub_id", sql.Named("sub_id", sub_id)).Scan(&content.ID, &content.Title, &content.Description)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (s *SubscriberService) GetAllSubscribers(plan_id int) ([]Subscription, error) {
	rows, err := s.db.Query("SELECT sub_id,Sub_name,descrptn FROM [dbo].[subscription_dtls] g inner join allsubscrptn_dtls m on m. [Plan_id]= g.[plan_id]  WHERE g.[Plan_id] = @Plan_id", sql.Named("Plan_id", plan_id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subscriptions := []Subscription{}

	for rows.Next() {
		var subscription Subscription
		err := rows.Scan(&subscription.ID, &subscription.Title, &subscription.Description)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil

}

func makeGetSubscriberDetailsEndpoint(svc SubscriberService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSubscriberDetailsRequest)
		content, err := svc.GetSubscriberDetails(req.sub_id)
		if err != nil {
			return nil, err
		}
		return content, nil
	}
}

func makeGetAllSubscribersEndpoint(svc SubscriberService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllSubscribersRequest)
		content, err := svc.GetAllSubscribers(req.plan_id)
		if err != nil {
			return nil, err
		}
		return content, nil
	}
}

func decodeGetSubscriberDetailsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	sub_id, err := strconv.Atoi(vars["sub_id"])
	if err != nil {
		return nil, err
	}
	return GetSubscriberDetailsRequest{sub_id: sub_id}, nil
}

func decodeGetAllSubscribersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	plan_id, err := strconv.Atoi(vars["plan_id"])
	if err != nil {
		return nil, err
	}
	return GetAllSubscribersRequest{plan_id: plan_id}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type GetSubscriberDetailsRequest struct {
	sub_id int `json:"sub_id"`
}

type GetAllSubscribersRequest struct {
	plan_id int `json:"plan_id"`
}

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", dbServer, dbUser, dbPassword, dbName)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	svc := SubscriberService{db: db}

	GetSubscriberDetailsHandler := httptransport.NewServer(
		makeGetSubscriberDetailsEndpoint(svc),
		decodeGetSubscriberDetailsRequest,
		encodeResponse,
	)

	GetAllSubscribersHandler := httptransport.NewServer(
		makeGetAllSubscribersEndpoint(svc),
		decodeGetAllSubscribersRequest,
		encodeResponse,
	)

	r := mux.NewRouter()

	r.Handle("/subscription/{sub_id}", GetSubscriberDetailsHandler).Methods("GET")
	r.Handle("/plan/{plan_id}", GetAllSubscribersHandler).Methods("GET")

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
