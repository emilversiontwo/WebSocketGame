FROM node:25.8.2
LABEL authors="soupik"

RUN npm install -g http-server

WORKDIR /home/node/app

COPY ./client .

RUN npm install

RUN npm i vite@latest -D

EXPOSE 3000

#ENTRYPOINT ["top", "-b"]