FROM node:18 AS builder
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

# Development stage
FROM node:18 AS dev
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
EXPOSE 5000
CMD ["npm", "run", "dev"]

# Production stage
# FROM node:18-slim AS prod
# WORKDIR /app
# COPY --from=builder /app/public ./public
# COPY --from=builder /app/package.json ./
# RUN npm install --production
# RUN npm install -g serve
# EXPOSE 5000
# CMD ["serve", "-s", "public", "-l", "5000"]