FROM alpine:latest

RUN apk add postgresql

ENV POSTGRES_PASSWORD="1234s"
ENV POSTGRES_USER="postgres"
ENV POSTGRES_PORT=5432

RUN (addgroup -S postgres && adduser -S postgres -G postgres || true)
RUN mkdir -p /var/lib/postgresql/data
RUN mkdir -p /run/postgresql/
RUN chown -R postgres:postgres /run/postgresql/
RUN chmod -R 777 /var/lib/postgresql/data
RUN chown -R postgres:postgres /var/lib/postgresql/data

USER postgres

RUN initdb /var/lib/postgresql/data
RUN  pg_ctl start -D /var/lib/postgresql/data  
RUN echo "host all  all    0.0.0.0/0  md5" >> /var/lib/postgresql/data/pg_hba.conf
RUN sed -i "s/^#listen_addresses = 'localhost'/listen_addresses = '*'/" /var/lib/postgresql/data/postgresql.conf;


EXPOSE 5432
WORKDIR /var/lib/postgresql/
COPY start.sh ./	
COPY db_init.sql ./
CMD ["sh","start.sh"]

