package c

import (
	"encoding/json"
	"fmt"
	"log"
)

type Rawer interface {
	SetRawData(rawData []byte)
	GetRawData() []byte
}

type RawDataContainer struct {
	rawData []byte `json:"-"`
}

func (r *RawDataContainer) SetRawData(rawData []byte) {
	if rawData == nil {
		log.Println("RawDataContainer.SetRawData set nil data")
		r.rawData = []byte("")
	}
	r.rawData = rawData
}
func (r *RawDataContainer) GetRawData() []byte {
	return r.rawData
}

func SetRawDataIntoSlice(rawData []byte, targets []Rawer, fromMapKeys ...string) error {
	var slices []interface{}
	if len(fromMapKeys) == 0 {
		err := json.Unmarshal(rawData, &slices)
		if err != nil {
			return err
		}
	} else {
		var mp map[string]interface{}
		err := json.Unmarshal(rawData, &mp)
		if err != nil {
			return err
		}
		length := len(fromMapKeys)
		for i, key := range fromMapKeys {
			value, ok := mp[key]
			if !ok {
				return fmt.Errorf("SetRawDataIntoSlice: key `%s` doesn't exist in rawData: %s", key, rawData)
			}
			if i < length-1 {
				subMp, ok := value.(map[string]interface{})
				if !ok {
					return fmt.Errorf("SetRawDataIntoSlice: expect value of key `%s` to be map in rawData: `%s`", key, rawData)
				}
				mp = subMp
			} else {
				s, ok := value.([]interface{})
				if !ok {
					return fmt.Errorf("SetRawDataIntoSlice: expect final value to be array in rawData: `%s`", rawData)
				}
				slices = s
			}
		}
	}

	if len(slices) != len(targets) {
		return fmt.Errorf("SetRawDataIntoSlice -> dismatched length between rawData slices and targets: %d <-> %d", len(slices), len(targets))
	}
	for i, s := range slices {
		chunk, err := json.Marshal(s)
		if err != nil {
			return fmt.Errorf("SetRawDataIntoSlice -> Marshal slice error: %v", err)
		}
		targets[i].SetRawData(chunk)
	}
	return nil
}

func SplitRawDataIntoMap(rawData []byte) (map[string][]byte, error) {
	var mp map[string]interface{}
	err := json.Unmarshal(rawData, &mp)
	if err != nil {
		return nil, err
	}
	output := map[string][]byte{}
	for key, s := range mp {
		chunk, err := json.Marshal(s)
		if err != nil {
			return nil, fmt.Errorf("SetRawDataIntoMap -> Marshal map error: %v", err)
		}
		output[key] = chunk
	}
	return output, nil
}
