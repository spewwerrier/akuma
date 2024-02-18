why hashmaps?


names of manga are absurd
so instead of using hardcoded names as links
we map the name to a unique identifier in database (sqlite)


name               | identifier (unique)
----------------------------------------
8938467505p4       | Vagabond
3289376fkjg0       | Sousou No Frieren

this way
next time we open link, it will be
8938467505p4 -> Vagabond


next time we user requests 8938467505p4 
we find the corresponding directory
we go to manga/Vagabond
list all the cbz of manga/Vagabond/*.cbz
