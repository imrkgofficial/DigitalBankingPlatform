# Dockerfile for Angular frontend 
# Step 1: Build the Angular app
FROM node:16-alpine AS build

# Set the working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the rest of the Angular application
COPY . .

# Build the Angular app
RUN npm run build --prod

# Step 2: Serve the Angular app using nginx
FROM nginx:alpine

# Copy the built Angular app to nginx's public folder
COPY --from=build /app/dist/your-angular-app /usr/share/nginx/html

# Expose the port that Nginx is listening on
EXPOSE 80

# Run Nginx
CMD ["nginx", "-g", "daemon off;"]
