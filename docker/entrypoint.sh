#!/bin/sh
#
# Authors:
#           TI INNOVACION - IJAM
#

java $JAVA_OPTS -Djavax.net.ssl.trustStorePassword=changeit -Djava.security.egd=file:/dev/./urandom -jar /app.jar
