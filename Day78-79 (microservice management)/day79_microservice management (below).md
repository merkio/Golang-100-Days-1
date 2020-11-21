# Microservice management (below)
**@author: Davie**
**Copyright: Beijing Qianfeng Internet Technology Co., Ltd.**
## One, microservice management-service registration and query
### 1.1, define a service
The definition of the service is defined by a .json json file, which uses the json format to define the relevant content of the service to be registered. The following is an example of the json format of the service:
```
{
"service": {
"id": "firstservice",
"name": "firstservice",
"tags": ["dev"],
"port": 80,
}
}
```
### 1.2, service registration
#### 1.2.1. Create a folder where the service files are stored
```
sudo mkdir /etc/consul.d
```
Note: The .d suffix indicates the storage directory of a series of configuration files
#### 1.2.2. Write the service content as in Chapter 1 and save it as a file
```
vim firstservice.json
```
The content is the content in the json part of the chapter.
Note: Each service is separately declared in a file in the form of a json file format, and then put together in a directory. It is read when consul starts.
#### 1.2.3, save the firstservice.json file to the specified directory
```
mv firstservice.json /etc/consul.d/
```
Use the above command to move the customized firstservice.json service file to the centralized storage service directory to be started when the consul cluster is started, namely /etc/consul.d
### 1.3、Service query
#### 1.3.1, start consul
Since we have added a service, the started service is configured in the form of a configuration file. Therefore, it is necessary to specify the directory corresponding to the service configuration file when starting, as shown below:
```
consul agent -dev -config-dir /etc/consul.d/
```
The above command means: start the service according to the files in the directory registered by the service specified by -config-dir.
#### 1.3.2、Service query
The query of the service supports two ways of query: **DNS** and **HTTP**
* **a. The first type: DNS**
```
dig @127.0.0.1 -p 8600 dev.firstservice.service.consul
```
![dns service query](./img/WX20190619-145024@2x.png)
【Description:】
* 1. dev.firstservice.service.consul is a fixed format combination, the specific format is: **tag.servicename.service.consul**, that is, **tag** and **servicename** are customized when the service is created Configuration content.
* 2. The port for DNS access is 8600
* **b, the second type: HTTP**
```
curl http://localhost:8500/v1/catalog/service/firstservice
```
![http mode query](./img/WX20190619-144852@2x.png)
【Description:】
* 1. HTTP access path: host:port/version number/service/service name.
* 2. Address: Used to specify the IP address of a specific Service. By default, the agent used by the service is used.
### 1.4, register multiple services
#### 1.4.1. Write a json file for each service
For example, to create a second service, sendservice, will create a secondservice.json file as follows:
```
{
"service": {
"id": "secondservice",
"name": "secondservice",
"tags": ["dev"],
"port": 80,
}
}
```
* Use http format access:
```
curl http://localhost:8500/v1/catalog/service/secondservice
```
#### 1.4.2, write multiple services to a json file
In addition to defining each service individually as a json configuration file, consul also allows multiple service combinations to be defined in a json file. The specific method is to slightly modify the format of the json file and modify the original service to services. Modify the object type corresponding to the original service to an array type. Examples are as follows:
```
{
"services": [
{
"id": "firstservice",
"name": "firstservice",
"tags": ["dev"],
"port": 80
},
{
"id": "secondservice",
"name": "secondservice",
"tags": ["dev"],
"port": 80
}
]
}
```
Description:
* 1. Put in multiple services instead of the original service
* 2. Multiple services are placed in an array and named services
In the actual development process, there are numerous microservices. If each file is placed in one file, there will be a lot of files; and if all services are placed in one file, the file is too large and inappropriate. Therefore, in practice, the two are often used in combination. For example, assuming there are 100 microservices, put them in a 10json file, and put 10 microservices in each file.
## Two, microservice management-Docker installation and operation of consul nodes
### 2.1. Build a cluster
In a real production environment, a real consul cluster needs to be deployed. To simulate the effect of multiple cluster deployments on one machine, there are two solutions: one is to use virtual machines, and the other is to use container technology.
In this series of courses, the latter container technology is used to implement cluster construction.
### 2.2、Introduction to Docker
Docker is an open source application container engine, based on the Go language and open source following the Apache 2.0 protocol.
Docker allows developers to package their applications and dependent packages into a lightweight, portable container, and then publish to any popular Linux machine, it can also be virtualized.
Containers completely use the sandbox mechanism, and there will not be any interfaces between them, and more importantly, the performance overhead of containers is very low.
### 2.3、Docker installation
#### 2.3.1, install under Windows system
For installation under the window system, please refer to the following link document: [https://www.runoob.com/docker/windows-docker-install.html](https://www.runoob.com/docker/windows-docker- install.html)
You can install Docker under the window system according to the above link.
#### 2.3.2、Installation under MacOS
There are two ways to install under MacOS. You can also refer to the following link: [https://www.runoob.com/docker/macos-docker-install.html](https://www.runoob.com/docker/macos-docker-install.html)
### 2.4、Docker test
After Docker is installed, you can test it. You can check whether Docker is installed successfully in the form of terminal commands:
```
docker version
```
![docker version view](./img/1561086326268.jpg)
### 2.5, install consul in Docker
Docker is just a container, an environment for building a cluster. If you want to build multiple clusters, you must also install the consul environment. Install consul environment in Docker, you can use instructions to achieve. Docker installation consul instructions are as follows:
* **docker search**
```
docker search consul
```
Use the above command to view the consul image file, and output the following content in the terminal:
![Find consul image file in docker](./img/WX20190621-114548@2x.png)
* **docker pull**
```
docker pull consul
```
After querying the relevant content of the consul image through search, you can use the above docker pull command to install the consul environment.
![Docker installation consul environment](./img/WX20190621-115304@2x.png)
### 2.6, verify that Docker installs Consul successfully
You can install consul in Docker. You can use the following command to check whether the installation of consul in docker is successful:
```
docker images
```
![Check if consul is installed in docker](./img/WX20190621-134602@2x.png)
or
```
docker run consul version
```
![View consul version in docker](./img/WX20190621-115457@2x.png)
### 2.7, start a single node consul agent in Docker
After installing consul in Docker, first try to start a server node. You can start a single node in docker with the following command:
```
$ docker run -p 8500:8500/tcp consul agent -server -ui -bootstrap-expect=1 -client=0.0.0.0
```
![Docker start single node server](./img/WX20190711-155259@2x.png)
In the above command, the parameter description is as follows:
* Ports are exposed, they are: HTTP port: 8500
* -h: The corresponding node1 is the name of the node
* -server: indicates that the node type to be started is server type
* -bootstrap-expect: The parameter used to elect the leader of the server node, which means that the election starts when several server nodes are reached
In the exposed http port, it is also mapped to the port of the host. Therefore, we can access the server information in the host. such as:
* curl access HTTP interface:
```
curl localhost:8500/v1/catalog/nodes
```
![http port mapping access](./img/WX20190621-165728@2x.png)
* dig to interact with DNS interface:
```
dig @0.0.0.0 -p 8600 node1.node.consul
```
![dns port mapping access](./img/WX20190621-165754@2x.png)
In addition, we can also visit with the browser, type in the browser of the host: [http://127.0.0.1:8500](http://127.0.0.1:8500) to visit, you can view the server node Related Information.
![Host's browser access server information](./img/WX20190621-171549@2x.png)
* View the number of nodes on the host
Consul has been installed on the host, you can use the command line to interact with the Consul Agent in the container:
```
consul members
```
![Interact with nodes in docker in the host](./img/WX20190621-172012@2x.png)
## Three, microservice management--Docker builds Consul cluster
### 3.1. Build and test Consul cluster on one host
With the help of the Docker container, a server node has been started and can communicate with it.
Next, I hope to use Docker to build a consul cluster. Take starting 3 consul cluster nodes as an example:
#### 3.1.1, start the first node
When starting the first node, the -bootstrap parameter is not used, but -bootstrap-expect 3 is used. With this parameter, the node will not start and become a usable cluster until all three terminals are connected together.
```go
$ docker run -d -p 8500:8500 -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_1 consul agent -server -bootstrap -ui -node=1 -client='0.0.0.0'
```
The above parameters are explained as follows:
* ui: Means to start the Web UI manager, the default port 8500 is open, which can be accessed in the browser.
* --Then
#### 3.1.2, view node IP
We need to know the internal IP of this container. Use the following command to put this IP in the environment variable JOIN_IP.
```go
$ JOIN_IP="$(docker inspect -f'{{ .NetworkSettings.IPAddress }}' node1)"
```
You can also directly view the IP of the container with the following command:
```go
$docker inspect -f'{{ .NetworkSettings.IPAddress }}' node1
```
#### 3.1.3, start the second node
Start node2 and tell him to join node1 via $JOIN_IP:
```go
$ docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_2 consul agent -server -node=2 -join='172.17.0.2'
```
Here you need to explain the parameters:
* CONSUL_BIND_INTERFACE is a few commonly used environment variables provided to us by consul mirroring. This constant has the same effect as -bind.
* name: name the node to be started
* node: Assign an id to the minimum order node
* agent: means to start the Agent process.
* server: means to start Consul Server mode.
* join: means to join a certain cluster.
#### 3.1.4, start the third node
In the same way we start node3:
```go
$ docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_3 consul agent -server -node=3 -join='172.17.0.2'
```
Now we have a cluster with 3 nodes running on one machine. Note that the container is named based on the name of the Consul Agent.
We have not exposed any port to access this cluster, but we can use the fourth agent node in client mode (not the -server parameter). This means that he does not participate in the election but can interact with the cluster. And this client mode agent does not require disk persistence.
```go
$ docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_4 consul agent -client -node=4 -join='172.17.0.2' -client='0.0.0.0'
$ docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_5 consul agent -client -node=5 -join='172.17.0.2' -client='0.0.0.0'
$ docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_6 consul agent -client -node=5 -join='172.17.0.2' -client='0.0.0.0'
```
If the above commands can be executed successfully, it means that our cluster is successfully built.
### 3.2, view the status of the cluster
After the cluster is built, we can view the status of the nodes in the cluster through a browser or terminal command line.
#### 3.2.1. Browser:
We can access the localhost:8500 port in the browser, and we can see the following effects:
* Three serve type node cluster nodes:
![HTTP://7System error.com1.At 0.Administration Department.cloudcomputer.com/WX20190712-152703@2小.PNG](./IMG/WX20190712-152703@2小.PNG)
* All Node nodes (server and client):
![HTTP://7System error.com1.At 0.Management Department.cloudcomputer.com/WX20190712-152716@2小.PNG](./IMG/WX20190712-152716@2小.PNG)
#### 3.2.2, command line to view node status:
Execute the following commands in the terminal:
```go
$consul members
```
or it could be:
```go
$docker exec consul_server_1 consul members
```
You can see the following output effects:
![HTTP://7System error.com1.At 0.Management Department.cloudcomputer.com/WX20190712-153156@2小.PNG](./IMG/WX20190712-153156@2小.PNG)
### 3.3, stop the node
* Active container status view
Use the docker ps command to output the currently running container:
```go
$docker ps
```
![HTTP://7System Error.com1.At 0.Management Department.cloudComputer.com/WX20190712-153542@2小.PNG](./IMG/WX20190712-153542@2小.PNG)
* Stop container activity
You can use the following command to stop the container that is currently active:
```go
$docker stop containerID
```
![HTTP://7System error.com1.At 0.Management Department.cloudcomputer.com/WX20190712-153839@2小.PNG](./IMG/WX20190712-153839@2小.PNG)
If you want to stop more than one, you can separate them with spaces.
* Remove container
If you want to completely remove the started node container, you can use the rm command to achieve:
```go
$docker rm containerID
```
![HTTP://7System error.com1.At 0.Management Department.cloudcomputer.com/WX20190712-154045@2小.PNG](./IMG/WX20190712-154045@2小.PNG)
## Four, microservice management--microservice definition
### 4.1, consul common commands and options
#### 4.1.1. Commonly used commands: command
The usage form of the consul command is:
```go
consul command [option]
```
* Agent: Consul nodes are divided into two types, client and server. These two types of nodes are collectively called agent nodes.
* join: The function of this command is to join the agent to the consul cluster. When a new agent node is started, it is often necessary to specify that the node needs to be added to a specific consul cluster. At this time, use the join command to specify.
* members: List all member node information in the consul cluster, including ip, port, status, type and other information.
#### 4.1.2. Common options: option
In addition to the command command, there are also option options for developers to use. Common and frequently used options are:
* -data-dir: This option is used to specify the data directory where the agent stores the state. This is required for all agents. It is especially important for the server because they must persist the state of the cluster.
* -config-dir: This option is used to specify the location of the service configuration file and check definition. Usually designated as "a certain path/consul.d" (usually, .d represents the directory where a series of configuration files are stored)
* -config-file: Specify a configuration file to be loaded. This option can be configured multiple times, and then multiple configuration files.
* -dev: This option is used to create a server node in a development environment. Under this parameter configuration, there will be no persistence operation, that is, no data will be written to disk. The dev mode is only used in development and test environments, and cannot be used in production environments.
* -bootstrap-expect: This option is used to notify consul server type nodes and specify the number of server nodes in the cluster. This parameter is to delay the election start until all nodes are started.
* -node: The node option is used to specify the name of the node in the cluster. The name needs to be unique in the cluster. It is recommended to use the IP of the machine directly.
* -bind: This option is used to specify the IP address of the node.
* -server: This option is used to indicate that the consul node type is server type. The recommended number of servers in each data center (DC) is 3 to 5. All server nodes must be elected after joining the cluster, and the raft consensus algorithm is used to ensure the consistency of data operations.
* -client: This parameter is used to specify consul as the client node type.
* -join: English means to join. The join option is used to specify which cluster to add the node to.
* -dc: dc is the abbreviation for datacenter, this option is used to specify the dc instance that the node joins.
### 4.2, microservice definition standards and options
In addition to command line options, the definition and configuration of microservices can also be placed in files. In some cases, this may be easier, such as when using a configuration management system configuration. The configuration files are in JSON format, making them easy to read and edit by people and computers. The configuration is formatted as a single JSON object that contains the configuration.
The configuration file is not only used to set up the agent, but also to provide inspection and service definition. These configuration files can also be recognized by other software and functions. They are recorded under check configuration and service configuration respectively. Service and inspection definitions support updating during reloading.
For example, the following JSON format configuration file:
```go
{
"datacenter": "east-aws",
"data_dir": "/opt/consul",
"log_level": "INFO",
"node_name": "foobar",
"server": true,
"watches": [
{
"type": "checks",
"handler": "/usr/bin/health-check-handler.sh"
}
],
"telemetry": {
"statsite_address": "127.0.0.1:2180"
}
}
```
The above json file format is a case. There are many options for consul's json file configuration, such as:
* addresses: This configuration option is used to set the binding address. In Consul 1.0 and later, you can set these as a list of addresses to bind to. It supports binding and setting multiple types of addresses, including: dns, http, https, grpc and other four types.
* bootstrap: This configuration is equivalent to adding the -bootstrap command line flag to the command line.
* bootstrap_expect: This configuration is equivalent to adding the -bootstrap_expect command line flag to the command line.
* bind_addr: This configuration is equivalent to adding the -bind command to the command line.
* ca_file: This configuration is used to specify the directory of the CA certificate file.
* ca_path: This configuration is used to specify the overall directory of the CA certificate.
* client_addr: This configuration has the same function as the -client command in the command line.
* config_entries: Under this configuration item, it is carried out by configuring the secondary configuration item. The secondary configuration item can configure the bootstrap option.
* connect: Some configuration items about the connection are set through this configuration, which is also completed through the secondary configuration items. The supported secondary configuration items are: enabled, ca_provider, and ca_config.
* datacenter: This configuration item has the same effect as the -datacenter command in the command line.
* data_dir: This configuration item has the same effect as the -data-dir command in the command line, and is used to specify the directory where the microservice json definition file is located.
* dns_config: This option is used to configure dns related parameters.
* domain: This configuration item has the same effect as the -domain command in the command line.
* node_id: This configuration item has the same effect as the -node-id command in the command line, and is used to customize the node-id.
* node_name: This configuration item has the same effect as the -node command in the command line and is used to specify a name for the node.
* ports: This configuration item is used to configure the port number of the node, which can be configured through the secondary configuration options: dns, http, https, grpc, serf_lan, serf_wan, server and other different types of ports.
* protocol: This configuration option has the same function as the -protocol command in the command line.
As above, only some of the configuration items of the json configuration file are listed. All configuration options are explained on the official website of consul. You can visit the following link to view: [https://www.consul.io/docs/agent/ options.html](https://www.consul.io/docs/agent/options.html)
