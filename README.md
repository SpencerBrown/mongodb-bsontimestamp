# mongodb-bsontimestamp - command for converting MongoDB BSON timestamp values

The [BSON timestamp](http://bsonspec.org/spec.html) is a special internal type used by MongoDB replication and sharding. It is a 64-bit value, which is a pair of 32-bit numbers. The first is a timestamp, encoded in the number of seconds since the Linux epoch (00:00 UTC January 1, 1970). The second is a value, called the increment, that distinguishes between timestamps within the same second.

Sometimes it is useful to convert between a 64-bit integer representation of a BSON timestamp, and a pair of 32-bit integers representing the number of seconds and the increment. THis command-line utility does that for you.

Usage:

```bash
# Convert from 64-bit integer (in decimal) to the number of seconds and increment (in decimal)
mongodb-bsontimestamp <integer>
# Convert from a pair of numbers (in decimal) to a 64-bit integer (in decimal) representing the BSON timestamp
mongodb-bsontimestamp <integer> <integer>
```

You will find releases of this command-line utility built for Windows, Linux, and macOS on the releases page. 