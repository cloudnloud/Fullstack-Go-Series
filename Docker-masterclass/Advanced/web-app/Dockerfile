FROM node:16-alpine

WORKDIR /app/web-app

COPY ./package*.json ./

RUN npm install

COPY ./ .

RUN npm run build

CMD ["npm","start"]