package rex

import (
	eos "github.com/shiqinfeng1/eos-go"
)

func NewConsolidate(owner eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: REXAN,
		Name:    ActN("consolidate"),
		Authorization: []eos.PermissionLevel{
			{Actor: owner, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(Consolidate{
			Owner: owner,
		}),
	}
}

type Consolidate struct {
	Owner eos.AccountName
}
