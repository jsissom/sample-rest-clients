FROM perl:5.20
RUN perl -MCPAN -e 'install JSON'
RUN perl -MCPAN -e 'install REST::Client'
COPY . /usr/src/myapp
WORKDIR /usr/src/myapp
CMD [ "perl", "./close.pl" ]

