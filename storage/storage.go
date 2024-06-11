package storage

import (
	pb "github.com/Mubinabd/RestaurantService/genproto"
)
type StorageI interface {
	Restaurant() RestaurantI
}

type RestaurantI interface {
	CreateRestaurant(req *pb.CreateRestaurantReq) (*pb.Void, error)
	UpdateRestaurant(req *pb.CreateRestaurantReq) (*pb.Void, error)
	DeleteRestaurant(id *pb.ById) (*pb.Void, error)
	GetRestaurant(req *pb.ById) (*pb.Restaurant, error)
	GetAllRestaurants(req *pb.AddressFilter) (*pb.Restaurants, error)
}
