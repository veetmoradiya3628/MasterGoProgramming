#!/bin/bash

mkdir -p sample_files

WORDS=(
  apple banana orange mango grape kiwi
  cloud docker kubernetes golang nodejs
  redis kafka postgres mysql rabbitmq
  system design microservice distributed
  concurrency thread mutex channel worker
  cache queue database network protocol
  api backend frontend server client
)

for i in {1..20}
do
    FILE="sample_files/file_${i}.txt"

    # Random word count between 50 and 200
    # 50 <-> 151
    WORD_COUNT=$((RANDOM % 151 + 50))

    > "$FILE"

    for ((j=1; j<=WORD_COUNT; j++))
    do
        WORD=${WORDS[$RANDOM % ${#WORDS[@]}]}
        echo -n "$WORD " >> "$FILE"
    done

    echo "" >> "$FILE"
done

echo "Generated 20 files in sample_files/"