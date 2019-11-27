This program is used to calculate pairwise distance, optimally of a small FASTA file vs. a Large FASTA database. Golang is used due to its 
near C level efficiency and its simple user readability. The Large database is first converted into a binary file using the FASTAtoBin.go 
program. The bin file is then used as an input in the main GoTheDistance.go program, which outputs labels and relative distance beween each 
sequence in the compared FASTA file and the BIN database.