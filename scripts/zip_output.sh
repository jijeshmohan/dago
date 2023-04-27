for file in ./out/*; do
  tar -czf "${file}.tar.gz" "$file"
done