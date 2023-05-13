FROM fauria/vsftpd

# Install the required packages
RUN apt-get update && apt-get install -y shadow

# Add a user and set a password
RUN useradd -m ftpuser && echo "ftpuser:ftppassword" | chpasswd

# Expose the FTP port
EXPOSE 21

# Start the FTP server
CMD ["/usr/sbin/vsftpd", "/etc/vsftpd/vsftpd.conf"]