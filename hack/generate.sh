if [ -z $1 ]; then
    echo Please supply a day number
    exit 1
fi

# year is the second argument, or 2022 if not supplied
year=${2:-2022}

dir=$year/day$1
mkdir -p $dir
touch $dir/input.txt $dir/instructions.txt $dir/testInput.txt

if [ ! -f $dir/main.go ]; then
    cat  <<EOF >$dir/main.go
package main 

import _ "embed"

//go:embed input.txt
var input string

//go:embed testInput.txt
var testInput string

func main() {

}
EOF
fi
