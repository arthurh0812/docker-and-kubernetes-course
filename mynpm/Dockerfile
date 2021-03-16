FROM node:14-alpine

RUN userdel -r node

ARG UID

ARG GID

RUN addgroup --gid ${GID} user

RUN adduser --disable-password --gecos '' --uid ${UID} --gid ${GID} user

WORKDIR /app

ENTRYPOINT [ "npm" ]