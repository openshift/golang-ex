Golang Sample App on OpenShift
============================

This is a sample Golang web application for OpenShift v3 that use the [beego](http://beego.me/)
framework. This sample use official Docker [golang](https://registry.hub.docker.com/u/library/golang/) image for the build
without any modification to the image or sample itself.

This example was copied from the official [beego samples](https://github.com/beego/samples) repository.

If you'd like to install it, follow [these directions](https://github.com/openshift/golang-ex/blob/master/README.md#installation).  

The steps in this document assume that you have access to an OpenShift deployment that you can deploy applications on.

###Installation: 

1. Fork a copy of [golang-ex](https://github.com/openshift/golang-ex)
2. Clone your repository to your development machine and cd to the repository directory
3. Add a Golang application from the beego template:

		$ oc process -f openshift/templates/beego.json -v=SOURCE_REPOSITORY_URL=https://github.com/yourusername/golang-ex | oc create -f - 

4. Note that creating from a template will automatically start a new build. Watch your build progress:

		$ oc build-logs beego-example-1

5. Wait for frontend pods to start up (this can take a few minutes):  

		$ oc get pods -w


	Sample output:  

    	NAME                     READY     REASON       RESTARTS   AGE
    	beego-example-1-build    0/1       ExitCode:0   0          24m
    	beego-frontend-1-879rd   1/1       Running      0          21m


6. Check the IP and port the frontend service is running on:  

		$ oc get svc


	Sample output:  

		NAME             LABELS                              SELECTOR              IP(S)           PORT(S)
    	beego-frontend   template=beego-example   name=beego-frontend   172.30.214.52   8080/TCP

In this case, the IP for frontend is 172.30.161.15 and it is on port 8080.  
*Note*: you can also get this information from the web console.


###License
This code is dedicated to the public domain to the maximum extent permitted by applicable law, pursuant to [CC0](http://creativecommons.org/publicdomain/zero/1.0/).

The sample beego application code follows the original [Apache License](https://github.com/beego/samples/blob/master/LICENSE)
