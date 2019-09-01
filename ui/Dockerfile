FROM node:latest

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run tsc
EXPOSE 8080
CMD ["npm", "start"]
