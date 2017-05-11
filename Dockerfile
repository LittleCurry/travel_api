FROM busybox
COPY resource/ /resource
CMD ["/resource/main", "-env", "prod", "-config", "/resource/api.yml"]
EXPOSE 80
