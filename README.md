# Shifter

## Reasoning
While I'm developing apps I often need to receive webhook events from outside services. I don't want to expose my local machine to the internet. So, I created this tool to receive webhook events and forward them to my local machine.

You can just create a tunnel to this server and start receiving webhook events etc. You can shift requests to any local server running on your machine.

In my case, I use [Cloudflare Argo Tunnel](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps) to create a tunnel to this server. I've domain `example.com` and I've created a tunnel to this server. Now, I can use `example.com` to receive webhook events.

You can use any tunneling service to create a tunnel to this server.

## Usage

- Clone the repository
- You can check configuration in `config.json` file. You can change the configuration as per your need.
- Run the following command to build the program.
- So if you've a domain `example.com` you can use like this: 'https://example.com/test_webhook'
- 
### Configuration

- `port` - Port on which the server will run.
- `name` - Name of the configuration.
- `routes` - You can define multiple routes. Each route has the following properties.
    - `key` - Unique key for the route.
    - `from` - Configuration for the incoming request.
        - `method` - HTTP method for the incoming request.
    - `to` - Configuration for the outgoing request.
        - `method` - HTTP method for the outgoing request.
        - `url` - URL to which the request will be forwarded.

```json
{
  "key": "test_webhook",
  "from": {
    "method": "POST"
  },
  "to": {
    "method": "POST",
    "url": "http://localhost:8080/test"
  }
}
```



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