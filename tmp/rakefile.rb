task :default do
	sh "flex a.I"
	sh "gcc lex.yy.c -lfl"
	
end	