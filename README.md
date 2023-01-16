# Technical test Scality

### How would you deploy a stateless REST API in Kubernetes? How about a stateful one?

Stateless : [Run a Stateless Application Using a Deployment](https://kubernetes.io/docs/tasks/run-application/run-stateless-application-deployment/)
Statefull : [Run a Single-Instance Stateful Application](https://kubernetes.io/docs/tasks/run-application/run-single-instance-stateful-application/)

### Pod "A" sends a request to "my-service.my-namespace.svc": what happens behind the scenes?

- 

### What kind of Kubernetes object do you use to deploy a Pod on every node?

[DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)

### List of users in YAML file

```sh
cd users && go run main.go
```

### Explain two approaches to network load balancing, their benefits and drawbacks

[Different types of load balancing pros and cons](https://dealna.com/en/Article/Post/693/Load-Balancing-Do-the-Cons-Balance-the-Pros)

### What command(s) do you use to list storage devices?

```sh
lsblk
```

### An LVM LV is used by a process and we need to double its size. What is the procedure?

-

### How to go from Case 1 to Case 2 (one git command; HEAD on F)

```sh
  git merge
```

### How to go from Case 1 to Case 3 (one git command; HEAD on D)

```sh
  git cherry-pick E
```