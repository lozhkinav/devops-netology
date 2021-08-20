#!/bin/bash
# display command line options

count=1
for param in "$@"; do
<<<<<<< HEAD
<<<<<<< HEAD
    echo "\$@ Parameter #$count = $param"
=======
    echo "Parameter: $param"
>>>>>>> 2f356db (git-rebase 1)
=======
    echo "Next parameter: $param"
>>>>>>> 1a12cea (git-rebase 2)
    count=$(( $count + 1 ))
done

echo "====="
