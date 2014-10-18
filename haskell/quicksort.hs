f [] = []

-- Given a list xs that starts with the value x
-- ys: Take all the elements a that are less than or equal to x
-- zs: Take all the elements b that are greater than x
-- Recrusively sort ys with the function call f ys
-- Recursively sort zs with the function call f zs
-- All values less than x are sorted, and all the values greater or equal to x are sorted
-- Concatenate them
f (x:xs) = f  ++ [x] ++ f zs
		where
			ys = [a | a <- xs, a <= x]
			zs = [b | b <- xs, b > x]