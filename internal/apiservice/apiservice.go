package apiservice

import (
	"github.com/ssoready/ssoready/internal/gen/ssoready/v1/ssoreadyv1connect"
	"github.com/ssoready/ssoready/internal/google"
	"github.com/ssoready/ssoready/internal/store"
)

type Service struct {
	Store        *store.Store
	GoogleClient *google.Client
	ssoreadyv1connect.UnimplementedSSOReadyServiceHandler
}