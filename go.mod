module github.com/ssoready/ssoready

go 1.26.0

// Keep the runtime defaults of Go 1.21, which built the images before. The
// go line moved to 1.26 only because updated dependencies require it. Newer
// defaults (e.g. rejecting certificates with negative serial numbers, which
// some SAML identity providers still use) need their own review and testing.
godebug default=go1.21

require (
	connectrpc.com/connect v1.16.2
	connectrpc.com/vanguard v0.2.0
	github.com/alecthomas/participle/v2 v2.1.1
	github.com/aws/aws-sdk-go-v2 v1.47.8
	github.com/aws/aws-sdk-go-v2/config v1.27.12
	github.com/aws/aws-sdk-go-v2/service/s3 v1.114.9
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.28.7
	github.com/aws/smithy-go v1.28.8
	github.com/cloudflare/cloudflare-go v0.109.0
	github.com/cyrusaf/ctxlog v1.3.2
	github.com/getsentry/sentry-go v0.28.0
	github.com/go-jose/go-jose/v3 v3.0.5
	github.com/go-jose/go-jose/v4 v4.1.9
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/golang-migrate/migrate/v4 v4.17.1
	github.com/google/go-cmp v0.7.0
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	github.com/hasura/go-graphql-client v0.13.0
	github.com/jackc/pgx/v5 v5.12.0
	github.com/resend/resend-go/v2 v2.6.0
	github.com/rs/cors v1.11.0
	github.com/segmentio/analytics-go/v3 v3.3.0
	github.com/ssoready/conf v0.0.0-20240508183332-dbc356674c9e
	github.com/ssoready/prettyuuid v0.0.0-20241023163822-285da46017b3
	github.com/stretchr/testify v1.11.1
	github.com/ucarion/cli v0.2.0
	golang.org/x/crypto v0.58.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260817212433-ac3dfec99bb1
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260817212433-ac3dfec99bb1
	google.golang.org/grpc v1.85.0
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.26 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.12 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.5.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.8.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.5.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.11.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.14.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.20.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.23.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.7 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/segmentio/backo-go v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.62.0 // indirect
	golang.org/x/sync v0.25.0 // indirect
	golang.org/x/sys v0.50.0 // indirect
	golang.org/x/text v0.44.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	nhooyr.io/websocket v1.8.17 // indirect
)
