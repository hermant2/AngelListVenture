FROM node:14.13.1-buster

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.3 -g --silent

COPY . ./

CMD ["yarn", "start"]
