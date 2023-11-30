ARG BUILDER_IMAGE="unit:1.31.0-go1.21"
ARG RUNNER_IMAGE="unit:1.31.0-go1.21"

## Phase I: Create the build image
FROM ${BUILDER_IMAGE} AS BUILDER

ARG BUILD_DIR="/go/src/{{REPOSITORY_NAME}}"
ARG DEFAULT_PORT="8989"

# install build dependencies
RUN apt-get update -y && apt-get install -y build-essential git \
    && apt-get clean && rm -f /var/lib/apt/lists/*_*

WORKDIR ${BUILD_DIR}

COPY . .

RUN go install ${BUILD_DIR}/cmd/main/main.go

# build project
RUN go build -o ${BUILD_DIR}/http-server ${BUILD_DIR}/cmd/main/main.go

## Phase II: Create the deployable image
FROM ${RUNNER_IMAGE}

ARG BUILD_DIR="/go/src/{{REPOSITORY_NAME}}"
ARG CONFIG_DIR="/docker-entrypoint.d"
ARG RUN_DIR="/tmp"

WORKDIR ${RUN_DIR}

COPY --from=BUILDER ${BUILD_DIR}/http-server .
COPY --from=BUILDER ${BUILD_DIR}/unit/config.json ${CONFIG_DIR}/http-server.json

EXPOSE ${DEFAULT_PORT}
