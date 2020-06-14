# Usage: ruby compare_benchmarks.rb CURRENT_SHA PARENT_SHA

require 'json'

current = JSON.parse(File.read("bench-#{ARGV[0]}.json"))
parent = JSON.parse(File.read("bench-#{ARGV[1]}.json"))
keys = (current.keys + parent.keys).compact.uniq

max_key_length = keys.map(&:length).max + 2

puts "Comparing #{ARGV[0]} vs #{ARGV[1]}"
puts [
  "Bench".ljust(max_key_length),
  "Ops/sec".ljust(10),
  "Parent".ljust(10),
  "Delta".ljust(10),
  "Delta %".ljust(10),
].join('')

keys.each do |key|
  delta = (parent[key].to_f - current[key].to_f)/parent[key].to_f.round(1) if parent[key]
  delta_percentage = ((delta/parent[key].to_f) * 100).round(1) if delta
  puts [
    key.to_s.ljust(max_key_length),
    current[key].to_s.ljust(10),
    parent[key].to_s.ljust(10),
    delta.to_s.ljust(10),
    "#{delta_percentage}%".rjust(10)
  ].join(" ")
end
