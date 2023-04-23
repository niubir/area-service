package grpc

import (
	"context"
	"encoding/json"

	"github.com/niubir/area-service/grpc/server"
	"github.com/niubir/area-service/models"
	"github.com/niubir/area-service/services"
)

type AreaService struct {
	server.UnsafeAreaServiceServer
}

// 刷新区域
func (as *AreaService) FreshAreas(ctx context.Context, req *server.Empty) (*server.Empty, error) {
	var ack server.Empty
	if err := services.FreshAreas(); err != nil {
		return &ack, err
	}
	return &ack, nil
}

// 获取指定区域
func (as *AreaService) GetArea(ctx context.Context, req *server.GetAreaReq) (*server.GetAreaAck, error) {
	var ack server.GetAreaAck
	area, err := services.GetArea(req.Code)
	if err != nil {
		return &ack, err
	}
	ack.Area, err = as.areaExchange(area)
	if err != nil {
		return &ack, err
	}
	return &ack, nil
}

// 获取省份列表
func (as *AreaService) GetProvinces(ctx context.Context, req *server.GetProvincesReq) (*server.GetProvincesAck, error) {
	var ack server.GetProvincesAck
	provinces, err := services.GetProvinces()
	if err != nil {
		return &ack, err
	}
	ack.Provinces, err = as.areasExchange(provinces)
	if err != nil {
		return &ack, err
	}
	return &ack, nil
}

// 获取城市列表
func (as *AreaService) GetCities(ctx context.Context, req *server.GetCitiesReq) (*server.GetCitiesAck, error) {
	var ack server.GetCitiesAck
	cities, err := services.GetCities(req.ProvinceCode)
	if err != nil {
		return &ack, err
	}
	ack.Cities, err = as.areasExchange(cities)
	if err != nil {
		return &ack, err
	}
	return &ack, nil
}

// 获取区县列表
func (as *AreaService) GetDistricts(ctx context.Context, req *server.GetDistrictsReq) (*server.GetDistrictsAck, error) {
	var ack server.GetDistrictsAck
	districts, err := services.GetDistricts(req.CityCode)
	if err != nil {
		return &ack, err
	}
	ack.Districts, err = as.areasExchange(districts)
	if err != nil {
		return &ack, err
	}
	return &ack, nil
}

// 获取街道列表
func (as *AreaService) GetStreets(ctx context.Context, req *server.GetStreetsReq) (*server.GetStreetsAck, error) {
	var ack server.GetStreetsAck
	streets, err := services.GetStreets(req.DistrictCode)
	if err != nil {
		return &ack, err
	}
	ack.Streets, err = as.areasExchange(streets)
	if err != nil {
		return &ack, err
	}
	return &ack, nil
}

func (as AreaService) areasExchange(areas models.Areas) ([]*server.Area, error) {
	b, err := json.Marshal(areas)
	if err != nil {
		return nil, err
	}
	var resp []*server.Area
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (as AreaService) areaExchange(area *models.Area) (*server.Area, error) {
	b, err := json.Marshal(area)
	if err != nil {
		return nil, err
	}
	var resp *server.Area
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
