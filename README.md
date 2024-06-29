# CapabilityPE (capipe)

<p align="center">
 <img src="images/logo.png?raw=true" alt="Logo" width="50%" height="50%" />
</p>

You know what’s missing? An easy tool that just installs capibilities (ArgoCD, Datadog, KubeVirt, etc.) in an easy way. One command to do it all.

A tool that easily gets whatever capabilities you want your kubernetes cluster to have deployed in production.

`capipe`, which stands for Capibility Platform Engineering, allows you to specify capabilities that you want to install within your Platform Engineering environment.

## Why?

The two biggest questions I get are:
1. What tools should I use?
2. How can I easily deploy a production environment?

Those questions are why I made CapiPE.

Easily deploy what you need in a Kubernetes cluster with one command.

## Dependencies

1. Helm

## Install

Before an actual version is built and released, you can build the CLI/binary/artifact by running the following command in the directory/repo.

```
go build
```

## Command Examples

![](images/help.png)

Add a GitOps Controller

```
capipe argocd
```

```
capipe flux
```

Use flags
```
./capipe datadog --apikey "" --clustername ""
```

## What's Coming...

- One command to install multiple Platform Capabilities
    - Platform Capabilities installed based on a particular stack you choose. Here are some examples:
        - App Stack 1:
            - ArgoCD
            - Crossplane
            - OPA
            - Datadog

        - App Stack 2:
            - Flux CD
            - Kyverno
            - Crossplane
            - Grafana/Prometheus/Tempo/Loki

        - App Stack 3:
            - ArgoCD
            - Crossplane
            - OPA
            - Signoz

        - NetSec App Stack 1:
            - ArgoCD
            - Crossplane
            - OPA
            - Grafana/Prometheus/Tempo/Loki
            - Cilium

        - Virtualized App Stack 1:
            - KubeVirt
            - Cilium
            - ArgoCD
            - Crossplane
            - OPA
            - Grafana/Prometheus/Tempo/Loki

            

![](images/capipe.png)

