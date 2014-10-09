divisors :: Integer -> [Integer]
divisors 1 = [1]
divisors x = 1:[y | y <- [2..(x `div` 2)], x `mod` y == 0] ++ [x]

isPrime :: Integer -> Bool
isPrime x = divisors x == [1, x]

primeNums :: [Integer]
primeNums = take 100 [ x | x <- [2..], isPrime x]

main = putStrLn $ show $ primeNums