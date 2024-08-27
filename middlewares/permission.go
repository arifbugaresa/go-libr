package middlewares

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go-libr/utils/common"
	"go-libr/utils/logger"
	"go-libr/utils/redis"
)

func Permission(permission map[string]string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessGranted bool
		for accessCode, grantPermission := range permission {
			allowedPermissions := GetPermission(ctx, accessCode)
			for _, allowedPermission := range allowedPermissions {
				if allowedPermission == grantPermission {
					accessGranted = true
					break
				}
			}
		}

		if !accessGranted {
			common.GenerateErrorResponse(ctx, "you do not have access to this menu")
			return
		}

		ctx.Next()
	}
}

const (
	PermissionCreate = "c"
	PermissionRead   = "r"
	PermissionUpdate = "u"
	PermissionDelete = "d"
)

func GetPermission(ctx *gin.Context, accessCode string) (allowedPermission []string) {
	var (
		isExistAccessCode bool
		accessGrant       string
		mappingPermission = map[string][]string{
			PermissionDelete: {PermissionRead, PermissionCreate, PermissionUpdate, PermissionDelete},
			PermissionUpdate: {PermissionRead, PermissionCreate, PermissionUpdate},
			PermissionCreate: {PermissionRead, PermissionCreate},
			PermissionRead:   {PermissionRead},
		}
	)

	jwtToken, err := GetJwtTokenFromHeader(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, "failed get jwt from header")
		return
	}

	redisSessionStr, err := redis.RedisClient.Get(ctx, jwtToken).Result()
	if err != nil {
		err = errors.New("redis error")
		logger.ErrorWithCtx(ctx, nil, err)
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	var redisSession RedisSession
	err = json.Unmarshal([]byte(redisSessionStr), &redisSession)
	if err != nil {
		common.GenerateErrorResponse(ctx, "failed unmarshal redis session")
		return
	}

	for _, permission := range redisSession.Permission {
		if permission.AccessCode == accessCode {
			isExistAccessCode = true
			accessGrant = permission.AccessGrant
			break
		}
	}

	if !isExistAccessCode {
		common.GenerateErrorResponse(ctx, "you do not have access to this menu")
		return
	}

	allowedPermission = mappingPermission[accessGrant]

	return
}
