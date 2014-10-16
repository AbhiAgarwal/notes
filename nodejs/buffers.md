- Pure javascript, while great with unicode-encoded strings, does not handle straight binary data very well.
- Strings are built-in to V8, and allocate memory within the VM. Buffers were added not to make all string operations faster, but to represent binary data, where as Strings are unicode.
	- When writing large amounts of data to a socket, it's much more efficient to have that data in binary format, vs having to convert from unicode.
	- So for common operations, like concat, I'm not surprised Strings are faster"
-

# Resources

- http://stackoverflow.com/questions/4901570/buffer-vs-string-speed-why-is-string-faster