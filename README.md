# Tacklebox

**Tacklebox** is a self-hosted solution for receiving an analyzing webhooks. It works similarly to services like RequestBin, but puts you in complete control of your data.

This repository contains the code for the API server and UI. Also included is a `docker-compose` file for easy deployment - just edit the `.env_template`, rename it to `.env`, and do `docker-compose up -d`.

For a live demo, [click here!](https://tacklebox.willbaker.dev/)

## Basic Usage

- Deploy the application per the instructions above
- Access the UI using your browser at the host you specified in the `.env` file
- Create a user and sign in
- Create new bins to serve as endpoints for webhooks
- Analyze the event data

## Components

Tacklebox uses both a relational database (Postgres) and a document database (MongoDB) for data persistence. The document database allows for hook event data to be stored in a flexible way, without concern for the size of the payload. The relational database represents the associations between the users, bins, and hooks.

The React/Redux UI is served via an nginx webserver out of a docker container. The API server is written in Go and also lives in a docker container. Caddy is used as a reverse proxy in front of the API and UI containers, providing automatic TLS certificate generation and termation, as well as proxying the requests to the correct endpoint.
