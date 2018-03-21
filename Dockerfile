FROM scratch
ADD bin/devopsdays-cli_*_linux_amd64 /devopsdays-cli
CMD ["/devopsdays-cli"]
