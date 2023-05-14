FROM postgres

ENV POSTGRES_PASSWORD=password1234
ENV POSTGRES_USER=postgres

ADD ./init.sql /docker-entrypoint-initdb.d/

ENTRYPOINT [ "docker-entrypoint.sh"]

CMD ["postgres"]