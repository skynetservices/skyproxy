from ubuntu

run echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
run apt-get update
run apt-get install -y nginx
run apt-get install -y vim
run rm /etc/nginx/sites-enabled/default
add site.conf /etc/nginx/sites-enabled/
add server.crt /etc/nginx/certificates/
add server.key /etc/nginx/certificates/
expose 80 443
cmd nginx