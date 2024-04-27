FROM node:21-alpine3.18

WORKDIR /app

# Install dependencies based on the preferred package manager
COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml* ./
RUN npm install

ENV NODE_ENV production

COPY . .
RUN npm run build

CMD npm run start