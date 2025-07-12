package response_yh

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK) // 强制全部为200 避免客户端太复杂

	errorResponse := map[string]any{
		"error":   true,
		"code":    statusCode,
		"message": message,
		"data":    nil,
	}

	json.NewEncoder(w).Encode(errorResponse)
}

// response_yh.SendJSON 发送JSON响应
func SendJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	dataNew := map[string]any{
		"error":   false,
		"code":    1,
		"data":    data,
		"message": "ok",
	}

	if err := json.NewEncoder(w).Encode(dataNew); err != nil {
		// 如果编码失败，发送错误响应
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
