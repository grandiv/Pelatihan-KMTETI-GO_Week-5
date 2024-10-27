package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	"Pelatihan-KMTETI-GO_Week-5/src/db"
	"Pelatihan-KMTETI-GO_Week-5/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}

type ProductResponse struct {
	Data []*Product `json:"data"`
}

type ProductRequest struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}

func GetAllProduct() (*ProductResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("product")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	var prodList []*Product

	for cur.Next(context.TODO()) {
		var prod model.Product
		cur.Decode(&prod)
		prodList = append(prodList, &Product{
			Name: prod.Name,
			Price: prod.Price,
			Stock: prod.Stock,
		})
	}
	return &ProductResponse{
		Data: prodList,
	}, nil
}

func CreateProduct(req io.Reader) error {
	var prodReq ProductRequest
	err := json.NewDecoder(req).Decode(&prodReq)
	if err != nil {
		return errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("product")
	_, err = coll.InsertOne(context.TODO(), model.Product {
		ID: primitive.NewObjectID(),
		Name: prodReq.Name,
		Price: prodReq.Price,
		Stock: prodReq.Stock,
	})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	return nil
}