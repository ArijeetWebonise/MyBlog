FROM ubuntu
WORKDIR /application

RUN apt-get update && apt-get install gnupg curl -y
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt-get update && apt-get install yarn -y

CMD [ "yarn", "start" ]
