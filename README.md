# DevOps Test
This is a brief guide on how to run DevOps test on your local machine.

### Prerequisites
- Docker must be installed on your machine. To download Docker, visit [docker.com](https://www.docker.com/products/docker-desktop).

### Download the Image
To download the Docker image, open a terminal and run the following command:

```shell
docker pull giovannirossini/devops-test:latest
```

### Run the Image
To run the Docker image, run the following command in your terminal:

```shell
docker run -d giovannirossini/devops-test:latest
```

### Verify the Test
To verify that the test is running, open a web browser and go to http://localhost. You should see the test running.

---

### Expected Outcome

Understanding of Linux Package Manager (Alpine): The developer should demonstrate their ability to use the Alpine package manager to install, update and manage packages within the Alpine environment.

Basic Network Understanding (Netstat): The developer should show their knowledge of using the netstat command to view network connections.

Basic Knowledge of Docker: The developer should be able to recreate the Docker container and expose its ports for external access, demonstrating their understanding of Docker's basic functionality and networking capabilities.

Decrypting a Message (using base64 -d): The developer should demonstrate their ability to use the base64 command to decode an encrypted message, showcasing their understanding of basic linux tools.

Overall, the DevOps test aims to evaluate the developer's ability to apply a combination of technical skills and knowledge in a sandbox scenario, providing valuable insights into their overall proficiency in DevOps practices.


### Resolution Steps

First you need o run the docker:

```shell
docker run -name devops-test -d giovannirossini/devops-test:latest
```

This will run the container in the background, but will not expose the application.

Second you need to find the port that the application is running on.
A Go binary code is running inside the container, so the developer must have a basic understanding of network to identify the port on which the application is running. For this step you can run `netstat` to help you identify the port:

```shell
docker exec -it devops-test sh

# Once inside the container it must to install netstat to identify the port

/opt # apk add net-tools
OK: 6 MiB in 16 packages

# Now using the netstat command you can identify the port
/opt # netstat -tulpn
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 :::8573                 :::*                    LISTEN      1/qemu-x86_64
```

Once the port is identified, it is necessary to restart the container and map the port for external access:

```shell
docker stop devops-test
docker rm devops-test
docker run -name devops-test -p 80:8573 -d giovannirossini/devops-test:latest
```

Now you can access the application on port 80 of your machine.
Upon accessing the first page, you will be presented with a message indicating which page to visit next.

```md
Congratulations you pass the first test! Now go to /challenge to get the next challenge.
```

On the `/challenge` page, a portion of the message is encrypted.

```md
Congratulations again! Now go to L2Rldm9wczIwMjM= to start your final challenge.
```

The developer must have a basic understanding of cryptography to decrypt the message.You can decoding the message using base64, this will reveal the next page to visit.

```shell
echo L2Rldm9wczIwMjM= | base64 -d
/devops2023
```

Upon accessing `/devops2023`, you will receive the final message.

```md
Congratulations you pass the last test! This is a template of a DevOps test using Docker and Go. Project: https://github.com/giovannirossini/devops-test.
```

_Note that this is a template and can be customized to meet your needs._