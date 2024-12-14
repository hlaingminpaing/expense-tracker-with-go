# Use the official Nginx image as the base image
FROM nginx:alpine

# Set the working directory in the container to the default nginx html directory
WORKDIR /usr/share/nginx/html

# Remove the default Nginx index.html file
RUN rm -rf ./*

# Copy your frontend files into the container's Nginx html directory
COPY . .
# COPY ./frontend/style.css ./style.css
# COPY ./frontend/app.js ./app.js

# Expose port 80 so that the web server can serve the files
EXPOSE 80

# Start the Nginx server
CMD ["nginx", "-g", "daemon off;"]
