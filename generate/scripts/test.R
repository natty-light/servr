library("jsonlite")
args <- commandArgs(trailingOnly = TRUE)
print(args)
cat(args[1], file = "data.txt")
