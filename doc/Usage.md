# How to use

Mercury is a trustless, peer to peer decentralized communication protocol that makes entities be able to securely transmit messages, verifiable credentials, and verifiable presentations with each other.

## Build binary file

1. Clone this project,
2. Use command  ``` ./make```.

## Run agent

Use CLI:

```
./agent-mercury
```

CLI parameters:

```
GLOBAL OPTIONS:
   --loglevel <level>  Set the log level to <level> (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel (default: 1)
   --rest-ip value     Set http rest ip addr default:127.0.0.1 (default: "127.0.0.1")
   --http-port value   Set http rest port default:8080 (default: "8080")
   --chain-addr value  Set block chain rpc addr default:127.0.0.1:20334 (default: "http://polaris2.ont.io:20336")
   --https-port value  Set https rest port default:8443 (default: "8443")
   --enable-https      start https restful service
   --enable-package    start package msg
   --help, -h          show help

```

By default, the agent will connect polaris (ontology testnet) for querying DID, you can change   ```chain-addr```  to connect mainnet node or you local sync node.


## Tools

We provide some CLI commands to help developers create DID and other functions.

More details can be found at: [Tools CLI](../cmd/manual.md).


## Restful API

The agent also provides restful APIs for clients.

More details can be found at: [Restful API](RestAPI_Document.md).
