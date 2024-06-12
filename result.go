package zubao

import (
	"encoding/json"
	"strconv"
)

type Result struct {
	Code int    `json:"result"` // 1: success, 0: failed
	Msg  string `json:"msg"`    // error message
}

func (r *Result) UnmarshalJSON(data []byte) error {
	type Alias Result
	aux := &struct {
		Result string `json:"result"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	result, err := strconv.Atoi(aux.Result)
	if err != nil {
		return err
	}
	r.Code = result
	return nil
}
