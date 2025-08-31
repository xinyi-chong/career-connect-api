package role

import (
	"context"

	v1 "gf_demo/api/role/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) GetRoleByID(ctx context.Context, roleID string) (*entity.Role, error) {
	var role *entity.Role
	err := dao.Role.Ctx(ctx).With(
		entity.Permission{},
		entity.Permission{}.Feature,
	).Where(do.Role{
		Id: roleID,
	}).Scan(&role)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Role By ID:", roleID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Role By ID: " + err.Error())
	}

	return role, err
}

func (s *sRole) PostCreateRole(ctx context.Context, req *v1.PostCreateRoleReq) (*string, error) {
	id := uuid.New().String()
	_, err := dao.Role.Ctx(ctx).Data(do.Role{
		Id: id,
		Name: req.Name,
		Status: consts.ACTIVE,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Role", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Role: " + err.Error())
	}

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "Role: ", id)
	return &id, err
}

func (s *sRole) PatchUpdateRoleByID(ctx context.Context, req *v1.PatchUpdateRoleByIDReq, roleID string) (error) {
	_, err := dao.Role.Ctx(ctx).Data(do.Role{
		Name: req.Name,
		Status: req.Status,
	}).Where(do.Role{
		Id: roleID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Role By ID", roleID, err) 
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Role By ID: " + err.Error())
		return err
	}

	// Remove Caches
	companyaccounts, _ := service.Company().GetCompanyAccountsByRoleID(ctx, roleID)
	service.Company().RemoveCompanyAccountsCaches(ctx, companyaccounts)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Role By ID", roleID) 
	return nil
}

func (s *sRole) DeleteRoleByID(ctx context.Context, roleID string) (error) {
	err := gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		if err := dao.Role.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err := dao.Role.Ctx(ctx).Where(do.Role{
				Id: roleID,
			}).Delete()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Delete Role By ID", roleID, err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Role By ID: " + err.Error()) 
			return err
		}

		// Remove User Account Session that are Associated with the Deleted Role
		companyaccounts, err := service.Company().GetCompanyAccountsByRoleID(ctx, roleID)
		if err != nil {
			return err
		}
		err = service.Session().RemoveCompanyAccountsSessions(ctx, companyaccounts)
		if err != nil {
			return err
		}
		
		// Remove Caches
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_PERMISSIONS_BY_ROLE_ID + roleID)
		service.Company().RemoveCompanyAccountsCaches(ctx, companyaccounts)

		g.Log().Info(ctx, consts.SUCCESS_DELETE, "Role By ID", roleID)
		return nil
	})
	
	return err
}