FROM node:boron-alpine

ARG VERSION
ENV VERSION ${VERSION}

RUN apk add --no-cache make gcc g++ python

RUN yarn global add webpack

RUN mkdir -p /app
WORKDIR /app

COPY package.json /app/
COPY yarn.lock /app/
RUN yarn install

RUN apk del make gcc g++ python

COPY . /app

ENV NODE_ENV production
RUN yarn run build

EXPOSE 8080 8443

CMD [ "yarn", "run", "server" ]
