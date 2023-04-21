package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/niubir/area-service/models"
)

const (
	area_file_path          = "./data"
	area_file_name_template = "area_%s.json"
	area_file_name_layout   = "20060102150405"
	area_top_code           = "100000"
)

var (
	areaCacheMu    sync.RWMutex
	areaCache      = make(map[string]models.Area)
	provincesCache = make(map[string]models.Areas)
	citiesCache    = make(map[string]models.Areas)
	districtsCache = make(map[string]models.Areas)
	streetsCache   = make(map[string]models.Areas)
)

func initArea() error {
	fileInfos, err := ioutil.ReadDir(area_file_path)
	if err != nil {
		return err
	}
	var (
		lastModTime     time.Time
		lastModFileName string
	)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		fileName := fileInfo.Name()
		if !strings.HasPrefix(fileName, "area_") || !strings.HasSuffix(fileName, ".json") {
			continue
		}
		if newModTime := fileInfo.ModTime(); newModTime.After(lastModTime) {
			lastModTime = newModTime
			lastModFileName = fileName
		}
	}

	if lastModFileName != "" {
		body, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", area_file_path, lastModFileName))
		if err != nil {
			return err
		}
		var areas models.Areas
		if err := json.Unmarshal(body, &areas); err != nil {
			return err
		}
		if err := saveAreasCache(areas); err != nil {
			return err
		}
	} else {
		if err := FreshAreas(); err != nil {
			return err
		}
	}

	return nil
}

func FreshAreas() error {
	areas, err := getAmapAreas()
	if err != nil {
		return err
	}
	if err := saveAreasFile(areas); err != nil {
		return err
	}
	if err := saveAreasCache(areas); err != nil {
		return err
	}
	return nil
}

func GetArea(code string) (*models.Area, error) {
	areaCacheMu.RLock()
	defer areaCacheMu.RUnlock()

	area, ok := areaCache[code]
	if !ok {
		return nil, fmt.Errorf("code not found: %s", code)
	}

	return &area, nil
}

func GetProvinces() (models.Areas, error) {
	areaCacheMu.RLock()
	defer areaCacheMu.RUnlock()

	return provincesCache[area_top_code], nil
}

func GetCities(provinceCode string) (models.Areas, error) {
	areaCacheMu.RLock()
	defer areaCacheMu.RUnlock()

	return citiesCache[provinceCode], nil
}

func GetDistricts(cityCode string) (models.Areas, error) {
	areaCacheMu.RLock()
	defer areaCacheMu.RUnlock()

	return districtsCache[cityCode], nil
}

func GetStreet(districtCode string) (models.Areas, error) {
	areaCacheMu.RLock()
	defer areaCacheMu.RUnlock()

	return streetsCache[districtCode], nil
}

func saveAreasFile(areas models.Areas) error {
	body, err := json.Marshal(areas)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf(
		area_file_path+"/"+area_file_name_template,
		time.Now().Format(area_file_name_layout),
	))
	if err != nil {
		return err
	}
	if _, err := file.Write(body); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}

func saveAreasCache(areas models.Areas) error {
	areaCacheMu.Lock()
	defer areaCacheMu.Unlock()

	areaCache = make(map[string]models.Area)
	provincesCache = make(map[string]models.Areas)
	citiesCache = make(map[string]models.Areas)
	districtsCache = make(map[string]models.Areas)
	streetsCache = make(map[string]models.Areas)

	for _, province := range areas {

		areaCache[province.Code] = models.Area{
			Name:   province.Name,
			Level:  province.Level,
			Code:   province.Code,
			Center: province.Center,
		}
		provincesCache[area_top_code] = append(provincesCache[area_top_code], models.Area{
			Name:   province.Name,
			Level:  province.Level,
			Code:   province.Code,
			Center: province.Center,
		})

		for _, city := range province.Subs {

			areaCache[city.Code] = models.Area{
				Name:   city.Name,
				Level:  city.Level,
				Code:   city.Code,
				Center: city.Center,
			}
			citiesCache[province.Code] = append(citiesCache[province.Code], models.Area{
				Name:   city.Name,
				Level:  city.Level,
				Code:   city.Code,
				Center: city.Center,
			})

			for _, district := range city.Subs {

				areaCache[district.Code] = models.Area{
					Name:   district.Name,
					Level:  district.Level,
					Code:   district.Code,
					Center: district.Center,
				}
				districtsCache[city.Code] = append(districtsCache[city.Code], models.Area{
					Name:   district.Name,
					Level:  district.Level,
					Code:   district.Code,
					Center: district.Center,
				})

				for _, street := range district.Subs {

					areaCache[street.Code] = models.Area{
						Name:   street.Name,
						Level:  street.Level,
						Code:   street.Code,
						Center: street.Center,
					}
					streetsCache[district.Code] = append(streetsCache[district.Code], models.Area{
						Name:   street.Name,
						Level:  street.Level,
						Code:   street.Code,
						Center: street.Center,
					})

				}
			}
		}
	}

	return nil
}
