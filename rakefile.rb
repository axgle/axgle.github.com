task :default do
	sh "git commit -a -m '.'"
	sh "git push origin master"
end

task :lf do
  Dir["*.htm"].each do |f|
    d=IO.read(f).gsub(/\r\n/,"\n").gsub("\r","\n")
    puts f
    open(f,"w"){|f|f<<d}
  end  
end  