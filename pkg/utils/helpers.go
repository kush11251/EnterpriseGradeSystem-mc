package utils

import (
	"encoding/json"
	"fmt"
)

// JsonResponse generates a JSON response
(func JsonResponse(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
})
