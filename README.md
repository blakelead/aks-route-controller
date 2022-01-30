# AKS Route Controller

**Early stage of development**

Kubernetes controller that manages an Azure Route Table by adding/deleting routes when Nodes are created/deleted.

## Why this controller?

[AGIC](https://docs.microsoft.com/en-us/azure/application-gateway/ingress-controller-overview) allows us to use Azure Application Gateway as an Ingress Controller to route requests to pods.

But when AKS is deployed with kubenet network plugin, pods are not reachable. A Route Table needs to be attached to the Application Gateway Subnet and Routes for each AKS Nodes need to be created.

AGIC cannot do that at the time of writing this.