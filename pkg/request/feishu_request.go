package request

import (
	"checkme/config"
	"checkme/internal/dto"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

var (
	ErrorNoType = errors.New("不存在的推送类型")
)

func NotifyOnFeishu(baseURL string, msg string) (datatypes.JSON, error) {
	req, err := http.NewRequest("POST", baseURL, strings.NewReader(msg))
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func NotifyDataOfStudy(cfg *config.Config, c *gin.Context, msg string) (datatypes.JSON, error) {
	body := `
{
  "msg_type": "interactive",
  "card": {
    "header": {
      "title": {
        "tag": "plain_text",
        "content": "提醒你去学习了"
      }
    },
    "elements": [
      {
        "tag": "div",
		"collapsible": true,
        "text": {
          "tag": "lark_md",
          "content": "**Sender IP**: *%s* \n**Sender UA**: *%s*"
        }
      },
      {
        "tag": "action",
        "actions": [
          {
            "tag": "button",
            "text": {
              "tag": "plain_text",
              "content": "查看详情"
            },
            "type": "primary",
            "url": "https://me.bluore.top"
          }
        ]
      }
    ]
  }
}
`
	body = fmt.Sprintf(body, c.ClientIP(), c.GetHeader("User-Agent"))

	resp, err := NotifyOnFeishu(cfg.Notify.FeishuBot, body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NotifyDataOfGame(cfg *config.Config, c *gin.Context, msg string) (datatypes.JSON, error) {
	body := `
{
  "msg_type": "interactive",
  "card": {
    "header": {
      "title": {
        "tag": "plain_text",
        "content": "一起上号"
      }
    },
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "lark_md",
          "content": "%s"
        }
      },
      {
        "tag": "div",
		"collapsible": true,
        "text": {
          "tag": "lark_md",
          "content": "**Sender IP**: *%s* \n**Sender UA**: *%s*"
        }
      },
      {
        "tag": "action",
        "actions": [
          {
            "tag": "button",
            "text": {
              "tag": "plain_text",
              "content": "查看详情"
            },
            "type": "primary",
            "url": "https://me.bluore.top"
          }
        ]
      }
    ]
  }
}
`
	body = fmt.Sprintf(body, msg, c.ClientIP(), c.GetHeader("User-Agent"))

	resp, err := NotifyOnFeishu(cfg.Notify.FeishuBot, body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NotifyDataOfEncourage(cfg *config.Config, c *gin.Context, msg string) (datatypes.JSON, error) {
	body := `
{
  "msg_type": "interactive",
  "card": {
    "header": {
      "title": {
        "tag": "plain_text",
        "content": "加油哦!!!"
      }
    },
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "lark_md",
          "content": "%s"
        }
      },
      {
        "tag": "div",
		"collapsible": true,
        "text": {
          "tag": "lark_md",
          "content": "**Sender IP**: *%s* \n**Sender UA**: *%s*"
        }
      },
      {
        "tag": "action",
        "actions": [
          {
            "tag": "button",
            "text": {
              "tag": "plain_text",
              "content": "查看详情"
            },
            "type": "primary",
            "url": "https://me.bluore.top"
          }
        ]
      }
    ]
  }
}
`
	body = fmt.Sprintf(body, msg, c.ClientIP(), c.GetHeader("User-Agent"))

	resp, err := NotifyOnFeishu(cfg.Notify.FeishuBot, body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NotifyDataOfDropLine(cfg *config.Config, c *gin.Context, msg string) (datatypes.JSON, error) {
	body := `
{
  "msg_type": "interactive",
  "card": {
    "header": {
      "title": {
        "tag": "plain_text",
        "content": "%s"
      }
    },
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "lark_md",
          "content": "稍一句话~"
        }
      },
      {
        "tag": "div",
		"collapsible": true,
        "text": {
          "tag": "lark_md",
          "content": "**Sender IP**: *%s* \n**Sender UA**: *%s*"
        }
      },
      {
        "tag": "action",
        "actions": [
          {
            "tag": "button",
            "text": {
              "tag": "plain_text",
              "content": "查看详情"
            },
            "type": "primary",
            "url": "https://me.bluore.top"
          }
        ]
      }
    ]
  }
}
`
	body = fmt.Sprintf(body, msg, c.ClientIP(), c.GetHeader("User-Agent"))

	resp, err := NotifyOnFeishu(cfg.Notify.FeishuBot, body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NotifyDataOfTapTap(cfg *config.Config, c *gin.Context, msg string) (datatypes.JSON, error) {
	body := `
{
  "msg_type": "interactive",
  "card": {
    "header": {
      "title": {
        "tag": "plain_text",
        "content": "戳了戳你~"
      }
    },
    "elements": [
      {
        "tag": "div",
        "text": {
          "tag": "lark_md",
          "content": "%s"
        }
      },
      {
        "tag": "div",
		"collapsible": true,
        "text": {
          "tag": "lark_md",
          "content": "**Sender IP**: *%s* \n**Sender UA**: *%s*"
        }
      },
      {
        "tag": "action",
        "actions": [
          {
            "tag": "button",
            "text": {
              "tag": "plain_text",
              "content": "查看详情"
            },
            "type": "primary",
            "url": "https://me.bluore.top"
          }
        ]
      }
    ]
  }
}
`
	body = fmt.Sprintf(body, msg, c.ClientIP(), c.GetHeader("User-Agent"))

	resp, err := NotifyOnFeishu(cfg.Notify.FeishuBot, body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NotifyFeishu(cfg *config.Config, c *gin.Context, data dto.CreateNotifyRequest) error {
	switch data.Type {
	case "study":
		if _, err := NotifyDataOfStudy(cfg, c, data.Msg); err != nil {
			return err
		}
	case "game":
		if _, err := NotifyDataOfGame(cfg, c, data.Msg); err != nil {
			return err
		}
	case "encourage":
		if _, err := NotifyDataOfEncourage(cfg, c, data.Msg); err != nil {
			return err
		}
	case "DropLine":
		if _, err := NotifyDataOfDropLine(cfg, c, data.Msg); err != nil {
			return err
		}
	case "TapTap":
		if _, err := NotifyDataOfTapTap(cfg, c, data.Msg); err != nil {
			return err
		}
	default:
		return ErrorNoType
	}
	return nil
}
