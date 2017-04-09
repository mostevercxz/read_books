# ML Notes

## 1.1 Introduction

### What is ml?
Arthur Samuel described it as: 
>the field of study that gives computers the ability to learn without being explicitly programmed.

That is an older, informal definition.

Tom Mitchell provides a more modern definition:
>A computer program is said to learn from experience E with respect to some class of tasks T and performance measure P, if its performance at tasks in T, as measured by P, improves with experience E.

### Supervised Learning:
Supervised learning is the machine learning task of inferring a function from labeled training data. The training data consist of a set of training examples. In supervised learning, each example is a pair consisting of an input object (typically a vector) and a desired output value (also called the supervisory signal).

In supervised learning, we are given a data set and already know what our correct output should look like, having the idea that there is a relationship between the input and the output.

Supervised learning problems are categorized into **regression** and **classification** problems. 

In a **regression** problem, we are trying to predict results within a **continuous** output, meaning that we are trying to map input variables to some continuous function. 

In a **classification** problem, we are instead trying to predict results in a **discrete** output. In other words, we are trying to map input variables into discrete categories.


### Unsupervised Learning
Unsupervised learning is a type of machine learning algorithm used to draw inferences from datasets consisting of input data without labeled responses. The most common unsupervised learning method is **cluster analysis**, which is used for exploratory data analysis to find hidden patterns or grouping in data.

Example:

* Clustering: Take a collection of 1,000,000 different genes, and find a way to automatically group these genes into groups that are somehow similar or related by different variables, such as lifespan, location, roles, and so on.
* Non-clustering: The "Cocktail Party Algorithm", allows you to find structure in a chaotic environment. (i.e. identifying individual voices and music from a mesh of sounds at a cocktail party).

## 1.2 Linear Regression with one variable
What is linear regression?
>In statistics, linear regression is an approach for modeling the relationship between a scalar dependent variable y(因变量) and one or more explanatory variables (or independent variables自变量) denoted X. The case of one explanatory variable is called simple linear regression.
### 1.2.1 Model and cost function
Recap the output of supervised learning: 
* Regression : real-valued output
* Classification : discrete-valued output

m = Number of training examples
x = input variables / features
y = output vars / target vars
(x(i),y(i)) = one training example

![alt text](http://pic.yupoo.com/mostevercxz/GmhXnLIc/U9Yru.png "input notation")

    Trainging set -> Learning algorithm -> h(hypothesis)
    Size of house(x) -> h -> Estimated price(y)

Why not call h a function ? historical reason...

#### single-variable/univariate linear regression.
Cost function(also called Squared error function/Mean squared error)
![costfunc math notation](http://pic.yupoo.com/mostevercxz/GmhZXtD7/medish.jpg
 "costfunc math notation")

Let theta0 = 0, a simple cost function(use concrete data sets to demostrate)

#### Contour plots/figures
A contour plot is a graphical technique for representing a 3-dimensional surface by plotting constant z slices, called contours, on a 2-dimensional format. That is, given a value for z, lines are drawn for connecting the (x,y) coordinates where that z value occurs.
A contour line of a two variable function has a constant value at all points of the same line.
![contour plot](http://pic.yupoo.com/mostevercxz/Gmi1FWAO/medish.jpg)
### 1.2.2 Parameter learning
**Gradient descent** is a more general algorithm, 
and is used not only in linear regression, but also used all over the place in ml.

    a:=b//assignment
    a=b//Truth assertion

![gradient descent algorithm](http://pic.yupoo.com/mostevercxz/Gmi2MrHu/medish.jpg)
alpha is learning rate.

**Batch gradient descent**, looks at every example in the entire training set on every step. 
#### Gradient Descent for linear regression
![theta0, theta1 formula](http://pic.yupoo.com/mostevercxz/Gmi4hkNN/medish.jpg)

![partial derivative](http://pic.yupoo.com/mostevercxz/Gmi4hoIX/medish.jpg)

## 1.3 Linear Algebra review
### Matrices and Vectors
Rectangualr array of numbers.(two-demension array)
4x2 matrix.(rows x columns)
A(i,j) refers i-th row, j-th column.

What is vector ? 
>An n * 1 matrix. (0-indexed vs 1-indexed, assume we are using 1-indexed vectors.)

Upper captical letters to denote matrix, low-case y to denote vectors/numbers/scalars... etc.
### Addition and Scalar Mulitplication(A * x)
so easy...
### Matrix Vector Multiplication
so easy..
### Matrix Multiplication Properties
* not commutative : A * B != B * A
* associative 结合性 : A * (B * C) = (A * B) * C
* Identity Matrix, I(n,n) 单位矩阵,对角为1, ones along the diagonals, A * I = I * A = A
### Inverse and Transpose
the inverse of A is A(-1) only when:
>A * A(-1) = A(-1) * A = I

Sometimes there is no inverse at all. singular matrix, zero matrix.

If A is square and determinant of A, det(A) != 0, then A has inverse.

Let A be an m * n matrix, Let B(n * m) = A(T) , B(i,j)=A(j,i).
