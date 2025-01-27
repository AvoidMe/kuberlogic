// Code generated by go-swagger; DO NOT EDIT.

package cmd

import (
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	errors "github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"

	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/app"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/config"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/models"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/restapi"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/restapi/operations"

	apiBackup "github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/restapi/operations/backup"

	apiRestore "github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/restapi/operations/restore"

	apiService "github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/generated/restapi/operations/service"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/logging"
	apiserverMiddleware "github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/pkg/net/middleware"
	"github.com/kuberlogic/kuberlogic/modules/dynamic-apiserver/util/k8s"
	cloudlinuxv1alpha1 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/api/v1alpha1"
	sentry2 "github.com/kuberlogic/kuberlogic/modules/dynamic-operator/sentry"
	"k8s.io/client-go/kubernetes"
	k8scheme "k8s.io/client-go/kubernetes/scheme"
)

var (
	// version of package, substitute via ldflags
	ver string
)

func Main(args []string) {
	mainLog := logging.WithComponentLogger("main")
	cfg, err := config.InitConfig("kuberlogic", logging.WithComponentLogger("config"))
	if err != nil {
		mainLog.Fatalw("", "error", err)
		os.Exit(1)
	}

	// init sentry
	if dsn := cfg.SentryDsn; dsn != "" {
		sentryTags := &sentry2.SentryTags{
			Component:    "apiserver",
			Version:      ver,
			DeploymentId: cfg.DeploymentId,
		}
		logging.UseSentry(dsn, sentryTags)

		err := sentry2.InitSentry(dsn, sentryTags)
		if err != nil {
			mainLog.Errorw("Sentry initialization failed", "error", err)
		}

		// Flush buffered events before the program terminates.
		defer sentry.Flush(2 * time.Second)
	}

	logging.DebugLevel(cfg.DebugLogs)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		mainLog.Fatalw("swagger does not loaded", "error", err)
	}

	err = cloudlinuxv1alpha1.AddToScheme(k8scheme.Scheme)
	if err != nil {
		mainLog.Fatalw("could not add to scheme", "error", err)
	}

	k8sconf, err := k8s.GetConfig(cfg)
	if err != nil {
		mainLog.Fatalw("could not get config", "error", err)
	}

	crdClient, err := k8s.GetKuberLogicClient(k8sconf)
	if err != nil {
		mainLog.Fatalw("could not generate rest client", "error", err)
	}

	baseClient, err := kubernetes.NewForConfig(k8sconf)
	if err != nil {
		mainLog.Fatalw("could not get base client", "error", err)
	}

	srv := app.New(cfg, baseClient, crdClient, logging.WithComponentLogger("server"))
	api := operations.NewKuberlogicAPI(swaggerSpec)
	// Applies when the "x-token" header is set
	api.KeyAuth = func(token string) (*models.Principal, error) {
		if token == os.Getenv("KUBERLOGIC_APISERVER_TOKEN") {
			prin := models.Principal(token)
			return &prin, nil
		}
		api.Logger("Access attempt with incorrect api key auth: %s", token)
		return nil, errors.New(401, "incorrect api key auth")
	}

	api.BackupBackupAddHandler = apiBackup.BackupAddHandlerFunc(srv.BackupAddHandler)
	api.BackupBackupDeleteHandler = apiBackup.BackupDeleteHandlerFunc(srv.BackupDeleteHandler)
	api.BackupBackupListHandler = apiBackup.BackupListHandlerFunc(srv.BackupListHandler)
	api.RestoreRestoreAddHandler = apiRestore.RestoreAddHandlerFunc(srv.RestoreAddHandler)
	api.RestoreRestoreDeleteHandler = apiRestore.RestoreDeleteHandlerFunc(srv.RestoreDeleteHandler)
	api.RestoreRestoreListHandler = apiRestore.RestoreListHandlerFunc(srv.RestoreListHandler)
	api.ServiceServiceAddHandler = apiService.ServiceAddHandlerFunc(srv.ServiceAddHandler)
	api.ServiceServiceCredentialsUpdateHandler = apiService.ServiceCredentialsUpdateHandlerFunc(srv.ServiceCredentialsUpdateHandler)
	api.ServiceServiceDeleteHandler = apiService.ServiceDeleteHandlerFunc(srv.ServiceDeleteHandler)
	api.ServiceServiceEditHandler = apiService.ServiceEditHandlerFunc(srv.ServiceEditHandler)
	api.ServiceServiceGetHandler = apiService.ServiceGetHandlerFunc(srv.ServiceGetHandler)
	api.ServiceServiceListHandler = apiService.ServiceListHandlerFunc(srv.ServiceListHandler)
	//api.BearerAuth = srv.BearerAuthentication
	api.Logger = logging.WithComponentLogger("api").Infof
	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "KuberLogic API"
	parser.LongDescription = "This is a KuberLogic API"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			mainLog.Fatalw("could not add group", "error", err)
		}
	}

	if _, err := parser.ParseArgs(args); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	h := api.Serve(nil)
	r := chi.NewRouter()
	r.Use(apiserverMiddleware.NewLoggingMiddleware)
	r.Use(middleware.Recoverer)
	r.Use(apiserverMiddleware.NewCorsMiddleware(cfg))
	r.Use(apiserverMiddleware.SentryLogPanic)
	r.Use(apiserverMiddleware.SetSentryRequestScope)

	r.Mount("/", h)

	server.ConfigureAPI()
	server.SetHandler(r)

	server.Port = cfg.HTTPBindPort
	server.Host = cfg.BindHost
	if err := server.Serve(); err != nil {
		mainLog.Fatalw("problem with serve server", "error", err)
	}
}
