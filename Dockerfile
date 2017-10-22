# Copyright 2015 The Kubernetes Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# ----- Go Dev Image ------
#
FROM golang:1.9 AS godev

# set working directory
RUN mkdir -p /go/src/github.com/sdminonne/workflow-controller
WORKDIR /go/src/github.com/sdminonne/workflow-controller

# copy sources
COPY . .

# set entrypoint to bash
ENTRYPOINT ["/bin/bash"]

#
# ------ Go Test Runner ------
#
FROM godev AS tester
# run test and calculate coverage
RUN make test
# upload coverage reports to Codecov.io: pass CODECOV_TOKEN as build-arg
ARG CODECOV_TOKEN
# default codecov bash uploader (sometimes it's worth to use GitHub version or custom one, to avoid bugs)
ARG CODECOV_BASH_URL=https://codecov.io/bash
# set Codecov expected env
ARG VCS_COMMIT_ID
ARG VCS_BRANCH_NAME
ARG VCS_SLUG
ARG CI_BUILD_URL
ARG CI_BUILD_ID
RUN if [ "$CODECOV_TOKEN" != "" ]; then curl -s $CODECOV_BASH_URL | bash -s; fi

#
# ------ Go Builder ------
#
FROM godev AS builder

# build workflow-controller binary
RUN make

#
# ------ Workflow Controller image ------
#
FROM busybox
COPY --from=builder /go/src/github.com/sdminonne/workflow-controller/workflow-controller /workflow-controller
ENTRYPOINT ["/workflow-controller"]
CMD ["--v=2"]

ARG VCS_COMMIT_ID=dev
LABEL org.label-schema.vcs-ref=$VCS_COMMIT_ID \
      org.label-schema.vcs-url="https://github.com/codefresh-io/workflow-controller" \
      org.label-schema.description="" \
      org.label-schema.vendor="Open Source" \
      org.label-schema.url="https://github.com/codefresh-io/worflow-controller" \
      org.label-schema.version="0.1.1" \
      org.label-schema.docker.cmd="docker run -d --rm -v $HOME/.kube/config:/home/root/.kube/config codefreshio/workflow-controller --kubeconfig=/home/root/.kube/config --v=2" \
      org.label-schema.docker.cmd.help="docker run -it --rm codefreshio/workflow-controller --help"