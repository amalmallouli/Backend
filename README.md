# Backend

To start phpmyadmin

sudo service apache2 restart
 sudo service nginx stop
sudo service nginx disable

Edit: Do this first

sudo -H gedit /etc/apache2/apache2.conf
add this line somewhere

Include /etc/phpmyadmin/apache.conf
and finally restart apache.

sudo service apache2 restart

sudo apt get update

to run this project
change .env name Database and password

then RUN :
go run server.go
