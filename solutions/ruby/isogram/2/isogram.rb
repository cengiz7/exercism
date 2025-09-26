module Isogram
  def self.isogram?(input)
    uniq_letters = input.downcase.chars.reject { |c| ' -'.include? c }
    uniq_letters == uniq_letters.uniq
  end
end