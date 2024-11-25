# Set the base image
FROM nginx:latest

# Copy Nginx config files from the repo to the container
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./nginx/sites-available/matcher.conf /etc/nginx/sites-available/matcher.conf

# Create a symlink for sites-enabled
RUN ln -s /etc/nginx/sites-available/matcher.conf /etc/nginx/sites-enabled/

# Expose port 80 and 443
EXPOSE 80 443

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
