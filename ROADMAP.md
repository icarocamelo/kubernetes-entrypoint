# Roadmap

This is a living list of what's planned for kubernetes-entrypoint, aimed at keeping it useful for the kind of long-lived OpenStack-on-Kubernetes deployments (OpenStack-Helm and similar) that depend on it. It's maintained on an ongoing basis, with releases cut roughly every 1-3 months as items below are finished. Priorities can shift — open an issue if something here should move up or if you need something that isn't listed.

## Next release

Small, self-contained fixes that unblock everything else below.

- [x] Migrate to Go modules, drop the vendor/Godeps setup, update to Go 1.26 and current `k8s.io/client-go` (done).
- [ ] Fix the `dependencies/socket` test that only fails when run as root (filesystem permission checks are bypassed for root, so the "no permissions" case can't be exercised in containers that run as root by default).
- [ ] Replace `.travis.yml` with a GitHub Actions workflow that actually runs (`go build`, `go test`, `go vet`) on every push and PR.
- [ ] Add Dependabot/Renovate for `go.mod` so dependencies don't silently drift years out of date again.

## Within ~1-3 months

Reliability and observability gaps that matter most once this runs unattended in production clusters.

- [ ] **Configurable timeout per dependency.** `entrypoint.Resolve()` currently polls forever with no way to fail fast — a dependency that will never resolve (typo'd service name, RBAC issue) hangs the init container indefinitely instead of surfacing an error. Add a `DEPENDENCY_TIMEOUT` (or per-type override) that exits non-zero with a clear message once exceeded.
- [ ] **Configurable poll interval / backoff.** The 2-second `resolverSleepInterval` is hardcoded. Make it configurable, and back off on repeated API errors instead of hammering the API server at a fixed rate.
- [ ] **Graceful shutdown.** No SIGTERM/SIGINT handling today — a pod terminated mid-init leaves the resolve loop running until the process is killed outright. Handle signals and exit cleanly.
- [ ] **Structured logging.** Swap the plain `log.Logger` wrapper in `logger/` for JSON-capable structured logging, so entrypoint output is easy to filter and correlate in the ELK/Loki-style stacks most OpenStack-Helm deployments already run.
- [ ] **Multi-arch container images.** Publish an official minimal (distroless-style) image for amd64/arm64 via GitHub Actions — the `Makefile` already cross-compiles both, this just wires up publishing.

## Within ~3-6 months

Dependency types and API migrations that keep the tool aligned with where Kubernetes itself is heading.

- [ ] **StatefulSet dependency.** Galera/MariaDB, RabbitMQ, and Ceph in OpenStack-Helm deployments are commonly StatefulSets, not just Deployments/DaemonSets/Jobs. Add a `DEPENDENCY_STATEFULSET` type analogous to the existing Daemonset one.
- [ ] **Migrate the Service dependency off the deprecated Endpoints API** to `discovery.k8s.io/v1` EndpointSlices, which is what modern clusters actually populate at scale.
- [ ] **CRD condition dependency.** A generic "wait until this custom resource's status condition is true" dependency, useful for waiting on cert-manager `Certificate` objects, Rook/Ceph CRDs, etc., without hardcoding each one.
- [ ] Expand the `Config` dependency's template variables beyond `{{ .IP }}` / `{{ .HOSTNAME }}` (e.g. pod name, namespace) for more complex templating needs.

## Longer term / opportunistic

- [ ] Optional Prometheus metrics (dependency wait duration, resolved/failed counts) for the cases where entrypoint runs long enough for it to matter.
- [ ] SBOM generation and vulnerability scanning as part of the release pipeline.
- [ ] Revisit whether the `Container` dependency's polling model could use the Kubernetes API's native watch support instead of repeated `List` calls, to cut down on API server load in large clusters.

## Maintenance track (every release)

- Keep `go.mod` dependencies current — don't let another multi-year gap happen.
- Re-run `go vet` / `golangci-lint` and clear any new findings.
- Confirm `README.md` build/usage instructions still match reality.
