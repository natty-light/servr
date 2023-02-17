FROM node:lts-alpine

WORKDIR /usr/src/web

COPY . .

RUN yarn install && yarn build

EXPOSE 5000

CMD [ "node", "build/index.js" ]