FROM postgres:14

ADD create-extension.sh /docker-entrypoint-initdb.d/
RUN chmod 755 /docker-entrypoint-initdb.d/create-extension.sh