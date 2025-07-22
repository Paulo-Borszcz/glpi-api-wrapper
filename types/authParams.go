package types

type AuthParams struct {
	Type   string `json:"token_type"`
	Expire string `json:"expires_in"`
	Token  string `json:"access_token"`
}
