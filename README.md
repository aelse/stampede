# stampede

Stampede is a cache stampede avoidance library implementing the XFetch algorithm.
Use this package to optimally determine when any cacheable item should be refreshed.

## What is a cache stampede?

A typical pattern for accessing an item from cache looks something like:

```
GetFromCache(key, ttl):
    item = cache.get(key)
    if !item:
        item = ExpensiveGenerateValue()
        cache.set(item, ttl)
    return item
```

The problem with this pattern is that when a cached item expires it results in many
expensive operations as threads/processes/systems independently repopulate the
cache. This is known as a [cache stampede](https://en.wikipedia.org/wiki/Cache_stampede).
One mitigation is to use locking around the expensive operation, but this
has its own hazards particularly in distributed systems.

A better solution is probabilistic cache regeneration. As the item's expiration time
approaches it becomes increasingly likely that the cached item will be refreshed.
This means that the cache will be refreshed earlier than required but the overall
efficiency is improved as fewer calls are made to the expensive generation code.

We implement a variation of the XFetch algorithm to determine when an item should be refreshed.

## Why do I need this?

Stampede will help you easily choose an optimal time to refresh items in your cache,
whatever the purpose of the cache may be.

You can use this any time you have some value that is accessed concurrently by
multiple threads/processes/systems. This might be as simple as an authentication
token in your running process that requires regeneration periodically, or maybe
frequently used data in memcached or redis shared by multiple systems.

## More information

This package uses ideas from a paper titled [Optimal Probabilistic Cache Stampede Prevention](https://dl.acm.org/citation.cfm?id=2757813).

* Vattani, Chierichetti, Lowenstein, 'Optimal Probabilistic Cache Stampede Prevention', 2015 8(8) _Proceedings of the VLDB Endowment_ pp886-897.
