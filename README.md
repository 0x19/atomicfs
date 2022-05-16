# AtomicFS

Decentralized General Storage Protocol.

Basically, I'd like to build the blockchain database for storage. Similar ideas are Swarn or IPFS. I'm targeting from following practice to obtain a deeper understanding into blockchain protocol internals.

This is a proof of concept, study of my own. 

**NOT FOR A PRODUCTION USE. NOT FOR ANY SORT OF USE RIGHT NOW.**

Eventually, once built to any sane level will make notice here and change the project scope of use.


## Project Structure

In order not to reinvent the wheel at this moment, going to use the https://github.com/golang-standards/project-layout as a base pointer towards how the architecture is going to be built with few caviats that I'm going to explain here in the future.

### Entrypoint

Path where main go file is located. I trully dislike having main.go in cmd/ path as it just makes whole cmd folder look bad. Therefore, I've choosen to name entrypoint path where main will reside and will be used to spawn and build the entire project. Borrowed naming idea from Docker entrypoint.

### Protos

All of the protocol buffer definitions will go inside of this path. Within, there's as well Makefile that is being called by the main root folder Makefile. It's basically a helper script to get us going with compiling the protosets and protos.

## Table Of Content

- [Cli]
  - [Wallet]


## External Package List

List of the external projects that AtomicFS is using. Many thanks to all of incredible people that
made those projects. Won't list all of the projects from the beginning but in the future, perhaps all will be here.

- [Zap]
- [Cobra]
- [PromptUI]
- [Viper]
- [Go Ethereum]
- [Protoc]
- [FsNotify]
- [gRPC]


[Zap]: <https://pkg.go.dev/go.uber.org/zap>
[Cobra]: <https://github.com/spf13/cobra>
[PromptUI]: <https://github.com/manifoldco/promptui>
[Go Ethereum]: <https://github.com/ethereum/go-ethereum>
[Viper]: <https://github.com/spf13/viper>
[Protoc]: <https://github.com/protocolbuffers/protobuf>
[FsNotify]: <https://github.com/fsnotify/fsnotify>
[gRPC]: <https://grpc.io/>
[Cli]: <https://github.com/0x19/atomicfs/blob/main/docs/readme/cli.md>
[Wallet]: <https://github.com/0x19/atomicfs/blob/main/docs/readme/cli-wallet.md>