package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_CreateShortURL(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		requestBody  string
		mockFunc     func(m *mocks.Client)
		expectedCode int
		expectedBody string
	}{
		{
			name:         `Invalid JSON`,
			method:       http.MethodPost,
			requestBody:  `invalid-json`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid body request",
		},
		{
			name:         "wrong HTTP Method",
			method:       http.MethodGet,
			requestBody:  `{"OriginalURL":"https://www.youtube.com/"}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Method not allowed",
		},
		{
			name:         "Empty URL",
			method:       http.MethodPost,
			requestBody:  `{"OriginalURL":""}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "URL cannot be empty",
		},
		{
			name:        "Create Failure - db error",
			method:      http.MethodPost,
			requestBody: `{"OriginalURL":"https://www.youtube.com/"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("CreatePostgresRow", "dba51bcc", "https://www.youtube.com/").Return(errors.New("db error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Failed to create row: db error",
		},
		{
			name:        "Create Success",
			method:      http.MethodPost,
			requestBody: `{"OriginalURL":"https://www.youtube.com/"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("CreatePostgresRow", "dba51bcc", "https://www.youtube.com/").Return(nil).Times(1)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"ShortURL":"http://localhost:8080/dba51bcc","OriginalURL":"https://www.youtube.com/"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			handler := &Http{client: mockClient}

			req := httptest.NewRequest(tt.method, "/create", bytes.NewBuffer([]byte(tt.requestBody)))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			handler.CreateShortURL(rec, req)

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}

func Test_DeleteURL(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		requestBody  string
		mockFunc     func(m *mocks.Client)
		expectedCode int
		expectedBody string
	}{
		{
			name:         `Invalid JSON`,
			method:       http.MethodDelete,
			requestBody:  `invalid-json`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid json body",
		},
		{
			name:         "Wrong HTTP Method",
			method:       http.MethodGet,
			requestBody:  `{"ShortURL":"http://localhost:8080/dba51bcc"}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Method not allowed",
		},
		{
			name:         "Empty Short URL",
			method:       http.MethodDelete,
			requestBody:  `{"ShortURL":""}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "URL cannot be empty", //Short URL cannot be empty
		},
		{
			name:        "Delete Failure - db error",
			method:      http.MethodDelete,
			requestBody: `{"ShortURL":"http://localhost:8080/dba51bcc"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("DeletePostgresRow", "dba51bcc").Return(errors.New("db error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Failed to delete the row: db error\n",
		},
		{
			name:        "Delete Success",
			method:      http.MethodDelete,
			requestBody: `{"ShortURL":"http://localhost:8080/dba51bcc"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("DeletePostgresRow", "dba51bcc").Return(nil).Times(1)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"message":"Short URL 'dba51bcc' deleted successfully"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			handler := &Http{client: mockClient}

			req := httptest.NewRequest(tt.method, "/create", bytes.NewBuffer([]byte(tt.requestBody)))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			handler.DeleteShortUrl(rec, req)

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}

func Test_GetOriginalURL(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		queryParameter string
		mockFunc       func(m *mocks.Client)
		expectedCode   int
		expectedBody   string
	}{
		{
			name:           "Wrong HTTP Method",
			method:         http.MethodDelete,
			queryParameter: "short=dba51bcc",
			mockFunc:       func(m *mocks.Client) {},
			expectedCode:   http.StatusMethodNotAllowed,
			expectedBody:   "Method not allowed",
		},
		{
			name:           "Missing Query Paramter",
			method:         http.MethodGet,
			queryParameter: "",
			mockFunc:       func(m *mocks.Client) {},
			expectedCode:   http.StatusBadRequest,
			expectedBody:   "Short url cannot be empty", //Short URL cannot be empty
		},
		{
			name:           "Get Failure - db error",
			method:         http.MethodGet,
			queryParameter: "short=dba51bcc",
			mockFunc: func(m *mocks.Client) {
				m.On("GetPostgresRow", "dba51bcc").Return("", errors.New("db error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Failed to get the row: db error\n",
		},
		{
			name:           "Get Success",
			method:         http.MethodGet,
			queryParameter: "short=dba51bcc",
			mockFunc: func(m *mocks.Client) {
				m.On("GetPostgresRow", "dba51bcc").Return("https://www.youtube.com/", nil).Times(1)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"Original":"https://www.youtube.com/","ShortCode":"dba51bcc"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			handler := &Http{client: mockClient}

			req := httptest.NewRequest(tt.method, "/get?"+tt.queryParameter, nil)

			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			handler.GetOriginal(rec, req)

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}
