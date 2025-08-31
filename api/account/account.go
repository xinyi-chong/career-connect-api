// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	"gf_demo/api/account/v1"
)

type IAccountV1 interface {
	PatchUpdateEmailMe(ctx context.Context, req *v1.PatchUpdateEmailMeReq) (res *v1.PatchUpdateEmailMeRes, err error)
	PatchUpdatePasswordMe(ctx context.Context, req *v1.PatchUpdatePasswordMeReq) (res *v1.PatchUpdatePasswordMeRes, err error)
}
