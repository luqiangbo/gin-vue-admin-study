package response

import (
	"go-class/model/system/tables"
)

type SysAuthorityResponse struct {
	Authority tables.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      tables.SysAuthority `json:"authority"`
	OldAuthorityId string              `json:"old_authority_id"`
}
