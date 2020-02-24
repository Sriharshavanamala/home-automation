// Code generated by jrpc. DO NOT EDIT.

package handler

import (
	"net/http"

	"github.com/jakewright/home-automation/libraries/go/errors"
	"github.com/jakewright/home-automation/libraries/go/request"
	"github.com/jakewright/home-automation/libraries/go/response"
	"github.com/jakewright/home-automation/libraries/go/router"
	"github.com/jakewright/home-automation/libraries/go/slog"
	def "github.com/jakewright/home-automation/service.scene/def"
)

// SceneRouter wraps router.Router to provide a convenient way to set handlers
type SceneRouter struct {
	*router.Router
	CreateScene func(*def.CreateSceneRequest) (*def.CreateSceneResponse, error)
	ReadScene   func(*def.ReadSceneRequest) (*def.ReadSceneResponse, error)
	ListScenes  func(*def.ListScenesRequest) (*def.ListScenesResponse, error)
	DeleteScene func(*def.DeleteSceneRequest) (*def.DeleteSceneResponse, error)
	SetScene    func(*def.SetSceneRequest) (*def.SetSceneResponse, error)
}

// NewRouter returns a router that is ready to add handlers to
func NewRouter() *SceneRouter {
	rr := &SceneRouter{
		Router: router.New(),
	}

	rr.Router.Handle("POST", "/scenes", func(w http.ResponseWriter, r *http.Request) {
		if rr.CreateScene == nil {
			slog.Panicf("No handler exists for POST /scenes")
		}

		body := &def.CreateSceneRequest{
			Request: r,
			Context: r.Context(),
		}
		if err := request.Decode(r, body); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		rsp, err := rr.CreateScene(body)
		if err != nil {
			err = errors.WithMessage(err, "failed to handle request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		response.WriteJSON(w, rsp)
	})

	rr.Router.Handle("GET", "/scene", func(w http.ResponseWriter, r *http.Request) {
		if rr.ReadScene == nil {
			slog.Panicf("No handler exists for GET /scene")
		}

		body := &def.ReadSceneRequest{
			Request: r,
			Context: r.Context(),
		}
		if err := request.Decode(r, body); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		rsp, err := rr.ReadScene(body)
		if err != nil {
			err = errors.WithMessage(err, "failed to handle request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		response.WriteJSON(w, rsp)
	})

	rr.Router.Handle("GET", "/scenes", func(w http.ResponseWriter, r *http.Request) {
		if rr.ListScenes == nil {
			slog.Panicf("No handler exists for GET /scenes")
		}

		body := &def.ListScenesRequest{
			Request: r,
			Context: r.Context(),
		}
		if err := request.Decode(r, body); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		rsp, err := rr.ListScenes(body)
		if err != nil {
			err = errors.WithMessage(err, "failed to handle request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		response.WriteJSON(w, rsp)
	})

	rr.Router.Handle("DELETE", "/scene", func(w http.ResponseWriter, r *http.Request) {
		if rr.DeleteScene == nil {
			slog.Panicf("No handler exists for DELETE /scene")
		}

		body := &def.DeleteSceneRequest{
			Request: r,
			Context: r.Context(),
		}
		if err := request.Decode(r, body); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		rsp, err := rr.DeleteScene(body)
		if err != nil {
			err = errors.WithMessage(err, "failed to handle request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		response.WriteJSON(w, rsp)
	})

	rr.Router.Handle("POST", "/scene/set", func(w http.ResponseWriter, r *http.Request) {
		if rr.SetScene == nil {
			slog.Panicf("No handler exists for POST /scene/set")
		}

		body := &def.SetSceneRequest{
			Request: r,
			Context: r.Context(),
		}
		if err := request.Decode(r, body); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = errors.Wrap(err, errors.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		rsp, err := rr.SetScene(body)
		if err != nil {
			err = errors.WithMessage(err, "failed to handle request")
			slog.Error(err)
			response.WriteJSON(w, err)
			return
		}

		response.WriteJSON(w, rsp)
	})

	return rr
}