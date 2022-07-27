package xutil

import (
	"fmt"
	"encoding/json"
)

func ConvertToStringArray[T any](src []T) []string {
	res := make([]string, 0, len(src))
	for i := range src {
		res = append(res, fmt.Sprintf("%v", src[i]))
	}

	return res
}

func ConvertByJson(src any, dest any) error {
	marshal, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(marshal, dest)
	if err != nil {
		return err
	}

	return nil
}
