Contributing
============

Design choices
--------------
* *get* functions must return values only of `string`, `int` and `struct`
* *get* functions who return a `struct` include only `string` and `int` fields.
* *get* functions return `int` values or `struct` field only for percentages values (to be used by the gauges).
* *get* functions return an empty string or 0 value if errors have occured.
