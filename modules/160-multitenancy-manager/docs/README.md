---
title: "Module multitenancy-manager"
search: multitenancy
description: Multitenancy and Projects in Kubernetes. The multitenancy-manager module in Deckhouse allows creating projects for various development teams with the ability to subsequently deploy applications in them.
---
## Description

The module enables the creation of projects in a Kubernetes cluster. **Project** is an isolated environment where applications can be deployed.

## Why is this needed?

The standard `Namespace` resource, used for logical resource separation in Kubernetes, does not provide necessary functionalities, hence it is not an isolated environment:
* [Resource consumption by pods](https://kubernetes.io/docs/concepts/policy/resource-quotas/) is not limited by default;
* [Network communication](https://kubernetes.io/docs/concepts/services-networking/network-policies/) with other pods works by default from any point in the cluster;
* Unrestricted access to node resources: address space, network space, mounted host directories.

The configuration capabilities of `Namespace` do not fully meet modern development requirements. By default, the following features are not included for `Namespace`:
* Log collection;
* Audit;
* Vulnerability scanning.

The functionality of projects allows addressing these issues.

{% alert level="warning" %}
The `secret-copier` module [cannot be used together](../secret-copier/) with `multitenancy-manager` module.
{% endalert %}

## Advantages of the module

For platform administrators:
* **Consistency**: Administrators can create projects using the same template, ensuring consistency and simplifying management.
* **Security**: Projects provide isolation of resources and access policies between different projects, supporting a secure multitenant environment.
* **Resource Consumption**: Administrators can easily set quotas on resources and limitations for each project, preventing excessive resource usage.

For platform users:
* **Quick Start**: Developers can request projects created from ready-made templates from administrators, allowing for a quick start to developing a new application.
* **Isolation**: Each project provides an isolated environment where developers can deploy and test their applications without impacting other projects.

## Internal Logic

### Creating a project

To create projects, the following [Custom Resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) are used:
* [ProjectTemplate](cr.html#projecttemplate) — a resource that describes the project template. It defines a list of resources to be created in the project and a schema for parameters that can be passed when creating the project;
* [Project](cr.html#project) — a resource that describes a specific project.

When creating a [Project](cr.html#project) resource from a specific [ProjectTemplate](cr.html#projecttemplate), the following happens:
1. The [parameters](cr.html#project-v1alpha2-spec-parameters) passed are validated against the OpenAPI specification (the [openAPI](cr.html#projecttemplate-v1alpha1-spec-parametersschema) field of [ProjectTemplate](cr.html#projecttemplate));
1. Rendering of the [resources template](cr.html#projecttemplate-v1alpha1-spec-resourcestemplate) is performed using [Helm](https://helm.sh/docs/). Values for rendering are taken from the [parameters](cr.html#projecttemplate-v1alpha1-spec-parametersschema) field of the [Project](cr.html#project) resource;
1. A `Namespace` is created with a name matching the name of [Project](cr.html#project);
1. All resources described in the template are created in sequence.

> **Attention!** When changing the project template, all created projects will be updated according to the new template.

### Isolating a project

The project is based on the `Namespace` resource mechanism. Namespaces group pods, services, secrets, and other objects but do not provide complete isolation. The project functionality enhances namespaces by offering additional tools to improve control and security levels. To manage project isolation, Kubernetes features can be leveraged, such as:

- Access control resources (`AuthorizationRule` / `RoleBinding`) — manage interaction with objects within a `Namespace`. Define rules and assign roles to precisely control who can perform actions in your project.
- Resource quotas (`ResourceQuota`) — set limits on resource usage, such as CPU time, RAM, and object counts within a `Namespace`. These quotas help prevent excessive load and maintain control over applications within the project.
- Network connectivity control resources  (`NetworkPolicy`) — control incoming and outgoing network traffic within a `Namespace`. Configure allowed connections between pods to enhance security and manage network interactions effectively.

These tools can be combined to configure the project according to the requirements of your application.

