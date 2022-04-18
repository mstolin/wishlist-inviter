FROM postgres:14.2-alpine

COPY load-extensions.sh /docker-entrypoint-initdb.d/
RUN chmod 755 /docker-entrypoint-initdb.d/load-extensions.sh
