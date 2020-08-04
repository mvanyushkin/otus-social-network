## stage1 as builder
#FROM node:10-alpine as builder
#
#COPY package.json package-lock.json ./
#RUN npm install && mkdir /app && mv ./node_modules ./app
#WORKDIR /app
#COPY . .
#RUN npm run build
#
#FROM nginx:alpine
#COPY ./nginx.conf /etc/nginx/nginx.conf
#RUN rm -rf /usr/share/nginx/html/*
#COPY --from=builder /app/build /usr/share/nginx/html
#
#EXPOSE 8080
#
#ENTRYPOINT ["nginx", "-g", "daemon off;"]


# Stage 0, "build-stage", based on Node.js, to build and compile the frontend
FROM tiangolo/node-frontend:10 as build-stage
WORKDIR /app
COPY package*.json /app/
RUN npm install
COPY ./ /app/
RUN npm run build
# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM nginx:1.15
COPY --from=build-stage /app/build/ /usr/share/nginx/html
# Copy the default nginx.conf provided by tiangolo/node-frontend
COPY --from=build-stage /nginx.conf /etc/nginx/conf.d/default.conf