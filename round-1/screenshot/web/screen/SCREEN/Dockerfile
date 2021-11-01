FROM polinux/httpd-php:latest
COPY composer.json /
COPY php.ini /etc/
COPY tmp /data/tmp
RUN chmod 777 -R /data/tmp/php
COPY ./htdocs/ /var/www/html/
RUN chown -R apache:apache /var/www/html/
WORKDIR /
COPY --from=composer /usr/bin/composer /usr/bin/composer
RUN composer install
RUN wget https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-linux-x86_64.tar.bz2
RUN tar xvjf phantomjs-2.1.1-linux-x86_64.tar.bz2 -C /usr/local/share/
RUN ln -sf /usr/local/share/phantomjs-2.1.1-linux-x86_64/bin/phantomjs /bin/
RUN chmod 777 /bin/phantomjs
