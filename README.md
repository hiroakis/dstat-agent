# dstat-agent

dstat-agent. You can check your server resource via REST.

## pre-requirement

This tool uses dstat. You need to install dstat to your server where this agent works on.

* CentOS, Redhat

```
yum install dstat
```

* Debian, Ubuntu

```
apt-get install dstat
```

## Installation

```
go get github.com/hiroakis/dstat-agent
```

or clone this repository and build by yourself.

## How to use

* start agent

```
dstat-agent
```

The agent use 8888 port by default. If you would like to choose listen address as you like, you can use `-host` option and `-port` option.

* access via http

You can get dstat values as json from remote.

```
curl SERVER:8888
```

* image

![](dstat-agent.png?raw=true)

## Docker

If you would like to see demo quickly, you can also use docker. Just type following command after clone this repository.

```
docker build -t hiroakis/dstat-agent .
docker run -p 8888:8888 hiroakis/dstat-agent
curl http://localhost:8888
```

## License

MIT