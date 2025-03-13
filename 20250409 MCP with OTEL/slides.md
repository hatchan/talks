---
title: MCP with OTEL
author: Benno van den Berg
options:
  implicit_slide_ends: true
---

Who am I?
=========

<!-- speaker_note: Who am I? -->

* Benno van den Berg
* Worked in various observability roles for about 10 years
* Currently working at Fiberplane

<!-- end_slide -->

What is MCP?
============

<!-- speaker_note: MCP -->

> The Model Context Protocol (MCP) is an open protocol that enables seamless integration between LLM applications and external data sources and tools.

<!-- pause -->

```
+-----+     +-----+     +---------------------+
| LLM | --> | MCP | --> | data sources / tool |
+-----+     +-----+     +---------------------+
```

MCP
===

# Transport

- SSE
- STDIO

<!-- pause -->

# Data

- Resources
- Prompts
- Tools

Fiberplane
==========

```
                         +----------+
                         | Hono API |
                         +----------+
                              |
                              V
+-----+     +-----+     +-------------+
| LLM | --> | MCP | --> | otel-worker |
+-----+     +-----+     +-------------+
```

<!-- pause -->

```typescript
import { instrument } from "@fiberplane/hono-otel";

const app = new Hono<{ Bindings: Bindings }>();

// ...

export default instrument(app);
```

Demo
====

```
                           +-------------+
                           | Fashion API |
                           +-------------+
                                  |
                                  V
+--------+     +-----+     +-------------+
| Cursor | --> | MCP | --> | otel-worker |
+--------+     +-----+     +-------------+
```
