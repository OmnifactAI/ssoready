// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: queries.sql

package queries

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const authGetInitData = `-- name: AuthGetInitData :one
select idp_redirect_url, sp_entity_id
from saml_connections
where saml_connections.id = $1
`

type AuthGetInitDataRow struct {
	IdpRedirectUrl *string
	SpEntityID     *string
}

func (q *Queries) AuthGetInitData(ctx context.Context, id uuid.UUID) (AuthGetInitDataRow, error) {
	row := q.db.QueryRow(ctx, authGetInitData, id)
	var i AuthGetInitDataRow
	err := row.Scan(&i.IdpRedirectUrl, &i.SpEntityID)
	return i, err
}

const authGetSAMLFlow = `-- name: AuthGetSAMLFlow :one
select id, saml_connection_id, access_code, state, create_time, expire_time, subject_idp_id, subject_idp_attributes, update_time, auth_redirect_url, get_redirect_time, initiate_request, initiate_time, assertion, app_redirect_url, receive_assertion_time, redeem_time, redeem_response
from saml_flows
where id = $1
`

func (q *Queries) AuthGetSAMLFlow(ctx context.Context, id uuid.UUID) (SamlFlow, error) {
	row := q.db.QueryRow(ctx, authGetSAMLFlow, id)
	var i SamlFlow
	err := row.Scan(
		&i.ID,
		&i.SamlConnectionID,
		&i.AccessCode,
		&i.State,
		&i.CreateTime,
		&i.ExpireTime,
		&i.SubjectIdpID,
		&i.SubjectIdpAttributes,
		&i.UpdateTime,
		&i.AuthRedirectUrl,
		&i.GetRedirectTime,
		&i.InitiateRequest,
		&i.InitiateTime,
		&i.Assertion,
		&i.AppRedirectUrl,
		&i.ReceiveAssertionTime,
		&i.RedeemTime,
		&i.RedeemResponse,
	)
	return i, err
}

const authGetValidateData = `-- name: AuthGetValidateData :one
select saml_connections.sp_entity_id,
       saml_connections.idp_entity_id,
       saml_connections.idp_x509_certificate,
       environments.redirect_url
from saml_connections
         join organizations on saml_connections.organization_id = organizations.id
         join environments on organizations.environment_id = environments.id
where saml_connections.id = $1
`

type AuthGetValidateDataRow struct {
	SpEntityID         *string
	IdpEntityID        *string
	IdpX509Certificate []byte
	RedirectUrl        *string
}

func (q *Queries) AuthGetValidateData(ctx context.Context, id uuid.UUID) (AuthGetValidateDataRow, error) {
	row := q.db.QueryRow(ctx, authGetValidateData, id)
	var i AuthGetValidateDataRow
	err := row.Scan(
		&i.SpEntityID,
		&i.IdpEntityID,
		&i.IdpX509Certificate,
		&i.RedirectUrl,
	)
	return i, err
}

const createAppOrganization = `-- name: CreateAppOrganization :one
insert into app_organizations (id, google_hosted_domain)
values ($1, $2)
returning id, google_hosted_domain
`

type CreateAppOrganizationParams struct {
	ID                 uuid.UUID
	GoogleHostedDomain *string
}

func (q *Queries) CreateAppOrganization(ctx context.Context, arg CreateAppOrganizationParams) (AppOrganization, error) {
	row := q.db.QueryRow(ctx, createAppOrganization, arg.ID, arg.GoogleHostedDomain)
	var i AppOrganization
	err := row.Scan(&i.ID, &i.GoogleHostedDomain)
	return i, err
}

const createAppSession = `-- name: CreateAppSession :one
insert into app_sessions (id, app_user_id, create_time, expire_time, token)
values ($1, $2, $3, $4, $5)
returning id, app_user_id, create_time, expire_time, token
`

type CreateAppSessionParams struct {
	ID         uuid.UUID
	AppUserID  uuid.UUID
	CreateTime time.Time
	ExpireTime time.Time
	Token      string
}

func (q *Queries) CreateAppSession(ctx context.Context, arg CreateAppSessionParams) (AppSession, error) {
	row := q.db.QueryRow(ctx, createAppSession,
		arg.ID,
		arg.AppUserID,
		arg.CreateTime,
		arg.ExpireTime,
		arg.Token,
	)
	var i AppSession
	err := row.Scan(
		&i.ID,
		&i.AppUserID,
		&i.CreateTime,
		&i.ExpireTime,
		&i.Token,
	)
	return i, err
}

const createAppUser = `-- name: CreateAppUser :one
insert into app_users (id, app_organization_id, display_name, email)
values ($1, $2, $3, $4)
returning id, app_organization_id, display_name, email
`

type CreateAppUserParams struct {
	ID                uuid.UUID
	AppOrganizationID uuid.UUID
	DisplayName       string
	Email             *string
}

func (q *Queries) CreateAppUser(ctx context.Context, arg CreateAppUserParams) (AppUser, error) {
	row := q.db.QueryRow(ctx, createAppUser,
		arg.ID,
		arg.AppOrganizationID,
		arg.DisplayName,
		arg.Email,
	)
	var i AppUser
	err := row.Scan(
		&i.ID,
		&i.AppOrganizationID,
		&i.DisplayName,
		&i.Email,
	)
	return i, err
}

const createOrganization = `-- name: CreateOrganization :one
insert into organizations (id, environment_id, external_id)
values ($1, $2, $3)
returning id, environment_id, external_id
`

type CreateOrganizationParams struct {
	ID            uuid.UUID
	EnvironmentID uuid.UUID
	ExternalID    *string
}

func (q *Queries) CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error) {
	row := q.db.QueryRow(ctx, createOrganization, arg.ID, arg.EnvironmentID, arg.ExternalID)
	var i Organization
	err := row.Scan(&i.ID, &i.EnvironmentID, &i.ExternalID)
	return i, err
}

const createOrganizationDomain = `-- name: CreateOrganizationDomain :one
insert into organization_domains (id, organization_id, domain)
values ($1, $2, $3)
returning id, organization_id, domain
`

type CreateOrganizationDomainParams struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Domain         string
}

func (q *Queries) CreateOrganizationDomain(ctx context.Context, arg CreateOrganizationDomainParams) (OrganizationDomain, error) {
	row := q.db.QueryRow(ctx, createOrganizationDomain, arg.ID, arg.OrganizationID, arg.Domain)
	var i OrganizationDomain
	err := row.Scan(&i.ID, &i.OrganizationID, &i.Domain)
	return i, err
}

const createSAMLConnection = `-- name: CreateSAMLConnection :one
insert into saml_connections (id, organization_id, sp_entity_id)
values ($1, $2, $3)
returning id, organization_id, idp_redirect_url, idp_x509_certificate, idp_entity_id, sp_entity_id
`

type CreateSAMLConnectionParams struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	SpEntityID     *string
}

func (q *Queries) CreateSAMLConnection(ctx context.Context, arg CreateSAMLConnectionParams) (SamlConnection, error) {
	row := q.db.QueryRow(ctx, createSAMLConnection, arg.ID, arg.OrganizationID, arg.SpEntityID)
	var i SamlConnection
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.IdpRedirectUrl,
		&i.IdpX509Certificate,
		&i.IdpEntityID,
		&i.SpEntityID,
	)
	return i, err
}

const createSAMLFlowGetRedirect = `-- name: CreateSAMLFlowGetRedirect :one
insert into saml_flows (id, saml_connection_id, access_code, expire_time, state, create_time, update_time,
                        auth_redirect_url, get_redirect_time)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
returning id, saml_connection_id, access_code, state, create_time, expire_time, subject_idp_id, subject_idp_attributes, update_time, auth_redirect_url, get_redirect_time, initiate_request, initiate_time, assertion, app_redirect_url, receive_assertion_time, redeem_time, redeem_response
`

type CreateSAMLFlowGetRedirectParams struct {
	ID               uuid.UUID
	SamlConnectionID uuid.UUID
	AccessCode       uuid.UUID
	ExpireTime       time.Time
	State            string
	CreateTime       time.Time
	UpdateTime       time.Time
	AuthRedirectUrl  *string
	GetRedirectTime  *time.Time
}

func (q *Queries) CreateSAMLFlowGetRedirect(ctx context.Context, arg CreateSAMLFlowGetRedirectParams) (SamlFlow, error) {
	row := q.db.QueryRow(ctx, createSAMLFlowGetRedirect,
		arg.ID,
		arg.SamlConnectionID,
		arg.AccessCode,
		arg.ExpireTime,
		arg.State,
		arg.CreateTime,
		arg.UpdateTime,
		arg.AuthRedirectUrl,
		arg.GetRedirectTime,
	)
	var i SamlFlow
	err := row.Scan(
		&i.ID,
		&i.SamlConnectionID,
		&i.AccessCode,
		&i.State,
		&i.CreateTime,
		&i.ExpireTime,
		&i.SubjectIdpID,
		&i.SubjectIdpAttributes,
		&i.UpdateTime,
		&i.AuthRedirectUrl,
		&i.GetRedirectTime,
		&i.InitiateRequest,
		&i.InitiateTime,
		&i.Assertion,
		&i.AppRedirectUrl,
		&i.ReceiveAssertionTime,
		&i.RedeemTime,
		&i.RedeemResponse,
	)
	return i, err
}

const deleteOrganizationDomains = `-- name: DeleteOrganizationDomains :exec
delete
from organization_domains
where organization_id = $1
`

func (q *Queries) DeleteOrganizationDomains(ctx context.Context, organizationID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteOrganizationDomains, organizationID)
	return err
}

const getAPIKeyBySecretValue = `-- name: GetAPIKeyBySecretValue :one
select id, app_organization_id, secret_value
from api_keys
where secret_value = $1
`

func (q *Queries) GetAPIKeyBySecretValue(ctx context.Context, secretValue string) (ApiKey, error) {
	row := q.db.QueryRow(ctx, getAPIKeyBySecretValue, secretValue)
	var i ApiKey
	err := row.Scan(&i.ID, &i.AppOrganizationID, &i.SecretValue)
	return i, err
}

const getAppOrganizationByGoogleHostedDomain = `-- name: GetAppOrganizationByGoogleHostedDomain :one
select id, google_hosted_domain
from app_organizations
where google_hosted_domain = $1
`

func (q *Queries) GetAppOrganizationByGoogleHostedDomain(ctx context.Context, googleHostedDomain *string) (AppOrganization, error) {
	row := q.db.QueryRow(ctx, getAppOrganizationByGoogleHostedDomain, googleHostedDomain)
	var i AppOrganization
	err := row.Scan(&i.ID, &i.GoogleHostedDomain)
	return i, err
}

const getAppSessionByToken = `-- name: GetAppSessionByToken :one
select app_sessions.app_user_id, app_users.app_organization_id
from app_sessions
         join app_users on app_sessions.app_user_id = app_users.id
where token = $1
  and expire_time > $2
`

type GetAppSessionByTokenParams struct {
	Token      string
	ExpireTime time.Time
}

type GetAppSessionByTokenRow struct {
	AppUserID         uuid.UUID
	AppOrganizationID uuid.UUID
}

func (q *Queries) GetAppSessionByToken(ctx context.Context, arg GetAppSessionByTokenParams) (GetAppSessionByTokenRow, error) {
	row := q.db.QueryRow(ctx, getAppSessionByToken, arg.Token, arg.ExpireTime)
	var i GetAppSessionByTokenRow
	err := row.Scan(&i.AppUserID, &i.AppOrganizationID)
	return i, err
}

const getAppUserByEmail = `-- name: GetAppUserByEmail :one
select id, app_organization_id, display_name, email
from app_users
where email = $1
`

func (q *Queries) GetAppUserByEmail(ctx context.Context, email *string) (AppUser, error) {
	row := q.db.QueryRow(ctx, getAppUserByEmail, email)
	var i AppUser
	err := row.Scan(
		&i.ID,
		&i.AppOrganizationID,
		&i.DisplayName,
		&i.Email,
	)
	return i, err
}

const getAppUserByID = `-- name: GetAppUserByID :one
select id, app_organization_id, display_name, email
from app_users
where app_organization_id = $1
  and id = $2
`

type GetAppUserByIDParams struct {
	AppOrganizationID uuid.UUID
	ID                uuid.UUID
}

func (q *Queries) GetAppUserByID(ctx context.Context, arg GetAppUserByIDParams) (AppUser, error) {
	row := q.db.QueryRow(ctx, getAppUserByID, arg.AppOrganizationID, arg.ID)
	var i AppUser
	err := row.Scan(
		&i.ID,
		&i.AppOrganizationID,
		&i.DisplayName,
		&i.Email,
	)
	return i, err
}

const getEnvironment = `-- name: GetEnvironment :one
select id, redirect_url, app_organization_id, display_name, auth_url
from environments
where app_organization_id = $1
  and id = $2
`

type GetEnvironmentParams struct {
	AppOrganizationID uuid.UUID
	ID                uuid.UUID
}

func (q *Queries) GetEnvironment(ctx context.Context, arg GetEnvironmentParams) (Environment, error) {
	row := q.db.QueryRow(ctx, getEnvironment, arg.AppOrganizationID, arg.ID)
	var i Environment
	err := row.Scan(
		&i.ID,
		&i.RedirectUrl,
		&i.AppOrganizationID,
		&i.DisplayName,
		&i.AuthUrl,
	)
	return i, err
}

const getEnvironmentByID = `-- name: GetEnvironmentByID :one
select id, redirect_url, app_organization_id, display_name, auth_url
from environments
where id = $1
`

func (q *Queries) GetEnvironmentByID(ctx context.Context, id uuid.UUID) (Environment, error) {
	row := q.db.QueryRow(ctx, getEnvironmentByID, id)
	var i Environment
	err := row.Scan(
		&i.ID,
		&i.RedirectUrl,
		&i.AppOrganizationID,
		&i.DisplayName,
		&i.AuthUrl,
	)
	return i, err
}

const getOrganization = `-- name: GetOrganization :one
select organizations.id, organizations.environment_id, organizations.external_id
from organizations
         join environments on organizations.environment_id = environments.id
where environments.app_organization_id = $1
  and organizations.id = $2
`

type GetOrganizationParams struct {
	AppOrganizationID uuid.UUID
	ID                uuid.UUID
}

func (q *Queries) GetOrganization(ctx context.Context, arg GetOrganizationParams) (Organization, error) {
	row := q.db.QueryRow(ctx, getOrganization, arg.AppOrganizationID, arg.ID)
	var i Organization
	err := row.Scan(&i.ID, &i.EnvironmentID, &i.ExternalID)
	return i, err
}

const getOrganizationByID = `-- name: GetOrganizationByID :one
select id, environment_id, external_id
from organizations
where id = $1
`

func (q *Queries) GetOrganizationByID(ctx context.Context, id uuid.UUID) (Organization, error) {
	row := q.db.QueryRow(ctx, getOrganizationByID, id)
	var i Organization
	err := row.Scan(&i.ID, &i.EnvironmentID, &i.ExternalID)
	return i, err
}

const getSAMLAccessCodeData = `-- name: GetSAMLAccessCodeData :one
select saml_flows.id             as saml_flow_id,
       saml_flows.subject_idp_id,
       saml_flows.subject_idp_attributes,
       saml_flows.state,
       organizations.id          as organization_id,
       organizations.external_id as organization_external_id,
       environments.id           as environment_id
from saml_flows
         join saml_connections on saml_flows.saml_connection_id = saml_connections.id
         join organizations on saml_connections.organization_id = organizations.id
         join environments on organizations.environment_id = environments.id
where environments.app_organization_id = $1
  and saml_flows.access_code = $2
`

type GetSAMLAccessCodeDataParams struct {
	AppOrganizationID uuid.UUID
	AccessCode        uuid.UUID
}

type GetSAMLAccessCodeDataRow struct {
	SamlFlowID             uuid.UUID
	SubjectIdpID           *string
	SubjectIdpAttributes   []byte
	State                  string
	OrganizationID         uuid.UUID
	OrganizationExternalID *string
	EnvironmentID          uuid.UUID
}

func (q *Queries) GetSAMLAccessCodeData(ctx context.Context, arg GetSAMLAccessCodeDataParams) (GetSAMLAccessCodeDataRow, error) {
	row := q.db.QueryRow(ctx, getSAMLAccessCodeData, arg.AppOrganizationID, arg.AccessCode)
	var i GetSAMLAccessCodeDataRow
	err := row.Scan(
		&i.SamlFlowID,
		&i.SubjectIdpID,
		&i.SubjectIdpAttributes,
		&i.State,
		&i.OrganizationID,
		&i.OrganizationExternalID,
		&i.EnvironmentID,
	)
	return i, err
}

const getSAMLConnection = `-- name: GetSAMLConnection :one
select saml_connections.id, saml_connections.organization_id, saml_connections.idp_redirect_url, saml_connections.idp_x509_certificate, saml_connections.idp_entity_id, saml_connections.sp_entity_id
from saml_connections
         join organizations on saml_connections.organization_id = organizations.id
         join environments on organizations.environment_id = environments.id
where environments.app_organization_id = $1
  and saml_connections.id = $2
`

type GetSAMLConnectionParams struct {
	AppOrganizationID uuid.UUID
	ID                uuid.UUID
}

func (q *Queries) GetSAMLConnection(ctx context.Context, arg GetSAMLConnectionParams) (SamlConnection, error) {
	row := q.db.QueryRow(ctx, getSAMLConnection, arg.AppOrganizationID, arg.ID)
	var i SamlConnection
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.IdpRedirectUrl,
		&i.IdpX509Certificate,
		&i.IdpEntityID,
		&i.SpEntityID,
	)
	return i, err
}

const getSAMLConnectionByID = `-- name: GetSAMLConnectionByID :one
select id, organization_id, idp_redirect_url, idp_x509_certificate, idp_entity_id, sp_entity_id
from saml_connections
where id = $1
`

func (q *Queries) GetSAMLConnectionByID(ctx context.Context, id uuid.UUID) (SamlConnection, error) {
	row := q.db.QueryRow(ctx, getSAMLConnectionByID, id)
	var i SamlConnection
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.IdpRedirectUrl,
		&i.IdpX509Certificate,
		&i.IdpEntityID,
		&i.SpEntityID,
	)
	return i, err
}

const getSAMLFlow = `-- name: GetSAMLFlow :one
select saml_flows.id, saml_flows.saml_connection_id, saml_flows.access_code, saml_flows.state, saml_flows.create_time, saml_flows.expire_time, saml_flows.subject_idp_id, saml_flows.subject_idp_attributes, saml_flows.update_time, saml_flows.auth_redirect_url, saml_flows.get_redirect_time, saml_flows.initiate_request, saml_flows.initiate_time, saml_flows.assertion, saml_flows.app_redirect_url, saml_flows.receive_assertion_time, saml_flows.redeem_time, saml_flows.redeem_response
from saml_flows
         join saml_connections on saml_flows.saml_connection_id = saml_connections.id
         join organizations on saml_connections.organization_id = organizations.id
         join environments on organizations.environment_id = environments.id
where environments.app_organization_id = $1
  and saml_flows.id = $2
`

type GetSAMLFlowParams struct {
	AppOrganizationID uuid.UUID
	ID                uuid.UUID
}

func (q *Queries) GetSAMLFlow(ctx context.Context, arg GetSAMLFlowParams) (SamlFlow, error) {
	row := q.db.QueryRow(ctx, getSAMLFlow, arg.AppOrganizationID, arg.ID)
	var i SamlFlow
	err := row.Scan(
		&i.ID,
		&i.SamlConnectionID,
		&i.AccessCode,
		&i.State,
		&i.CreateTime,
		&i.ExpireTime,
		&i.SubjectIdpID,
		&i.SubjectIdpAttributes,
		&i.UpdateTime,
		&i.AuthRedirectUrl,
		&i.GetRedirectTime,
		&i.InitiateRequest,
		&i.InitiateTime,
		&i.Assertion,
		&i.AppRedirectUrl,
		&i.ReceiveAssertionTime,
		&i.RedeemTime,
		&i.RedeemResponse,
	)
	return i, err
}

const getSAMLRedirectURLData = `-- name: GetSAMLRedirectURLData :one
select environments.auth_url
from saml_connections
         join organizations on saml_connections.organization_id = organizations.id
         join environments on organizations.environment_id = environments.id
where environments.app_organization_id = $1
  and saml_connections.id = $2
`

type GetSAMLRedirectURLDataParams struct {
	AppOrganizationID uuid.UUID
	ID                uuid.UUID
}

func (q *Queries) GetSAMLRedirectURLData(ctx context.Context, arg GetSAMLRedirectURLDataParams) (*string, error) {
	row := q.db.QueryRow(ctx, getSAMLRedirectURLData, arg.AppOrganizationID, arg.ID)
	var auth_url *string
	err := row.Scan(&auth_url)
	return auth_url, err
}

const listEnvironments = `-- name: ListEnvironments :many
select id, redirect_url, app_organization_id, display_name, auth_url
from environments
where app_organization_id = $1
  and id > $2
order by id
limit $3
`

type ListEnvironmentsParams struct {
	AppOrganizationID uuid.UUID
	ID                uuid.UUID
	Limit             int32
}

func (q *Queries) ListEnvironments(ctx context.Context, arg ListEnvironmentsParams) ([]Environment, error) {
	rows, err := q.db.Query(ctx, listEnvironments, arg.AppOrganizationID, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Environment
	for rows.Next() {
		var i Environment
		if err := rows.Scan(
			&i.ID,
			&i.RedirectUrl,
			&i.AppOrganizationID,
			&i.DisplayName,
			&i.AuthUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOrganizationDomains = `-- name: ListOrganizationDomains :many
select id, organization_id, domain
from organization_domains
where organization_id = any ($1::uuid[])
`

func (q *Queries) ListOrganizationDomains(ctx context.Context, dollar_1 []uuid.UUID) ([]OrganizationDomain, error) {
	rows, err := q.db.Query(ctx, listOrganizationDomains, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrganizationDomain
	for rows.Next() {
		var i OrganizationDomain
		if err := rows.Scan(&i.ID, &i.OrganizationID, &i.Domain); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOrganizations = `-- name: ListOrganizations :many
select id, environment_id, external_id
from organizations
where environment_id = $1
  and id > $2
order by id
limit $3
`

type ListOrganizationsParams struct {
	EnvironmentID uuid.UUID
	ID            uuid.UUID
	Limit         int32
}

func (q *Queries) ListOrganizations(ctx context.Context, arg ListOrganizationsParams) ([]Organization, error) {
	rows, err := q.db.Query(ctx, listOrganizations, arg.EnvironmentID, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Organization
	for rows.Next() {
		var i Organization
		if err := rows.Scan(&i.ID, &i.EnvironmentID, &i.ExternalID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSAMLConnections = `-- name: ListSAMLConnections :many
select id, organization_id, idp_redirect_url, idp_x509_certificate, idp_entity_id, sp_entity_id
from saml_connections
where organization_id = $1
  and id > $2
order by id
limit $3
`

type ListSAMLConnectionsParams struct {
	OrganizationID uuid.UUID
	ID             uuid.UUID
	Limit          int32
}

func (q *Queries) ListSAMLConnections(ctx context.Context, arg ListSAMLConnectionsParams) ([]SamlConnection, error) {
	rows, err := q.db.Query(ctx, listSAMLConnections, arg.OrganizationID, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SamlConnection
	for rows.Next() {
		var i SamlConnection
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.IdpRedirectUrl,
			&i.IdpX509Certificate,
			&i.IdpEntityID,
			&i.SpEntityID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSAMLFlows = `-- name: ListSAMLFlows :many
select id, saml_connection_id, access_code, state, create_time, expire_time, subject_idp_id, subject_idp_attributes, update_time, auth_redirect_url, get_redirect_time, initiate_request, initiate_time, assertion, app_redirect_url, receive_assertion_time, redeem_time, redeem_response
from saml_flows
where saml_connection_id = $1
  and id > $2
order by id
limit $3
`

type ListSAMLFlowsParams struct {
	SamlConnectionID uuid.UUID
	ID               uuid.UUID
	Limit            int32
}

func (q *Queries) ListSAMLFlows(ctx context.Context, arg ListSAMLFlowsParams) ([]SamlFlow, error) {
	rows, err := q.db.Query(ctx, listSAMLFlows, arg.SamlConnectionID, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SamlFlow
	for rows.Next() {
		var i SamlFlow
		if err := rows.Scan(
			&i.ID,
			&i.SamlConnectionID,
			&i.AccessCode,
			&i.State,
			&i.CreateTime,
			&i.ExpireTime,
			&i.SubjectIdpID,
			&i.SubjectIdpAttributes,
			&i.UpdateTime,
			&i.AuthRedirectUrl,
			&i.GetRedirectTime,
			&i.InitiateRequest,
			&i.InitiateTime,
			&i.Assertion,
			&i.AppRedirectUrl,
			&i.ReceiveAssertionTime,
			&i.RedeemTime,
			&i.RedeemResponse,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEnvironment = `-- name: UpdateEnvironment :one
update environments
set display_name = $1,
    redirect_url = $2
where id = $3
returning id, redirect_url, app_organization_id, display_name, auth_url
`

type UpdateEnvironmentParams struct {
	DisplayName *string
	RedirectUrl *string
	ID          uuid.UUID
}

func (q *Queries) UpdateEnvironment(ctx context.Context, arg UpdateEnvironmentParams) (Environment, error) {
	row := q.db.QueryRow(ctx, updateEnvironment, arg.DisplayName, arg.RedirectUrl, arg.ID)
	var i Environment
	err := row.Scan(
		&i.ID,
		&i.RedirectUrl,
		&i.AppOrganizationID,
		&i.DisplayName,
		&i.AuthUrl,
	)
	return i, err
}

const updateOrganization = `-- name: UpdateOrganization :one
update organizations
set external_id = $1
where id = $2
returning id, environment_id, external_id
`

type UpdateOrganizationParams struct {
	ExternalID *string
	ID         uuid.UUID
}

func (q *Queries) UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) (Organization, error) {
	row := q.db.QueryRow(ctx, updateOrganization, arg.ExternalID, arg.ID)
	var i Organization
	err := row.Scan(&i.ID, &i.EnvironmentID, &i.ExternalID)
	return i, err
}

const updateSAMLConnection = `-- name: UpdateSAMLConnection :one
update saml_connections
set idp_entity_id        = $1,
    idp_redirect_url     = $2,
    idp_x509_certificate = $3
where id = $4
returning id, organization_id, idp_redirect_url, idp_x509_certificate, idp_entity_id, sp_entity_id
`

type UpdateSAMLConnectionParams struct {
	IdpEntityID        *string
	IdpRedirectUrl     *string
	IdpX509Certificate []byte
	ID                 uuid.UUID
}

func (q *Queries) UpdateSAMLConnection(ctx context.Context, arg UpdateSAMLConnectionParams) (SamlConnection, error) {
	row := q.db.QueryRow(ctx, updateSAMLConnection,
		arg.IdpEntityID,
		arg.IdpRedirectUrl,
		arg.IdpX509Certificate,
		arg.ID,
	)
	var i SamlConnection
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.IdpRedirectUrl,
		&i.IdpX509Certificate,
		&i.IdpEntityID,
		&i.SpEntityID,
	)
	return i, err
}

const updateSAMLFlowAppRedirectURL = `-- name: UpdateSAMLFlowAppRedirectURL :one
update saml_flows
set app_redirect_url = $1
where id = $2
returning id, saml_connection_id, access_code, state, create_time, expire_time, subject_idp_id, subject_idp_attributes, update_time, auth_redirect_url, get_redirect_time, initiate_request, initiate_time, assertion, app_redirect_url, receive_assertion_time, redeem_time, redeem_response
`

type UpdateSAMLFlowAppRedirectURLParams struct {
	AppRedirectUrl *string
	ID             uuid.UUID
}

func (q *Queries) UpdateSAMLFlowAppRedirectURL(ctx context.Context, arg UpdateSAMLFlowAppRedirectURLParams) (SamlFlow, error) {
	row := q.db.QueryRow(ctx, updateSAMLFlowAppRedirectURL, arg.AppRedirectUrl, arg.ID)
	var i SamlFlow
	err := row.Scan(
		&i.ID,
		&i.SamlConnectionID,
		&i.AccessCode,
		&i.State,
		&i.CreateTime,
		&i.ExpireTime,
		&i.SubjectIdpID,
		&i.SubjectIdpAttributes,
		&i.UpdateTime,
		&i.AuthRedirectUrl,
		&i.GetRedirectTime,
		&i.InitiateRequest,
		&i.InitiateTime,
		&i.Assertion,
		&i.AppRedirectUrl,
		&i.ReceiveAssertionTime,
		&i.RedeemTime,
		&i.RedeemResponse,
	)
	return i, err
}

const updateSAMLFlowRedeem = `-- name: UpdateSAMLFlowRedeem :one
update saml_flows
set update_time = $1,
    redeem_time = $2,
    redeem_response = $3
where id = $4
returning id, saml_connection_id, access_code, state, create_time, expire_time, subject_idp_id, subject_idp_attributes, update_time, auth_redirect_url, get_redirect_time, initiate_request, initiate_time, assertion, app_redirect_url, receive_assertion_time, redeem_time, redeem_response
`

type UpdateSAMLFlowRedeemParams struct {
	UpdateTime     time.Time
	RedeemTime     *time.Time
	RedeemResponse []byte
	ID             uuid.UUID
}

func (q *Queries) UpdateSAMLFlowRedeem(ctx context.Context, arg UpdateSAMLFlowRedeemParams) (SamlFlow, error) {
	row := q.db.QueryRow(ctx, updateSAMLFlowRedeem,
		arg.UpdateTime,
		arg.RedeemTime,
		arg.RedeemResponse,
		arg.ID,
	)
	var i SamlFlow
	err := row.Scan(
		&i.ID,
		&i.SamlConnectionID,
		&i.AccessCode,
		&i.State,
		&i.CreateTime,
		&i.ExpireTime,
		&i.SubjectIdpID,
		&i.SubjectIdpAttributes,
		&i.UpdateTime,
		&i.AuthRedirectUrl,
		&i.GetRedirectTime,
		&i.InitiateRequest,
		&i.InitiateTime,
		&i.Assertion,
		&i.AppRedirectUrl,
		&i.ReceiveAssertionTime,
		&i.RedeemTime,
		&i.RedeemResponse,
	)
	return i, err
}

const updateSAMLFlowSubjectData = `-- name: UpdateSAMLFlowSubjectData :one
update saml_flows
set subject_idp_id         = $1,
    subject_idp_attributes = $2
where id = $3
returning id, saml_connection_id, access_code, state, create_time, expire_time, subject_idp_id, subject_idp_attributes, update_time, auth_redirect_url, get_redirect_time, initiate_request, initiate_time, assertion, app_redirect_url, receive_assertion_time, redeem_time, redeem_response
`

type UpdateSAMLFlowSubjectDataParams struct {
	SubjectIdpID         *string
	SubjectIdpAttributes []byte
	ID                   uuid.UUID
}

func (q *Queries) UpdateSAMLFlowSubjectData(ctx context.Context, arg UpdateSAMLFlowSubjectDataParams) (SamlFlow, error) {
	row := q.db.QueryRow(ctx, updateSAMLFlowSubjectData, arg.SubjectIdpID, arg.SubjectIdpAttributes, arg.ID)
	var i SamlFlow
	err := row.Scan(
		&i.ID,
		&i.SamlConnectionID,
		&i.AccessCode,
		&i.State,
		&i.CreateTime,
		&i.ExpireTime,
		&i.SubjectIdpID,
		&i.SubjectIdpAttributes,
		&i.UpdateTime,
		&i.AuthRedirectUrl,
		&i.GetRedirectTime,
		&i.InitiateRequest,
		&i.InitiateTime,
		&i.Assertion,
		&i.AppRedirectUrl,
		&i.ReceiveAssertionTime,
		&i.RedeemTime,
		&i.RedeemResponse,
	)
	return i, err
}

const upsertSAMLFlowInitiate = `-- name: UpsertSAMLFlowInitiate :one
insert into saml_flows (id, saml_connection_id, access_code, expire_time, state, create_time, update_time,
                        initiate_request, initiate_time)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
on conflict (id) do update set update_time      = excluded.update_time,
                               initiate_request = excluded.initiate_request,
                               initiate_time    = excluded.initiate_time
returning id, saml_connection_id, access_code, state, create_time, expire_time, subject_idp_id, subject_idp_attributes, update_time, auth_redirect_url, get_redirect_time, initiate_request, initiate_time, assertion, app_redirect_url, receive_assertion_time, redeem_time, redeem_response
`

type UpsertSAMLFlowInitiateParams struct {
	ID               uuid.UUID
	SamlConnectionID uuid.UUID
	AccessCode       uuid.UUID
	ExpireTime       time.Time
	State            string
	CreateTime       time.Time
	UpdateTime       time.Time
	InitiateRequest  *string
	InitiateTime     *time.Time
}

func (q *Queries) UpsertSAMLFlowInitiate(ctx context.Context, arg UpsertSAMLFlowInitiateParams) (SamlFlow, error) {
	row := q.db.QueryRow(ctx, upsertSAMLFlowInitiate,
		arg.ID,
		arg.SamlConnectionID,
		arg.AccessCode,
		arg.ExpireTime,
		arg.State,
		arg.CreateTime,
		arg.UpdateTime,
		arg.InitiateRequest,
		arg.InitiateTime,
	)
	var i SamlFlow
	err := row.Scan(
		&i.ID,
		&i.SamlConnectionID,
		&i.AccessCode,
		&i.State,
		&i.CreateTime,
		&i.ExpireTime,
		&i.SubjectIdpID,
		&i.SubjectIdpAttributes,
		&i.UpdateTime,
		&i.AuthRedirectUrl,
		&i.GetRedirectTime,
		&i.InitiateRequest,
		&i.InitiateTime,
		&i.Assertion,
		&i.AppRedirectUrl,
		&i.ReceiveAssertionTime,
		&i.RedeemTime,
		&i.RedeemResponse,
	)
	return i, err
}

const upsertSAMLFlowReceiveAssertion = `-- name: UpsertSAMLFlowReceiveAssertion :one
insert into saml_flows (id, saml_connection_id, access_code, expire_time, state, create_time, update_time,
                        assertion, receive_assertion_time)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
on conflict (id) do update set update_time            = excluded.update_time,
                               assertion              = excluded.assertion,
                               receive_assertion_time = excluded.receive_assertion_time
returning id, saml_connection_id, access_code, state, create_time, expire_time, subject_idp_id, subject_idp_attributes, update_time, auth_redirect_url, get_redirect_time, initiate_request, initiate_time, assertion, app_redirect_url, receive_assertion_time, redeem_time, redeem_response
`

type UpsertSAMLFlowReceiveAssertionParams struct {
	ID                   uuid.UUID
	SamlConnectionID     uuid.UUID
	AccessCode           uuid.UUID
	ExpireTime           time.Time
	State                string
	CreateTime           time.Time
	UpdateTime           time.Time
	Assertion            *string
	ReceiveAssertionTime *time.Time
}

func (q *Queries) UpsertSAMLFlowReceiveAssertion(ctx context.Context, arg UpsertSAMLFlowReceiveAssertionParams) (SamlFlow, error) {
	row := q.db.QueryRow(ctx, upsertSAMLFlowReceiveAssertion,
		arg.ID,
		arg.SamlConnectionID,
		arg.AccessCode,
		arg.ExpireTime,
		arg.State,
		arg.CreateTime,
		arg.UpdateTime,
		arg.Assertion,
		arg.ReceiveAssertionTime,
	)
	var i SamlFlow
	err := row.Scan(
		&i.ID,
		&i.SamlConnectionID,
		&i.AccessCode,
		&i.State,
		&i.CreateTime,
		&i.ExpireTime,
		&i.SubjectIdpID,
		&i.SubjectIdpAttributes,
		&i.UpdateTime,
		&i.AuthRedirectUrl,
		&i.GetRedirectTime,
		&i.InitiateRequest,
		&i.InitiateTime,
		&i.Assertion,
		&i.AppRedirectUrl,
		&i.ReceiveAssertionTime,
		&i.RedeemTime,
		&i.RedeemResponse,
	)
	return i, err
}
