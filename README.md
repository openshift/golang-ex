Golang Sample App on OpenShift
==============================

This is a sample Golang web application for OpenShift v3 that use the [beego](http://beego.me/)
framework. This sample uses official Docker [golang](https://registry.hub.docker.com/u/library/golang/) image for the build
without any modification to the image or sample itself.

This example was copied from the official [beego samples](https://github.com/beego/samples) repository.

If you'd like to install it, follow [these instructions](https://github.com/openshift/golang-ex/blob/master/README.md#installation).

The steps in this document assume that you have access to an OpenShift deployment that you can deploy applications on, including the default set of ImageStreams defined.  Instructions for installing the default ImageStreams are available [here](http://docs.openshift.org/latest/admin_guide/install/first_steps.html).  If you are defining the set of ImageStreams now, remember to pass in the proper cluster-admin credentials and to create the ImageStreams in the 'openshift' namespace.

### Installation

1. Fork a copy of [golang-ex](https://github.com/openshift/golang-ex)
2. Clone your repository to your development machine and cd to the repository directory
3. Add a Golang application from the beego template:

        $ oc new-app openshift/templates/beego.json -p SOURCE_REPOSITORY_URL=https://github.com/yourusername/golang-ex

4. A build will be started automatically. It might take some time.  
You can run the command below to watch for builds:

        $ oc get builds -w

5. Once the build is running, follow the build logs:

        $ oc logs -f bc/beego-example

6. Wait for the beego-example pod to start up (may take some time):

        $ oc get pods -w

    Sample output:

        NAME                     READY     REASON       RESTARTS   AGE
        beego-example-1-6c23l     1/1       Running        0          2m
        beego-example-1-build     0/1       ExitCode:0     0          4m

7. Check the IP and port the beego-example service is running on:

        $ oc get svc

    Sample output:

        NAME             LABELS                              SELECTOR              IP(S)           PORT(S)
        beego-example     template=beego-example     name=beego-example           172.30.210.29    8080/TCP

In this case, the IP for beego-example is 172.30.210.29 and it is on port 8080.  
*Note*: you can also get this information from the web console.

### Building

After you create the application from a template, you can follow the build progress
with the `oc logs` command:

```console
$ oc logs -f bc/beego-example
Step 0 : FROM golang:1.5
 ---> 06918e33c280
Step 1 : USER nobody
 ---> Using cache
 ---> 135bdc55c498
Step 2 : ENV GO15VENDOREXPERIMENT 1
 ---> Using cache
 ---> 404e4094569e
Step 3 : RUN mkdir -p /go/src/github.com/openshift/golang-ex
 ---> Running in 684a99f755fe
 ---> 00ef6de11cda
Removing intermediate container 684a99f755fe
Step 4 : WORKDIR /go/src/github.com/openshift/golang-ex
 ---> Running in e9190f59da2c
 ---> 9a7d93430e72
Removing intermediate container e9190f59da2c
Step 5 : COPY . /go/src/github.com/openshift/golang-ex
 ---> ba13589484a7
Removing intermediate container 7cebbee54378
Step 6 : RUN go-wrapper install
 ---> Running in 26f41423327a
+ exec go get -v -d
+ exec go install -v
github.com/openshift/golang-ex/vendor/github.com/astaxie/beego/config
github.com/openshift/golang-ex/vendor/github.com/astaxie/beego/utils
github.com/openshift/golang-ex/vendor/github.com/astaxie/beego/grace
github.com/openshift/golang-ex/vendor/github.com/astaxie/beego/session
github.com/openshift/golang-ex/vendor/github.com/astaxie/beego/logs
github.com/openshift/golang-ex/vendor/github.com/astaxie/beego/context
github.com/openshift/golang-ex/vendor/github.com/astaxie/beego/toolbox
github.com/openshift/golang-ex/vendor/github.com/Unknwon/goconfig
github.com/openshift/golang-ex/vendor/github.com/astaxie/beego
github.com/openshift/golang-ex/vendor/github.com/beego/i18n
github.com/openshift/golang-ex/vendor/github.com/gorilla/websocket
github.com/openshift/golang-ex/models
github.com/openshift/golang-ex/controllers
github.com/openshift/golang-ex
 ---> f91517380de3
Removing intermediate container 26f41423327a
Step 7 : CMD go-wrapper run
 ---> Running in c99ec8507c23
 ---> 01c4d1f5e508
Removing intermediate container c99ec8507c23
...
Successfully built 1d2138f43eab
I0113 14:19:10.406586       1 docker.go:95] Pushing image 172.30.149.59:5000/demo/beego-example:latest ...
I0113 14:19:12.002300       1 docker.go:99] Push successful
```

### Accessing the application

If you have the OpenShift router running, you should be able to access the
application just by typing the beego-example route DNS
(beego-example.openshiftapps.com) into your browser. However, you will have to
run you own DNS server first, but we can cheat it by modifying the `/etc/hosts`
file on your host machine. Just append this line at the end of that file:

```
192.168.124.206 beego-example.openshiftapps.com
```

The `192.168.124.206` is the IP address of the machine you are running the
OpenShift on. Once you have this change, you can just type the application DNS
into browser and you should see the chat application login screen.

### License

This code is dedicated to the public domain to the maximum extent permitted by applicable law, pursuant to [CC0](http://creativecommons.org/publicdomain/zero/1.0/).

The sample beego application code follows the original [Apache License](https://github.com/beego/samples/blob/master/LICENSE)
