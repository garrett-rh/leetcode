if [ -z "$1" ]; then 
    target="./..."
else 
    target=${@:1}
fi

docker build -t leetcode . && \
docker run --rm leetcode ${target}

