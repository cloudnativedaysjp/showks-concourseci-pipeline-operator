# showks-concourseci-pipeline-operator

## Install

```
$ cat ./config/secret.env
CONCOURSECI_URL=xxxxxxxxxxz
CONCOURSECI_TEAM=xxxxxxxxxxz
CONCOURSECI_USERNAME=xxxxxxxxxxz
CONCOURSECI_PASSWORD=xxxxxxxxxxz
```

```
$ make deploy
```

## Usage

```yaml
apiVersion: showks.cloudnativedays.jp/v1beta1
kind: ConcourseCIPipeline
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
  name: concoursecipipeline-sample
spec:
  target: main
  pipeline: test
  manifest: |-
    jobs:
    - name: hello-world
      plan:
      - task: say-hello
        config:
          platform: linux
          image_resource:
            type: docker-image
            source: {repository: alpine}
          run:
            path: echo
            args: ["Hello, world!!!!"]
```

## Development

Concourse CIにアクセスするため、以下の通り環境変数をセットします。

```
export CONCOURSECI_URL=http://example.com/
export CONCOURSECI_TEAM=xxxx
export CONCOURSECI_USERNAME=xxxx
export CONCOURSECI_PASSWORD=xxxx
```

コントローラーを手元で実行します。

```
$ kubectl apply -f config/crds
$ make run
```



