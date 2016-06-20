is_author = false
is_title = false
File.open("index.txt", "r") do |file|
  contents = file.read
  contents = contents.gsub(/\A---(.|\n)*?---/, '')
  puts contents
  file.close
end
File.open("index.txt", "r") do |file|
  file.readlines.each do |line|
    is_author = true if line.match("author\.")
    if is_author
      author = line.split(": ")[-1]
      puts author
    end
    is_title = true if line.match("title\.")
    if is_title
      title = line.split(": ")[-1]
      slug = title.downcase.strip.gsub(' ', '-').gsub(/[^\w-]/, '')
      title = title.gsub('"', '')
      puts slug
      puts title
    end
    is_author = false
    is_title = false
  end
end
