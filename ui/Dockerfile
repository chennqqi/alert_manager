
FROM node:10.12.0-alpine

# set working directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

RUN npm install -g serve

ENV PATH /usr/src/app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY . /usr/src/app
RUN npm install

RUN npm run build

ENV NODE_ENV production
# start app

CMD ["serve", "-s", "build"]