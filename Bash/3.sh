# Task 3)

: '
make a swap function.
pass by value first.
prove that original variables are still unaffected.
pass by pointer/reference again.
prove that original variables are indeed affected.
https://stackoverflow.com/questions/540298/passing-arguments-by-reference
'
X=4
Y=8

echo -e "initial values of X is $X  & initial values of Y is $Y \n  " # -n to read input from same line

swapbyvalue() {  
    A=$X
    X=$Y
    Y=$A
}

swapbyreference(){
# as of Bash 4.3-alpha, you can use namerefs to pass function arguments by reference:
    local -n ref1=X
    local -n ref2=Y
    A=$ref1
    ref1=$ref2
    ref2=$A
}

swapbyvalue $X $Y
echo -e "values after swapping by values is X = $X & Y = $Y "

swapbyreference X Y
echo -e "values after swapping by reference is X = $X & Y = $Y "



# imperfect function. pass by value & reference resulting in same answer.