package v1

//func TestHandler_loadBalancer(t *testing.T) {
//	type mockBehavior func(r *servicemocks.MockDocumentService, document dto.Document)
//
//	updated := time.Now()
//	created := updated.Add(-time.Hour)
//	isTemplate := true
//	userID := uuid.New().String()
//	createdFileUUID := uuid.New()
//
//	tests := []struct {
//		name                 string
//		inputDocument        dto.Document
//		mockBehavior         mockBehavior
//		expectedStatusCode   int
//		expectedResponseBody string
//	}{
//		{
//			name: "Failed. Database. Duplicate Key",
//			inputDocument: dto.Document{
//				ID:     123,
//				TreeID: 1,
//				Name:   "This is a template document",
//			},
//			mockBehavior: func(r *servicemocks.MockDocumentService, document dto.Document) {
//				r.EXPECT().
//					Get(gomock.Any(), gomock.Any(), gomock.Any()).
//					Return(dto.Document{
//						ID:     123,
//						TreeID: 1,
//						Name:   "This is a template document",
//					}, gorm.ErrDuplicatedKey)
//			},
//			expectedStatusCode:   500,
//			expectedResponseBody: `{"reason":"duplicated key not allowed"}`,
//		},
//		{
//			name: "Failed. Database. Invalid Value",
//			inputDocument: dto.Document{
//				ID:     123,
//				TreeID: 1,
//				Name:   "This is a template document",
//			},
//			mockBehavior: func(r *servicemocks.MockDocumentService, document dto.Document) {
//				r.EXPECT().
//					Get(gomock.Any(), gomock.Any(), gomock.Any()).
//					Return(dto.Document{
//						ID:     123,
//						TreeID: 1,
//						Name:   "This is a template document",
//					}, gorm.ErrInvalidValue)
//			},
//			expectedStatusCode:   500,
//			expectedResponseBody: `{"reason":"invalid value, should be pointer to struct or slice"}`,
//		},
//		{
//			name: "Success. Update document",
//			inputDocument: dto.Document{
//				ID:        123,
//				UserID:    userID,
//				TreeID:    1,
//				UpdatedAt: updated,
//				Name:      "This is a template document",
//				Path:      createdFileUUID,
//				Template:  &isTemplate,
//			},
//			mockBehavior: func(r *servicemocks.MockDocumentService, document dto.Document) {
//				r.EXPECT().
//					Get(gomock.Any(), gomock.Any(), gomock.Any()).
//					Return(dto.Document{
//						ID:        123,
//						UserID:    userID,
//						TreeID:    1,
//						CreatedAt: created,
//						UpdatedAt: updated,
//						Name:      "This is a template document",
//						Path:      createdFileUUID,
//						Template:  &isTemplate,
//					}, nil)
//			},
//			expectedStatusCode: 200,
//			expectedResponseBody: fmt.Sprintf(`{"id":123,"createdAt":"%s","updatedAt":"%s","name":"This is a template document","path":"%s","template":true}`,
//				created.Format(time.RFC3339Nano),
//				updated.Format(time.RFC3339Nano),
//				createdFileUUID.String()),
//		},
//	}
//	for _, tt := range tests {
//
//		t.Run(tt.name, func(t *testing.T) {
//			// Init Dependencies
//			c := gomock.NewController(t)
//			defer c.Finish()
//
//			repo := servicemocks.NewMockDocumentService(c)
//			tt.mockBehavior(repo, tt.inputDocument)
//
//			services := &service.Services{DocumentService: repo}
//			handler := Handler{services, nil, nil}
//
//			// Init Endpoint
//			r := gin.New()
//			r.GET("/api/v1/tree/:treeID/document/:docID", handler.readDocument)
//
//			// Create Request
//			w := httptest.NewRecorder()
//			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/tree/%d/document/%d", 1, 1), nil)
//
//			// Make Request
//			r.ServeHTTP(w, req)
//
//			// Assert
//			assert.Equal(t, tt.expectedStatusCode, w.Code)
//			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
//		})
//	}
//}
