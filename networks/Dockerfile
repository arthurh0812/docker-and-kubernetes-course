FROM node:15.1-alpine

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

CMD [ "node" , "app.js" ]