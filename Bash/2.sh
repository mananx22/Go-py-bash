# Task 2)

: '
take user input as string
convert string to int
write fibonacci sequence till that
'


# to directly recieve input from cmd use -> echo "hello $1" 

echo -n "Please enter the destination number:" # -n to read input from same line
read IN
A=0
B=1

main() {    
     #  echo $(($A+$IN)) rounded brackets are used to do calculation in bash scripting
     for ((i=0; i<IN; i++)) do        
        echo $A
        K=$((A + B))  
        A=$B 
        B=$K
        done
}

main

