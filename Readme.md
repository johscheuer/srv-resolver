# SRV Resolver

A small tool to read an SRV entry over the [Mesos-DNS REST API](http://mesosphere.github.io/mesos-dns/docs/http.html) and generate a redis.conf file to allow the redis-slaves to connect to the redis-master.
