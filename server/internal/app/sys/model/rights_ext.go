package model

type RightsTree struct {
	Rights
	Children []*RightsTree `json:"children"`
}
