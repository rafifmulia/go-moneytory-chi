package router

import (
	"encoding/json"
	"net/http"
	"net/http/pprof"
	"restfulapi/conf"
	"restfulapi/handler"
	mid "restfulapi/middleware"

	"github.com/go-chi/chi/v5"
	chimid "github.com/go-chi/chi/v5/middleware"
)

const (
	basePath string = "/v1"
)

var (
	httpProfileFlag bool
	trx             handler.TransactionHandler
)

func setFlags() {
	httpProfileFlag = conf.GetHttpProfileFlag()
}

func initHandler() {
	trx = handler.NewTransactionHandler()
}

func panicRoute(w http.ResponseWriter, r *http.Request) {
	panic("Panic Route")
}

func notFoundRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(404)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"message": "404 Not Found",
	})
}

func ChiInitRouter() *chi.Mux {
	setFlags()
	initHandler()
	r := chi.NewMux()
	r = mid.ChiRootMiddleware(r)
	r.NotFound(notFoundRoute)
	r.Get("/panic", panicRoute)
	if httpProfileFlag {
		r.Mount("/debug", chimid.Profiler())
	}
	r.Group(func(r chi.Router) {
		r.Use(mid.AuthMiddleware)
		r.Get(basePath+"/transaction", trx.ListTransaction)
		r.Post(basePath+"/transaction", trx.CreateTransaction)
		r.Get(basePath+"/transaction/{trxId}", trx.GetTransaction)
		r.Put(basePath+"/transaction/{trxId}", trx.UpdateTransaction)
		r.Delete(basePath+"/transaction/{trxId}", trx.DeleteTransaction)
	})
	return r
}

func StdInitRouter() http.Handler {
	setFlags()
	initHandler()
	mux := http.NewServeMux()
	mux.Handle("/", mid.StdRootMiddleware(http.HandlerFunc(notFoundRoute)))
	mux.Handle("/panic", mid.StdRootMiddleware(http.HandlerFunc(panicRoute)))
	if httpProfileFlag {
		mux.Handle("/debug/pprof/", mid.StdRootMiddleware(http.HandlerFunc(pprof.Index)))
		mux.Handle("/debug/pprof/cmdline", mid.StdRootMiddleware(http.HandlerFunc(pprof.Cmdline)))
		mux.Handle("/debug/pprof/profile", mid.StdRootMiddleware(http.HandlerFunc(pprof.Profile)))
		mux.Handle("/debug/pprof/symbol", mid.StdRootMiddleware(http.HandlerFunc(pprof.Symbol)))
		mux.Handle("/debug/pprof/trace", mid.StdRootMiddleware(http.HandlerFunc(pprof.Trace)))
	}
	// Belum ada implementasi method GET, POST, PUT, DELETE.
	mux.Handle(basePath+"/transaction", mid.StdRootMiddleware(mid.AuthMiddleware(http.HandlerFunc(trx.ListTransaction))))
	// Belum ada implementasi mengambil path variable.
	mux.Handle(basePath+"/transaction/{trxId}", mid.StdRootMiddleware(mid.AuthMiddleware(http.HandlerFunc(trx.ListTransaction))))
	return mux
}
