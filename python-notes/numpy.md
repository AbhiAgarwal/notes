```
import numpy as np

# creating array as python array then
# converting it into a np array
a = [1, 2, 3, 4, 5]
b = np.array(a)

or

# creating array as np array
a = np.array([1, 2, 3, 4, 5])
# accessing eleent is easy as it is simply a[0], etc.

# n dimensions
a.ndim

# shape means how many elements and how many dimensions
# (number of elements, dimension per element)
a.shape

# reshape reshapes the array
b = a.reshape((5, 1))
# this would create a shape of (5, 1)
# where we have 2 dimensions as to access the array
# we must do b[0][0], etc. Therefore e.ndim is 2

# The funny thing starts when we realize just how much the NumPy package is
optimized.
# so if you want a true copy not a reference then when you copy do
c = b.copy()

# we can also apply things on the whole to each element
# if we want to multiply each element by 2
c = c*2
# would do the trick

# we can also filter data 
# if we do 
c > 4 
# then we get back a boolean array
# with indexes of elements that are greater than 4 as being
# true

# To get only the elements we can do
c[c > 4]

# we can also set those values if we are required to
# if you want to change the values of all above a
# certain range then
c[c > 4] = 3
# would do the trick

# There is also a clip function that you can give a lower and upper bound
c.clip(0, 4)
# which would set the lower and upper bound and cut the values 
# lower or higher towards that

# check if not a number
np.isnan(c) 
# this function checks whether or not something is not
# a number or NaN

# declering array with nan
c = np.array([1, 2, np,nan])

# to filter you can also do
c[~np.isnan(c)]
# which uses the filter method and only
# returns the elements that are not NaN

# there is a cool mean method, which returns
# the mean of the numbers in the array
np.mean(c[~np.isnan(c)])

# finds datatype of the functions
c.dtype
```
