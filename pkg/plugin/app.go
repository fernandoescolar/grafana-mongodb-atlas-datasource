package plugin

import (
	"context"
	"net/http"

	dserrors "github.com/fernandoescolar/grafana-mongodb-atlas-datasource/pkg/errors"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
)

// Make sure App implements required interfaces. This is important to do
// since otherwise we will only get a not implemented error response from plugin in
// runtime. Plugin should not implement all these interfaces - only those which are
// required for a particular task.
var (
	_ backend.CallResourceHandler   = (*App)(nil)
	_ backend.QueryDataHandler      = (*App)(nil)
	_ instancemgmt.InstanceDisposer = (*App)(nil)
	_ backend.CheckHealthHandler    = (*App)(nil)
)

// App is an example app plugin with a backend which can respond to data queries.
type App struct {
	backend.CallResourceHandler
	im instancemgmt.InstanceManager
}

// NewApp creates a new example *App instance.
func NewApp(_ context.Context, _ backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	app := &App{
		im: datasource.NewInstanceManager(newDataSourceInstance),
	}

	app.CallResourceHandler = httpadapter.New(app)

	return app, nil
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created.
func (a *App) Dispose() {
	//
}

// QueryData handles multiple queries and returns multiple responses.
func (a *App) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	h, err := a.im.Get(ctx, req.PluginContext)
	if err != nil {
		return nil, err
	}

	if val, ok := h.(*Instance); ok {
		return HandleQueryData(ctx, val, req)
	}
	return nil, dserrors.ErrorBadDatasource
}

// CheckHealth handles health checks sent from Grafana to the plugin.
func (a *App) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	h, err := a.im.Get(ctx, req.PluginContext)
	if err != nil {
		return nil, err
	}

	if val, ok := h.(*Instance); ok {
		return CheckHealth(ctx, val, req)
	}

	return &backend.CheckHealthResult{
		Status:      backend.HealthStatusError,
		Message:     dserrors.ErrorBadDatasource.Error(),
		JSONDetails: nil,
	}, dserrors.ErrorBadDatasource
}

// ServeHTTP is the main HTTP handler for serving resource calls
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pluginCtx := backend.PluginConfigFromContext(r.Context())

	h, err := a.im.Get(r.Context(), pluginCtx)
	if err != nil {
		panic(err)
	}

	if ds, ok := h.(*Instance); ok {
		GetRouter(ds.Handlers).ServeHTTP(w, r)
		return
	}

	panic(dserrors.ErrorBadDatasource)
}
