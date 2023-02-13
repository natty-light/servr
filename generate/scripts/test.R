library("rmarkdown")

args <- commandArgs(trailingOnly = TRUE)

out <- paste0(strsplit(args[1], ".csv")[1], ".pdf")

render(
  "/usr/src/server/scripts/test_markdown.Rmd",
  output_dir = "/usr/src/server/",
  output_file = out
)