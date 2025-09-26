module Pangram
  def self.pangram? phrase
    phrase
      .gsub(/[^a-zA-Z]/, '')
      .downcase
      .chars
      .uniq
      .sort == ('a'..'z').to_a
  end
end