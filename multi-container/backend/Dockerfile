FROM node:15.1-alpine

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

ARG DEFAULT_PORT=80

ENV PORT=${DEFAULT_PORT}

ENV MONGO_INITDB_ROOT_USERNAME=root

ENV MONGO_INITDB_ROOT_PASSWORD=secret

EXPOSE ${PORT}

VOLUME [ "/app/node_modules" ]

CMD [ "npm", "run", "start" ]