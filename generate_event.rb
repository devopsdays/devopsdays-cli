#!/usr/bin/env ruby

# require 'parse_date'
# we use parse_date by @shad https://gist.github.com/shad/114749

# ParseDate
#
# A helper for parsing dates of an unknown format.
#
# Takes a randomly formatted date and makes a best guess at the date, or
# throws an exception if there's no best guess.  Will take international date
# into account if you pass in a 'short form' date to help suggest a starting
# point for the search.
# Will also try Chronic parsing to get relative dates (yesterday, tomorrow
# in 3 days, etc.)
#
# Examples:
#   date = ParseDate.parse('08-09-1977')
#   date = ParseDate.parse('09-08-1977', :short_form => '%d/%m/%y')
#   date = ParseDate.parse('August 09, 1977')
#   date = ParseDate.parse('yesterday')
#
# ParseDate is free and open... you can redistribute it and/or modify
# it anyway you see fit.  No warrantee or guarantee provided :)

require 'date'
require 'rubygems'
require 'chronic'

module ParseDate
  # Parse a date of unknown format.
  # Optionally, you can pass in a "short_form" to suggest potential
  # characteristics expected.
  # It will be used to determine if the date is international format,
  # or US formatted
  # (Note: Date separators don't matter with short dates,
  # they will be stripped out)
  def self.parse( date_string, options={} )
    date = nil
    short_form = options[:short_form]

    if(day_before_month?(short_form))
      # Normalize separators
      shorted_date_string = date_string.gsub(/[\.-]/,"/")
      shorted_short_form = short_form.gsub(/[\.-]/,"/")

      if(long_year?(date_string))
        date ||= Date.strptime(shorted_date_string, shorted_short_form.gsub(/\%y/,"\%Y")) rescue nil
      else
        date ||= Date.strptime(shorted_date_string, shorted_short_form.gsub(/\%Y/, "\%y")) rescue nil
      end
    end
    date ||= Chronic.parse(date_string).to_date rescue nil
    date ||= Date.parse(date_string) # allow this exception through if we can't parse at all
    date
  end

  private

  # determine if this is an international date format (day before month before year)
  def self.day_before_month?(short_form)
    return false if short_form.nil?
    (/\%d/i =~ short_form) < (/\%m/i =~ short_form)
  end

  # Figure out if we have a long (4 digit) year
  def self.long_year?(date)
    # if no spaces...
    if(date.strip.index(' ').nil?)
      # split on -, /, or . (valid separators for dates I guess.)
      parts = date.split(/[-\/\.]/)
      # if the first bit, or last bit if 4 chars
      (parts.first.size == 4) || (parts.last.size == 4)
    end
  end

end

require 'yaml' # TODO: decide if I really need this

event_slug = '2015-paris'
city = 'Paris'

event_directory = "/Users/mattstratton/src/devopsdays-webby/site/content/events/#{event_slug}"
# TODO: take this as some kind of argument

# TODO: turn all these get date things into a module

def get_event_start(event_directory)
  file = File.open("#{event_directory}/_event_date_start.txt", "r")
  while !file.eof?
    line = file.readline
    event_start_date = line
    event_start_date = ParseDate.parse(event_start_date)
  end
  return event_start_date
end

def get_event_end(event_directory)
  file = File.open("#{event_directory}/_event_date_end.txt", "r")
  while !file.eof?
    line = file.readline
    event_end_date = line
    event_end_date = ParseDate.parse(event_end_date)
  end
  return event_end_date
end

def get_plat_sponsors(event_directory)
  # sponsor_start = false
  is_plat = false
  sponsor_array = ''
  File.open("#{event_directory}/_sponsors.txt", "r") do |file|
    file.readlines.each do |line|
      if line.match('@.sponsors')
        is_plat = true
      end
      if is_plat == true
        if line.match('%>')
          sponsor_array << line
          break
        end
        sponsor_array << line
      end
    end
  end
  sponsor_array.chomp!
  sponsor_array = sponsor_array[3...-2]
  eval(sponsor_array)
  # puts @psponsors
  # puts @psponsors.first
  a = Array.new
  @psponsors.each do |sponsor|
    sponsor.each do |key, value|
      if key == :image
        # puts value[0...-4]
        a.push value[0...-4]
      end
    end
  end
  puts a
  return a
end

def write_event_file(event_slug, city)
  unless File.exist?("/Users/mattstratton/src/probablyfine/#{event_slug}.yml")
    config = File.new("/Users/mattstratton/src/probablyfine/#{event_slug}.yml", 'w+')
    config.puts "name: #{event_slug}"
    config.close
  end
  event_data_file = YAML::load_file("/Users/mattstratton/src/probablyfine/#{event_slug}.yml")

  slug_array = event_slug.split('-') #TODO this need to not split becasue some slugs have dashes in them. Basically need to regex the stuff before and after the first dash. Will need to take the city as an argument
  year = slug_array.first
  city = city.capitalize
  status = 'past'
  startdate = get_event_start("/Users/mattstratton/src/devopsdays-webby/site/content/events/#{event_slug}")
  enddate = get_event_end("/Users/mattstratton/src/devopsdays-webby/site/content/events/#{event_slug}")
  cfp_date_start = ''
  cfp_date_end = ''
  cfp_date_announce = ''
  coordinates = ''
  nav_elements = Dir.entries("/Users/mattstratton/src/devopsdays-webby/site/content/events/#{event_slug}").select {|entry| File.directory? File.join("/Users/mattstratton/src/devopsdays-webby/site/content/events/#{event_slug}",entry) and !(entry =='.' || entry == '..' || entry == 'logos' || entry == 'images') }

  event_data_file = File.open("/Users/mattstratton/src/probablyfine/#{event_slug}.yml", 'w')
  event_data_file.puts "name: #{event_slug}"
  event_data_file.puts "year: \"#{year}\""
  event_data_file.puts "city: \"#{city}\""
  event_data_file.puts "friendly: \"#{event_slug}\""
  event_data_file.puts "status: \"#{status}\""
  event_data_file.puts "startdate: #{startdate}"
  event_data_file.puts "enddate: #{enddate}"
  event_data_file.puts "cfp_date_start: #{cfp_date_start}"
  event_data_file.puts "cfp_date_end: #{cfp_date_end}"
  event_data_file.puts "cfp_date_announce: #{cfp_date_announce}"
  event_data_file.puts "coordinates: #{coordinates}"
  event_data_file.puts "location: \"#{city}\""
  event_data_file.puts 'nav_elements:'
  nav_elements.each do |nav|
    event_data_file.puts "  - name: #{nav}"
  end











  # write to yaml for fun
  # event_data_file['year'] = year
  # event_data_file['city'] = city
  # nav_elements.each do |element|
  #   event_data_file['nav_elements'][element] = element
  # end
  # # event_data_file['nav_elements']['name'] = nav_elements
  # File.open("/Users/mattstratton/src/probablyfine/#{event_slug}.yml", 'w') do |h|
  #   h.write event_data_file.to_yaml
  # end

  puts "The year is #{year} and the city is #{city}"
end

# write_event_file(event_slug, city)
get_plat_sponsors("/Users/mattstratton/src/devopsdays-webby/site/content/events/#{event_slug}")


# Dir.glob("#{event_directory}/*.txt") do |item| # note one extra "*"
#   puts "working on: #{item}..."
# end


# this is some junk for doing crap with directories but ignoring the . and ..
# Dir.foreach(event_directory) do |item|
#   next if item == '.' or item == '..'
#   puts item
# end

# for subdirectories
# Dir.glob("**/*.txt") do |my_text_file| # note one extra "*"
#   puts "working on: #{my_text_file}..."
# end
