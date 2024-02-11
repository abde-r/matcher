FROM node:16

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install --production

# Install additional dependencies
RUN npm install bcrypt jsonwebtoken cookie-parser

# Copy the rest of the application code
COPY . .

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["npm", "start"]
