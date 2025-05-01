![Laelapa/GoHome](./gohome-light.png)

*Why go hard when you can GoHome?*


A simple personal **webpage** + **blog** + **portfolio** built from the ground up using *Go*, *templ* and *TailwindCSS*. 
---

[![CodeQL](https://github.com/Laelapa/GoHome/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/Laelapa/GoHome/actions/workflows/github-code-scanning/codeql) [![CodeFactor](https://www.codefactor.io/repository/github/laelapa/gohome/badge)](https://www.codefactor.io/repository/github/laelapa/gohome) [![Go Report Card](https://goreportcard.com/badge/github.com/Laelapa/GoHome)](https://goreportcard.com/report/github.com/Laelapa/GoHome)

---
## Overview

**Laelapa/GoHome** is a Go HTTP server using the standard [net/http](https://pkg.go.dev/net/http) library, [uber/zap](https://github.com/uber-go/zap) for structured logging, and [a-h/templ](https://github.com/a-h/templ?tab=readme-ov-file) for HTML templating. The project provides a foundation for building and hosting personal websites with minimal dependencies.

You can easily adapt it to host your own website by customizing:
- the **endpoints** ( in `/internal/routes/router.go` ), 
- the **static files** ( in `/static/` ), 
- and the **content** ( in `/internal/interface/templates/` ).

## Features

### Deployment Modes

GoHome can be run in different modes to suit your current development phase or deployment needs:



- **Bakend Development Mode**: for compiling, docker building and deploying locally during development, configured with colorized, human-readable logging output for easier debugging.

    >*Use this mode when making changes to actual Go code & application logic.*

    ---
- **Frontend Development Mode with Live Reloading**: leveraging the functionality of *templ* and *TailwindCSS* to rebuild on-the-fly whenever a change is made on a template, this mode gives the developer the ability to locally host a server with live-updating content and see the changes that they make in their templates & styling be reflected in real time on their browser.

    >*Use this mode when making changes to the content & styling of the webpage*
    ---
- **Testing Mode** simulates production behavior while keeping debugging-friendly features. It goes through the exact same process of deploying in production: compiling, building the templates & css, building the docker image and deploying the containerized server in a very similar configuration to production, locally. 

    ---
- **Production Mode** is configured for online deployment with features like log sampling during high-traffic scenarios and minimal debug logs to reduce clutter in the output.

    ---

For a more complete list of features and configuration details about each mode read the [deployment modes documentation](https://github.com/Laelapa/GoHome/wiki/Deployment-Modes).

## Public Facing Measures

GoHome implements several security measures to protect against common web vulnerabilities and user-generated content issues:

- **Security Headers**: Implements headers like `X-Content-Type-Options: nosniff` to prevent MIME type sniffing.
- **Log Sanitization**: Protects against log injection attacks by sanitizing HTTP headers and other user-controlled inputs before logging.
- **Static File Protection**: Blocks access to sensitive files inside the `/static/` folder through the expandable [middleware/fileserverBlacklist.go](internal/middleware/fileserverBlacklist.go).

### Logging Features

- **Structured & Sensible Logging** of application state and incoming HTTP requests. Human-friendly colorized string output in development mode, JSON output in production/testing.

- **Header Sanitization**: Protection against log injection & flooding attacks through user-controlled HTTP headers when running in production & testing modes.

- **Easily Expandable** with standardized const log fields

For more details and directions on how to expand/modify logging features read the [logging documentation](https://github.com/Laelapa/GoHome/wiki/Logging).


## Getting Started - Quick Setup

To begin building your own website on top of this project:

### ‣ Step 1: Clone the Repository

```sh
git clone https://github.com/Laelapa/GoHome.git
cd GoHome
```
---
### ‣ Step 2: Install Dependencies

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
- If you are on a different system, find which [file](https://github.com/tailwindlabs/tailwindcss/releases) suits your operating system & architecture and, while taking care to change the filename to the one that fits your case, follow [these instructions](https://tailwindcss.com/blog/standalone-cli) to install it.

then verify the installation:

```sh
./tailwindcss --version
```

#### Install Go Dependencies

```sh
go mod download
```
---
### ‣ Step 3: Set Up the Development / Live Updating Development Environment

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
---
### ‣ Step 4: Deploy in a Docker Container Locally

When you are done with development and want to see & test the final product in a state similar to how it would be deployed containerized online, do:

```sh
make docker-test
```
This will build a docker container and launch it with some default environment variables set. You can manipulate those by modifying the Dockerfile or by running the target's commands manually.

Afterwards you can connect your terminal to the log output of the docker container with:

```sh
make docker-logs
```
---
### ‣ Step 5: Cleanup

You can stop and delete the docker container with:
```sh
make docker-stop
```
and in case you have built multiple containers without removing previous ones you can look into [docker pruning](https://docs.docker.com/engine/manage-resources/pruning/), keeping in mind that these commands can interfere with docker objects irrelevant to this project that you might have on your machine.

---
### ‣ Step 6: Create a Production Docker Image
to upload to an image registry or deploy to the service of your choice, just build it from the Dockerfile:

```sh
docker build -t your-img-name:tag .
```

> When deploying make sure to configure the environment variables on the service you are using. You can read about them in `dotenv.example` or in the [environment variables documentation](https://github.com/Laelapa/GoHome/wiki/Environment-Variables).


# Important Considerations

## Project Status

GoHome is currently a work-in-progress project that, in its current state, prioritizes public-facing security and functionality over developer experience. As I'm building this primarily to host my own personal website, current development focuses on making the visitor-facing parts secure and functional, rather than creating a polished configuration experience. Many settings are currently hardcoded in the templates and configuration options are minimal and mostly live in the code instead of configuration files (a lot of them adressed in the `TODO` & `FIXME` comments). The project is intentionally being developed with an outside-in approach, securing the perimeter first, then gradually improving the internal developer experience as it matures.

## CONFIGURATION RESPONSIBILITY

While GoHome provides [these safeguards](#public-facing-measures), the project, for now, has **limited protection against misconfiguration** by the deployer. You should exercise caution when **setting environment variables**. The application has been set up with somewhat sensible defaults but it has no protections in place against bad input yet (it will get them at some point in the future). Invalid values might lead to unexpected behavior or security issues. See the [environment variables documentation](https://github.com/Laelapa/GoHome/wiki/Environment-Variables) for information on what their default, their valid and expected values are along with how each of them affects the configuration of the application.

