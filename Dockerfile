FROM kitabisa/debian-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/kitabisa/golang-poc"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/golang-poc

WORKDIR /opt/golang-poc

COPY ./bin/golang-poc /opt/golang-poc/
COPY ./migrations /opt/golang-poc/migrations/

RUN chmod +x /opt/golang-poc/golang-poc

# Create appuser
RUN adduser --disabled-password --gecos '' golang-poc
USER golang-poc

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/golang-poc/golang-poc"]
