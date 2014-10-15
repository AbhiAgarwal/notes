fib x
	| x < 2 = x
	| otherwise = fib (x - 1) + fib (x - 2)

-- or 

foo 1 = 1
foo 2 = 2
foo x = foo(x - 1) + foo(x - 2)

main = print $ foo 2