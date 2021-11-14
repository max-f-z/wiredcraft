### part 5 

provide k8s manifests files.

## Overview

persistent storage using redis. However it does not scale at the momemnt.
Single instance does not provide high avialiablity.
Persistent storage using storage class "local-storage", pv & pvcs

Web-app configured with hpa, that will scale between 2-6 replicas based on CPU usage.
BUT a better idea would be have custom metrics and scale based on custom metrics defined in go code.

Also provide deployment, service, ingress manifests