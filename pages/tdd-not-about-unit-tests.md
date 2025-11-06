---
layout: top-title
color: slate
align: l
---

:: title ::

# Test-Driven Development - not about Unit Tests

:: content ::

These are subjective opinions rooted in **_years of experience_**. 

- Don't test single units in isolation.
- Test how units collaborate to provide the _behaviours_ of the application.
  - Don't call a _Controller Method_
  - Send an HTTP request, verify the response, and side effects in the system.
- Only mock at layer boundaries.

## The primary measure of the quality of a test suite

- Does it provide fast feedback?
- Does it enable safe refactoring?


