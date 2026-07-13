FROM node:23-alpine

WORKDIR /app

# Copy package files and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the rest of the application code
COPY . .

# Expose the port Vite runs on
EXPOSE 5173

# Start the Vite development server and expose it to the network
CMD ["npm", "run", "dev", "--", "--host"]
