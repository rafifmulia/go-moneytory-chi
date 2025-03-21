package middleware

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restfulapi/conf"
	"restfulapi/exception"
	"restfulapi/helper"
	"runtime/debug"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
)

var (
	debugMode bool
)

func setFlags() {
	debugMode = conf.GetDebugFlag()
}

type wInjector struct {
	writtenHeader, writtenBody bool
	statusCode                 int
	http.ResponseWriter
}

func (w *wInjector) Write(b []byte) (int, error) {
	w.writtenBody = true
	return w.ResponseWriter.Write(b)
}

func (w *wInjector) WriteHeader(statusCode int) {
	w.writtenHeader = true
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Catch panic in handlers or middlewares and return http error response.
// Log incoming requests and its status code.
func panicHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		iw := &wInjector{
			statusCode:     200,
			ResponseWriter: w,
		}
		log.Printf("Incoming request from %s %s\n", r.Method, r.URL.Path)
		defer catchPanic(iw, func() {
			log.Printf("Request from %s %s has been responded %d\n", r.Method, r.URL.Path, iw.statusCode)
		})
		next.ServeHTTP(iw, r)
	}
	return http.HandlerFunc(fn)
}

func catchPanic(w http.ResponseWriter, end func()) {
	msg := recover()
	if msg != nil {
		val := reflect.ValueOf(msg)
		tp := val.Type()
		switch err := msg.(type) {
		case *exception.BadRequestException:
			helper.RespBadRequest(w, getErrMessage(err))
		case *form.InvalidDecoderError:
			helper.RespBadRequest(w, getErrMessage(err))
		case form.DecodeErrors:
			helper.RespBadRequest(w, getErrMessage(err))
		case *exception.UnauthorizedException:
			helper.RespUnauthorized(w, getErrMessage(err))
		case *exception.NotFoundException:
			helper.RespNotFound(w, getErrMessage(err))
		case *exception.UnprocessableEntityException:
			helper.RespUnprocessableEntity(w, getErrMessage(err))
		case *validator.InvalidValidationError:
			helper.RespUnprocessableEntity(w, getErrMessage(err))
		case validator.ValidationErrors:
			helper.RespUnprocessableEntity(w, getErrMessage(err))
		case validator.FieldError:
			helper.RespUnprocessableEntity(w, getErrMessage(err))
		case error:
			helper.RespInternalServerError(w, getErrMessage(err))
		default:
			log.Printf("panicHandler type:%s\npanicHandler val:%s\npanicHandler stack trace:%s\n", tp.Name(), msg, debug.Stack())
			if debugMode {
				helper.RespInternalServerError(w, fmt.Sprintf("%s", msg))
			} else {
				helper.RespInternalServerError(w, "")
			}
		}
	}
	end()
}

func getErrMessage(err error) string {
	if debugMode {
		return err.Error()
	}
	return ""
}

func ChiRootMiddleware(r *chi.Mux) *chi.Mux {
	setFlags()
	r.Use(panicHandler)
	return r
}

func StdRootMiddleware(next http.Handler) http.Handler {
	setFlags()
	return panicHandler(next)
}
