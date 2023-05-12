FROM fauria/vsftpd

# Add a user and set a password
RUN echo "ftpuser:ftppassword" | chpasswd

# Expose the FTP port
EXPOSE 21

# Start the FTP server
CMD ["/usr/sbin/vsftpd", "/etc/vsftpd/vsftpd.conf"]
