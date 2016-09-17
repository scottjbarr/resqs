# Resqs

## Why?

I want a simple, but reliable queue service. Redis is awesome and does
most of what I need, but does not have reliable delivery baked in. For
example, a consumer reads a messages from a queue endpoint but
crashes. Most implementations I have seen so far do not deal with this
very well.

## It does what, now?

Nothing yet.

## License

```json
{
    "id": "c502539c-7bfd-11e6-b717-bf6757fa14d5",
    "attempt": 0,
    "max_attempts": 200,
    "created_at": "2016-09-16 21:16:53+10:00",
    "last_attempt": "",
    "payload": "{\"id\": \"123\"}"
}
```
