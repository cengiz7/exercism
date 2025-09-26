module Transpose
  def self.transpose(input)
    lines = input.split("\n")
    max = lines.map(&:length).max
    lines.map {|line| line.ljust(max, "\0") }
         .map(&:chars)
         .transpose
         .map(&:join)
         .join("\n")
         .gsub(/\0+$/, "")
         .gsub(/\0/, " ")
  end
end
