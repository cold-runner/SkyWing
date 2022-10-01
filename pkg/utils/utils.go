package utils

//func GetClaims(c *gin.Context) (*models.RoleCharacter, error) {
//	token := c.Request.Header.Get("Authorization")
//	// 按.分割
//	parts := strings.SplitN(token, ".", 3)
//	if !(len(parts) == 3) {
//		zap.L().Error("无效的Token")
//		response.ResponseError(c, response.CodeInvalidToken)
//		c.Abort()
//		return nil, nil
//	}
//	claims, err := jwt.ParseToken(token)
//	if err != nil {
//		zap.L().Error("token解析失败", zap.Error(err))
//		response.ResponseError(c, response.CodeInvalidToken)
//		c.Abort()
//		return nil, nil
//	}
//	return claims, err
//}

//// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
//func GetUserID(c *gin.Context) uint {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return 0
//		} else {
//			return cl.ID
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.ID
//	}
//}
//
//// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
//func GetUserUuid(c *gin.Context) uuid.UUID {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return uuid.UUID{}
//		} else {
//			return cl.UUID
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.UUID
//	}
//}
//
//// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
//func GetUserAuthorityId(c *gin.Context) uint {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return 0
//		} else {
//			return cl.AuthorityId
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.AuthorityId
//	}
//}
//
//// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
//func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return nil
//		} else {
//			return cl
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse
//	}
//}
