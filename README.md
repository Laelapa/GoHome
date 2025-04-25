# Laelapa/GoHome
A simple personal **webpage** + **blog** + **portfolio** built from the ground up using *Go*, *templ* and *TailwindCSS*. 


## Overview

**Laelapa/GoHome** is a Go HTTP server using the standard [net/http](https://pkg.go.dev/net/http) library, [uber/zap](https://github.com/uber-go/zap) for structured logging, and [a-h/templ](https://github.com/a-h/templ?tab=readme-ov-file) for HTML templating. The project provides a foundation for building and hosting personal websites with minimal dependencies.

By customizing the **endpoints** ( in `/internal/routes/router.go` ), the **static files** ( in `/static/` ), and the **content** ( in `/internal/interface/templates/` ) you can easily adapt it to host your own website.

## Features

### Deployment Modes

GoHome can be run in different modes to suit your current development phase or deployment needs:



- **Development Mode**: for compiling, docker building and deploying locally during development, featuring colorized, human-readable logging output for easier debugging.

    >*Use this mode when making changes to actual Go code & logic.*

    ---
- **Development Mode with Live Reloading**: leveraging the functionality of *templ* and *TailwindCSS* to rebuild on-the-fly whenever a change is made on a template, this mode gives the developer the ability to locally host a server with live-updating content and see the changes that they make in their templates & styling be reflected in real time on their browser.

    >*Use this mode when making changes to the content & styling of the webpage, essentially when dealing with the frontend.*
    ---
- **Testing Mode**: simulates production behavior while keeping debugging-friendly features. Essentially goes through the exact same process of deploying in production: compiling, building the templates & css, building the docker container and deploying the server in a very similar configuration to production, locally. 

    ---
- **Production Mode**: is configured for online deployment with features like log sampling during high-traffic scenarios turned on and minimal debug logs to reduce clutter in the output.

    ---

For a more complete list of features and configuration details about each mode read the [deployment modes documentation](https://github.com/Laelapa/GoHome/wiki/Deployment-Modes).

### Logging Features

- **Structured & Sensible Logging**: of application state and incoming HTTP requests. Human-friendly colorized string output in development mode, JSON output in production/testing.

- **Header Sanitization**: Protection against log injection & flooding attacks through user-controlled HTTP headers when running in production & testing modes.

- **Easily Expandable**: with standardized const log fields

For more details and directions on how to expand/modify logging features read the [logging documentation](https://github.com/Laelapa/GoHome/wiki/Logging).


## Getting Started - Quick Setup

To begin building your own website on top of this project:

### Step 1: Clone the Repository

```sh
git clone https://github.com/Laelapa/GoHome.git
cd GoHome
```

### Step 2: Install Dependencies

 The included **Makefile** can make(*hhhehehe*) the setup quite simple. If you dont have [Make](https://www.gnu.org/software/make/) installed you can just execute the relevant target's commands manually.
#### Install templ CLI
with:

```sh
go install github.com/a-h/templ/cmd/templ@latest
```

verify the installation:
```
templ --version
```

#### Install TailwindCSS CLI

- For linux-x64 systems:

    ```sh
    make install-tailwind-linux
    ```
- If you are on a different system, find which [file](https://github.com/tailwindlabs/tailwindcss/releases) suits your operating system & architecture and follow [these instructions](https://tailwindcss.com/blog/standalone-cli) to install it, taking care to change the filename to the one that fits your case.

then verify the installation:

```sh
./tailwindcss --version
```

#### Install Go Dependencies

```sh
go mod download
```

### Step 3: Set Up the Development / Live Updating Development Environment

At this point you are all set to dive into development.

#### Compile and run once

If you are editing Go code, application & server logic and want to  compile to check the effects of your interventions on the application you can:

```sh
make run
```
This will compile the Go code, build the templ templates & tailwind's css, and launch a server at http://localhost:8080

#### Development with Live Reloading

For front-end development, to be able to see how your edits reflect on the website live on a browser, do:

```sh
make run-watch
```

### Step 4: Deploy in a Docker Container Locally

When you are done with development and want to see & test the final product in a state similar to how it would be deployed containerized online, do:

```sh
make docker-test
```
This will build a docker container and launch it with some default environment variables set. You can manipulate those by modifying the Dockerfile or by running the target's commands manually.

Afterwards you can connect your terminal to the log output of the docker container with:

```sh
make docker-logs
```

### Step 5: Cleanup

You can stop and delete the docker container with:
```sh
make docker-stop
```
and in case you have built multiple containers without removing previous ones you can look into [docker pruning](https://docs.docker.com/engine/manage-resources/pruning/), keeping in mind that these commands can interfere with docker objects irrelevant to this project that you might have on your machine.

### Step 6: Create a Production Docker Image
to upload to an image registry or deploy to the service of your choice, just build it from the Dockerfile:

```sh
docker build -t your-img-name:tag .
```

> When deploying make sure to configure the environment variables on the service you are using. You can read about them in `dotenv.example` or in the [environment variables documentation](https://github.com/Laelapa/GoHome/wiki/Environment-Variables).

## NOTE on fly.io files

This repository includes configuration files for deploying to [fly.io](https://fly.io/):

- `fly.toml`: Contains Fly.io application configuration,
- `.github/workflows/fly-deploy.yml`: GitHub Actions workflow for Fly.io CD.



**Note:**  
These files are used for my personal deployment setup. They are **not required** for general usage or local development of this project. You can safely ignore/delete them and you don't even need to keep them around as a reference in case you are interested in deploying to fly.io yourself. They are both auto-generated by their [command-line tool](https://github.com/superfly/flyctl) as part of the deployment process.