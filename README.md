# simple-operator
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites

#### 开发机
- go version go1.23.1
- docker version 20.10.12. （只安装客户端，服务端有限制）
- kubectl version v1.29.3
- Access to a minikube cluster.
- eval $(minikube docker-env)
- kubebuilder 4.2.0 (生成脚手架代码)

#### 加一些配置项，包有可能下载不下来
- go env -w GOPROXY=https://proxy.golang.org,direct
- go env -w GOSUMDB=sum.golang.google.cn

#### minikube
- minikube version: v1.32.0
- Docker Engine - Community 24.0.7 (安装minikube时自动安装)



### Development steps
#### 1. Create a directory for your project:
```
    mkdir simple-operator
    cd simple-operator
```
#### 2. Initialize kubebuilder project
```
    kubebuilder init --domain example.com --repo github.com/wxquare/simple-operator
```

#### 3. Create a new API (Custom Resource Definition): To create an API and controller scaffolding for your Operator, run the following:
```
    kubebuilder create api --group batch --version v1 --kind CronJob
```
- --group batch: The API group name.
- --version v1: The version of the API.
- --kind CronJob: The name of the resource you’re managing.

#### Step 4: Implement Operator Logic
- Define the CRD: api/v1/cronjob_types.go
- Implement the Controller Logic: controllers/cronjob_controller.go


#### Step5: Build the Operator Image
```
    make docker-build IMG=wxquare/simple-operator:latest
```

#### Step6: Push Image to docker hub
```
    make docker-push IMG=wxquare/simple-operator:latest
```

#### Step7: Deploy the CRD (Custom Resource Definition)
```
    make install
```

#### Step8: Deploy Operator
```
    make deploy IMG=docker.io/wxquare/simple-operator:latest
```

#### Step9: Verify The Operator is Running and check logs
```
    kubectl get pods -n simple-operator-system
    kubectl logs simple-operator-controller-manager-788ff8d877-w6t6t -n simple-operator-system -f
```


#### step10: Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

#### step11. To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```
**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```
**UnDeploy the controller from the cluster:**

```sh
make undeploy
```