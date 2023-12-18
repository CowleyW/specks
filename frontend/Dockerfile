FROM node:20.10.0-alpine AS build

WORKDIR /app

COPY . .

RUN npm ci
RUN npm run build --omit=dev


FROM nginx:latest

COPY --from=build /app/dist/frontend/browser /usr/share/nginx/html
COPY /nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
