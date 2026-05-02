# Civicore — Seeded Findings (25 total)

## SAST (11)
| ID | CWE | File |
|---|---|---|
| CIV-SAST-001 | CWE-89   | `internal/handlers/citizens.go` (fmt.Sprintf into sql.Query) |
| CIV-SAST-002 | CWE-208  | `internal/auth/totp.go` (string == comparison on TOTP code — timing) |
| CIV-SAST-003 | CWE-798  | `internal/config/config.go` (hardcoded OIDC client secret) |
| CIV-SAST-004 | CWE-79   | `internal/handlers/profile.go` (raw HTML write of citizen name) |
| CIV-SAST-005 | CWE-22   | `internal/handlers/docs.go` (filepath.Join with user input) |
| CIV-SAST-006 | CWE-918  | `internal/services/registry.go` (http.Get on caller URL) |
| CIV-SAST-007 | CWE-330  | `internal/auth/session.go` (math/rand for session IDs) |
| CIV-SAST-008 | CWE-326 | `internal/transport/tls.go` (TLS InsecureSkipVerify=true) |
| CIV-SAST-009 | CWE-639  | `internal/handlers/citizens.go` (no auth check on /citizens/{id}) |
| CIV-SAST-010 | CWE-352  | `internal/middleware/cors.go` (`AllowAllOrigins=true` + credentials) |
| CIV-SAST-011 | CWE-117  | `internal/middleware/log.go` (request body logged unredacted) |

## IaC (7)
- CIV-IAC-001 RDS publicly accessible (`infra/terraform/rds.tf`)
- CIV-IAC-002 RDS no encryption (`infra/terraform/rds.tf`)
- CIV-IAC-003 RDS deletion_protection=false (`infra/terraform/rds.tf`)
- CIV-IAC-004 SG ingress 5432 (`infra/terraform/main.tf`)
- CIV-IAC-005 KMS key with `kms:*` to `*` (`infra/terraform/kms.tf`)
- CIV-IAC-006 Container runs as root (`Dockerfile`)
- CIV-IAC-007 K8s `automountServiceAccountToken: true` + privileged (`infra/k8s/deployment.yaml`)

## SCA (3)
| ID | Module | Version | CVE |
|---|---|---|---|
| CIV-SCA-001 | github.com/golang-jwt/jwt/v4    | 4.4.1   | CVE-2024-51744 |
| CIV-SCA-002 | github.com/lestrrat-go/jwx      | 1.2.16  | CVE-2024-22195 (transitive jx) |
| CIV-SCA-003 | golang.org/x/crypto             | 0.0.0-20220525230936-793ad666bf5e | CVE-2022-27191 |

## Pipeline (4)
- CIV-CI-001 Hardcoded OIDC client secret in workflow env
- CIV-CI-002 No permissions block
- CIV-CI-003 `actions/checkout@v3` + `pull_request_target` + checkout PR head SHA
- CIV-CI-004 Image push using mutable `latest` to public registry
