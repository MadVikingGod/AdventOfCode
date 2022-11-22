if [ -z $1 ]; then
    echo Please supply a day number
    exit 1
fi

# year is the second argument, or 2022 if not supplied
year=${2:-2022}

dir=$year/day$1
mkdir -p $dir
touch $dir/input.txt $dir/instructions.txt
cat  <<EOF >$dir/main.go
package main 

func main() {

}
EOF
