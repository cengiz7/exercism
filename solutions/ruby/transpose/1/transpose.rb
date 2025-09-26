module Transpose
  def self.transpose(input)
    lines = input.split("\n")
    max = lines.map(&:length).max
    lines.map {|line| line.ljust(max).chars }
         .transpose
         .map(&:join)
         .join("\n")
         .rstrip
  end
end
