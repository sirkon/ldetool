# Rationale

There's a traditional solution for this kind of tasks: regular expression with capture groups. But it has numerous generic and Go-specific disadvantages:

1. Regexes are hard to read and debug.
2. Regexes don't help with string → numeric conversions
3. You can't generally say what caused an error with regexes, so +1 to debug complexity.
4. Speed. While simple non-capturing regular expressions can be speedy, they quickly becomes slow as the complexity of the regular expression grows
5. They are overpowered for simple log parsing. In our experience with log processing we are not looking for patterns within the line. Usually our data is well structured and it is easier to think (and compute!) in terms of bounds and separators. And if the data is not well structured then it is a good idea to make it one, just for the sake of readability.
6. Go regular expressions are slow. Go regular expressions with group capture are even slower.
7. There are no cheap way in Go regexes what would give us a convenient way to access a group's value, we must use arrays instead of access to captured value by group name, thus it is hard for reading and comprehension.

There is another traditional approach: manual data extraction. We manually command to find a symbol or substring and pass
it or take everything before it and put into variable, it also has his share of generic disadvantages:

1. It is annoying as hell to write it
2. It can be hard to read

Still, the major advantage is:
1. It can be fast

In my previous job we had lack of funding for hardware while the data was constantly growing. We moved from Python into Go to deal with it. And a prototype of this tool was created to simplify our tasks, as we were writing line data decomposers manually. In the end our parsing scripts became even faster than they were before, as the code generator is not afraid of making even the most boring yet speedy things again and again.
