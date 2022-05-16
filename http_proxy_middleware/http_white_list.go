package http_proxy_middleware

//func HTTPWhiteListMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		serverInterface, ok := c.Get("service")
//		if !ok {
//			middleware.ResponseError(c, 2001, errors.New("service not found"))
//			c.Abort()
//			return
//		}
//		serviceDetail := serverInterface.(*dao.ServiceDetail)
//
//		iplist := []string{}
//		if serviceDetail.AccessControl.WhiteList != "" {
//			if !public.InStringSlice(iplist, c.ClientIP()) {
//				middleware.ResponseError(c, 3001, errors.New(fmt.Sprintf("%s not in white ip list", c.ClientIP())))
//				c.Abort()
//				return
//			}
//		}
//		c.Next()
//	}
//}
