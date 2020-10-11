package handler

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Search_Error(t *testing.T) {
	cm := &ctxMock{}
	tm := &treeMock{}
	val := 123
	cm.On("Get", "params").Return(Params{Val: &val}).Once()
	cm.On("JSON", http.StatusOK, false).Return(errors.New("something wrong")).Once()
	tm.On("Search", 123).Return(false).Once()

	h := NewHandler(tm)
	assert.NotNil(t, h.Search(cm))

	cm.AssertExpectations(t)
	tm.AssertExpectations(t)
}

func TestHandler_Search_OK(t *testing.T) {
	cm := &ctxMock{}
	tm := &treeMock{}
	val := 123
	cm.On("Get", "params").Return(Params{Val: &val}).Once()
	cm.On("JSON", http.StatusOK, false).Return(nil).Once()
	tm.On("Search", 123).Return(false).Once()

	h := NewHandler(tm)
	assert.Nil(t, h.Search(cm))

	cm.AssertExpectations(t)
	tm.AssertExpectations(t)
}

func TestHandler_Insert_Error(t *testing.T) {
	cm := &ctxMock{}
	tm := &treeMock{}
	val := 123
	cm.On("Get", "params").Return(Params{Val: &val}).Once()
	cm.On("JSON", http.StatusOK, false).Return(errors.New("something wrong")).Once()
	tm.On("Insert", 123).Return(false).Once()

	h := NewHandler(tm)
	assert.NotNil(t, h.Insert(cm))

	cm.AssertExpectations(t)
	tm.AssertExpectations(t)
}

func TestHandler_Insert_OK(t *testing.T) {
	cm := &ctxMock{}
	tm := &treeMock{}
	val := 123
	cm.On("Get", "params").Return(Params{Val: &val}).Once()
	cm.On("JSON", http.StatusOK, false).Return(nil).Once()
	tm.On("Insert", 123).Return(false).Once()

	h := NewHandler(tm)
	assert.Nil(t, h.Insert(cm))

	cm.AssertExpectations(t)
	tm.AssertExpectations(t)
}

func TestHandler_Delete_Error(t *testing.T) {
	cm := &ctxMock{}
	tm := &treeMock{}
	val := 123
	cm.On("Get", "params").Return(Params{Val: &val}).Once()
	cm.On("JSON", http.StatusOK, false).Return(errors.New("something wrong")).Once()
	tm.On("Delete", 123).Return(false).Once()

	h := NewHandler(tm)
	assert.NotNil(t, h.Delete(cm))

	cm.AssertExpectations(t)
	tm.AssertExpectations(t)
}

func TestHandler_Delete_OK(t *testing.T) {
	cm := &ctxMock{}
	tm := &treeMock{}
	val := 123
	cm.On("Get", "params").Return(Params{Val: &val}).Once()
	cm.On("JSON", http.StatusOK, false).Return(nil).Once()
	tm.On("Delete", 123).Return(false).Once()

	h := NewHandler(tm)
	assert.Nil(t, h.Delete(cm))

	cm.AssertExpectations(t)
	tm.AssertExpectations(t)
}

type treeMock struct {
	mock.Mock
}

func (m *treeMock) Search(value int) (found bool) {
	args := m.Called(value)
	return args.Bool(0)
}

func (m *treeMock) Insert(value int) (found bool) {
	args := m.Called(value)
	return args.Bool(0)
}

func (m *treeMock) Delete(value int) (found bool) {
	args := m.Called(value)
	return args.Bool(0)
}

type ctxMock struct {
	mock.Mock
}

func (m *ctxMock) Get(key string) interface{} {
	args := m.Called(key)
	return args.Get(0)
}

func (m *ctxMock) JSON(code int, i interface{}) error {
	args := m.Called(code, i)
	return args.Error(0)
}

func (m *ctxMock) Request() *http.Request {
	panic("implement me")
}

func (m *ctxMock) SetRequest(r *http.Request) {
	panic("implement me")
}

func (m *ctxMock) SetResponse(r *echo.Response) {
	panic("implement me")
}

func (m *ctxMock) Response() *echo.Response {
	panic("implement me")
}

func (m *ctxMock) IsTLS() bool {
	panic("implement me")
}

func (m *ctxMock) IsWebSocket() bool {
	panic("implement me")
}

func (m *ctxMock) Scheme() string {
	panic("implement me")
}

func (m *ctxMock) RealIP() string {
	panic("implement me")
}

func (m *ctxMock) Path() string {
	panic("implement me")
}

func (m *ctxMock) SetPath(p string) {
	panic("implement me")
}

func (m *ctxMock) Param(name string) string {
	panic("implement me")
}

func (m *ctxMock) ParamNames() []string {
	panic("implement me")
}

func (m *ctxMock) SetParamNames(names ...string) {
	panic("implement me")
}

func (m *ctxMock) ParamValues() []string {
	panic("implement me")
}

func (m *ctxMock) SetParamValues(values ...string) {
	panic("implement me")
}

func (m *ctxMock) QueryParam(name string) string {
	panic("implement me")
}

func (m *ctxMock) QueryParams() url.Values {
	panic("implement me")
}

func (m *ctxMock) QueryString() string {
	panic("implement me")
}

func (m *ctxMock) FormValue(name string) string {
	panic("implement me")
}

func (m *ctxMock) FormParams() (url.Values, error) {
	panic("implement me")
}

func (m *ctxMock) FormFile(name string) (*multipart.FileHeader, error) {
	panic("implement me")
}

func (m *ctxMock) MultipartForm() (*multipart.Form, error) {
	panic("implement me")
}

func (m *ctxMock) Cookie(name string) (*http.Cookie, error) {
	panic("implement me")
}

func (m *ctxMock) SetCookie(cookie *http.Cookie) {
	panic("implement me")
}

func (m *ctxMock) Cookies() []*http.Cookie {
	panic("implement me")
}

func (m *ctxMock) Set(key string, val interface{}) {
	panic("implement me")
}

func (m *ctxMock) Bind(i interface{}) error {
	panic("implement me")
}

func (m *ctxMock) Validate(i interface{}) error {
	panic("implement me")
}

func (m *ctxMock) Render(code int, name string, data interface{}) error {
	panic("implement me")
}

func (m *ctxMock) HTML(code int, html string) error {
	panic("implement me")
}

func (m *ctxMock) HTMLBlob(code int, b []byte) error {
	panic("implement me")
}

func (m *ctxMock) String(code int, s string) error {
	panic("implement me")
}

func (m *ctxMock) JSONPretty(code int, i interface{}, indent string) error {
	panic("implement me")
}

func (m *ctxMock) JSONBlob(code int, b []byte) error {
	panic("implement me")
}

func (m *ctxMock) JSONP(code int, callback string, i interface{}) error {
	panic("implement me")
}

func (m *ctxMock) JSONPBlob(code int, callback string, b []byte) error {
	panic("implement me")
}

func (m *ctxMock) XML(code int, i interface{}) error {
	panic("implement me")
}

func (m *ctxMock) XMLPretty(code int, i interface{}, indent string) error {
	panic("implement me")
}

func (m *ctxMock) XMLBlob(code int, b []byte) error {
	panic("implement me")
}

func (m *ctxMock) Blob(code int, contentType string, b []byte) error {
	panic("implement me")
}

func (m *ctxMock) Stream(code int, contentType string, r io.Reader) error {
	panic("implement me")
}

func (m *ctxMock) File(file string) error {
	panic("implement me")
}

func (m *ctxMock) Attachment(file string, name string) error {
	panic("implement me")
}

func (m *ctxMock) Inline(file string, name string) error {
	panic("implement me")
}

func (m *ctxMock) NoContent(code int) error {
	panic("implement me")
}

func (m *ctxMock) Redirect(code int, url string) error {
	panic("implement me")
}

func (m *ctxMock) Error(err error) {
	panic("implement me")
}

func (m *ctxMock) Handler() echo.HandlerFunc {
	panic("implement me")
}

func (m *ctxMock) SetHandler(h echo.HandlerFunc) {
	panic("implement me")
}

func (m *ctxMock) Logger() echo.Logger {
	panic("implement me")
}

func (m *ctxMock) SetLogger(l echo.Logger) {
	panic("implement me")
}

func (m *ctxMock) Echo() *echo.Echo {
	panic("implement me")
}

func (m *ctxMock) Reset(r *http.Request, w http.ResponseWriter) {
	panic("implement me")
}
