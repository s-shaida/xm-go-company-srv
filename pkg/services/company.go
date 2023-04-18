package services

import (
	"context"
	"net/http"

	"github.com/s-shaida/xm-go-company-svc/pkg/db"
	"github.com/s-shaida/xm-go-company-svc/pkg/models"
	pb "github.com/s-shaida/xm-go-company-svc/pkg/pb"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	var company models.Company

	company.Name = req.Name
	company.Description = req.Description
	company.AmountOfEmployees = req.Amount
	company.Registered = req.Registered
	company.Type = pb.CompanyType_name[int32(req.Type.Number())]

	if result := s.H.DB.Create(&company); result.Error != nil {
		return &pb.CreateCompanyResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateCompanyResponse{
		Status: http.StatusCreated,
		Id:     company.Id,
	}, nil
}

func (s *Server) GetOneCompany(ctx context.Context, req *pb.GetOneCompanyRequest) (*pb.GetOneCompanyResponse, error) {
	var company models.Company

	if result := s.H.DB.First(&company, req.Id); result.Error != nil {
		return &pb.GetOneCompanyResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.GetOneCompanyData{
		Id:          company.Id,
		Name:        company.Name,
		Description: company.Description,
		Amount:      company.AmountOfEmployees,
		Registered:  company.Registered,
		Type:        pb.CompanyType(pb.CompanyType_value[company.Type]),
	}

	return &pb.GetOneCompanyResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) PatchCompany(ctx context.Context, req *pb.PatchCompanyRequest) (*pb.PatchCompanyResponse, error) {
	var company models.Company

	if result := s.H.DB.First(&company, req.Id); result.Error != nil {
		return &pb.PatchCompanyResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	company.Name = req.Name
	company.Description = req.Description
	company.AmountOfEmployees = req.Amount
	company.Registered = req.Registered
	company.Type = pb.CompanyType_name[int32(req.Type.Number())]

	if result := s.H.DB.Save(&company); result.Error != nil {
		return &pb.PatchCompanyResponse{
			Status: http.StatusBadRequest,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.PatchCompanyResponse{
		Status: http.StatusOK,
		Id:     company.Id,
	}, nil
}

func (s *Server) DeleteCompany(ctx context.Context, req *pb.DeleteCompanyRequest) (*pb.DeleteCompanyResponse, error) {
	var company models.Company

	if result := s.H.DB.Delete(&company, req.Id); result.Error != nil {
		return &pb.DeleteCompanyResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.DeleteCompanyResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil
}
