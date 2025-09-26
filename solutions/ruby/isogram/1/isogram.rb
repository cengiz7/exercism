module Isogram
  def self.isogram?(input)
    cleaned_string = input.downcase.chars.reject { |c| ' -'.include? c }
    cleaned_string.count == cleaned_string.uniq.count
  end
end