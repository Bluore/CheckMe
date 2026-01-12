package change

import (
	"checkme/internal/dto"
	"checkme/pkg/request"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/datatypes"
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

	// 异步查询IP属地
	ch := request.GetIPLocationWithCtx(ctx, "162.141.131.115")

	for key, val := range conf {
		data[key] = val
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
	case <-time.After(1500 * time.Millisecond):
		// IP属地查询超时
		fmt.Println(fmt.Sprintf("IP 查询超时: %s", ip))
	}

	res, err := MapToJSON(data)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func MapToJSON(m map[string]interface{}) (datatypes.JSON, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return datatypes.JSON(b), nil
}
