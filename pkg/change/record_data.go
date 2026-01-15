package change

import (
	"checkme/internal/dto"
	"checkme/pkg/request"
	mytime "checkme/pkg/time"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	time "time"

	"gorm.io/datatypes"
)

var (
	DataDisable = errors.New("disable 该应用")
)

func ChangeData(ctx context.Context, cov *map[string]map[string]interface{}, req *dto.UploadRecordRequest, ip string) (datatypes.JSON, error) {

	conf, ok := (*cov)[req.Application]
	if !ok {
		return req.Data, nil
	}

	var data map[string]interface{}
	err := json.Unmarshal(req.Data, &data)
	if err != nil {
		return nil, err
	}

	if data["is_disable"] == true {
		return nil, DataDisable
	}

	// 异步查询IP属地
	ch := request.GetIPLocationWithCtx(ctx, ip)

	for key, val := range conf {
		switch key {
		case "description":
			data[key] = descGet(cov, val)
		default:
			data[key] = val
		}
	}

	if val, ok := data["is_hide_music"]; ok || val == true {
		delete(data, "music_name")
	}

	select {
	case res := <-ch:
		if res.Err == nil {
			// 成功查询IP属地
			data["location"] = res.Location
		}
	case <-time.After(3000 * time.Millisecond):
		// IP属地查询超时
		fmt.Println(fmt.Sprintf("IP 查询超时: %s", ip))
	}

	res, err := MapToJSON(data)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func descGet(cov *map[string]map[string]interface{}, desc interface{}) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Printf("Type: %T", desc)
	switch data := desc.(type) {
	case string:
		return data
	case []interface{}:
		choose := r.Intn(len(data))
		return descGet(cov, data[choose])
	case map[interface{}]interface{}:
		// 日期、权重模式

		var sum int = 0
		var beforeData interface{}
		var afterData interface{}
		var startTime time.Time
		var endTime time.Time

		now := time.Now()

		for key, val := range data {
			switch key := key.(type) {
			case int:
				sum += key
			case string:
				switch key {
				case "<-":
					beforeData = val
				case "->":
					afterData = val
				default:
					tr, err := mytime.NewTimeRange(key)
					if err != nil {
						return "ERR: 配置文件时间格式错误"
					}
					if tr.Contains(now) {
						return descGet(cov, val)
					} else {
						if HHMMToMM(tr.Start) < HHMMToMM(startTime) {
							startTime = tr.Start
						}
						if HHMMToMM(tr.End) > HHMMToMM(endTime) {
							endTime = tr.End
						}
					}
				}
			}
		}

		if beforeData != nil && HHMMToMM(now) < HHMMToMM(startTime) {
			return descGet(cov, beforeData)
		} else if afterData != nil && HHMMToMM(now) > HHMMToMM(endTime) {
			return descGet(cov, afterData)
		}

		choose := rand.Intn(sum)
		for w, d := range data {
			if _, ok := w.(int); !ok {
				continue
			}

			choose -= w.(int)

			if choose < 0 {
				return descGet(cov, d)
			}
		}

		return "ERR: 配置文件错误"
	case map[string]interface{}:
		var res string
		for key, val := range data {
			if _, ok := val.(string); !ok {
				res += descGet(cov, val)
				continue
			}
			switch key {
			case "copy_by":
				if _, ok := (*cov)[val.(string)]; !ok {
					return "ERR: 配置文件出现不存在的映射"
				}
				res += descGet(cov, (*cov)[val.(string)]["description"])
			default:
				res += val.(string)
			}
		}
		return res
	default:
		return "ERR: 配置文件错误"
	}
}

func checkStrTime(str string) (bool, error) {
	tr, err := mytime.NewTimeRange(str)
	if err != nil {
		return false, err
	}
	now := time.Now()

	return tr.Contains(now), nil
}

func HHMMToMM(t time.Time) int {
	return t.Hour()*60 + t.Minute()
}

func MapToJSON(m map[string]interface{}) (datatypes.JSON, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return datatypes.JSON(b), nil
}
