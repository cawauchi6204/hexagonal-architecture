# set up the container.
FROM mysql:8.0-debian

# set utf-8 encoding and mysql root password.
ENV LANG=ja_JP.UTF-8
ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=db
ENV MYSQL_USER=user

# copy the init sql file to the init dir.
COPY init_db.sql /docker-entrypoint-initdb.d/