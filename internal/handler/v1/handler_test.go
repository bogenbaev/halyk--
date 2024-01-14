package v1

//func TestHandler_Init(t *testing.T) {
//	type fields struct {
//		services *service.Services
//		keycloak keycloak.IKeycloak
//	}
//	type args struct {
//		api *gin.RouterGroup
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		{
//			name: "Success.",
//			fields: fields{
//				services: &service.Services{},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// Init Dependencies
//			c := gomock.NewController(t)
//			defer c.Finish()
//			_, engine := gin.CreateTestContext(nil)
//
//			router := engine.Group("/api")
//
//			infoS := servicemocks.NewMockInformationService(c)
//
//			h := &Handler{
//				services: &service.Services{InformationService: infoS},
//				keycloak: implementation.Keycloak("", ""),
//			}
//
//			h.Init(router)
//		})
//	}
//}
