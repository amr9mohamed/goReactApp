FROM node:alpine

# Add Maintainer info
LABEL maintainer="Amr Mohamed <amrmohamed53@gmail.com>"

# Set the current working directory inside the container 
WORKDIR /app

# Copy package.json and yarn.lock files 
COPY package.json yarn.lock ./

RUN ["yarn", "inistall"]

COPY . .
CMD [ "yarn", "build" ]

