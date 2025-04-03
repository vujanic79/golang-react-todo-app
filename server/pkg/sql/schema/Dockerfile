FROM gomicro/goose

WORKDIR /pkg/sql/schema

ADD *.sql ./
ADD migrations.sh ./

RUN chmod +x /pkg/sql/schema/migrations.sh

ENTRYPOINT ["/pkg/sql/schema/migrations.sh"]
