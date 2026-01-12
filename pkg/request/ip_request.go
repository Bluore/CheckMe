package request

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPLocationResult struct {
	IP       string
	Location string
	Err      error
}

func QueryIP(ip string) (map[string]interface{}, error) {
	url := fmt.Sprintf("http://ip-api.com/json?ip=%s", ip)

	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 将JSON解析为map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetIPLocation(ip string) (string, error) {
	data, err := QueryIP(ip)
	if err != nil {
		return "归属地查询失败", err
	}
	res := fmt.Sprintf("%s %s %s", data["country"], data["regionName"], data["city"])

	return res, err
}

func GetIPLocationWithCtx(ctx context.Context, ip string) <-chan IPLocationResult {
	ch := make(chan IPLocationResult, 1)

	go func() {
		defer close(ch)

		done := make(chan struct{})
		var loc string
		var err error

		go func() {
			loc, err = GetIPLocation(ip)
			close(done)
		}()

		select {
		case <-ctx.Done():
			ch <- IPLocationResult{IP: ip, Location: "", Err: ctx.Err()}
		case <-done:
			ch <- IPLocationResult{IP: ip, Location: loc, Err: err}
		}
	}()

	return ch
}
