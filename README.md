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
Step 0 : FROM docker.io/library/golang:1.4.2-onbuild
# Executing 3 build triggers
Trigger 0, COPY . /go/src/app
Step 0 : COPY . /go/src/app
Trigger 1, RUN go-wrapper download
Step 0 : RUN go-wrapper download
 ---> Running in 407a110288b5
+ exec go get -v -d
github.com/astaxie/beego (download)
github.com/beego/i18n (download)
github.com/Unknwon/goconfig (download)
github.com/beego/samples (download)
github.com/gorilla/websocket (download)
Trigger 2, RUN go-wrapper install
Step 0 : RUN go-wrapper install
 ---> Running in d44245edf645
+ exec go install -v
github.com/astaxie/beego/config
github.com/astaxie/beego/utils
github.com/astaxie/beego/grace
github.com/astaxie/beego/session
github.com/astaxie/beego/logs
github.com/astaxie/beego/context
github.com/astaxie/beego/toolbox
github.com/Unknwon/goconfig
github.com/astaxie/beego
github.com/beego/i18n
github.com/beego/samples/WebIM/models
github.com/gorilla/websocket
github.com/beego/samples/WebIM/controllers
app
 ---> ce01cedcb031
Removing intermediate container d44245edf645
Removing intermediate container c0b88982ecd9
Removing intermediate container 407a110288b5
Step 1 : ENTRYPOINT ./run
 ---> Running in 23b409f8c11d
 ---> e584df4684e8
Removing intermediate container 23b409f8c11d
Successfully built e584df4684e8
I0625 10:40:56.628610       1 sti.go:96] Using provided push secret for pushing 172.30.106.31:5000/goex/beego-example image
I0625 10:40:56.628720       1 sti.go:99] Pushing 172.30.106.31:5000/goex/beego-example image ...
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
