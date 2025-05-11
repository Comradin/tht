# Thesis
Many applications follow a bad logging practice. Lots of lines of
log output are written to StdOut on application startup. This was
and even is an okay practice for long running server processes,
but it has the potential to become a problem for short lived and
often invoked services.

With the advent of the modern SaaS platforms for observability like
Dynatrace, Datadog, or Grafana these practices can quickly cost
large amounts of money.

# Simulation
For demonstration purposes this `cli` will write such a typical
startup log and propose possible solutions for for better logging
practices.
