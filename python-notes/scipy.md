```
import scipy as sp

# getting data from file
data = sp.genfromtxt('filename', delimiter='\t')
# you have to specify the filename and the delimiter
# in order to be able to get the data from the file

# to think about the arrangement of the data we can do
# data.shape and see what it consists of

# to split the dimensions we can do
x = data[:, 0]
y = data[:, 1]

# the more the number of dimensions the more
# we are able to split
g = data[:, 2] 
# if we have 3 elements per sub-array

# to check the number of invalid data that has nan
# we can do
sp.sum(sp.isnan(y))
# we can also do sp.sum(np.isnan(y))

# then to filter we simply do
x = x[~sp.isnan(y)]
y = y[~sp.isnan(y)]

# to use plots and things we can use
# matplotlib.pyplot
# import matplotlib.pyplot as plt

```