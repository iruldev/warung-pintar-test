package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/iruldev/warung-pintar-test/shipping-service/proto/shipping"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"reflect"
)

type shippingService struct {
	pb.UnimplementedShippingServer
}

func (service *shippingService) CalculateShipping(ctx context.Context, empty *emptypb.Empty) (*pb.CalculateShippingResponse, error) {
	baseURL := fmt.Sprintf("http://%s", os.Getenv("CART_SERVICE_HOST"))
	res, err := http.Get(baseURL + "/api/carts")
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Cart not found")
	}

	body, _ := io.ReadAll(res.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	data := responseBody["data"].(map[string]interface{})
	carts := data["carts"]

	totalPrice := 0
	var totalWeight float64

	for i := 0; i < reflect.ValueOf(carts).Len(); i++ {
		cart := carts.([]interface{})[i].(map[string]interface{})
		log.Println(cart)
		q := 0
		p := 0
		w := 0
		for key, value := range cart {
			if key == "quantity" {
				q += int(value.(float64))
			}

			if  key == "product" {
				for pkey, pval := range value.(map[string]interface{}) {
					if pkey == "price" {
						p += int(pval.(float64))
					}

					if pkey == "weight" {
						w += int(pval.(float64))
					}
				}
			}
		}

		totalPrice += q * p
		totalWeight += math.Ceil(float64(q * w) / 1000)
	}

	totalPrice = (int(totalWeight) * 3000) + totalPrice

	log.Println("Calculating shipping cost for :", data["id"])
	cost := &pb.CalculateShippingResponse{Cost: int32(totalPrice)}
	return cost, nil
}

func main() {
	log.Println("Starting shipping service.")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Error in listen", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterShippingServer(grpcServer, &shippingService{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Error when server grpc", err.Error())
	}
}