### Docker Masterclass - Basic


#### Prerequisties

```
Docker installation

Please check this link : https://docs.docker.com/engine/install/

Install Docker based on your system configurations


docker --version

```


#### Docker commands

**List out all the images**
```
docker images
```

**List out all containers**

```
docker ps -a
```

**Search images**
```
docker search centos
```

**Pull image**

```
docker pull centos
```


**Build Docker images**

```
docker build -t <image-name> <path-to-dockerfile>
```

**Running Docker Containers**

```
docker run <image-name>
```

You can also specify options to customize the container, such as:

* `-p <host-port>:<container-port>`: maps a host port to a container port
* `-v <host-path>:<container-path>`: mounts a host directory as a volume in the container
* `-e <env-variable>=<value>`: sets an environment variable in the container


**Stop Container**

```
docker stop <container-id>
```

**Remove Container**

```
docker rm <container-id>
```


**Inspect Container**

```
docker container inspect 9abcc2cfa22e
```

**Inspect Image**

```
docker image inspect nginx
```

**Get into Docker Container file system**

```
docker exec -it 4471 sh
```