# How to Use The Knowledge Base Microservice

This project is intended to be used in tandem with the [Notification Microservice](https://github.com/bezmoradi/notification-microservice)

## Intentions

I use both this project and the [Notification Microservice](https://github.com/bezmoradi/notification-microservice) on a **daily** basis to review tips about the programming languages, frameworks, libraries, and especially insights from books I've previously read. By review, I mean reviewing the notes I have taken, the snippets which I have found useful, and all in all the tips which help me do my day-to-day job faster.

## How Do They Work?

The way that I have created this personal **Knowledge Base** is that I use a GitHub repository like [Tutorials](https://github.com/bezmoradi/tutorials) repo of mine by adding all my notes in a markdown format then it's the responsibility this repo to access [Tutorials](https://github.com/bezmoradi/tutorials) repo, fetching one document randomly then pass it to Kafka (Of course the number of documents that can be fetched is configurable).  
Then it's the [Notification Microservice](https://github.com/bezmoradi/notification-microservice)'s job to listen to Kafka for new events and as soon as a new one is fired (Let's say every 6 hours or so), that document will be reformatted to HTML then sent to my personal email.

## Installation

Before installing this repo, as mentioned previously, it heavily relies on Kafka so refer to the [Notification Microservice Kafka Installation](https://github.com/bezmoradi/notification-microservice?tab=readme-ov-file#kafka-installation)'s README file to see how to install Kafka and configure it. To clone this repo, run the following command:

```text
$ git clone git@github.com:bezmoradi/knowledge-base-microservice.git
```

Make a copy of the `.env.example` file and call it `.env` which includes the following keys:

```text
MONGODB_URL=
DATABASE=

CONTENT_COUNT=

GITHUB_TOKEN=
GITHUB_USERNAME=
GITHUB_REPO=

KAFKA_BROKER=
```

To keep a record of which files of the [Tutorials](https://github.com/bezmoradi/tutorials) repo is pulled off and sent to my email, I've used MongoDB Atlas which has a free tier (This [link](https://www.mongodb.com/database/free) walks you through the process of creating a free instance of MongoDB which is more than enough for this microservice's needs).  
After creating your Atlas database, you will be provided with a URI something similar to the following one:

```text
mongodb+srv://<username>:<password>@cluster0.z8xjp.mongodb.net/<database_name>?retryWrites=true&w=majority
```

I personally have created two databases one for development and another one for production as follows:

```text
dev_knowledge_base_microservice
prod_knowledge_base_microservice
```

At this point, we can populate the first two keys inside the `.env` file as follows:

```text
MONGODB_URL=mongodb+srv://<username>:<password>@cluster0.z8xjp.mongodb.net/dev_knowledge_base_microservice?retryWrites=true&w=majority
DATABASE=dev_knowledge_base_microservice
```

The `CONTENT_COUNT` key is the number of let's say tips/posts/documents you name it, we want to be sent to our email based on the CRONJOB schedule (I've set it to one for myself; meaning if the CRONJOB schedule is set to every six hours, I'll receive four tips from the [Tutorials](https://github.com/bezmoradi/tutorials) repo daily to review and expand my knowledge).
Next is some metadata about the repository we want to pull data from:

```text
GITHUB_TOKEN=YXxTi2GmFNGE3Ga2AUZLKLB4IvgJSBw
GITHUB_USERNAME=bezmoradi
GITHUB_REPO=knowledge-base
```

My repository, which is private, is called `knowledge-base`. As an example, if you are interested in Go programming language, you can clone my [Go Tutorials](https://github.com/bezmoradi/tutorials/tree/master/go) repository or create your own one and call it whatever you want.  
In order to get a GitHub Access Token, visit [Fine-grained personal access tokens](https://github.com/settings/tokens?type=beta) and create one. While creating the token, the minimum permissions this microservice needs to be able to interact with your repo is as follows:

-   Read access to metadata
-   Read and Write access to code  

The finalized `.env` file should be as follows:

```text
MONGODB_URL=mongodb+srv://<username>:<password>@cluster0.z8xjp.mongodb.net/dev_knowledge_base_microservice?retryWrites=true&w=majority
DATABASE=dev_knowledge_base_microservice

CONTENT_COUNT=1

GITHUB_TOKEN=YXxTi2GmFNGE3Ga2AUZLKLB4IvgJSBw
GITHUB_USERNAME=<your_github_username>
GITHUB_REPO=knowledge-base

KAFKA_BROKER=localhost:9092
```

## How to Create A Docker Container

A `Dockerfile` is also included in this repo for those (including myself) who like to run the app via Docker. First create an image:

```text
$ docker build -t knowledge-base-microservice-image
```

As the `.env` file is ignored inside the `.dockerignore` file, while creating a new container we have to pass it to docker:  

```text
$ docker run -d --env-file ./.env knowledge-base-microservice-image
```

## How to Connect The Dockerized App to Dockerized Kafka

If interested following this scenario, please visit
[How to Connect The Dockerized App to Dockerized Kafka](https://github.com/bezmoradi/notification-microservice?tab=readme-ov-file#how-to-connect-the-dockerized-app-to-dockerized-kafka) inside the Notification Microservice companion repository.

## CI/CD & Deployment

For setting up CI/CD and deployment, visit my walk-through in the other companion microservice at [CI/CD & Deployment](https://github.com/bezmoradi/notification-microservice?tab=readme-ov-file#cicd--deployment)
