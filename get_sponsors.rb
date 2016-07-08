# encoding: utf-8

require 'nokogiri'
require 'erb'

@eventhome = ""

path = File.join("../devopsdays-webby/site/content/events/2015-paris/_sponsors.txt")
doc = Nokogiri::HTML(ERB.new(File.read(path)).result(binding))

links = doc.xpath("//a[./img]")

export = links.each_with_object({}) do |element, h|
  h[element["href"]] = (element.first_element_child["src"])[/.*\/(.*?)\./,1]
end

puts export.class
