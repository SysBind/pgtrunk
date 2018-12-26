FROM postgres:11-alpine

MAINTAINER Asaf Ohayon <asaf@sysbind.co.il>

ENV PGPOOL_VERSION=4.0.2
ENV PGPOOL_SHA256 a9324dc84e63961476cd32e74e66b6fdebc5ec593942f8710a688eb88e50dcc1

RUN set -ex \
	\
	&& apk add --no-cache --virtual .fetch-deps \
		tar \
    && wget -O pgpool-II.tar.gz "http://www.pgpool.net/mediawiki/images/pgpool-II-$PGPOOL_VERSION.tar.gz" \
    && echo "$PGPOOL_SHA256 *pgpool-II.tar.gz" | sha256sum -c - \
    && mkdir -p /usr/src/pgpool-II \
    && tar \
    		--extract \
    		--file pgpool-II.tar.gz \
    		--directory /usr/src/pgpool-II \
    		--strip-components 1 \
    && rm pgpool-II.tar.gz \
    && apk add --no-cache --virtual .build-deps \
       gcc \
       libc-dev \
       linux-headers \
       make

COPY pgpool/fix_compile_alpine38.patch /

RUN set -ex \
    \
    && cd /usr/src/pgpool-II \
    && patch -p1 < /fix_compile_alpine38.patch \
    && ./configure \
    && make \
    && make install \
    && make -C src/sql/pgpool-recovery \
    && make -C src/sql/pgpool-recovery  install \
    && apk del .fetch-deps .build-deps \
    && rm /*.patch \
    && rm -rf /usr/src/pgpool-II


FROM golang
WORKDIR /go/src/github.com/sysbind/pgtrunk
COPY . .
RUN go get -d -v golang.org/x/sys/unix github.com/spf13/cobra/cobra
RUN CGO_ENABLED=0 GOOS=linux go build -a .

# Final image, official postgres + pgpool's pcp_* binaries and libs + pgtrunk executable
FROM postgres:11-alpine

COPY --from=0 /usr/local/bin/pcp_* /usr/local/bin/
COPY --from=0 /usr/local/lib/libpcp.so.1.0.0 /usr/local/lib/
RUN ln -s /usr/local/lib/libpcp.so.1.0.0 //usr/local/lib/libpcp.so.1 \
    && ln -s /usr/local/lib/libpcp.so.1.0.0 /usr/local/lib/libpcp.so

COPY --from=1 /go/src/github.com/sysbind/pgtrunk/pgtrunk /usr/local/bin/

COPY docker-entrypoint-pgtrunk.sh /usr/local/bin/

ENTRYPOINT ["docker-entrypoint-pgtrunk.sh"]

CMD ["postgres"]
