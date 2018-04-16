# Rationale

There's a traditional solution for this kind of tasks: regular expression with capture groups. But it has numerous generic and Go-specific disadvantages:

1. Regexes are hard to read and debug.
2. Speed. While simple non-capturing regular expressions can be speedy, they quickly becomes slow as the complexity of the regular expression grows
3. They are overpowered for simple log parsing. In our experience with log processing we are not looking for patterns within the line. Usually our data is well structured and it is easier to think (and compute!) in terms of bounds and separators. And if the data is not well structured then it is a good idea to make it one, just for the sake of readability.
4. Go regular expressions are slow. Go regular expressions with group capture are even slower.
5. There are no cheap way in Go regexes what would give us a convenient way to access a group's value, we must use arrays instead of access to captured value by group name, thus it is hard for reading and comprehension.

There is another traditional approach: manual data extraction. We manually command to find a symbol or substring and pass
it or take everything before it and put into variable, it also has his share of generic disadvantages:

1. It is annoying as hell to write it
2. It can be hard to read

Still, the major advantage is:
1. It can be fast

We had severe shortage of resources at my last job, we couldn't just buy some more power, so we had no choice and were writing all this manually.
It turned out most of things to retrieve data are repetitive and we are writing nearly the same things again and again and it would be perfect to automate the process.

So, we wrote a code generator for this purpose. The code turned to be even faster than one we used to write, since we actually
were trying to reduce amount of code introducing helper abstractions what have some cost while the generator just puts raw code.

