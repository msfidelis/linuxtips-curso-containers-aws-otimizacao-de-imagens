FROM node:22

RUN apt-get update && apt-get install -y \
    curl \
    git \
    vim

RUN apt-get install -y nodejs
RUN apt-get install -y build-essential

RUN rm -rf /var/lib/apt/lists/*

ADD . /app

WORKDIR /app

RUN npm install

EXPOSE 3000

CMD ["node", "index.js"]
