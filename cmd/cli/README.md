# Thesis
Many applications follow a bad logging practice. Lots of lines of
log output are written to StdOut on application startup. This was
and even is an okay practice for long running server processes,
but it has the potential to become a problem for short lived and
often invoked services.

With the advent of the modern SaaS platforms for observability like
Dynatrace, Datadog, or Grafana these practices can quickly cost
large amounts of money.

# Not so good logging practices
For demonstration purposes this `cli` will write such a typical
startup log and propose possible solutions for for better logging
practices.

```bash
❯ ./cli
Initiating Startup Sequence

Powering up zero-point energy core
Stabilizing quantum vacuum fluctuations
Unfolding subspace emitter field
Subspace emitter field fully unfolded
Subspace emitter field is now active
...
```

This is an example of a typical application startup log intended
for huamn consumption. The information is not machine readable by
default and the purpose is to inform the user about the current state
of the application. As described above, this can be an okay practice
for long running server processes, or when it is necessary to output
information to the user.

However, this is not the case for short lived and often invoked
services. Another non beneficial side effect of a logging practice like
this is fact, that these logs are not easily machine readable.
