FROM node:14 as build
ARG REACT_APP_ENDPOINT
ENV REACT_APP_ENDPOINT $REACT_APP_ENDPOINT
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY package.json ./
COPY yarn.lock ./
RUN npm install react-scripts@3.4.1 -g --silent
RUN yarn install
COPY . ./
RUN yarn run build

FROM nginx:stable-alpine
COPY --from=build /app/build /usr/share/nginx/html
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
