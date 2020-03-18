package main

import (
	pb "caesar-go/consignment-service/proto/consignment"
	"context"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
}

// 模拟数据库
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// service中要实现pb中所有的方法
type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// 保存我们的consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {

}
