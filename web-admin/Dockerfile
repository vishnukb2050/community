# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files
COPY package.json package-lock.json ./
RUN npm install

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Final stage
FROM node:20-alpine

WORKDIR /app

# Install serve to run the static files
RUN npm install -g serve

# Copy build files from builder
COPY --from=builder /app/dist ./dist

EXPOSE 3000

# Start the application
CMD ["serve", "-s", "dist", "-l", "3000"]
