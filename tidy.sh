# A shell script to clean-up the index.html
# Skip the meta-generator tag

tidy -m --tidy-mark no webroot/index.html

