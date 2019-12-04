# Third Party Service Deployment Microservice for Mattermost Private Cloud
## This Document
This document proposes and outlines a general design for a new microservice responsible for creating instances of third party software utilized by Mattermost Private Cloud for various purposes.

## Scope of Responsibilities
At this time this microservice is proposed to be responsible for deploying third party services (e.g. Prometheus and Fluent-Bit) and its responsibilities shall not be expanded without further design or discussion.

## Design
This microservice will be a stateless Go application that presents a REST API for input from the `cloud` provisioning server and outputs deployments in AWS with the specifications provided by the `cloud server`. The `cloud server` API will serve as a proxy for handling requests from users and distributing the necessary requests to this service for provisioning of third party services used to monitor Mattermost Private Cloud.

Deployment of individual services by the microservice will start out being performed the same way as it is currently performed by `cloud server` for existing functionality, with an eye towards isolating Helm-based deployments to be replaced with Operator-based deployments. Any new services added to this deployment should avoid external dependencies if possible.

Interfaces will be designed such that new services that we wish to be created using this service can be added on using the same patterns defined by the first services in order to encourage modularity.

### Code layout
The microservice's main entrypoint will be delivered as a compiled Go binary alongside the `cloud` command as part of the `mattermost-cloud` project and repositories.

The entrypoint will live in `cmd/` with the `cloud` entrypoint and `main` package for that binary, and shall draw on the same libraries from `internal/` in this repository.

Shared code between the existing `cloud` server code and the new microservice currently tightly integrated with `cloud` will be abstracted into its own package or a relevant package in `internal/` for use by both.

A new `Makefile` target will be added for the new executable and `make build` will be altered to build both `cloud` as well as the new binary. A new target to build only `cloud` will be added as well.

### Deployment

Since this microservice will be stateless (state will remain in the `cloud server`'s database and should be passed between that service and this one as necessary) `n` copies of the microservice will be able to be deployed in order to solve any HA problems that may arise. Rolling upgrades of this service will be possible as a result, as well, if we run >1 instances of this microservice.

Otherwise it should be deployable like any other microservice and we should just be consistent with how `cloud server` is deployed.

### Deliverable

The deliverable will either be a compiled binary or a Docker image. Currently we have a dependency on Helm which will not be resolved in the first version. I am still investigating whether it is better to include the binary we are dependent on in the deliverable artifact by delivering the microservice as a Docker image or using something like `go-bindata` or `xar` to ship an executable archive containing both the new executable and any executables upon which the microservice depends, e.g. Helm.

### API

The REST API presented by the service should allow a REST client to:
1) Queue the creation of new deployments by passing all necessary metadata as JSON to this service
3) Determine the status of deployments still being created
2) Terminate ongoing deployments in-progress

## Prometheus
Prometheus is in the process of being removed from the `cloud` command and migrated to the new microservice. It is the first service to be moved.

## Fluent Bit
Fluent Bit will be investigated after Prometheus.

## Future
Other 3rd party services including the NGINX Ingress Controller may also be deployed through this microservice in the future.

## Open Questions

1. Is a binary or a Dockerfile a preferable deliverable artifact?

2. Routing/load balancing between `cloud server`(s) and `n` instances of this microservice (probably premature optimization to think about this now)

3. Naming: Generic names are used for this new service as well as the `cloud server` service that is apart of the `cloud` binary. Better names would probably make this topic easier to talk about.
