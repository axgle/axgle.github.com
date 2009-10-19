domain=$_.chomp #ruby -n page_size.rb
puts "=>"+`curl -m 3 http://#{domain}`.size.to_s
