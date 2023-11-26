package ali

import (
	"dilu/modules/ai/service/dto"
	"testing"
)

const token = "sk-bda12fce76524d98ac5912d934ea9ce6"

func TestQwen(t *testing.T) {
	msg := []dto.Message{
		{
			Role:    "user",
			Content: "龙井茶园的茶和其他的龙井茶有什么区别",
		},
	}
	params := Parameters{
		// RepetitionPenalty: 1.0,
		ResultFormat: "message",
		// EnableSearch:      true,
		// MaxTokens:         100,
		// TopP:              0.9,
		// Seed:              0,
		// TopK:              1,
		// IncrementalOutput: true,
	}
	resp, err := Qwen(token, QwenMax, msg, params)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp)
	}
}
