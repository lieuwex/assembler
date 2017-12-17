; setup values
	LDA X
	LDB Y

	SHRA
	NOTA
	ADDAB
	LDB 1
	JC .finish
	LDB 0
finish:
	NOTA
	ADDAB
	MAB
