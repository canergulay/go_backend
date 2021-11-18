package groups_api

type GetGroupsRequest struct {
	Locale string `json:"locale"`
	Name   string `json:"name"`
}
