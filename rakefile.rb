task :default do
	sh "git commit -a -m '.'"
	sh "git push origin master"
end
