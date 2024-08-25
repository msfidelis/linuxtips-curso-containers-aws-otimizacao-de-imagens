FROM node:22

RUN apt-get update && apt-get install -y \
    curl \
    nodejs \
    build-essential \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY package*.json ./

RUN npm install --only=production

COPY . .

EXPOSE 3000

CMD ["node", "index.js"]
