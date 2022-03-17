printf = function(s, ...)
	return io.write(s:format(...))
end

printf("%s\n", "hello")
printf("my name is %s and age is %d\n", "tom", 12)
