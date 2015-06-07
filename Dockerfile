FROM google/golang-runtime
ADD ["./","~/"]
WORDIR ~/
EXPOSE 8082
CMD ["./team-ZJU-FT"]
