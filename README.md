revel-docker
============

Docker integration for Revel

Usage
-----

Add the `docker.endpoint` configuration option to `app.conf`. For example, the docker socket on the local machine:

```
docker.endpoint = unix:///var/run/docker.sock
```

Then, ask revel to init the docker client on app start:

```go
revel.OnAppStart(reveldocker.Init)
```

Embed into your controller:

```go
type App struct {
	*revel.Controller
	reveldocker.DockerController
}
```

And add the intercept:

```go
revel.InterceptMethod((*reveldocker.DockerController).Before, revel.BEFORE)
```
