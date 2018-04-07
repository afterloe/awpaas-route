FROM java:8-jre
MAINTAINER afterloe <lm6289511@gmail.com>

EXPOSE 8080
VOLUME /opt/ascs-soa
COPY . /opt/ascs-soa/
COPY docker-entrypoint.sh /usr/local/bin/
RUN ln -s usr/local/bin/docker-entrypoint.sh / # backwards compat

ENTRYPOINT ["docker-entrypoint.sh"]
