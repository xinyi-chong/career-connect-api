package permission

import (
	"context"
	v1 "gf_demo/api/permission/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

func (s *sPermission) GetPermissionByID(ctx context.Context, permissionID string) (*entity.Permission, error) {
	var permission *entity.Permission
	err := dao.Permission.Ctx(ctx).With(
		entity.Feature{},
	).Where(do.Permission{
		Id: permissionID,
	}).Scan(&permission)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Permission By ID: ", permissionID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Permission By ID: " + err.Error())
	}
		
	return permission, err
}

func (s *sPermission) GetPermissionsByRoleID(ctx context.Context, roleID string) ([]*entity.Permission, error) {
	cachePermissions := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_PERMISSIONS_BY_ROLE_ID + roleID, &[]*entity.Permission{})
	if cachePermissions != nil {
		return *cachePermissions.(*[]*entity.Permission), nil
	}

	var permissions []*entity.Permission
	err := dao.Permission.Ctx(ctx).With(
		entity.Feature{},
	).Where(do.Permission{
		RoleId: roleID,
	}).Scan(&permissions)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Permissions By Role ID: ", roleID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Permissions By Role ID: " + err.Error())
		return nil, err
	}
	
	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_PERMISSIONS_BY_ROLE_ID + roleID, permissions)

	return permissions, nil
}

func (s *sPermission) GetPermissionsByFeatureID(ctx context.Context, featureID string) ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	err := dao.Permission.Ctx(ctx).Where(do.Permission{
		FeatureId: featureID,
	}).Scan(&permissions)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Permissions By Feature ID: ", featureID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Permissions By Feature ID: " + err.Error())
	}
		
	return permissions, err
}

func (s *sPermission) GetPermissionByRoleIDFeatureID(ctx context.Context, roleID string, featureID string) (*entity.Permission, error) {
	var permission *entity.Permission
	err := dao.Permission.Ctx(ctx).Where(do.Permission{
		RoleId: roleID,
		FeatureId: featureID,
	}).Scan(&permission)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Permission By Role ID & Feature ID: ", featureID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Permission By Role ID & By Feature ID: " + err.Error())
	}
		
	return permission, err
}

func (s *sPermission) GetCompanyPermissionsByUserID(ctx context.Context, userID string) ([]*model.Permission, error) {
	permissions := []*model.Permission{}
	companyaccounts, err := service.Company().GetCompanyAccountsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// If user has been assigned any role
	for _, companyaccount := range companyaccounts {
		rolePermissions, err := s.GetPermissionsByRoleID(ctx, companyaccount.RoleId)
		if err != nil {
			return nil, err
		}

		var featureIDs []string
		for _, rolePermission := range rolePermissions {
			if rolePermission.Allow > 0 {
				featureIDs = append(featureIDs, rolePermission.FeatureId)
			}
		}
		permission := model.Permission{
			CompanyID:  companyaccount.CompanyId,
			RoleID:     companyaccount.RoleId,
			FeatureIDs: featureIDs,
		}
		permissions = append(permissions, &permission)
	}
	return permissions, nil
}

func (s *sPermission) PostCreatePermission(ctx context.Context, req *v1.PostCreatePermissionReq) (*string, error) {
	id := uuid.New().String()
	_, err := dao.Permission.Ctx(ctx).Data(do.Permission{
		Id: id,
		RoleId: req.RoleID,
		FeatureId: req.FeatureID,
		Allow: req.Allow,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Permission: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Permission: " + err.Error())
		return nil, err
	}

	// Remove User Session By Role ID
	service.Session().RemoveUserSessionsByRoleID(ctx, req.RoleID)

	// Remove Caches
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_PERMISSIONS_BY_ROLE_ID + req.RoleID)

	// Remove CompanyAccounts Caches By Role ID
	companyaccounts, _ := service.Company().GetCompanyAccountsByRoleID(ctx, req.RoleID)
	service.Company().RemoveCompanyAccountsCaches(ctx, companyaccounts)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Permission: ", id)
	return &id, nil
}

func (s *sPermission) PatchUpdatePermissionByRoleIDFeatureID(ctx context.Context, req *v1.PatchUpdatePermissionByRoleIDFeatureIDReq, roleID string, featureID string) (error) {
	_, err := dao.Permission.Ctx(ctx).Data(do.Permission{
		Allow: req.Allow,
	}).Where(do.Permission{
		RoleId: roleID,
		FeatureId: featureID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Permission By Role ID: ", roleID, "& Feature ID: ", featureID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Permission By Role ID & Feature ID: " + err.Error())
		return err
	}
	
	// Remove User Session By Role ID
	service.Session().RemoveUserSessionsByRoleID(ctx, roleID)

	// Remove Caches
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_PERMISSIONS_BY_ROLE_ID + roleID)

	// Remove CompanyAccounts Caches By Role ID
	companyaccounts, _ := service.Company().GetCompanyAccountsByRoleID(ctx, roleID)
	service.Company().RemoveCompanyAccountsCaches(ctx, companyaccounts)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Permission By Role ID: ", roleID, "& Feature ID: ", featureID)
	return nil
}

func (s *sPermission) DeletePermissionByID(ctx context.Context, permissionID string) (error) {
	permission, err := s.GetPermissionByID(ctx, permissionID)
	if err != nil {
		return err
	}

	_, err = dao.Permission.Ctx(ctx).Where(do.Permission{
		Id: permissionID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Permission By ID: ", permissionID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Permission By ID: " + err.Error())
		return err
	}

	// Remove User Session By Role ID
	service.Session().RemoveUserSessionsByRoleID(ctx, permission.RoleId)

	// Remove Caches
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_PERMISSIONS_BY_ROLE_ID + permission.RoleId)

	// Remove CompanyAccounts Caches By Role ID
	companyaccounts, _ := service.Company().GetCompanyAccountsByRoleID(ctx, permission.RoleId)
	service.Company().RemoveCompanyAccountsCaches(ctx, companyaccounts)
		
	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Permission By ID: ", permissionID)
	return nil
}