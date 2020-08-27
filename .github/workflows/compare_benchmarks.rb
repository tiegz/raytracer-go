# Usage: ruby compare_benchmarks.rb CURRENT_SHA PARENT_SHA

require 'json'

current = JSON.parse(File.read("bench-#{ARGV[0]}.json"))
parent = JSON.parse(File.read("bench-#{ARGV[1]}.json"))
keys = (current.keys + parent.keys).compact.uniq

max_key_length = keys.map(&:length).max + 5

puts "Comparing #{ARGV[0]} vs #{ARGV[1]}"
puts [
  "Name".ljust(max_key_length),
  "ns/op^".ljust(10),
  "ns/op".ljust(10),
  "delta".ljust(10),
  "delta %".rjust(10),
].join('')

def red(txt); "\033[0;31m#{txt}\033[0m"; end
def green(txt); "\033[0;32m#{txt}\033[0m"; end

keys.each do |key|
  parent_ns_per_op, current_ns_per_op = parent[key].to_f, current[key].to_f
  delta, delta_text = 0, "N/A"
  if parent[key]
    delta = (parent_ns_per_op - current_ns_per_op).round(3)
    delta_percentage = ((delta/parent_ns_per_op) * 100).round(1)
    # Highlight anything that's 50% slower or faster.
    delta_text = if delta_percentage < -10
      red("#{delta_percentage}%")
    elsif delta_percentage > 10
      green("#{delta_percentage}%")
    else
      "#{delta_percentage}%"
    end
  end
  puts [
    key.to_s.ljust(max_key_length),
    parent_ns_per_op.to_s.ljust(10),
    current_ns_per_op.to_s.ljust(10),
    delta.to_s.ljust(10),
    delta_text.rjust(10)
  ].join(" ")
end
