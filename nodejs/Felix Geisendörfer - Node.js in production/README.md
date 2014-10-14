http://vimeo.com/43380478

# Notes

- `stream-cache` npm module
- `knox` s3
- `spawn` in child_processes
- `sudo port install ImageMagick` if not have ImageMagick

# Staying alive

- Upstart, Monit, etc.
- Useful during run-time errors, etc.
- Detects process crashes, restarts them
- Make sure to log this!!

# Metrics

- node-measured, etc.
- Graphite, Librato, etc.
	- http://graphite.wikidot.com/
- Measure early, measure often
- http://12factor.net/

# NodeJS good for

- long running workloads (web sockets)
- I/O bound workloads
	- streaming connections
- Code sharing between server/client