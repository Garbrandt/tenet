package servers

func (s *Server) initializeRoutes() {
	s.Router.NoRoute(Interpreter)

	v1 := s.Router.Group("")
	{
		v1.GET("/dashboard/*param", Interpreter)
		v1.GET("/generate/*param", Interpreter)
	}

	api := s.Router.Group("/api")
	{
		api.POST("/user/sign_in", s.SignIn)
		api.GET("/user/refresh", s.Refresh)

		api.GET("/dashboard", AnalyticsMiddleware(), s.GetDashboard)

		api.GET("/content/:section/:env/:id/attachments", s.GetContentAttachments)
		api.POST("/contents", s.CreateContent)
		api.DELETE("/contents/:id", s.DeleteContent)

		api.POST("/attachments/:section/:env", s.UploadAttachments)
		api.DELETE("/attachments/", s.DeleteAttachments)

		api.GET("/options/:section/:env", s.GetAllOptions)
	}
}
