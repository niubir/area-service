package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/niubir/area-service/config"
	"github.com/niubir/area-service/models"
)

type AmapAreas []AmapArea

type AmapArea struct {
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Adcode    string    `json:"adcode"`
	Center    string    `json:"center"`
	Districts AmapAreas `json:"districts"`
}

func getAmapAreas() (models.Areas, error) {
	provinces, err := getAmapAreasQuery(area_top_code, 1)
	if err != nil {
		return nil, err
	}
	for i, province := range provinces {
		cities, err := getAmapAreasQuery(province.Code, 3)
		if err != nil {
			return nil, err
		}
		provinces[i].Subs = cities
	}
	return provinces, nil
}

func getAmapAreasQuery(code string, subdistrict int) (models.Areas, error) {
	response, err := http.Get(fmt.Sprintf(
		"https://restapi.amap.com/v3/config/district?key=%s&keywords=%s&subdistrict=%d&extensions=base&filter=adcode",
		config.Config.AmapKey,
		code,
		subdistrict,
	))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err := checkAmapBody(body); err != nil {
		return nil, err
	}
	var resp struct {
		Districts AmapAreas `json:"districts"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	if len(resp.Districts) == 0 {
		return nil, fmt.Errorf("code not found: %s", code)
	}
	if resp.Districts[0].Adcode != code {
		return nil, fmt.Errorf("code not found: %s", code)
	}
	return resp.Districts[0].Districts.ToAreas(), nil
}

func checkAmapBody(body []byte) error {
	var resp struct {
		Status   string `json:"status"`
		Info     string `json:"info"`
		Infocode string `json:"infocode"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}
	if resp.Status != "1" || resp.Infocode != "10000" {
		if resp.Info == "" || resp.Info == "OK" {
			resp.Info = "Unknow error"
		}
		return errors.New(resp.Info)
	}
	return nil
}

func (ds AmapAreas) ToAreas() models.Areas {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Adcode < ds[j].Adcode
	})
	var areas models.Areas
	for _, sub := range ds {
		areas = append(areas, sub.ToArea())
	}
	return areas
}

func (d AmapArea) ToArea() models.Area {
	var level models.Constant
	switch d.Level {
	case "country":
		level = models.AREA_LEVEL_COUNTRY
	case "province":
		level = models.AREA_LEVEL_PROVINCE
	case "city":
		level = models.AREA_LEVEL_CITY
	case "district":
		level = models.AREA_LEVEL_DISTRICT
	case "street":
		level = models.AREA_LEVEL_STREET
	}
	sort.Slice(d.Districts, func(i, j int) bool {
		return d.Districts[i].Adcode < d.Districts[j].Adcode
	})
	var subs models.Areas
	for _, sub := range d.Districts {
		subs = append(subs, sub.ToArea())
	}
	area := models.Area{
		Name:   d.Name,
		Level:  level,
		Code:   d.Adcode,
		Center: d.Center,
		Subs:   subs,
	}
	return area
}
