# Shifter

## Reasoning
While I'm developing apps I often need to receive webhook events from outside services. I don't want to expose my local machine to the internet. So, I created this tool to receive webhook events and forward them to my local machine.

You can just create a tunnel to this server and start receiving webhook events etc.

In my case, I use [Cloudflare Argo Tunnel](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps) to create a tunnel to this server. I've domain `webhook.example.com` and I've created a tunnel to this server. Now, I can use `webhook.example.com` to receive webhook events.

You can use any tunneling service to create a tunnel to this server.

## Usage

- Clone the repository
- You can check configuration in `config.json` file. You can change the configuration as per your need.
- Run the following command to bu

```shell
$ make build
```

- Run the following command to run the program (it'll start the server with the configuration provided in `config.json` file)

```shell
$ make run
```

- Run the following command to run the program with custom configuration file

```shell
$ make run config=custom_config.json
```