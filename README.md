# Figment

![Figment](https://vignette1.wikia.nocookie.net/disney/images/7/77/Figment.gif/revision/latest?cb=20140727213403)

**Figment** (working title) is a **caching package** for arbitrary data in Go.

## Motivation

There are a lot of caching packages available that focus on minimal memory allocation,
speed, and expiration policies. Figment aims to supply other caching policies beyond
expiration like stale while revalidate and stale if error.

## Features

- Standard Get, Set,  Delete, and List inferface
- Expiration and eviction policies to remove items after a given duration
- [Stale while revalidate and stale if error](https://tools.ietf.org/html/rfc5861) policies not specific to HTTP
- Self warming items on cache misses
- Self revalidating stale items
- Backup and restore cache from file (maybe)