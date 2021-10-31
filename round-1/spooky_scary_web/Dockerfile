# Run with --restart=always
FROM node:lts

WORKDIR /spookyweb

COPY package*.json ./

RUN npm ci

COPY . .

EXPOSE 9010

CMD [ "node", "app.js" ]