# Kubernetes Entrypoint

Kubernetes-entrypoint enables complex deployments on top of Kubernetes.

## Overview

Kubernetes-entrypoint is meant to be used as a container entrypoint, which means it has to bundled in the container.
Before launching the desired application, the entrypoint verifies and waits for all specified dependencies to be met.

Kubernetes-entrypoint queries the Kubernetes API directly, and each container is self-aware of its dependencies and their states.
Therefore, no centralized orchestration layer is required to manage deployments, and scenarios (such as failure recovery or pod migration) become easy.

## Usage

Kubernetes-entrypoint reads the dependencies out of environment variables passed into a container.
There is only one required environment variable `COMMAND` which specifies a command (arguments delimited by whitespace) which has to be executed when all dependencies are resolved:

`COMMAND="sleep inf"`

## Supported types of dependencies

All dependencies are passed as environment variables with the format `DEPENDENCY_<NAME>`, delimited by a colon.
For dependencies to be effective please use [readiness probes](https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/) for all containers.

Most dependencies also support specifying a namespace using the `namespace:name` format. If `namespace` is omitted, the dependency is assumed to be running in the same namespace as kubernetes-entrypoint. This is not supported for the Container, Config, or Socket dependencies, since a different namespace is irrelevant for those cases.

For instance:

`DEPENDENCY_SERVICE=mysql:mariadb,keystone-api`

resolves `mariadb` in the `mysql` namespace and `keystone-api` in the same namespace as kubernetes-entrypoint was deployed in.

### Service
Checks whether given kubernetes service has at least one endpoint.
Example:

`DEPENDENCY_SERVICE=mariadb,keystone-api`

### Container
Within a pod composed of multiple containers, kubernetes-entrypoint waits for the containers specified by their names to start.
This dependency requires a `POD_NAME` environment variable which can be easily passed through the [downward API](https://kubernetes.io/docs/concepts/workloads/pods/downward-api/).
Example:

`DEPENDENCY_CONTAINER=nova-libvirt,virtlogd`

### Daemonset
Checks if a specified daemonset is already running on the same host.
This dependency requires a `POD_NAME` environment variable which can be easily passed through the [downward API](https://kubernetes.io/docs/concepts/workloads/pods/downward-api/).
The `POD_NAME` variable is mandatory and is used to resolve dependencies.
Example:

`DEPENDENCY_DAEMONSET=openvswitch-agent`

### Job
Checks if a given job or set of jobs with matching name and/or labels succeeded at least once.
In order to use labels, `DEPENDENCY_JOBS_JSON` must be used.
`DEPENDENCY_JOBS` is supported as well for backward compatibility.
Examples:

`DEPENDENCY_JOBS_JSON='[{"namespace": "foo", "name": "nova-init"}, {"labels": {"initializes": "neutron"}}]'`
`DEPENDENCY_JOBS=nova-init,neutron-init`

### Config
This dependency performs a container level templating of configuration files. It can template an ip address `{{ .IP }}` and hostname `{{ .HOSTNAME }}`.
Templated config has to be stored in an arbitrary directory `/configmaps/<name_of_file>/<name_of_file>`.
This dependency requires an `INTERFACE_NAME` environment variable to know which interface to use to obtain the ip address.
Example:

`DEPENDENCY_CONFIG=/etc/nova/nova.conf`

Kubernetes-entrypoint will look for the configuration file `/configmaps/nova.conf/nova.conf`, template the `{{ .IP }}` and `{{ .HOSTNAME }}` tags, and then save the file as `/etc/nova/nova.conf`.

### Socket
Checks whether a given file exists and that the container has rights to read it.
Example:

`DEPENDENCY_SOCKET=/var/run/openvswitch/ovs.socket`

### Pod
Checks if at least one pod matching the specified labels is already running, by default anywhere in the cluster, or use `"requireSameNode": true` to require a pod on the same node.
Labels are specified using JSON, as seen in the example below.
This dependency requires a `POD_NAME` env which can be easily passed through the [downward API](https://kubernetes.io/docs/concepts/workloads/pods/downward-api/).
The `POD_NAME` variable is mandatory and is used to resolve dependencies.
Example:

`DEPENDENCY_POD_JSON='[{"namespace": "foo", "labels": {"k1": "v1", "k2": "v2"}}, {"labels": {"k1": "v1", "k2": "v2"}, "requireSameNode": true}]'`

## Building

Kubernetes-entrypoint is built with Go modules and requires Go 1.26 or later (see `go.mod`).

Build a binary for your platform:

```
go build -o kubernetes-entrypoint .
```

Or use the provided `Makefile` targets to cross-compile for Linux amd64/arm64:

```
make linux-amd64
make linux-arm64
```

Run the test suite with:

```
make test
```
