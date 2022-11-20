package utils

import (
	"fmt"
	"github.com/jaimera/poc-services/domain/dto"
	"reflect"
)

// ConvertPortJson converts ports.json into a list of Port struct
func ConvertPortJson(jsonMap map[string]map[string]interface{}) ([]dto.PortDto, int, error) {
	var err error

	ports := make([]dto.PortDto, 0, len(jsonMap))
	for code, obj := range jsonMap {
		port := dto.PortDto{
			Slug:     code,
			Name:     fmt.Sprint(obj["name"]),
			City:     fmt.Sprint(obj["city"]),
			Province: fmt.Sprint(obj["province"]),
			Country:  fmt.Sprint(obj["country"]),
			Timezone: fmt.Sprint(obj["timezone"]),
		}

		if obj["code"] != nil {
			codeValue := fmt.Sprint(obj["code"])
			port.Code = &codeValue
		}

		if obj["coordinates"] != nil {
			coordinates := unpackArray(obj["coordinates"])

			port.Latitude, err = getFloat(coordinates[0])
			if err != nil {
				return nil, 0, err
			}
			port.Longitude, err = getFloat(coordinates[1])
			if err != nil {
				return nil, 0, err
			}
		}

		alias := unpackArray(obj["alias"])
		if len(alias) > 0 {
			for _, v := range alias {
				port.Alias = append(port.Alias, fmt.Sprint(v))
			}
		}

		regions := unpackArray(obj["regions"])
		if len(regions) > 0 {
			for _, v := range regions {
				port.Regions = append(port.Regions, fmt.Sprint(v))
			}
		}

		unlocs := unpackArray(obj["unlocs"])
		if len(unlocs) > 0 {
			for _, v := range unlocs {
				port.Unlocs = append(port.Unlocs, fmt.Sprint(v))
			}
		}

		ports = append(ports, port)
	}

	return ports, len(jsonMap), nil
}

var floatType = reflect.TypeOf(float64(0))

// unpack an unknown interface into array
func unpackArray(s any) []any {
	v := reflect.ValueOf(s)
	r := make([]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v.Index(i).Interface()
	}
	return r
}

// get float form an unknown interface
func getFloat(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}
