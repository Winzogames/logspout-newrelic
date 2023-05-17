ARG LOGSPOUT_VERSION
FROM gliderlabs/logspout:$LOGSPOUT_VERSION
ENV LOGSPOUT=ignore
COPY modules.go /src/