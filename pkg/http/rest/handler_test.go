package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"drexel.edu/voter-api/pkg/process"
	"drexel.edu/voter-api/pkg/retrieve"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

var testHandler *fiber.App

func init() {

	processService := process.NewService(&process.MockRepository{})
	retrievalService := retrieve.NewService(&retrieve.MockRepository{})

	router := Handler(3000, processService, retrievalService)

	testHandler = router
}

func TestHealth(t *testing.T) {
	r := httptest.NewRequest("GET", "/voters/health", nil)
	resp, _ := testHandler.Test(r, -1)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetAllVoters(t *testing.T) {
	r := httptest.NewRequest("GET", "/voters", nil)
	resp, _ := testHandler.Test(r, -1)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetVoterById(t *testing.T) {
	r := httptest.NewRequest("GET", "/voters/1", nil)
	resp, _ := testHandler.Test(r, -1)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestPostToCreateVoterById(t *testing.T) {

	requestBody := []byte(`{"Name": "Miguel","Email": "mad32@drexel.edu"}`)

	ctx := &fasthttp.RequestCtx{}

	ctx.Request.SetRequestURI("/voters/1")

	ctx.Request.Header.SetMethod("POST")

	ctx.Request.Header.SetContentType("application/json")

	ctx.Request.SetBody(requestBody)

	testHandler.Handler()(ctx)

	assert.Equal(t, http.StatusCreated, ctx.Response.StatusCode())
}

func TestGetAllHistoryForId(t *testing.T) {
	r := httptest.NewRequest("GET", "/voters/1/polls", nil)
	resp, _ := testHandler.Test(r, -1)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetSinglePollById(t *testing.T) {
	r := httptest.NewRequest("GET", "/voters/1/polls/1", nil)
	resp, _ := testHandler.Test(r, -1)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestPostToCreateSinglePollById(t *testing.T) {
	requestBody := []byte(`{"vote_date":"2024-02-22T06:00:47.948774-05:00"}`)

	ctx := &fasthttp.RequestCtx{}

	ctx.Request.SetRequestURI("/voters/1/polls/1")

	ctx.Request.Header.SetMethod("POST")

	ctx.Request.Header.SetContentType("application/json")

	ctx.Request.SetBody(requestBody)

	testHandler.Handler()(ctx)

	assert.Equal(t, http.StatusCreated, ctx.Response.StatusCode())
}

func TestUpdateVoterById(t *testing.T) {
	requestBody := []byte(`{"Name": "Miguel","Email": "mad32@drexel.edu"}`)

	ctx := &fasthttp.RequestCtx{}

	ctx.Request.SetRequestURI("/voters/1")

	ctx.Request.Header.SetMethod("PUT")

	ctx.Request.Header.SetContentType("application/json")

	ctx.Request.SetBody(requestBody)

	testHandler.Handler()(ctx)

	assert.Equal(t, http.StatusOK, ctx.Response.StatusCode())
}

func TestUpdateSinglePollById(t *testing.T) {
	requestBody := []byte(`{"vote_date":"2024-02-22T06:00:47.948774-05:00"}`)

	ctx := &fasthttp.RequestCtx{}

	ctx.Request.SetRequestURI("/voters/1/polls/1")

	ctx.Request.Header.SetMethod("PUT")

	ctx.Request.Header.SetContentType("application/json")

	ctx.Request.SetBody(requestBody)

	testHandler.Handler()(ctx)

	assert.Equal(t, http.StatusOK, ctx.Response.StatusCode())
}

func TestDeleteVoterById(t *testing.T) {
	r := httptest.NewRequest("DELETE", "/voters/1", nil)
	resp, _ := testHandler.Test(r, -1)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteSinglePollById(t *testing.T) {
	r := httptest.NewRequest("DELETE", "/voters/1/polls/1", nil)
	resp, _ := testHandler.Test(r, -1)
	assert.Equal(t, 200, resp.StatusCode)
}
