# Usage: ruby convert_benchmarks_to_json.rb FILENAME

require 'json'

puts JSON.pretty_generate(File.read(ARGV[0])
  .lines
  .select { |l| l.start_with?('Benchmark') }
  .reduce({}) { |dict, l|
    name, ops, ns_per_op, _ = l.split(' ')
      .map(&:strip)
    dict[name] = ns_per_op.to_f
    dict
  })
