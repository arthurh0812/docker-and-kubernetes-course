FROM node:14-alpine

WORKDIR /app

ENV PORT=8080
ENV HOST=0.0.0.0

COPY package.json .

RUN npm install

COPY . .

EXPOSE ${PORT}

CMD [ "node" , "app.js" ]