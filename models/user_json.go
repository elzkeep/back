package models

import "encoding/json"

// 수동 작성 파일 (model.json 생성 대상 아님)
// 비밀번호는 어떤 JSON 응답에도 포함하지 않음

func (c User) MarshalJSON() ([]byte, error) {
	type alias User

	item := alias(c)
	item.Passwd = ""

	return json.Marshal(item)
}

func (c Userlist) MarshalJSON() ([]byte, error) {
	type alias Userlist

	item := alias(c)
	item.Passwd = ""

	return json.Marshal(item)
}
