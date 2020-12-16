if [ -z $1 ]; then
    echo Please supply a day number
    exit 1
fi

input=2020/day$1/input

awk 'BEGIN { print "package main\n\nvar input = [][]int{" } { print "\t{"$0"},"} END { print "}" }' $input.txt > $input.go