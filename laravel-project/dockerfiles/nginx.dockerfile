FROM nginx:stable-alpine

WORKDIR /etc/nginx/conf.d

# use nginx.conf from env folder
COPY ./env/nginx.conf .

# rename to default.conf
RUN mv nginx.conf default.conf

WORKDIR /var/www/html

COPY src .