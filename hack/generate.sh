if [ -z $1 ]; then
    echo Please supply a day number
    exit 1
fi

dir=2021/day$1
mkdir -p $dir
touch $dir/input.txt $dir/instructions.txt
cat  <<EOF >$dir/main.go
package main 

func main() {

}
EOF
