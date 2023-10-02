# DocumentAI Expense Bot

## Pre requisites
- you should have a gloud user credentials file generated after performing authentication using `gloud auth login`, it should be located at `~/.config/gcloud/application_default_credentials.json`

- You either should have mongodb installed or docker if you want to use mongo driver to store documents if not you can use memory database which can be defined from environment variables

- `.env` file with data, see example in `.env.example` file


## Usage

```shell
make install # installs the dependencies defined in go.mod

make mongo # creates a mongodb docker container (obviously needs docker to be installed)

make env # creates a .env file based on .env.example

make test # you've guessed it, tests *_test.go files

make run # runs the project

```

## How can this be improved ?

To make this project ready to be deployed in a production environment it needs some adjustments and obviously more test coverage. In prod environment usually we would have a defined storage such as S3, NFS or something among those lines, perhaps even both. I think it would make sense to implement this logic using pub sub mechanism, for example we can upload documents to the service which then would publish an event to NATS and then the second service that listens to the upload events would replicate this data to different storages and also trigger scanning service so that the json is also stored along with the original documents.

It would be also nice to have prometheus metrics for those microservices after they are split so that we keep track of metrics and use tools then to visualize those metrics and monitor them.

Perhaps it is a good idea to also write a different transport layers if we talk about the microservices, we could have gRPC communication between some internal microservices and it makes sense to have both HTTP and gRPC in place so that it could be used by all services that implement those transport layers

## Notes

For sake of this task I went with the simple flat file structure which I think is fine but its purely subjective. I've used interfaces almost everywhere so that if we would have multiple implementations of the logic but with different drivers,configs or etc.

Regarding the unit tests, I did not spend too much time to have 100% coverage becuase this task is for demonstration purposes.

Logging is another thing to point out, currently I am using the logger provided by echo but would be better to have `slog` or `uber-go/zap` for structurized logging

### Author: Davit.K