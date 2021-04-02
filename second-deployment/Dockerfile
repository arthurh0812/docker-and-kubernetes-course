FROM node:14-alpine

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

RUN npm run build

ENV PORT=8080

EXPOSE ${PORT}

CMD [ "npm", "run", "start" ]