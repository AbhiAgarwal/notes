General notes

- Don't use `const` in strict mode.
- There exists a harmony mode:
	- `node --harmony`
- Any unhandled exception thrown in Node will halt the process. The exception output shows the offending file, and the line number and position of the exception.

From this talk: https://www.youtube.com/watch?v=jsiqvXi3qSA

Common uses for Node

- Template generation
- REST API calls
- Static content concatenation
- Metrics reporting

Common Problems

- Node applications are becoming more CPU bound
- Delivering completely static content - crazy
  - Don't deliver static content through Node.. you have CDNs!
  - HTTP parser in Node is not cheap.. 
  - Requesting the favicon is not cheap
  - Wasting CPU usage
- External API calls are the bottleneck
  - If 5% is being used by Node even when 1,000s of request
    - then you know it's not your NodeJS but an external API
    - you can't really do anything about this
- Simulataneous requests are allowed until process falters
  - Do performance metrics!! Find a sweet spot and find limits of your application
